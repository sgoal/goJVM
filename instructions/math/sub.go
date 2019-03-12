package math
import "jvmgo/instructions/base"
import "jvmgo/rtda"

type DSUB struct{ base.Index8Instruction }
type FSUB struct{ base.Index8Instruction }
type ISUB struct{ base.Index8Instruction }
type LSUB struct{ base.Index8Instruction }

func (self *DSUB) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	res := v1 - v2
	stack.PushDouble(res)
}

func (self *FSUB) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	res := v1- v2
	stack.PushFloat(res)
}

func (self *LSUB) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	res := v1- v2
	stack.PushLong(res)
}

func (self *ISUB) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	res := v1- v2
	stack.PushInt(res)
}