package math
import "jvmgo/instructions/base"
import "jvmgo/rtda"

type DNEG struct{ base.Index8Instruction }
type FNEG struct{ base.Index8Instruction }
type INEG struct{ base.Index8Instruction }
type LNEG struct{ base.Index8Instruction }

func (self *DNEG) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	res := -v1
	stack.PushDouble(res)
}

func (self *FNEG) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	v1 := stack.PopFloat()
	res := -v1
	stack.PushFloat(res)
}

func (self *LNEG) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	res := -v1
	stack.PushLong(res)
}

func (self *INEG) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	res := -v1
	stack.PushInt(res)
}