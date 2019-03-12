package math
import "jvmgo/instructions/base"
import "jvmgo/rtda"

type DMUL struct{ base.Index8Instruction }
type FMUL struct{ base.Index8Instruction }
type IMUL struct{ base.Index8Instruction }
type LMUL struct{ base.Index8Instruction }

func (self *DMUL) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	res := v1* v2
	stack.PushDouble(res)
}

func (self *FMUL) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	res := v1* v2
	stack.PushFloat(res)
}

func (self *LMUL) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	res := v1* v2
	stack.PushLong(res)
}

func (self *IMUL) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	res := v1* v2
	stack.PushInt(res)
}