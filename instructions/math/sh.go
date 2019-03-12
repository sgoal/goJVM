package math
import "jvmgo/instructions/base"
import "jvmgo/rtda"

type ISHL struct{ base.NoOperandsInstruction } // int左位移
type ISHR struct{ base.NoOperandsInstruction } // int算术右位移
type IUSHR struct{ base.NoOperandsInstruction } // int逻辑右位移
type LSHL struct{ base.NoOperandsInstruction } // long左位移
type LSHR struct{ base.NoOperandsInstruction } // long算术右位移
type LUSHR struct{ base.NoOperandsInstruction } // long逻辑右位移

func (self* ISHL) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	shift := uint32(v2) & 0x1f //0x 0001 1111
	res := v1<<shift
	stack.PushInt(res)
}

func (self* ISHR) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	shift := uint32(v2) & 0x1f //0x 0001 1111  max shift right 32
	res := v1>>shift
	stack.PushInt(res)
}

func (self* IUSHR) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	shift := uint32(v2) & 0x1f //0x 0001 1111  max shift right 32
	res := int32(uint32(v1)>>shift)  //go has not usgined shift
	stack.PushInt(res)
}

func (self* LSHL) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	shift := uint32(v2) & 0x3f //0x 0011 1111
	res := v1<<shift
	stack.PushLong(res)
}

func (self* LSHR) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	shift := uint32(v2) & 0x3f //0x 0001 1111  max shift right 32
	res := v1>>shift
	stack.PushLong(res)
}

func (self* LUSHR) Execute(frame *rtda.Frame)  {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	shift := uint32(v2) & 0x3f //0x 0001 1111  max shift right 32
	res := int64(uint64(v1)>>shift)
	stack.PushLong(res)
}