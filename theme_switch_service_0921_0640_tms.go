// 代码生成时间: 2025-09-21 06:40:39
 * This service allows clients to switch between different themes through a gRPC interface.
 */

package main

import (
    "context"
    "fmt"
    "log"
    "net"

    "google.golang.org/grpc"
# NOTE: 重要实现细节
    "google.golang.org/grpc/reflection"
    "google.golang.org/protobuf/types/known/emptypb"

    "path/to/your/protos" // Replace with the actual path to your generated protobuf package
# 扩展功能模块
)

// ThemeServiceServer is the server API for ThemeService
type ThemeServiceServer struct {
    Themes []string
    CurTheme string
}
# 扩展功能模块

// NewThemeServiceServer creates a new ThemeServiceServer with an initial theme
func NewThemeServiceServer(initTheme string) *ThemeServiceServer {
    return &ThemeServiceServer{
        Themes: []string{"light", "dark"}, // List of available themes
# 扩展功能模块
        CurTheme: initTheme,
    }
}
# 改进用户体验

// SwitchTheme changes the current theme
func (s *ThemeServiceServer) SwitchTheme(ctx context.Context, req *protos.SwitchThemeRequest) (*emptypb.Empty, error) {
    if req.Theme == nil || req.Theme.Value == s.CurTheme {
        return nil, fmt.Errorf("no new theme provided or already using the provided theme")
    }
# 添加错误处理

    // Check if the theme is available
    for _, theme := range s.Themes {
        if theme == req.Theme.Value {
            s.CurTheme = req.Theme.Value
# 增强安全性
            return &emptypb.Empty{}, nil
        }
    }
    return nil, fmt.Errorf("theme '%s' is not available", req.Theme.Value)
}
# 增强安全性

// GetCurTheme returns the current theme
func (s *ThemeServiceServer) GetCurTheme(ctx context.Context, req *emptypb.Empty) (*protos.ThemeResponse, error) {
    return &protos.ThemeResponse{Value: s.CurTheme}, nil
}

// StartServer starts the gRPC server
# 添加错误处理
func StartServer(port int, initTheme string) {
# FIXME: 处理边界情况
    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
# 添加错误处理
    }
    fmt.Printf("server listening on port %d
", port)

    s := grpc.NewServer()
    themeService := NewThemeServiceServer(initTheme)

    protos.RegisterThemeServiceServer(s, themeService)
# FIXME: 处理边界情况
    reflection.Register(s) // Enables gRPC reflection for debugging
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
# FIXME: 处理边界情况
}

func main() {
    // Default theme is 'light'
    StartServer(50051, "light")
}
