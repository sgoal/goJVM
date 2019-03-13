package conversions
import "jvmgo/instructions/base"
import "jvmgo/rtda"

type I2F struct{ base.Index8Instruction }
type I2L struct{ base.Index8Instruction }
type I2D struct{ base.Index8Instruction }

func _popi(frame *rtda.Frame) int32{
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	return v1
}

func (self *I2F) Execute(frame *rtda.Frame)  {
	frame.OperandStack().PushFloat(float32(_popi(frame)))
}

func (self *I2L) Execute(frame *rtda.Frame)  {
	frame.OperandStack().PushLong(int64(_popi(frame)))
}


func (self *I2D) Execute(frame *rtda.Frame)  {
	frame.OperandStack().PushDouble(float64(_popi(frame)))
}
