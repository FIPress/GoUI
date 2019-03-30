// +build linux

package goui

/*
#cgo linux openbsd freebsd pkg-config: gtk+-3.0 webkit2gtk-4.0

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdint.h>
#include <gtk/gtk.h>
#include <webkit2/webkit2.h>
#include "bridge.c"
//#include "_cgo_export.h"

//extern void handleTest();
extern void handleClientReq(_GoString_ s);

WebKitWebView* webview;

static void evalJSDone(GObject *object, GAsyncResult *result,
                       gpointer arg) {
    fprintf(stderr,"eval js finished\n");
}

static int evalJS(const char *js) {
    fprintf(stderr,"eval js:%s\n",js);
    webkit_web_view_run_javascript(webview, js, NULL,
                                  evalJSDone, NULL);
    return 0;
}

//menu
typedef void (*actionPtr)(void);

typedef struct Action {
    actionPtr func;
    const char* name;
} Action;

const int actionCount = 1;
const Action actionMap[] = {
        {&gtk_main_quit,"quit"}
};


void menuAction(GtkMenuItem *menuitem,
                gpointer     arg) {
    fprintf(stderr,"menu clicked:%s\n",(char *)arg);
    //
}

void standardMenuAction(GtkMenuItem *menuitem,
                gpointer     arg) {
    const char* name = (const char *)arg;
    fprintf(stderr,"standard menu clicked:%s\n",name);

    for(int i=0;i<actionCount;i++) {
        Action action = actionMap[i];
        if(strcmp(name,action.name) ==0 ) {
            action.func();
            return;
        }
    }
}

GtkWidget* createMenu(MenuDef menuDef) {
    GtkWidget* item;
    switch (menuDef.menuType) {
        case container: {
            item = gtk_menu_item_new_with_label (menuDef.title);
            GtkWidget* menu = gtk_menu_new();
            for(int i=0;i<menuDef.childrenCount;i++) {
                GtkWidget* child = createMenu(menuDef.children[i]);
                gtk_menu_shell_append(GTK_MENU_SHELL (menu), child);
            }
            gtk_menu_item_set_submenu (GTK_MENU_ITEM (item), menu);
            break;
        }
        case standard: {
            item = gtk_menu_item_new_with_label (menuDef.title);
            g_signal_connect(G_OBJECT(item), "activate",
                             G_CALLBACK(standardMenuAction), (char *)menuDef.action);
            break;
        }
        case custom: {
            item = gtk_menu_item_new_with_label (menuDef.title);
            g_signal_connect(G_OBJECT(item), "activate",
                             G_CALLBACK(menuAction),(char *) menuDef.action);
            break;
        }
        case separator: {
            item = gtk_separator_menu_item_new();
            break;
        }
    }
    return item;
}

GtkWidget* buildMenu(MenuDef* menuDefs, int count) {
    GtkWidget* bar = gtk_menu_bar_new ();

    for(int i=0;i<count;i++) {
        GtkWidget* item = createMenu(menuDefs[i]);
        gtk_menu_shell_append(GTK_MENU_SHELL(bar), item);
    }

    return bar;
}

static void messageReceived(WebKitUserContentManager *manager,
                            WebKitJavascriptResult   *js_result,
                            gpointer                  arg) {
    fprintf(stderr,"message received\n");
    JSCValue *value = webkit_javascript_result_get_js_value (js_result);
    if (jsc_value_is_string (value)) {
        JSCException *exception;
        gchar        *str_value;

        str_value = jsc_value_to_string (value);
        exception = jsc_context_get_exception (jsc_value_get_context (value));
        if (exception) {
            fprintf(stderr,"Error running javascript: %s", jsc_exception_get_message (exception));
        } else {
            fprintf(stderr,"Script result: %s\n", str_value);
            handleClientReq((_GoString_){str_value, strlen(str_value)});
        }

        g_free (str_value);
    } else {
        fprintf(stderr,"Error running javascript: unexpected return value");
    }
    webkit_javascript_result_unref (js_result);
}

static void loadChanged(WebKitWebView *webview,
                        WebKitLoadEvent event, gpointer arg) {
    fprintf(stderr,"load changed:%d\n",event);
    if (event == WEBKIT_LOAD_FINISHED) {
        fprintf(stderr,"load finished\n");
    }
}

static void windowDestroyed(gpointer arg) {
    fprintf(stderr,"destroy\n");
    gtk_main_quit();
}

static void exitApp() {
	gtk_main_quit();
}

static int create(WindowSettings settings,MenuDef* menuDefs,int menuCount) {
    if (gtk_init_check(0, NULL) == FALSE) {
        return -1;
    }
    fprintf(stderr, "new window\n");
    logging("new window");
    GtkWidget* window = gtk_window_new(GTK_WINDOW_TOPLEVEL);

    gtk_window_set_title((GtkWindow*)window, settings.title);
    gtk_window_set_position((GtkWindow*)window, GTK_WIN_POS_CENTER);
    gtk_window_set_default_size((GtkWindow*)window, settings.width, settings.height);

    WebKitUserContentManager *m = webkit_user_content_manager_new();
    webkit_user_content_manager_register_script_message_handler(m, "goui");
    g_signal_connect(m, "script-message-received",
                     G_CALLBACK(messageReceived), window);
    webview = WEBKIT_WEB_VIEW(webkit_web_view_new_with_user_content_manager(m));

    GtkWidget* box = gtk_box_new(GTK_ORIENTATION_VERTICAL,1);
    if(menuCount != 0) {
        GtkWidget* menuBar = buildMenu(menuDefs,menuCount);
        gtk_box_pack_start(GTK_BOX(box), menuBar,FALSE,FALSE,0);
    }

    gtk_box_pack_start(GTK_BOX(box), GTK_WIDGET(webview),TRUE,TRUE,0);
    gtk_container_add(GTK_CONTAINER(window), box);

    WebKitSettings *webKitSettings = webkit_web_view_get_settings(webview);
    webkit_settings_set_enable_javascript(webKitSettings,true);
    webkit_settings_set_allow_file_access_from_file_urls(webKitSettings,true);
    webkit_settings_set_allow_universal_access_from_file_urls(webKitSettings,true);

	//char path[255];
	//strncpy(path, settings.dir, sizeof(path));
	//strncat(path, settings.index, sizeof(path));
	//logging("sizeof:%d",sizeof(path));
	logging("url:%s",settings.url);
    webkit_web_view_load_uri(webview,settings.url);

    g_signal_connect(G_OBJECT(webview), "load-changed",
                     G_CALLBACK(loadChanged), NULL);

    fprintf(stderr,"add webview\n");

    if(settings.debug) {
        webkit_settings_set_enable_write_console_messages_to_stdout(webKitSettings, true);
        webkit_settings_set_enable_developer_extras(webKitSettings, true);
    }

    gtk_widget_show_all(window);
    g_signal_connect_swapped(G_OBJECT(window), "destroy",
                             G_CALLBACK(windowDestroyed), NULL);
    gtk_main();
    return 0;
}
*/
import "C"

import "unsafe"

type window struct {
}

func (w *window) create(settings Settings, menuDefs []MenuDef) {
	//C.Create((*C.WindowSettings)(unsafe.Pointer(settings)))
	cs := convertSettings(settings)
	cMenuDefs, count := convertMenuDefs(menuDefs)
	C.create(cs, cMenuDefs, count)
}

func (w *window) activate() {

}

func (w *window) invokeJS(js string) {
	cJs := C.CString(js)
	defer C.free(unsafe.Pointer(cJs))
	C.evalJS(cJs)
}

func (w *window) exit() {
	C.exitApp()
}
