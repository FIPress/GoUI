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
            //handleTest();
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

static int createWindow(WindowSettings settings) {
    if (gtk_init_check(0, NULL) == FALSE) {
        return -1;
    }
    fprintf(stderr, "new window\n");
    GtkWidget* window = gtk_window_new(GTK_WINDOW_TOPLEVEL);
    fprintf(stderr,"set title load:%s\n",settings.url);
    gtk_window_set_title((GtkWindow*)window, settings.title);
    gtk_window_set_position((GtkWindow*)window, GTK_WIN_POS_CENTER);
    gtk_window_set_default_size((GtkWindow*)window, settings.width, settings.height);

    WebKitUserContentManager *m = webkit_user_content_manager_new();
    webkit_user_content_manager_register_script_message_handler(m, "goui");
    g_signal_connect(m, "script-message-received",
                     G_CALLBACK(messageReceived), window);
    webview = WEBKIT_WEB_VIEW(webkit_web_view_new_with_user_content_manager(m));

    gtk_container_add(GTK_CONTAINER(window), GTK_WIDGET(webview));

    WebKitSettings *webKitSettings = webkit_web_view_get_settings(webview);
    webkit_settings_set_enable_javascript(webKitSettings,true);
    webkit_settings_set_allow_file_access_from_file_urls(webKitSettings,true);
    webkit_settings_set_allow_universal_access_from_file_urls(webKitSettings,true);

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

func (w *window) create(settings Settings) {
	//C.Create((*C.WindowSettings)(unsafe.Pointer(settings)))
	cs := convertSettings(settings)

	C.createWindow(cs)
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

//export handleTest
func handleTest() {
	println("test test")
}
