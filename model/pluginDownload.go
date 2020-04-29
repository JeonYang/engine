package model

import (
	"engine/log"
	"engine/util"
	"os"
	"path/filepath"
)

const MaxAttempts = 2

type PluginDownloadTask struct {
	pluginName string
	path       string
	url        string

	success  bool
	countTry int
}

func newPluginDownloadTask(path, name, urlPrefix string) *PluginDownloadTask {
	return &PluginDownloadTask{pluginName: name, path: filepath.Join(path, name), url: urlPrefix + "/" + name}
}

func (task *PluginDownloadTask) download() {
	task.countTry++
	err := util.WGet(task.url, task.path)
	if err != nil {
		log.Error("download plugin: %s, plugin url: %s, fail: %s", task.pluginName, task.url, err.Error())
	} else {
		log.Info("download plugin(%s) success", task.pluginName)
		task.success = true
	}
}

type PluginDownloader struct {
	tasks map[string]*PluginDownloadTask
	taskC chan *PluginDownloadTask
}

type DownloadDoneCallback func(string)

func (d *PluginDownloader) Start(doneCallback DownloadDoneCallback) {
	d.tasks = map[string]*PluginDownloadTask{}
	d.taskC = make(chan *PluginDownloadTask)
	taskC := make(chan *PluginDownloadTask)
	doneC := make(chan *PluginDownloadTask)
	go func() {
		log.Info("start plugin download worker")
		for {
			t := <-taskC
			t.download()
			doneC <- t
		}
	}()
	scheduleTask := func() {
		for k, v := range d.tasks {
			select {
			case taskC <- v:
				log.Info("schedule plugin(%s) download task success", k)
			default:
				log.Info("the previous plugin download task is still running, skip")
			}
			break
		}
	}
	go func() {
		for {
			select {
			case done := <-doneC:
				pluginTask, exit := d.tasks[done.pluginName]
				if done == nil || !exit {
					log.Error("done: %+v, tasks: %+v, pluginTask: %+v", done, d.tasks, pluginTask)
					continue
				}
				if (done.success && pluginTask == done) || done.countTry >= MaxAttempts {
					delete(d.tasks, done.pluginName)
					// 调用通知回调.
					doneCallback(done.pluginName)
				}
				scheduleTask()
			case task := <-d.taskC:
				d.tasks[task.pluginName] = task
				scheduleTask()
			}
		}
	}()
}

func (d *PluginDownloader) NewTask(path, name, urlPrefix string) {
	task := newPluginDownloadTask(path, name, urlPrefix)
	if _, err := os.Stat(task.path); os.IsNotExist(err) {
		d.taskC <- task
	}
}
