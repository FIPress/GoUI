// +build darwin
// +build amd64,!sim

package goui

/*
#cgo darwin CFLAGS: -x objective-c
#cgo darwin LDFLAGS: -framework Cocoa -framework WebKit

#include <stdlib.h>
#include <string.h>
#include <Cocoa/Cocoa.h>
#include <mach-o/dyld.h>
#include <WebKit/WebKit.h>
#include <objc/runtime.h>
#include "provider.h"

extern void menuClicked(_GoString_ s);
extern void handleClientReq(const char* s);
//extern void goLog(_GoString_ s);

@interface GoUIMessageHandler : NSObject <WKScriptMessageHandler> {
}
@end

@implementation GoUIMessageHandler
- (void)userContentController:(WKUserContentController *)userContentController
      didReceiveScriptMessage:(WKScriptMessage *)message {
      goUILog("didReceiveScriptMessage: %s\n",[message.name UTF8String]);
    if ([message.name isEqualToString:@"goui"]) {
    	const char* str = [message.body UTF8String];
        goUILog("Received event %s\n", str);
        handleClientReq(str);
        //(_GoString_){str, strlen(str)}
    }
}
@end

//menu

@interface CustomAction : NSObject
@end

@implementation CustomAction
- (void)action:(id)sender {
    //NSLog(@"click menu:%@",[sender representedObject]);
    const char* str = [[sender representedObject] UTF8String];
    //menuClicked((_GoString_){str, strlen(str)});
    menuClicked((_GoString_){str, strlen(str)});
    goUILog("click menu: %s\n",str);
}
@end

@interface GoUIMenu : NSObject
@end

@implementation GoUIMenu
static CustomAction* customAction;
static NSDictionary* actionMap;
+ (void)initialize {
    customAction = [[CustomAction alloc] init];
    actionMap = [NSDictionary dictionaryWithObjectsAndKeys:
                 //app menu
                 [NSValue valueWithPointer:@selector(orderFrontStandardAboutPanel:)], @"about",
                 [NSValue valueWithPointer:@selector(hide:)], @"hide",
                 [NSValue valueWithPointer:@selector(hideOtherApplications:)], @"hideothers",
                 [NSValue valueWithPointer:@selector(unhideAllApplications:)], @"unhide",
                 [NSValue valueWithPointer:@selector(terminate:)], @"quit",
                 nil];
}

NSString * utf8(const char* cs) {
    NSString *ns = @"";
    if(cs) {
        ns = [NSString stringWithUTF8String:cs];
    }
    return ns;
}

NSMenuItem* createMenuItem(MenuDef def) {
    NSString *title = utf8(def.title);
    NSString *key = utf8(def.key);

    SEL act = NULL;
    if(def.menuType == standard) {
        id pointer = [actionMap objectForKey:[NSString stringWithUTF8String:def.action]];
        if(pointer) {
            act = [pointer pointerValue];
        }
    } else if(def.menuType == custom) {
        act = @selector(action:);
    }
    NSMenuItem *item = [[[NSMenuItem alloc] initWithTitle:title
                                                   action: act
                                            keyEquivalent:key]
                        autorelease];
    if(def.menuType == custom) {
        [item setRepresentedObject:[NSString stringWithUTF8String:def.action]];
        [item setTarget:customAction];
    }
    return item;
}

NSMenuItem* createMenu(MenuDef def) {
    NSString *title = utf8(def.title);

    NSMenu *menu = [[[NSMenu alloc] initWithTitle:title] autorelease];
    NSMenuItem *item = createMenuItem(def);
    [item setSubmenu:menu];

    for(int i=0;i<def.childrenCount;i++) {
        buildSubMenu(menu,def.children[i]);
    }

    return item;
}

void buildSubMenu(NSMenu* parent,MenuDef def) {
	goUILog("build menu, type: %d\n",def.menuType);
    if(def.menuType == container) {
        [parent addItem:createMenu(def)];
    } else if(def.menuType == separator) {
        [parent addItem:[NSMenuItem separatorItem]];
    } else {
        [parent addItem:createMenuItem(def)];
    }
}

+(void)buildMenu:(MenuDef[])defs count: (int)count {
    NSMenu *menubar = [[[NSMenu alloc] initWithTitle:@"menu bar"] autorelease];
    [NSApp setMainMenu:menubar];

	goUILog("buildMenu count: %d\n",count);
    for(int i=0; i<count;i++) {
        @autoreleasepool {
            buildSubMenu(menubar,defs[i]);
        }
    }
}
@end

//window

@interface WindowDelegate : NSObject <NSWindowDelegate> {
@private
    NSView* _view;
}
@property (nonatomic, assign) NSView* view;
@end

@implementation WindowDelegate
@synthesize view = _view;
- (void)windowDidResize:(NSNotification *)notification {
    goUILog("windowDidResize\n");
}

- (void)windowDidMiniaturize:(NSNotification *)notification{
    goUILog("windowDidMiniaturize\n");
}
- (void)windowDidEnterFullScreen:(NSNotification *)notification {
    goUILog("windowDidEnterFullScreen\n");
}
- (void)windowDidExitFullScreen:(NSNotification *)notification {
    goUILog("windowDidExitFullScreen\n");
}
- (void)windowDidBecomeKey:(NSNotification *)notification {
    goUILog("Window: become key\n");
}

- (void)windowDidBecomeMain:(NSNotification *)notification {
    goUILog("Window: become main\n");
}

- (void)windowDidResignKey:(NSNotification *)notification {
    goUILog("Window: resign key\n");
}

- (void)windowDidResignMain:(NSNotification *)notification {
    goUILog("Window: resign main\n");
}

- (void)windowWillClose:(NSNotification *)notification {
    [NSAutoreleasePool new];
    goUILog("NSWindowDelegate::windowWillClose\n");
    [NSApp terminate:NSApp];
}
@end

@interface GoUIWindow : NSObject {
@private
    WKWebView* webView;
}
@end

@implementation GoUIWindow
//@synthesize webView = webView_;
//static WKWebView* webView;
- (void)create:(struct WindowSettings)settings {
    @autoreleasepool {
    	WindowDelegate* delegate = [[WindowDelegate alloc] init];
        NSRect rect = NSMakeRect(0, 0, settings.width, settings.height);
        id window = [[NSWindow alloc]
                     initWithContentRect:rect
                     styleMask:(NSWindowStyleMaskTitled |
                                NSWindowStyleMaskClosable |
                                NSWindowStyleMaskMiniaturizable |
                                NSWindowStyleMaskResizable |
                                NSWindowStyleMaskUnifiedTitleAndToolbar )
                     backing:NSBackingStoreBuffered
                     defer:NO];

        delegate.view = [window contentView];
        [window setDelegate:delegate];
        [window cascadeTopLeftFromPoint:NSMakePoint(settings.left,settings.top)];
        [window setTitle:[NSString stringWithUTF8String:settings.title]];

        GoUIMessageHandler* handler = [[GoUIMessageHandler alloc] init];

        WKUserContentController *userContentController = [[WKUserContentController alloc] init];
        [userContentController addScriptMessageHandler:handler name:@"goui"];
        WKWebViewConfiguration *configuration = [[WKWebViewConfiguration alloc] init];
        configuration.userContentController = userContentController;

        webView = [[WKWebView alloc] initWithFrame:rect configuration:configuration];
        [webView.configuration.preferences setValue:@YES forKey:@"developerExtrasEnabled"];
        NSString *index = [NSString stringWithUTF8String:settings.index];
        if([index hasPrefix:@"http"]) {
        	NSURL *nsURL = [NSURL URLWithString:index];
        	NSURLRequest *requestObj = [NSURLRequest requestWithURL:nsURL];
			[webView loadRequest:requestObj];
        } else {
			NSString *bundlePath = [[NSBundle mainBundle] resourcePath];
			NSString *dir = [bundlePath stringByAppendingPathComponent:[NSString stringWithUTF8String:settings.webDir]];
			index = [dir stringByAppendingPathComponent:index];
			goUILog("bundlePath:%s",[bundlePath UTF8String]);
			goUILog("dir:%s",[dir UTF8String]);
			goUILog("index:%s",[index UTF8String]);

			NSURL *nsURL = [NSURL fileURLWithPath:index isDirectory:false];
			NSURL *nsDir = [NSURL fileURLWithPath:dir isDirectory:true];
			[webView loadFileURL:nsURL allowingReadAccessToURL:nsDir];
		}

        [[window contentView] addSubview:webView];
        [window makeKeyAndOrderFront:nil];
    }
}

-(void) evaluateJS:(NSString*)script {
    goUILog("evalue:%s",[script UTF8String]);
    [webView evaluateJavaScript:script completionHandler:^(id _Nullable response, NSError * _Nullable error) {
        //goUILog("response:%s,error:%s",[response UTF8String],[error UTF8String]);
    }];
}
@end

//app
@interface ApplicationDelegate : NSObject <NSApplicationDelegate> {
@private
    MenuDef* _menuDefs;
    int _menuCount;
}
@property (nonatomic, assign) struct MenuDef* menuDefs;
@property (assign) int menuCount;
@end

@implementation ApplicationDelegate
-(void)applicationWillFinishLaunching:(NSNotification *)aNotification
{
    goUILog("applicationWillFinishLaunching\n");
    [NSApplication sharedApplication];
    goUILog("_menuCount: %d\n",_menuCount);
    if(_menuCount!=0) {
    	[GoUIMenu buildMenu:_menuDefs count:_menuCount ];
    }
    [NSApp setActivationPolicy:NSApplicationActivationPolicyRegular];
}

-(void)applicationDidFinishLaunching:(NSNotification *)notification
{
	goUILog("applicationDidFinishLaunching\n");
    [NSApplication sharedApplication];

    [NSApp activateIgnoringOtherApps:YES];
}
@end

@interface GoUIApp : NSObject

@end

@implementation GoUIApp
static GoUIWindow* window;

+(void)initialize {}
+(void)start:(WindowSettings)settings menuDefs:(struct MenuDef[])menuDefs menuCount: (int)menuCount {
    @autoreleasepool {
        [NSApplication sharedApplication];
        [NSApp setActivationPolicy:NSApplicationActivationPolicyRegular];

        ApplicationDelegate* appDelegate = [[ApplicationDelegate alloc] init];
        appDelegate.menuDefs = menuDefs;
        appDelegate.menuCount = menuCount;
        goUILog("menuCount: %d\n",menuCount);
        [NSApp setDelegate:appDelegate];

        window = [[GoUIWindow alloc] init];
        [window create:settings];
        [NSApp run];
        //run();
        }
}

+(void)evaluateJS:(NSString*)js {
	[window evaluateJS:js];
}

+(void)exit {
    [NSApp terminate:NSApp];
}
@end

void create(WindowSettings settings, MenuDef* menuDefs, int menuCount) {
	[GoUIApp start:settings menuDefs:menuDefs menuCount:menuCount];
}

void invokeJS(const char *js) {
	goUILog("invokeJS:%s",js);
	[GoUIApp evaluateJS:[NSString stringWithUTF8String:js]];
}

void exitApp() {
	[GoUIApp exit];
}
*/
import "C"

func cCreate(cs C.WindowSettings, cMenuDefs *C.MenuDef, count C.int) {
	C.create(cs, cMenuDefs, count)
}

func cActivate() {

}

func cInvokeJS(js *C.char) {
	C.invokeJS(js)
}

func cExit() {
	C.exitApp()
}
