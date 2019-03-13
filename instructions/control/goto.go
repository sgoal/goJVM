package comparisions
import "jvmgo/instructions/base"
import "jvmgo/rtda"

//always branch
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame)  {
	base.Branch(frame,self.Offset)
}
