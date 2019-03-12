package math
import "jvmgo/instructions/base"
import "jvmgo/rtda"

type IINC struct{ 
	Index		uint
	Const		int32
}

func (self *IINC) FetchOperands(reader *base.BytecodeReader)  {
	self.Index = uint(reader.ReadUint8())
	self.Const = int32(reader.ReadUint8())
}

func (self *IINC) Execute(frame *rtda.Frame)  {
	val := frame.LocalVars().GetInt(self.Index)
	frame.LocalVars().SetInt(self.Index,val + self.Const)

}
