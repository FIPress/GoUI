#include "textflag.h"
#include "funcdata.h"

TEXT ·InvokeMain(SB),$0-8
	MOVQ fn+0(FP), AX
	CALL AX
	RET
