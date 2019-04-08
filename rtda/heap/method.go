package heap
import "jvmgo/classfile"

type Method struct{
	ClassMember
	maxStack			uint
	maxLocals			uint
	code				[]byte
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i,cfMethod := range cfMethods{
		methods[i] = &Method{
			class : class,
		}
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttribute(cfMethod)
	}
}

func (self *Method) copyAttribute(cfMethod *classfile.MemberInfo){
	if codeAttr := cfMethod.CodeAttribute();codeAttr!=nil{
		self.maxLocals = codeAttr.MaxLocals()
		self.maxStack = codeAttr.MaxStack()
		self.code = codeAttr.Code()
	}
	
}