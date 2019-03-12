package math
import "jvmgo/instructions/base"
import "jvmgo/rtda"

type DADD struct{ base.Index8Instruction }
type FADD struct{ base.Index8Instruction }
type IADD struct{ base.Index8Instruction }
type LADD struct{ base.Index8Instruction }

func (self *DADD) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	res := v1+ v2
	stack.PushDouble(res)
}

func (self *FADD) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	res := v1+ v2
	stack.PushFloat(res)
}

func (self *LADD) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	res := v1+ v2
	stack.PushLong(res)
}

func (self *IADD) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	res := v1+ v2
	stack.PushInt(res)
}