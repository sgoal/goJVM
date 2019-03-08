package	rtda
import "math"
type OperandStack struct{
	size	uint
	slots	[]Slot
}

func newOperandStack(maxStack uint) *OperandStack{
	if maxStack >0{
		return &OperandStack{
			slots:		make([]Slot, maxStack),
		}
	}
	return nil
}

func (self *OperandStack) PushInt(val int32){
	self.slots[self.size].num = val
	self.size++
}

func (self *OperandStack) PopInt() int32{
	self.size--
	return self.slots[self.size].num
}

func (self *OperandStack) PushFloat(val float32){
	self.slots[self.size].num = int32(math.Float32bits(val))
	self.size++
}

func (self *OperandStack) PopFloat() float32{
	self.size--
	return math.Float32frombits(uint32(self.slots[self.size].num))
}


func (self *OperandStack) PushLong(val int64){
	self.slots[self.size].num = int32(val)
	self.slots[self.size+1].num = int32(val>>32)
	self.size+=2
}

func (self *OperandStack) PopLong() int64{
	high := uint32(self.slots[self.size].num)
	low := uint32(self.slots[self.size-1].num)
	self.size-=2
	return int64(high)<<32|int64(low)
}

func (self *OperandStack) PushDouble(val float64){
	self.PushLong(int64(math.Float64bits(val)))
}

func (self *OperandStack) PopDouble() float64{
	return math.Float64frombits(uint64(self.PopLong()))
}

func (self *OperandStack) PushRef(ref *Object){
	self.slots[self.size].ref = ref
	self.size ++
}

func (self *OperandStack) PopRef() *Object{
	self.size--
	ref := self.slots[self.size].ref
	self.slots[self.size].ref = nil
	return ref
}