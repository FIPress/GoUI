#include "provider.h"

#ifdef __cplusplus
extern "C" {
#endif

    typedef void (*fnGoUILog)(const char*);
    typedef void (*fnMenuClicked)(const char*);
    typedef void (*fnHandleClientReq)(const char*);

    __declspec(dllexport) void __cdecl seLogger(fnGoUILog fn);
    __declspec(dllexport) void __cdecl setHandleClientReq(fnHandleClientReq fn);

	__declspec(dllexport) void __cdecl create(WindowSettings settings, MenuDef* menuDefs, int menuCount);
	__declspec(dllexport) void __cdecl invokeJS(const char* js);
	__declspec(dllexport) void __cdecl exitWebview();

#ifdef __cplusplus
}
#endif