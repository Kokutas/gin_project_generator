package utils

import (
	"os"
)

func CreateFile(dir,file,content string) error {
	//以读写方式打开文件，如果不存在，则创建
	f, err := os.OpenFile(dir+"/"+file, os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString(content)
	return err
}
