package hot_reloading

import (
	"amp-templates/server/services/configuration"
	"os"
	"time"
)

type callback func()

var (
 	lastModTime = time.Unix(0, 0)
 	isRunning = false
	callBacks []callback
)

func ForceReload() {
	runCallBacks()
}

func Configure(cb callback) {
	callBacks = append(callBacks, cb)
	if !configuration.Config.Production && !isRunning {
		go func() {
			isRunning = true
			for range time.Tick(1 * time.Second) {
				checkFiles(configuration.Config.Static)
			}
		}()
	}
}

func runCallBacks() {
	for _, c := range callBacks { c() }
	callBacks = nil
}

func checkFiles(folder string) {
	needUpdate := false
	f, _ := os.Open(folder)
	fileInfos, _ := f.Readdir(-1)
	for _, fi := range fileInfos {
		if fi.ModTime().After(lastModTime) {
			lastModTime = fi.ModTime()
			needUpdate = true
		}
	}

	if needUpdate {
		runCallBacks()
	}
}