// +build android
// +build arm 386 amd64 arm64

package goui

/*
#cgo LDFLAGS: -landroid -llog

#include <jni.h>
#include <dlfcn.h>
#include <stdio.h>
#include <android/log.h>
#include "provider.h"

#define loge(...) __android_log_print(ANDROID_LOG_ERROR, "GoUI", __VA_ARGS__);
#define logd(...) __android_log_print(ANDROID_LOG_DEBUG, "GoUI", __VA_ARGS__);

static void alogd(const char *msg) {
	 __android_log_print(ANDROID_LOG_DEBUG, "GoUI", "%s",msg);
}

extern void handleClientReq(const char* s);
extern void invokeMain(uintptr_t ptr);

JavaVM* jvm=0;
jobject mainActivity = 0;
jmethodID evalJSID = 0;
jmethodID createWebViewID = 0;

JNIEXPORT jint JNICALL
JNI_OnLoad(JavaVM* vm, void* reserved) {
    JNIEnv* env;
    if ((*vm)->GetEnv(vm, (void**)&env, JNI_VERSION_1_6) != JNI_OK) {
        return -1;
    }

    return JNI_VERSION_1_6;
}

void callCreateWebView(WindowSettings settings) {
    logd("url:%s",settings.url);
    if(createWebViewID!=NULL) {
        JNIEnv *env;
        (*jvm)->AttachCurrentThread(jvm,&env,NULL);
        (*env)->CallVoidMethod(env,mainActivity, createWebViewID, (*env)->NewStringUTF(env,settings.url));
    }
}


JNIEXPORT jboolean JNICALL
Java_org_fipress_goui_android_GoUIActivity_invokeGoMain(
        JNIEnv *env,
        jobject  jobj) {
    (*env)->GetJavaVM(env,&jvm);
    mainActivity = (*env)->NewGlobalRef(env, jobj);
    jclass jcls = (*env)->GetObjectClass(env,jobj);
    evalJSID = (*env)->GetMethodID(env,jcls,"evalJavaScript","(Ljava/lang/String;)V");
    createWebViewID = (*env)->GetMethodID(env,jcls,"loadWebView","(Ljava/lang/String;)V");
	const char *dlsym_error = dlerror();

	uintptr_t goMainPtr = (uintptr_t)dlsym(RTLD_DEFAULT, "main.main");
    dlsym_error = dlerror();
    if(dlsym_error) {
        loge("dlsym_error:%s",dlsym_error);
        return 0;
    }
    invokeMain(goMainPtr);
	return 1;
}

//to invoke java
void invokeJS(const char* js) {
	if(evalJSID!=NULL) {
        JNIEnv *env;
        (*jvm)->AttachCurrentThread(jvm,&env,NULL);
        (*env)->CallVoidMethod(env,mainActivity, evalJSID, (*env)->NewStringUTF(env,js));
    }
}

JNIEXPORT void JNICALL
Java_org_fipress_goui_android_ScriptHandler_postMessage(
        JNIEnv *env,
        jobject jobj,
        jstring message) {
    const char *msg = (*env)->GetStringUTFChars(env,message,0);
	handleClientReq(msg);
	(*env)->ReleaseStringUTFChars(env,message,msg);
}

void create(WindowSettings settings) {
	callCreateWebView(settings);
}

void exitApp() {

}
*/
import "C"
import (
	"unsafe"
)

func cCreate(cs C.WindowSettings, cMenuDefs *C.MenuDef, count C.int) {
	C.create(cs)
}

func cActivate() {

}

func cInvokeJS(js *C.char) {
	C.invokeJS(js)
}

func cExit() {
	C.exitApp()
}

func Log(arg ...interface{}) {
	go func() {
		msg := fmt.Sprint(arg...)
		cMsg := C.CString(msg)
		defer C.free(unsafe.Pointer(cMsg))
		C.alogd(cMsg)
	}()
}
