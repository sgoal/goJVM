package comparisions
import "jvmgo/instructions/base"
import "jvmgo/rtda"

// access jump table by index and jump
type LOOKUP_SWITCH struct{ 
	defaultOffset			int32
	npairs					int32
	matchOffsets			[]int32
 }

 func (self *TABLE_SWITCH) FetchOperands(frame *rtda.Frame)  {
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.npairs = reader.ReadInt32()
	self.matchOffsets = reader.ReadInt32s(self.npairs *2)
}

func (self *TABLE_SWITCH) Execute(frame *rtda.Frame)  {
	key := frame.OperandStack().PopInt()
	var offset int
	len := self.npairs*2
	for i:=itn32(0);i<len;i+=2{
		if slef.matchOffsets[i] == key{
			offset = self.matchOffsets[i+1]
			base.Branch(frame, offset)
			return
		}
	}
	base.Branch(frame, self.defaultOffset)
}
