#include "util.h"
extern void goLog(const char *s);
static const int bufSize = 512;

inline void goUILog(const char *format, ...) {
	char buf[bufSize];
	va_list args;
    va_start(args,format);
	int len = vsnprintf(buf,bufSize, format,args);

	if(len < bufSize) {
		goLog(buf);
	} else {
		len++;
		char *tempBuf = 0;
		tempBuf = (char *)malloc(sizeof(char)*len);
		if(tempBuf != 0) {
		    vsnprintf(tempBuf,len, format,args);
		    goLog(tempBuf);
		    free(tempBuf);
		}
	}
	va_end(args);
}
