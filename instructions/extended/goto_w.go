package extended

import "jvmgo/instructions/base"
import "jvmgo/rtda"

type GOTO_W struct {
	offset	int
}

func (self *GOTO_W) FecthOperands(read *bease.BytecodeReader)  {
	self.offset = int(reader.ReadInt32())
}


func (self *GOTO_W) Execute(frame *rtda.Frame)  {
	base.Branch(frame, self.offset)
}
