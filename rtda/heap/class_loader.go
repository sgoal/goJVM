package heap

import "fmt"
import "jvmgo/classfile"
import "jvmgo/classpath"

type ClassLoader struct{
	cp			*classpath.Classpath
	classMap	map[string]*Class			//loaded class
}

func NewClassLoader(cp *classpath.Classpath) *ClassLoader{
	return &ClassLoader{
		cp:			cp,
		classMap:	make(map[string]*Class)
	}
}

func (self *ClassLoader) LoadClass(name string) *Class{
	if class, ok:=self.classMap[name]; ok{
		return class
	}
	return self.loadNonArrayClass(name)
}

func (self *ClassLoader) loadNonArrayClass(name string) *Class{
	data, entry := self.readClass(name)
	class := self.defineClass(data)
	link(class)
	fmt.Printf("[Loaded %s from %s]\n", name, entry)
	return class
}

func (self *ClassLoader) readClass(name string)  ([]byte, classpath.Entry){
	data, entry, err := self.cp.ReadClass(name)
	if err != nil{
		panic("java.lang.ClassNotFoundException: " + name)
	}
	return data, entry
}

func (self *ClassLoader) defineClass(data []byte) *Class{
	class := parseClass(data)
	class.loader = self
	resolveSuperClass(class)
	resolveInterfaces(class)
	self.classMap[class.name] = class
	return class
}