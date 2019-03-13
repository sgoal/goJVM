package conversions
import "jvmgo/instructions/base"
import "jvmgo/rtda"

type D2F struct{ base.NoOperandsInstruction }
type D2L struct{ base.NoOperandsInstruction }
type D2I struct{ base.NoOperandsInstruction }

func _popd(frame *rtda.Frame) float64{
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	return v1
}

func (self *D2F) Execute(frame *rtda.Frame)  {
	frame.OperandStack().PushFloat(float32(_popd(frame)))
}

func (self *D2L) Execute(frame *rtda.Frame)  {
	frame.OperandStack().PushLong(int64(_popd(frame)))
}


func (self *D2I) Execute(frame *rtda.Frame)  {
	frame.OperandStack().PushInt(int32(_popd(frame)))
}
