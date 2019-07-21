#include "textflag.h"
#include "funcdata.h"

TEXT ·InvokeMain(SB),$0-4
	MOVW fn+0(FP), R0
	BL (R0)
	RET
