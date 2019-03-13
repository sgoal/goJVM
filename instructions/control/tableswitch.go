package comparisions
import "jvmgo/instructions/base"
import "jvmgo/rtda"

// access jump table by index and jump
type TABLE_SWITCH struct{ 
	defaultOffset			int32
	low						int32
	high					int32
	jumpOffsets				[]int32
 }

 func (self *TABLE_SWITCH) FetchOperands(frame *rtda.Frame)  {
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.low = reader.ReadInt32()
	self.high = reader.ReadInt32()
	jumpOffsetsCount := self.high - self.low + 1
	self.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (self *TABLE_SWITCH) Execute(frame *rtda.Frame)  {
	index := frame.OperandStack().PopInt()
	var offset int
	if index >= self.low && index <= self.high{
		offset = int(self.jumpOffsets[index-self.low])
	}else{
		offset = int(self.defaultOffset)
	}
	base.Branch(frame, offset)
}
