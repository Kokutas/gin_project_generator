package utils

import (
	"os"
)

/**
 *  @Author: kokutas
 *  @Date: 2020/6/13 12:37
 *  @Description: 目录操作
**/
func CreateDir(dirs ...string) error {
	for _, dir := range dirs {
		err := os.MkdirAll(dir, 0777)
		if err != nil {
			return err
		}
		err = os.Chmod(dir, 0777)
		if err != nil {
			return err
		}
	}
	return nil
}
