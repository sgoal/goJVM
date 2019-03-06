package classfile
type SourceFileAttribute struct{
	cp					ConstantPool
	sourceFileindex		uint16
}

func (self* SourceFileAttribute) readInfo(reader *ClassReader)  {
	self.sourceFileindex = reader.readUint16()
}

func  (self* SourceFileAttribute) FileName() string {
	return self.cp.getUtf8(self.sourceFileindex)
}