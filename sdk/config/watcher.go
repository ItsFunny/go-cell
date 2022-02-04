package config

import (
	"errors"
	logrusplugin "github.com/itsfunny/go-cell/sdk/log/logrus"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
)

type RecursiveWatcher struct {
	*fsnotify.Watcher
	//Files   chan string
	//Folders chan string

	OnFileCreateOrModified func(str string)
	OnDirCreate            func(dir string)
}

func NewRecursiveWatcher(path string, onF func(str string), onD func(str string)) (*RecursiveWatcher, error) {
	folders := Subfolders(path)
	if len(folders) == 0 {
		return nil, errors.New("No folders to watch.")
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	rw := &RecursiveWatcher{Watcher: watcher}
	rw.OnFileCreateOrModified = onF
	rw.OnDirCreate = onD
	for _, folder := range folders {
		rw.AddFolder(folder, false)
	}

	return rw, nil
}

func (watcher *RecursiveWatcher) AddFolder(folder string, report bool) {
	err := watcher.Add(folder)
	if err != nil {
		log.Println("Error watching: ", folder, err)
	}
	if report {
		watcher.OnDirCreate(folder)
	}
}

func (watcher *RecursiveWatcher) run(debug bool) {
	for {
		select {
		case event := <-watcher.Events:
			// create a file/directory
			if event.Op&fsnotify.Create == fsnotify.Create {
				fi, err := os.Stat(event.Name)
				if err != nil {
					// eg. stat .subl513.tmp : no such file or directory
					logrusplugin.Error("failed", "err", err)
				} else if fi.IsDir() {
					if debug{
						logrusplugin.Info("detected new dir ", "name", fi.Name())
					}
					if !shouldIgnoreFile(filepath.Base(event.Name)) {
						watcher.AddFolder(event.Name, true)
					}
				} else {
					if debug{
						logrusplugin.Info("detected new file ", "name", event.Name)
					}
					watcher.OnFileCreateOrModified(event.Name)
				}
			}

			if event.Op&fsnotify.Write == fsnotify.Write {
				// modified a file, assuming that you don't modify folders
				if debug{
					logrusplugin.InfoF("Detected file modification %s", event.Name)
				}
				watcher.OnFileCreateOrModified(event.Name)
			}

		case err := <-watcher.Errors:
			logrusplugin.Error("err", err)
		}
	}
}

// Subfolders returns a slice of subfolders (recursive), including the folder provided.
func Subfolders(path string) (paths []string) {
	filepath.Walk(path, func(newPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			name := info.Name()
			// skip folders that begin with a dot
			if shouldIgnoreFile(name) && name != "." && name != ".." {
				return filepath.SkipDir
			}
			paths = append(paths, newPath)
		}
		return nil
	})
	return paths
}

// shouldIgnoreFile determines if a file should be ignored.
// File names that begin with "." or "_" are ignored by the go tool.
func shouldIgnoreFile(name string) bool {
	return strings.HasPrefix(name, ".") || strings.HasPrefix(name, "_")
}
