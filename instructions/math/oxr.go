package math

import "jvmgo/instructions/base"
import "jvmgo/rtda"

type IXOR struct{ base.NoOperandsInstruction }
type LXOR struct{ base.NoOperandsInstruction }

func (self *IXOR) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	v2, v1 := stack.PopInt(),stack.PopInt()
	res := v1 ^v2
	stack.PushInt(res)
}

func (self *LXOR) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	v2, v1 := stack.PopLong(),stack.PopLong()
	res := v1 ^v2
	stack.PushLong(res)
}