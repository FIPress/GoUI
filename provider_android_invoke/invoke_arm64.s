#include "textflag.h"
#include "funcdata.h"

TEXT ·InvokeMain(SB),$0-8
	MOVD fn+0(FP), R0
	BL (R0)
	RET
