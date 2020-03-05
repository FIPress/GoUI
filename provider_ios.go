// +build darwin
// +build arm arm64 sim

package goui

/*
#cgo darwin CFLAGS: -x objective-c
#cgo darwin LDFLAGS: -framework WebKit -framework Foundation -framework UIKit

#include <stdlib.h>
#include <string.h>
#import <UIKit/UIKit.h>
#include <WebKit/WebKit.h>
#include <objc/runtime.h>
#include "provider.h"

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

WKWebView *webView;

@interface ViewController : UIViewController

@end

@implementation ViewController

- (void)viewDidLoad {
    [super viewDidLoad];
    GoUIMessageHandler* handler = [[GoUIMessageHandler alloc] init];
    WKUserContentController *userContentController = [[WKUserContentController alloc] init];
    [userContentController addScriptMessageHandler:handler name:@"goui"];

    WKWebViewConfiguration *configuration = [[WKWebViewConfiguration alloc] init];
    configuration.userContentController = userContentController;
    webView = [[WKWebView alloc] initWithFrame:[[UIScreen mainScreen] bounds] configuration:configuration];
    [webView.configuration.preferences setValue:@YES forKey:@"developerExtrasEnabled"];

    NSBundle *bundle = [NSBundle mainBundle];
    NSString *path = [bundle pathForResource:@"ui/index" ofType:@"html"];
    NSURL *nsURL = [NSURL fileURLWithPath:path isDirectory:false];
    NSString *bundlePath = [[NSBundle mainBundle] resourcePath];
    NSString *webPath = [bundlePath stringByAppendingString:@"/ui"];
    NSURL *nsDir = [NSURL fileURLWithPath:webPath isDirectory:true];
    NSLog(@"dir:%@",nsDir);
    [webView loadFileURL:nsURL allowingReadAccessToURL:nsDir];
    [self.view addSubview:webView];
}

@end


@interface AppDelegate : UIResponder <UIApplicationDelegate>

@property (strong, nonatomic) UIWindow *window;

@end

@implementation AppDelegate
- (BOOL)application:(UIApplication *)application didFinishLaunchingWithOptions:(NSDictionary *)launchOptions {
    // Override point for customization after application launch.
    self.window = [[UIWindow alloc]initWithFrame:[[UIScreen mainScreen] bounds]];
    //UIView *view = [[UIView alloc]initWithFrame:[[UIScreen mainScreen] bounds]];
    ViewController *vc = [[ViewController alloc] init];
    [self.window setRootViewController:vc];
    [self.window makeKeyAndVisible];

    return YES;
}


- (void)applicationWillResignActive:(UIApplication *)application {
    // Sent when the application is about to move from active to inactive state. This can occur for certain types of temporary interruptions (such as an incoming phone call or SMS message) or when the user quits the application and it begins the transition to the background state.
    // Use this method to pause ongoing tasks, disable timers, and invalidate graphics rendering callbacks. Games should use this method to pause the game.
}


- (void)applicationDidEnterBackground:(UIApplication *)application {
    // Use this method to release shared resources, save user data, invalidate timers, and store enough application state information to restore your application to its current state in case it is terminated later.
    // If your application supports background execution, this method is called instead of applicationWillTerminate: when the user quits.
}


- (void)applicationWillEnterForeground:(UIApplication *)application {
    // Called as part of the transition from the background to the active state; here you can undo many of the changes made on entering the background.
}


- (void)applicationDidBecomeActive:(UIApplication *)application {
    // Restart any tasks that were paused (or not yet started) while the application was inactive. If the application was previously in the background, optionally refresh the user interface.
}


- (void)applicationWillTerminate:(UIApplication *)application {
    // Called when the application is about to terminate. Save data if appropriate. See also applicationDidEnterBackground:.
}


@end

void create() {
	@autoreleasepool {
		char* arg[] = {""};
        UIApplicationMain(1, arg, nil, NSStringFromClass([AppDelegate class]));
    }
}

void invokeJS(const char *js) {
	[webView evaluateJavaScript:[NSString stringWithUTF8String:js] completionHandler:^(id _Nullable response, NSError * _Nullable error) {
        //goUILog("response:%s,error:%s",[response UTF8String],[error UTF8String]);
    }];
}

void exitApp() {

}
*/
import "C"

func cCreate(cs C.WindowSettings, cMenuDefs *C.MenuDef, count C.int) {
	//cs := convertSettings(settings)
	C.create()
}

func cActivate() {

}

func cInvokeJS(js *C.char) {
	C.invokeJS(js)
}

func cExit() {
	//not supported
}
