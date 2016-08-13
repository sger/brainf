package main

const SIZE int = 65535

const (
	opEnd = iota
	opIncDp
	opDecDp
	opIncVal
	opDecVal
	opOut
	opIn
	opJmpFwd
	opJmpBck
)
