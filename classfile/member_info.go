package classfile

type MemberInfo struct{
	cp				ConstantPool
	accessFlags		uint16
	nameIndex		uint16
	descriptorIndex	uint16
	attributes		[]AttributeInfo
}

func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	
}

func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo{

}

func (self *MemberInfo) AccessFlags() uint16{
	return slef.accessFlags
}

func (self* MemberInfo) Name() string{
	return self.cp.getUtf8(self.nameIndex)
}

func (self* MemberInfo) Descripter() string{
	return self.cp.getUtf8(self.descriptorIndex)
}