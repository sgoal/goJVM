package base
import "jvmgo/rtda"

type Instruction interface{
	FetchOperands (reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}

type NoOperandsInstruction struct{}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader)  {
	//nothing to do
}

//jump
type BranchInstruction struct{
	Offset 	int
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader)  {
	self.Offset = int(reader.ReadInt16())
}

//local variable index
type Index8Instruction struct{
	Index 	uint
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader)  {
	self.Index = uint(reader.ReadInt8())
}

//local constant pool index
type Index16Instruction struct{
	Index 	uint
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader)  {
	self.Index = uint(reader.ReadInt16())
}