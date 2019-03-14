package conversions
import "jvmgo/instructions/base"
import "jvmgo/rtda"

type I2B struct{ base.NoOperandsInstruction }
type I2C struct{ base.NoOperandsInstruction }
type I2S struct{ base.NoOperandsInstruction }

type I2F struct{ base.NoOperandsInstruction }
type I2L struct{ base.NoOperandsInstruction }
type I2D struct{ base.NoOperandsInstruction }

func _popi(frame *rtda.Frame) int32{
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	return v1
}

func (self *I2B) Execute(frame *rtda.Frame)  {
	frame.OperandStack().PushInt(int32(int8(_popi(frame))))
}

func (self *I2C) Execute(frame *rtda.Frame)  {
	frame.OperandStack().PushInt(int32(uint16(_popi(frame))))
}

func (self *I2S) Execute(frame *rtda.Frame)  {
	frame.OperandStack().PushInt(int32(int16(_popi(frame))))
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
