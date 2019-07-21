#define UNICODE

#include <iostream>
#include <direct.h>
#include <functional>
#include <hstring.h>
#include <winrt/Windows.Foundation.h>
#include <winrt/Windows.Storage.h>
#include <winrt/Windows.Storage.Streams.h>
#include <winrt/Windows.Web.UI.Interop.h>
#include "../provider_windows.h"

#pragma comment(lib, "user32")
#pragma comment(lib, "windowsapp")

using namespace winrt;
using namespace Windows::Foundation;
using namespace Windows::Storage;
using namespace Windows::Storage::Streams;
using namespace Windows::Web;
using namespace Windows::Web::UI;
using namespace Windows::Web::UI::Interop;


HWND hWnd = nullptr;
fnHandleClientReq handleClientReq = nullptr;
WebViewControl webView = nullptr;
fnGoUILog fnLog = nullptr;

/*static void printLog(const char *format, ...) {
    if(fnLog != nullptr) {
        fnLog();
    }
}*/

static void goLog(const char* log) {
    //printf("windows provider \n");
    //printf(log);
    if(fnLog != nullptr) {
        fnLog(log);
    }
}

static void localLog(const char* log) {
    printf("local log: %s\n",log);

}

static void localLog(hstring log) {
    std::cout << log.c_str() << std::endl;
}

namespace {
	class UriToStreamResolver : public winrt::implements<UriToStreamResolver, IUriToStreamResolver>
	{
	private:
		static hstring webBase;
	public:
		UriToStreamResolver()
		{
		}

		static void InitWebPath() {
			char cwd[MAX_PATH];
			_getcwd(cwd, MAX_PATH);
			//localLog("cwd:%s");
			//goLog(cwd);
            goUILog("cwd:%s",cwd);
			//webBase = Uri(winrt::to_hstring(cwd), L"web").Path();
			webBase = winrt::to_hstring(cwd) + L"\\web";
			goUILog("webBase:%s",webBase);
		}

		IAsyncOperation<IInputStream> UriToStreamAsync(Uri uri) const
		{
			Uri localUri = Uri(webBase + uri.Path());
			auto fOp = StorageFile::GetFileFromPathAsync(localUri.AbsoluteUri());
			localLog("local uri:");
			localLog(localUri.Path());
			goUILog("local uri:%s",localUri.Path());
			StorageFile f = fOp.get();
			//goLog("got file");
			auto sOp = f.OpenAsync(FileAccessMode::Read);
			IRandomAccessStream stream = sOp.get();
			//goLog("got stream");

			co_return stream.GetInputStreamAt(0);
		}

	};

}
hstring UriToStreamResolver::webBase ;

Rect getRect() {
	if (hWnd == nullptr ) {
		return Rect();
	}

	RECT r;
	GetClientRect(hWnd, &r);
	return Rect(r.left, r.top, r.right - r.left, r.bottom - r.top);
}

void resize() {
	if (webView == nullptr) {
		return;
	}
	Rect rect = getRect();
	webView.Bounds(rect);
}

void invokeScript(const char* js) {
	webView.InvokeScriptAsync(
		L"eval", single_threaded_vector<winrt::hstring>({ winrt::to_hstring(js) }));
}

void createWebView(const char* url) {
    //goUILog("create webview:%s",url);
	init_apartment(winrt::apartment_type::single_threaded);
	UriToStreamResolver::InitWebPath();
	WebViewControlProcess wvProcess = WebViewControlProcess();

	auto op = wvProcess.CreateWebViewControlAsync(
		reinterpret_cast<int64_t>(hWnd), getRect());
	op.Completed([url](IAsyncOperation<WebViewControl> const& sender, AsyncStatus args) {
		webView = sender.GetResults();
		webView.Settings().IsScriptNotifyAllowed(true);
		webView.IsVisible(true);

		webView.ScriptNotify([](auto const& sender, auto const& args) {
			if(handleClientReq != 0) {
			    handleClientReq(to_string(args.Value()).c_str());
			} else {
			    goUILog("no handler");
			}

		});
		/*
		webView.NavigationStarting([](auto const& sender, auto const& args) {
		});
		*/
		auto uri = webView.BuildLocalStreamUri(L"GoUI", winrt::to_hstring(url));
		goUILog("uri:",uri.Path());
		auto resolver = winrt::make_self<UriToStreamResolver>();
		IUriToStreamResolver r = resolver.as<IUriToStreamResolver>();
		webView.NavigateToLocalStreamUri(uri, r);

		});
}

//The Microsoft Edge WebView2 control enables you to host web content in your application using Microsoft Edge(Chromium) as the rendering engine.
/*
void createWebView2(HWND hWnd) {
	ComPtr<IWebView2WebView> webviewWindow;
	// Known issue - app needs to run on PerMonitorV2 DPI awareness for WebView to look properly
	SetProcessDpiAwarenessContext(DPI_AWARENESS_CONTEXT_PER_MONITOR_AWARE_V2);
	// Locate the browser and set up the environment for WebView
	// Use CreateWebView2EnvironmentWithDetails if you need to specify browser location, user folder, etc.
	CreateWebView2Environment(
		Callback<IWebView2CreateWebView2EnvironmentCompletedHandler>(
			[&webviewWindow, hWnd](HRESULT result, IWebView2Environment * env) -> HRESULT {
				// Create a WebView, whose parent is the main window hWnd
				env->CreateWebView(hWnd, Callback<IWebView2CreateWebViewCompletedHandler>
					([&webviewWindow, hWnd](HRESULT result, IWebView2WebView * webview) -> HRESULT {
						if (webview != nullptr) {
							webviewWindow = webview;
						}
						// Add a few settings for the webview
						// this is a redundant demo step as they are the default settings values
						IWebView2Settings* Settings;
						webviewWindow->get_Settings(&Settings);
						Settings->put_IsScriptEnabled(TRUE);
						Settings->put_AreDefaultScriptDialogsEnabled(TRUE);
						Settings->put_IsWebMessageEnabled(TRUE);
						// Resize WebView to fit the bounds of the parent window
						RECT bounds;
						GetClientRect(hWnd, &bounds);
						webviewWindow->put_Bounds(bounds);
						// Schedule an async task to navigate to Bing
						webviewWindow->Navigate(L"https://www.bing.com/");
						// Step 4 - Navigation events
						// Step 5 - Scripting
						// Step 6 - Communication between host and web content
						return S_OK;
						}).Get());
				return S_OK;
			}).Get());

}*/



void closeWindow() {
	PostQuitMessage(0);
}

LRESULT CALLBACK WndProc(HWND hWnd, UINT message, WPARAM wParam, LPARAM lParam)
{
	switch (message)
	{
	case WM_SIZE:
		resize();
		break;
	case WM_CLOSE:
		DestroyWindow(hWnd);
		break;
	case WM_DESTROY:
		closeWindow();
		break;
	default:
		return DefWindowProc(hWnd, message, wParam, lParam);
		break;
	}

	return 0;
}

void createWindow(WindowSettings settings) {
	HINSTANCE hInst = winrt::check_pointer(GetModuleHandle(nullptr));
	const auto className = L"GoUIWindow";
	WNDCLASSEX wce;
	ZeroMemory(&wce, sizeof(WNDCLASSEX));
	wce.cbSize = sizeof(WNDCLASSEX);
	wce.style = CS_HREDRAW | CS_VREDRAW;
	wce.lpfnWndProc = WndProc;
	wce.cbClsExtra = 0;
	wce.cbWndExtra = 0;
	wce.hInstance = hInst;
	//wce.hIcon = LoadIcon(hInstance, IDI_APPLICATION);
	//wce.hCursor = LoadCursor(NULL, IDC_ARROW);
	//wcex.hbrBackground = (HBRUSH)(COLOR_WINDOW + 1);
	//wce.lpszMenuName = NULL;
	//wce.hIconSm = LoadIcon(wcex.hInstance, IDI_APPLICATION);
	wce.lpszClassName = className;
	check_bool(::RegisterClassExW(&wce));

	hWnd = check_pointer(CreateWindow(
		className,
		to_hstring(settings.title).c_str(),
		WS_OVERLAPPEDWINDOW,
		CW_USEDEFAULT, CW_USEDEFAULT,
		settings.width,
		settings.height,
		NULL,
		NULL,
		hInst,
		NULL
	));

	ShowWindow(hWnd,
		SW_SHOW);
	UpdateWindow(hWnd);
	SetFocus(hWnd);
}

using dispatch_fn_t = std::function<void()>;

void run() {
	MSG msg;
	BOOL res;
	while ((res = GetMessage(&msg, nullptr, 0, 0)) != -1) {
		if (msg.hwnd) {
			TranslateMessage(&msg);
			DispatchMessage(&msg);
			continue;
		}
		if (msg.message == WM_APP) {
			auto f = (dispatch_fn_t*)(msg.lParam);
			(*f)();
			delete f;
		}
		else if (msg.message == WM_QUIT) {
			return;
		}
	}
}

 void __cdecl seLogger(fnGoUILog fn) {
    std::cout << "set logger" << fn << std::endl;
    fnLog = fn;
 }

void __cdecl setHandleClientReq(fnHandleClientReq fn) {
    std::cout << "set client" << fn << std::endl;
    handleClientReq = fn;
    std::cout << "after set" << handleClientReq << std::endl;
}

void __cdecl create(WindowSettings settings, MenuDef* menuDefs, int menuCount) {
		goUILog("create:");
		createWindow(settings);
		createWebView(settings.url);
		run();
		//void* webviewPtr = (void*)(new GoUIWebView);
		//GoUIWebView* webview = static_cast<GoUIWebView*>(webviewPtr);
		//webview->create(settings, menuDefs, menuCount);
		//return webviewPtr;
}

void __cdecl invokeJS(const char* js) {
		invokeScript(js);
		//GoUIWebView* webview = static_cast<GoUIWebView*>(webviewPtr);
		//webviewinvokeScriptinvokeJS(js);
}

void __cdecl exitWebview() {
		closeWindow();
		//GoUIWebView* webview = static_cast<GoUIWebView*>(webviewPtr);
		//webview->exitWindow();
		//delete(webview);
}
