package math

import "jvmgo/instructions/base"
import "jvmgo/rtda"

type IAND struct{ base.NoOperandsInstruction }
type LAND struct{ base.NoOperandsInstruction }

func (self *IAND) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	v1, v2 := stack.PopInt(),stack.PopInt()
	res := v1 &v2
	stack.PushInt(res)
}

func (self *LAND) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	v1, v2 := stack.PopLong(),stack.PopLong()
	res := v1 &v2
	stack.PushLong(res)
}