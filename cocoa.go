// +build darwin

package goui

/*
#cgo darwin CFLAGS: -x objective-c
#cgo darwin LDFLAGS: -framework Cocoa -framework WebKit

#include "common.c"
#include <stdlib.h>
#include <string.h>
#include <Cocoa/Cocoa.h>
#include <mach-o/dyld.h>
#include <WebKit/WebKit.h>
#include <objc/runtime.h>

extern void handleClientReq(_GoString_ s);

//menu

@interface CustomAction : NSObject
@end

@implementation CustomAction
- (void)action:(id)sender {
    //NSLog(@"click menu:%@",[sender representedObject]);
    printf("click menu: %s",[[sender representedObject] UTF8String]);
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
    if(def.type == standard) {
        id pointer = [actionMap objectForKey:[NSString stringWithUTF8String:def.action]];
        if(pointer) {
            act = [pointer pointerValue];
        }
    } else if(def.type == custom) {
        act = @selector(action:);
    }
    NSMenuItem *item = [[[NSMenuItem alloc] initWithTitle:title
                                                   action: act
                                            keyEquivalent:key]
                        autorelease];
    if(def.type == custom) {
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
    if(def.type == container) {
        [parent addItem:createMenu(def)];
    } else if(def.type == separator) {
        [parent addItem:[NSMenuItem separatorItem]];
    } else {
        [parent addItem:createMenuItem(def)];
    }
}

+(void)buildMenu:(MenuDef[])defs size: (int)size {
    NSMenu *menubar = [[[NSMenu alloc] initWithTitle:@"menu bar"] autorelease];
    [NSApp setMainMenu:menubar];

    for(int i=0; i<size;i++) {
        @autoreleasepool {
            buildSubMenu(menubar,defs[i]);
        }
    }
}
@end

//window
@interface GoUIMessageHandler : NSObject <WKScriptMessageHandler> {
}
@end

@implementation GoUIMessageHandler
- (void)userContentController:(WKUserContentController *)userContentController
      didReceiveScriptMessage:(WKScriptMessage *)message {
      printf("didReceiveScriptMessage: %s\n",[message.name UTF8String]);
    if ([message.name isEqualToString:@"goui"]) {
    	const char* str = [message.body UTF8String];
        printf("Received event %s\n", str);
        handleClientReq((_GoString_){str, strlen(str)});
    }
}
@end

@interface WindowDelegate : NSObject <NSWindowDelegate> {
@private
    NSView* _view;
}
@property (nonatomic, assign) NSView* view;
@end

@implementation WindowDelegate
@synthesize view = _view;
- (void)windowDidResize:(NSNotification *)notification {
    printf("windowDidResize\n");
}

- (void)windowDidMiniaturize:(NSNotification *)notification{
    printf("windowDidMiniaturize\n");
}
- (void)windowDidEnterFullScreen:(NSNotification *)notification {
    printf("windowDidEnterFullScreen\n");
}
- (void)windowDidExitFullScreen:(NSNotification *)notification {
    printf("windowDidExitFullScreen\n");
}
- (void)windowDidBecomeKey:(NSNotification *)notification {
    printf("Window: become key\n");
}

- (void)windowDidBecomeMain:(NSNotification *)notification {
    printf("Window: become main\n");
}

- (void)windowDidResignKey:(NSNotification *)notification {
    printf("Window: resign key\n");
}

- (void)windowDidResignMain:(NSNotification *)notification {
    printf("Window: resign main\n");
}

- (void)windowWillClose:(NSNotification *)notification {
    [NSAutoreleasePool new];
    printf("NSWindowDelegate::windowWillClose\n");
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
        NSURL *nsURL = [NSURL fileURLWithPath:[NSString stringWithUTF8String:settings.url] isDirectory:false];
        NSURL *nsDir = [NSURL fileURLWithPath:[NSString stringWithUTF8String:settings.dir] isDirectory:true];
        [webView loadFileURL:nsURL allowingReadAccessToURL:nsDir];
        [[window contentView] addSubview:webView];
        [window makeKeyAndOrderFront:nil];
    }
}

-(void) evaluateJS:(NSString*)script {
    [webView evaluateJavaScript:script completionHandler:^(id _Nullable response, NSError * _Nullable error) {
        //printf("response:%s,error:%s",[response UTF8String],[error UTF8String]);
    }];
}
@end

//app
@interface ApplicationDelegate : NSObject <NSApplicationDelegate> {
@private
    MenuDef* _menuDefs;
    int _menuSize;
}
@property (nonatomic, assign) struct MenuDef* menuDefs;
@property (assign) int menuSize;
@end

@implementation ApplicationDelegate
-(void)applicationWillFinishLaunching:(NSNotification *)aNotification
{
    printf("applicationWillFinishLaunching\n");
    [NSApplication sharedApplication];
    if(_menuSize!=0) {
    	[GoUIMenu buildMenu:_menuDefs size:_menuSize ];
    }
    [NSApp setActivationPolicy:NSApplicationActivationPolicyRegular];
}

-(void)applicationDidFinishLaunching:(NSNotification *)notification
{
	printf("applicationDidFinishLaunching\n");
    [NSApplication sharedApplication];

    [NSApp activateIgnoringOtherApps:YES];
}
@end

@interface GoUIApp : NSObject

@end

@implementation GoUIApp
static GoUIWindow* window;
+(void)initialize {}
+(void)start:(WindowSettings)settings menuDefs:(struct MenuDef[])menuDefs menuSize: (int)menuSize {
    @autoreleasepool {
        [NSApplication sharedApplication];
        [NSApp setActivationPolicy:NSApplicationActivationPolicyRegular];

        ApplicationDelegate* appDelegate = [[ApplicationDelegate alloc] init];
        appDelegate.menuDefs = menuDefs;
        appDelegate.menuSize = menuSize;
        [NSApp setDelegate:appDelegate];

        window = [[GoUIWindow alloc] init];
        [window create:settings];
        [NSApp run];
        }
}

+(void)evaluateJS:(NSString*)js {
	[window evaluateJS:js];
}

+(void)exit {
    [NSApp terminate:NSApp];
}
@end

void create(WindowSettings settings,MenuDef[] menuDefs,int menuSize) {
	[GoUIApp start:settings menuDefs:menuDefs menuSize:menuSize];
}

void invokeJS(const char *js) {
	[GoUIApp evaluateJS:[NSString stringWithUTF8String:js]];
}

void exitApp() {
	[GoUIApp exit];
}
*/
import "C"
import (
	"unsafe"
)

type window struct {
}

func (w *window) create(settings Settings, menuDefs []MenuDef) {
	//C.Create((*C.WindowSettings)(unsafe.Pointer(settings)))
	cs := toCSettings(settings)
	cMenuDefs, size := toMenuDefs(menuDefs)
	C.create(cs, cMenuDefs, size)
}

func (w *window) activate() {

}

func (w *window) invokeJS(js string) {
	cJs := C.CString(js)
	defer C.free(unsafe.Pointer(cJs))
	C.invokeJS(cJs)
}

func (w *window) exit() {
	C.exitApp()
}
