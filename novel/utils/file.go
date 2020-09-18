/*
* 作者：刘时明
* 时间: 2019/9/30-10:41
* 作用：
 */
package utils

import "os"

// PathExists 目录是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
