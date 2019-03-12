package comparisions
import "jvmgo/instructions/base"
import "jvmgo/rtda"

type LCAMP struct{ base.NoOperandsInstruction }


func (self *LCAMP) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	v2, v1 := stack.PopLong(), stack.PopLong()
	if v1 > v2{
		stack.PushInt(1)
	}else if v1 == v2{
		stack.PushInt(0)
	}else{
		stack.PushInt(-1)
	}
}