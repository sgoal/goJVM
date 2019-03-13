package conversions
import "jvmgo/instructions/base"
import "jvmgo/rtda"

type L2F struct{ base.NoOperandsInstruction }
type L2D struct{ base.NoOperandsInstruction }
type L2I struct{ base.NoOperandsInstruction }

func _popl(frame *rtda.Frame) int64{
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	return v1
}

func (self *L2F) Execute(frame *rtda.Frame)  {
	frame.OperandStack().PushFloat(float32(_popl(frame)))
}

func (self *L2D) Execute(frame *rtda.Frame)  {
	frame.OperandStack().PushDouble(float64(_popl(frame)))
}


func (self *L2I) Execute(frame *rtda.Frame)  {
	frame.OperandStack().PushInt(int32(_popl(frame)))
}
