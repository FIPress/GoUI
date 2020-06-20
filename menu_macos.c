#include "menu_macos.h"

extern void menuClicked(const char * action);

@implementation CustomAction
- (void)action:(id)sender {
    //NSLog(@"click menu:%@",[sender representedObject]);
    const char* str = [[sender representedObject] UTF8String];
    //menuClicked((_GoString_){str, strlen(str)});
    menuClicked(str);
    goUILog("click menu: %s\n",str);
}
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