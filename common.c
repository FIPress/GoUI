typedef enum MenuType {
    container, //just a container item for sub items
    custom,
    standard,
    separator
} MenuType;

typedef struct MenuDef{
    MenuType type;
    const char* title;
    const char* action;
    const char* key;
    struct MenuDef* children;
    int childrenCount;
} MenuDef;


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


