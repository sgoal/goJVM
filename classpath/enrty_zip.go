package classpath
import "errors"
import "archive/zip"
import "io/ioutil"
import "path/filepath"
type ZipEntry struct {
	absDir 	string
}

func newZipEntry(path string)  *ZipEntry{
	absDir,err :=filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absDir}
}

func (self *ZipEntry) readClass(className string) ([]byte, Entry, error){
	r, err := zip.OpenReader(self.absDir)
	if err != nil{
		return nil, nil, err
	}
	defer r.Close()
	for _,f := range r.File{
		if f.Name == className{
			rc, err := f.Open()
			if nil != err{
				return nil, nil, err
			}
			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil{
				return nil, nil, err
			}
			return data, self, nil
		}
	}
	return nil,nil,errors.New("class not found:" + className)
}

func (self *ZipEntry) String() string{
	return self.absDir
}