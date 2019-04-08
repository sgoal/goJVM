package heap
import "jvmgo/classfile"

type FieldRef struct{
	MemberRef
	field		*Field
}

func newFieldRef(cp *ConstantPool,	refInfo *classfile.ConstantFieldrefInfo) *FieldRef{
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantFieldrefInfo)
	return ref
}

