package heap

import "jvmgo/classfile"

type MethodRef struct{
	MemberRef
	MethodRef	*Method
}

func newMehtodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodrefInfo) *MethodRef{
	ref := &MemberRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMethodrefInfo)
	return ref
}
