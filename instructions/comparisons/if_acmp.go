package comparisions
import "jvmgo/instructions/base"
import "jvmgo/rtda"
import ."jvmgo/rtda/heap"

//branch if ref comparision succeeds
type IF_ACMPEQ struct{ base.BranchInstruction }	//==
type IF_ACMPNE struct{ base.BranchInstruction }	//!=


func (self *IF_ACMPEQ) Execute(frame *rtda.Frame)  {
	val1, val2 := _acmp(frame)
	if val1 ==val2{
		base.Branch(frame,self.Offset)
	}
}

func (self *IF_ACMPNE) Execute(frame *rtda.Frame)  {
	val1, val2 := _acmp(frame)
	if val1 != val2 {
		base.Branch(frame,self.Offset)
	}
}

func _acmp(frame *rtda.Frame)(val1, val2 *Object){
	val2 = frame.OperandStack().PopRef()
	val1 = frame.OperandStack().PopRef()
	return 
}