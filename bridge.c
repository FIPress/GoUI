#ifndef _BRIDGE_
#define _BRIDGE_

#include <stdlib.h>


typedef enum MenuType {
    container, //just a container item for sub items
    custom,
    standard,
    separator
} MenuType;

typedef struct WindowSettings{
    const char* title;
    const char* url;
    const char* dir;
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

/*
typedef struct MenuArray MenuArray;
struct MenuDef;

struct MenuArray {
    MenuDef* array;
    int size;
} ;

//int menuCount;
//MenuDef* menuDefs;

static MenuArray newMenuArray(int count) {
    MenuArray ma = {};
    ma.size = count;
    if(count != 0) {
        ma.array = (MenuDef*)malloc(sizeof(MenuDef)*count);
    }

    return ma;
}

MenuDef* initMenuDefs(int count) {
    menuCount = count;
    menuDefs = allocMenuDefArray(count);
    return menuDefs;
}
static MenuDef newMenuDef(MenuType type,const char* title,const char* action,const char* key,int childrenCount) {
    MenuDef* children = allocMenuDefArray(childrenCount);
    MenuDef def = {title,action,key,children,childrenCount,type};
    return def;
}
*/




static void addChildMenu(MenuDef* children, MenuDef child, int index) {
    children[index] = child;
}
#endif
/*
class MenuDef {
public:
    const char* title;
    const char* action;
    const char* key;

    MenuDef(const char* _title,const char* _action,const char* _key,int _childrenCount);
    ~MenuDef();
    AddChild(MenuDef child);
private:
    const int childrenCount;
    std::vector<MenuDef> menuDefs;
};

MenuDef::MenuDef(const char* _title,const char* _action,const char* _key,int _childrenCount):
            title(_title), action(_action),key(_key), childrenCount(_childrenCount) {
    menuDefs.reserve(childrenCount);
}

MenuDef::~MenuDef() {
    delete title;
    delete action;
    delete key;

    menuDefs.clear();
    menuDefs.shrink_to_fit();
}

MenuDef::AddChild(MenuDef child) {
    menuDefs.push_back(child);
}*/




