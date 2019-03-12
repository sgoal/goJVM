package math
import "jvmgo/instructions/base"
import "jvmgo/rtda"

type DDIV struct{ base.Index8Instruction }
type FDIV struct{ base.Index8Instruction }
type IDIV struct{ base.Index8Instruction }
type LDIV struct{ base.Index8Instruction }

func (self *DDIV) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	res := v1/v2
	stack.PushDouble(res)
}

func (self *FDIV) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	res := v1/v2
	stack.PushFloat(float32(res))
}

func (self *LDIV) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0{
		panic("java.lang.ArithmeticException: /by zero")
	}
	res := v1/v2
	stack.PushLong(res)
}

func (self *IDIV) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0{
		panic("java.lang.ArithmeticException: /by zero")
	}
	res := v1/v2
	stack.PushInt(res)
}