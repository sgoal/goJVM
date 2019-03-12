package conversions
import "jvmgo/instructions/base"
import "jvmgo/rtda"

type F2D struct{ base.Index8Instruction }
type F2L struct{ base.Index8Instruction }
type F2I struct{ base.Index8Instruction }

func _popf(frame *rtda.Frame) float32{
	stack := frame.OperandStack()
	v1 := stack.PopFloat()
	return v1
}

func (self *F2D) Execute(frame *rtda.Frame)  {
	frame.OperandStack().PushDouble(float64(_popf(frame)))
}

func (self *F2L) Execute(frame *rtda.Frame)  {
	frame.OperandStack().PushLong(int64(_popf(frame)))
}


func (self *F2I) Execute(frame *rtda.Frame)  {
	frame.OperandStack().PushInt(int32(_popf(frame)))
}
