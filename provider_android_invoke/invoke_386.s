#include "textflag.h"
#include "funcdata.h"

TEXT ·InvokeMain(SB),$0-4
	MOVL fn+0(FP), AX
	CALL AX
	RET
