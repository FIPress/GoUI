#ifndef _BRIDGE_
#define _BRIDGE_

#include "c/common.h"

typedef enum MenuType {
    container, //just a container item for sub items
    custom,
    standard,
    separator
} MenuType;

typedef struct WindowSettings{
    const char* title;
    const char* webDir;
    const char* index;
    const char* url;
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


#endif





