package util

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"os"
)

func WGet(url, fileName string) (md5Str string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		err = fmt.Errorf("http get failed,err: %s.", err.Error())
		return
	}
	if resp.StatusCode != 200 {
		err = fmt.Errorf("http request failed,result code %d.", resp.Status)
		return
	}
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		err = fmt.Errorf("open file: %s failed,err: %v.", fileName, err)
		return
	}
	defer f.Close()
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		err = fmt.Errorf("io copy to file failed,err: %v.", err)
		return
	}
	hash := md5.New()
	if _, err = io.Copy(hash, f); err != nil {
		err = fmt.Errorf("io copy to md5 hash failed,err: %v.", err)
		return
	}
	md5Str = string(hash.Sum(nil))
	return
}
