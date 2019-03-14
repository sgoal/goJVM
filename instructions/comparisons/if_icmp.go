package comparisions
import "jvmgo/instructions/base"
import "jvmgo/rtda"

//branch if int comparision succeeds
type IF_ICMPEQ struct{ base.BranchInstruction }	//==
type IF_ICMPNE struct{ base.BranchInstruction }	//!=
type IF_ICMPLT struct{ base.BranchInstruction }	//<
type IF_ICMPLE struct{ base.BranchInstruction }	//<=
type IF_ICMPGT struct{ base.BranchInstruction }	//>
type IF_ICMPGE struct{ base.BranchInstruction }	//>=



func (self *IF_ICMPEQ) Execute(frame *rtda.Frame)  {
	val1, val2 := _icmp(frame)
	if val1 ==val2{
		base.Branch(frame,self.Offset)
	}
}

func (self *IF_ICMPNE) Execute(frame *rtda.Frame)  {
	val1, val2 := _icmp(frame)
	if val1 != val2 {
		base.Branch(frame,self.Offset)
	}
}

func (self *IF_ICMPLT) Execute(frame *rtda.Frame)  {
	val1, val2 := _icmp(frame)
	if val1 < val2 {
		base.Branch(frame,self.Offset)
	}
}

func (self *IF_ICMPLE) Execute(frame *rtda.Frame)  {
	val1, val2 := _icmp(frame)
	if val1 <= val2 {
		base.Branch(frame,self.Offset)
	}
}

func (self *IF_ICMPGT) Execute(frame *rtda.Frame)  {
	val1, val2 := _icmp(frame)
	if val1 > val2 {
		base.Branch(frame,self.Offset)
	}
}

func (self *IF_ICMPGE) Execute(frame *rtda.Frame)  {
	val1, val2 := _icmp(frame)
	if val1 >= val2 {
		base.Branch(frame,self.Offset)
	}
}

func _icmp(frame *rtda.Frame)(val1, val2 int32){
	val2 = frame.OperandStack().PopInt()
	val1 = frame.OperandStack().PopInt()
	return
}