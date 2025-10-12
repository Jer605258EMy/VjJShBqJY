// 代码生成时间: 2025-10-12 21:11:02
package main

import (
    "fmt"
    "log"
    "net"
    "os"
    "path/filepath"
    "strings"
    "time"

    "google.golang.org/grpc"
    "github.com/fsnotify/fsnotify"
    "google.golang.org/protobuf/types/known/emptypb"
)

const (
    defaultPort = ":50051"
)

// 文件监控服务定义
type FileWatcherService struct {
    watcher *fsnotify.Watcher
    events chan struct{} // 用于通知gRPC服务器发送文件变更事件
}

// 实现gRPC服务接口
type fileWatcherServiceServer struct {
    FileWatcherService
}

// 定义文件变更通知gRPC服务
type FileWatcherServiceServer interface {
    WatchFile(stream FileWatcherService_WatchFileServer) error
    GetFileEvent(event *FileEventRequest, stream FileWatcherService_GetFileEventServer) error
}

// 文件事件请求
type FileEventRequest struct {
    FileName string
}

// 文件事件响应
type FileEventResponse struct {
    FileName string
    Event    string // 文件事件类型
}

// 初始化文件监控服务
func NewFileWatcherService() (*FileWatcherService, error) {
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        return nil, err
    }
    return &FileWatcherService{
        watcher: watcher,
        events: make(chan struct{}, 1),
    }, nil
}

// 启动文件监控服务
func (s *FileWatcherService) Start() error {
    go func() {
        for {
            select {
            case event, ok := <-s.watcher.Events:
                if !ok {
                    return
                }
                if event.Op&fsnotify.Write == fsnotify.Write {
                    // 文件写入事件
                    s.events <- struct{}{}
                }
            case err, ok := <-s.watcher.Errors:
                if !ok {
                    return
                }
                log.Printf("file watcher error: %s", err)
            }
        }
    }()
    return nil
}

// 停止文件监控服务
func (s *FileWatcherService) Stop() error {
    return s.watcher.Close()
}

// WatchFile方法用于启动文件监控
func (s *fileWatcherServiceServer) WatchFile(stream FileWatcherService_WatchFileServer) error {
    fileName := ""
    if err := stream.RecvMsg(&fileName); err != nil {
        return err
    }
    if err := s.watcher.Add(fileName); err != nil {
        return err
    }
    for range s.events {
        if err := stream.Send(&FileEventResponse{FileName: fileName, Event: "modified"}); err != nil {
            return err
        }
    }
    return nil
}

// GetFileEvent方法用于获取文件事件
func (s *fileWatcherServiceServer) GetFileEvent(event *FileEventRequest, stream FileWatcherService_GetFileEventServer) error {
    fileName := event.FileName
    if err := s.watcher.Add(fileName); err != nil {
        return err
    }
    for range s.events {
        if err := stream.Send(&FileEventResponse{FileName: fileName, Event: "modified"}); err != nil {
            return err
        }
    }
    return nil
}

// main函数启动gRPC服务器
func main() {
    service, err := NewFileWatcherService()
    if err != nil {
        log.Fatalf("failed to create file watcher service: %v", err)
    }
    if err := service.Start(); err != nil {
        log.Fatalf("failed to start file watcher service: %v", err)
    }
    defer service.Stop()

    listener, err := net.Listen("tcp", defaultPort)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    defer listener.Close()

    srv := grpc.NewServer()
    FileWatcherServiceServer(&fileWatcherServiceServer{*service}).Register(srv)
    if err := srv.Serve(listener); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
