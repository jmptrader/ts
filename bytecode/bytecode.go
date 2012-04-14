package bytecode

import ()

const (
	NOP = iota 
	JUMP
	BRANCH
	VALUE
	BOUND
	FREE
	GLOBAL
	BOX
	UNDEFINE
	UNBOX
	UPDATE
	DEFINE
	PUSH
	FRAME
	SHUFFLE
	RETURN
	RETRACT
	CALL
	CLOSE
	CLOSEM
	PROLOG
	PROLOG_OPT
	PROLOG_REST
	EXTEND
	EXTENDA
	FINISH
	NEW
	GET
	GETM
	SET
	THIS
	LTHIS
	SUPER
	SOURCE
)

