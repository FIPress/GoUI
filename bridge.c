#ifndef _BRIDGE_
#define _BRIDGE_

#include <stdlib.h>
#include <stdio.h>
#include <stdarg.h>

typedef enum MenuType {
    container, //just a container item for sub items
    custom,
    standard,
    separator
} MenuType;

typedef struct WindowSettings{
    const char* title;
    const char* webDir;
    //const char* absPath;
    const char* index;
    int left;
    int top;
    int width;
    int height;
    int resizable;
    int debug;
} WindowSettings;

typedef struct MenuDef{
    const char* title;
    const char* action;
    const char* key;
    struct MenuDef* children;
    int childrenCount;
    MenuType menuType;
} MenuDef;

static MenuDef* allocMenuDefArray(int count) {
    if(count == 0) {
        return NULL;
    }
    return (MenuDef*)malloc(sizeof(MenuDef)*count);
}

static void addChildMenu(MenuDef* children, MenuDef child, int index) {
    children[index] = child;
}


extern void goLog(const char *s);
static const int bufSize = 512;

static void logging(const char *format, ...) {
	char buf[bufSize];
	va_list args;
    va_start(args,format);
	int len = vsnprintf(buf,bufSize, format,args);

	if(len < bufSize) {
		goLog(buf);
	} else {
		char tempBuf[len+1];
		vsnprintf(tempBuf,len+1, format,args);
		goLog(tempBuf);
	}
	va_end(args);
}

#endif





