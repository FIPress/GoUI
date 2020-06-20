#ifndef _GOUI_WINDOW_
#define _GOUI_WINDOW_

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

#endif