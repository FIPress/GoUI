#ifndef _GOUI_UTIL_OJ_
#define _GOUI_UTIL_OJ_

#include <Cocoa/Cocoa.h>
#include "util.h"

inline NSString * utf8WithLength(const char* cs, int len) {
    NSString *ns = @"";
    if(cs!=NULL && cs[0]!='\0') {
        if(len==0) {
            ns = [NSString stringWithUTF8String:cs];
        } else {
            NSString *s = [[NSString alloc] initWithBytes:cs length:len encoding:NSUTF8StringEncoding];
            ns = [s autorelease];
        }

    }
    return ns;
}

inline NSString * utf8(const char* cs) {
    return utf8WithLength(cs,0);
}

#endif