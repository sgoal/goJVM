package classpath
import "os"
import "path/filepath"
type ClassPath struct {
	bootClassPath Entry
	extClassPath Entry
	userClassPath Entry
}

func Parse(jreOption, cpOption string) *ClassPath{
	cp := &ClassPath{}
	cp.parseBootAndExtClassPath(jreOption)
	cp.pareseUserClassPath(cpOption)
	return cp
}

func (self *ClassPath) parseBootAndExtClassPath(jreOption string){
	jreDir := getJreDir(jreOption)
	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClassPath = newWildcardEntry(jreLibPath)
	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib","ext","*")
	self.extClassPath = newWildcardEntry(jreExtPath)
}
func getJreDir(jreOption string) string{
	if jreOption != "" && exists(jreOption){
		return jreOption
	}
	if exists("./jre"){
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME");jh!=""{
		return filepath.Join(jh,"jre")
	}
	panic("can not find jre folder!")
}

func exists(path string) bool{
	if _,err := os.Stat(path);err!=nil{
		if os.IsNotExist(err){
			return false	
		}
	}
	return true;
}

func (self *ClassPath) pareseUserClassPath(cpOption string){
	if cpOption == ""{
		cpOption = "."
	}
	self.userClassPath = newEntry(cpOption)
}

func (self *ClassPath) ReadClass(className string) ([]byte, Entry, error){
	className = className + ".class"
	if data, entry, err := self.bootClassPath.readClass(className); err == nil{
		return data, entry, nil
	}
	if data, entry, err := self.extClassPath.readClass(className); err == nil{
		return data, entry, nil
	}
	return self.userClassPath.readClass(className)
}


func (self *ClassPath) String() string {
	return self.userClassPath.String()
}