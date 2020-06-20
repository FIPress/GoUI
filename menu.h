#ifndef _GOUI_MENU_
#define _GOUI_MENU_

#include <stdlib.h>
#include "util.h"

typedef enum MenuType {
    container, //just a container item for sub items
    custom,
    standard,
    separator
} MenuType;

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
        return 0;
    }
    return (MenuDef*)malloc(sizeof(MenuDef)*count);
}

static void addChildMenu(MenuDef* children, MenuDef child, int index) {
    children[index] = child;
}

#endif