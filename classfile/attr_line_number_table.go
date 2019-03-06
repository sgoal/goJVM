package classfile

type LineNumberTableAttribute struct{
	lineNumberTable		[]*LineNumberTableEntry
}

type LineNumberTableEntry struct{
	startPc			uint16
	lineNumber		uint16
}

func (self *LineNumberTableAttribute) readInfo(reader *ClassReader)  {
	lineNumberTableLen := reader.readUint16()
	self.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLen)
	for i:= range self.lineNumberTable{
		self.lineNumberTable[i] = readLineNumberTableEntry(reader)
	}
}

func readLineNumberTableEntry(reader *ClassReader) *LineNumberTableEntry{
	entry := &LineNumberTableEntry{
		startPc : 		reader.readUint16(),
		lineNumber:		reader.readUint16(),
	}
	return entry
}