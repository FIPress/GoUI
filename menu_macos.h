#import <Cocoa/Cocoa.h>
#include "menu.h"
#include "util_cocoa.h"

@interface CustomAction : NSObject
@end

@interface GoUIMenu : NSObject
+(void)buildMenu:(struct MenuDef[])defs count: (int)count;
@end

