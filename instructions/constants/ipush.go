package constants
import "jvmgo/instructions/base"
import "jvmgo/rtda"

type BIPUSH struct{
	val	int8		//byte
}

type SIPUSH struct{
	val int16		//short
}

func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader)  {
	self.val = reader.ReadInt8()
}

func (self *BIPUSH) Execute(frame *rtda.Frame)  {
	frame.OperandStack().PushInt(int32(self.val))
}

func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader)  {
	self.val = reader.ReadInt16()
}

func (self *SIPUSH) Execute(frame *rtda.Frame)  {
	frame.OperandStack().PushInt(int32(self.val))
}