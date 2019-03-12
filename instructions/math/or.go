package math

import "jvmgo/instructions/base"
import "jvmgo/rtda"

type IOR struct{ base.NoOperandsInstruction }
type LOR struct{ base.NoOperandsInstruction }

func (self *IOR) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	v1, v2 := stack.PopInt(),stack.PopInt()
	res := v1 |v2
	stack.PushInt(res)
}

func (self *LOR) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	v1, v2 := stack.PopLong(),stack.PopLong()
	res := v1 |v2
	stack.PushLong(res)
}