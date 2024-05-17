package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
)

// Function to calculate the size of a directory
func getFolderSize(path string) (int64, error) {
    var size int64
    err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() {
            size += info.Size()
        }
        return nil
    })
    return size, err
}

// Struct to hold the state of the folder size monitoring
type FolderMonitor struct {
    path       string
    size       int64
    mu         sync.Mutex
    watcher    *fsnotify.Watcher
    quit       chan struct{}
}

// Function to start monitoring folder size changes
func (fm *FolderMonitor) Start() error {
    var err error
    fm.watcher, err = fsnotify.NewWatcher()
    if err != nil {
        return err
    }
    defer fm.watcher.Close()

    fm.quit = make(chan struct{})
    go fm.watchEvents()

    err = fm.watcher.Add(fm.path)
    if err != nil {
        return err
    }

    // Initial size calculation
    size, err := getFolderSize(fm.path)
    if err != nil {
        return err
    }
    fm.mu.Lock()
    fm.size = size
    fm.mu.Unlock()

    // Periodic size checking every 10 seconds
    go fm.periodicSizeCheck()

    <-fm.quit
    return nil
}

// Function to watch for file system events
func (fm *FolderMonitor) watchEvents() {
    for {
        select {
        case event, ok := <-fm.watcher.Events:
            if !ok {
                return
            }
            if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create || event.Op&fsnotify.Remove == fsnotify.Remove {
                fm.updateSize()
            }
        case err, ok := <-fm.watcher.Errors:
            if !ok {
                return
            }
            log.Println("error:", err)
        }
    }
}

// Function to periodically check the folder size
func (fm *FolderMonitor) periodicSizeCheck() {
    ticker := time.NewTicker(1 * time.Second)
    for {
        select {
        case <-ticker.C:
            fm.updateSize()
        case <-fm.quit:
            ticker.Stop()
            return
        }
    }
}

// Function to update the folder size
func (fm *FolderMonitor) updateSize() {
    size, err := getFolderSize(fm.path)
    if err != nil {
        log.Println("error getting folder size:", err)
        return
    }
    fm.mu.Lock()
    if fm.size != size {
        fm.size = size
        fmt.Printf("Folder size changed: %d bytes\n", size)
    }
    fm.mu.Unlock()
}

// Function to stop monitoring
func (fm *FolderMonitor) Stop() {
    close(fm.quit)
}

func main() {
    if len(os.Args) < 2 {
        log.Fatal("Please specify a folder to monitor")
    }
    folder := os.Args[1]

    fm := &FolderMonitor{path: folder}
    if err := fm.Start(); err != nil {
        log.Fatal(err)
    }
}
