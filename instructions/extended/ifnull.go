package extended

import "jvmgo/instructions/base"
import "jvmgo/rtda"

//branch if ref is null
type IFNULL struct {
	base.BranchInstruction
}

//branch if ref is not null
type IFNONNULL struct {
	base.BranchInstruction
}

func (self *IFNULL) Execute(frame *rtda.Frame)  {
	ref := frame.OperandStack().PopRef()
	if ref == nil{
		base.Branch(frame, self.Offset)
	}
}

func (self *IFNONNULL) Execute(frame *rtda.Frame)  {
	ref := frame.OperandStack().PopRef()
	if ref != nil{
		base.Branch(frame, self.Offset)
	}
}