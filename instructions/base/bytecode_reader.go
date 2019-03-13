package base
type BytecodeReader struct{
	code		[]byte
	pc			int
}

func (self *BytecodeReader) ReadUint8()  uint8{
	i:= self.code[self.pc]
	self.pc++
	return i
}

func (self *BytecodeReader) ReadInt8()  int8{
	return int8(self.ReadUint8())
}


func (self *BytecodeReader) ReadUint16()  uint16{
	high:= uint16(self.code[self.pc])
	low := uint16(self.code[self.pc+1])
	self.pc+=2
	return (high<<8)|low
}

func (self *BytecodeReader) ReadInt16() int16{
	return int16(self.ReadUint16())
}

func (self *BytecodeReader) ReadUint32()  uint32{
	high:= uint32(self.ReadUint16())
	low := uint32(self.ReadUint16())
	return (high<<16)|low
}

func (self *BytecodeReader) ReadInt32() int32{
	return int32(self.ReadUint32())
}

func (self *BytecodeReader) ReadInt32s(n int32) []int32{
	ints := make([]int32, n)
	for i := range ints{
		ints[i] = self.ReadInt32()
	}
	return ints
}
func (self *BytecodeReader) SkipPadding(){
	for self.pc %4 != 0{
		self.ReadUint8()
	}
}

