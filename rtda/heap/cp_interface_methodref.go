package heap

import "jvmgo/classfile"

type InterfaceMethodRef struct{
	MemberRef
	method *MemberRef
}

func newInterfaceMethodRef(cp *ConstantPool, refInfo *classfile.ConstantInterfaceMethodrefInfo)*InterfaceMethodRef{
	ref := InterfaceMethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}