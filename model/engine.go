package model

import (
	"engine/log"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"syscall"
)

var Engine = &engine{}

type engine struct {
}

func (engine *engine) Update(appPath, backupPath, downLoadPath string) error {
	defer os.Remove(downLoadPath)
	return engine.update(appPath, backupPath, downLoadPath)
}

func (engine *engine) update(appPath, backupPath, downLoadPath string) (err error) {
	// 拷贝engine到backup
	log.Info("backup engine..... ")
	err = engine.copyFile(appPath, backupPath)
	if err != nil {
		log.Errorf("backup engine fail.[%v]", err)
		return
	}

	defer os.Remove(backupPath)

	// 更换engine执行文件
	log.Info("rename engine ...")
	err = os.Rename(downLoadPath, appPath)
	if err != nil {
		log.Errorf("cover old files failed:%v", err)
		// restore
		os.Rename(backupPath, appPath)
		return
	}

	log.Info("new process ...")
	err = engine.newProcess()
	if err != nil {
		os.Rename(backupPath, appPath)
		return
	}
	return
}

func (engine *engine) newProcess() error {
	appPath, _ := filepath.Abs(os.Args[0])
	os.Chmod(appPath, 0755)
	env := os.Environ()
	log.Infof("starting new agent... ")
	err := syscall.Exec(appPath, os.Args, env)
	if err != nil {
		return fmt.Errorf("start new process :%v", err)
	}
	return nil
}

func (engine *engine) copyFile(src, dst string) error {
	sfi, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("src not exist , fail: %v", err)
	}
	sfile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("open %s fail:%v", src, err)
	}
	defer sfile.Close()

	dstDir := filepath.Dir(dst)
	_, err = os.Stat(dstDir)
	if err != nil {
		err = os.MkdirAll(dstDir, 0775)
		if err != nil {
			return fmt.Errorf("mkdir %s fail: %v", dstDir, err)
		}
	}
	dfile, err := os.OpenFile(dst, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, sfi.Mode())
	if err != nil {
		return fmt.Errorf("open %s fail:%v", dst, err)
	}
	defer dfile.Close()

	if _, err := io.Copy(dfile, sfile); err != nil {
		return fmt.Errorf("copy src: %s,dst: %v, fail: %v", src, dst, err)
	}

	return nil
}
