#ifndef _GOUI_UTIL_
#define _GOUI_UTIL_

#include <stdlib.h>
#include <stdio.h>
#include <stdarg.h>
#include <string.h>

void goUILog(const char *format, ...);

inline int notEmpty(const char* s) {
    return s!=0 && s[0]!='\0';
}


#endif