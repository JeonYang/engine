package util

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func WGet(url, fileName string) (md5Str string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		err = fmt.Errorf("http get failed,err: %v.", err)
		return
	}
	if resp.StatusCode != 200 {
		err = fmt.Errorf("http request failed,result code %d.", resp.StatusCode)
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

//  version1>version2 	1
//  version1=version2	0
//  version1<version2	-1
func CompareVersion(v1, v2 string) int {
	version1 := strings.Split(v1, ".")
	version2 := strings.Split(v2, ".")
	re := regexp.MustCompile(`\d+`)
	for i := 0; i < len(version1); i++ {
		if i >= len(version2) {
			return 1
		}
		num1, _ := strconv.Atoi(re.FindString(version1[i]))
		num2, _ := strconv.Atoi(re.FindString(version2[i]))
		if num1 > num2 {
			return 1
		}
		if num1 < num2 {
			return -1
		}
	}
	return 0
}
