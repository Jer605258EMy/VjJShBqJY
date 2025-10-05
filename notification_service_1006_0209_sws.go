// 代码生成时间: 2025-10-06 02:09:22
package main

import (
    "context"
    "fmt"
    "log"
    "net"

    "google.golang.org/grpc"
    p "path/to/your/protobuf/definitions"  // 假设你的proto文件定义在这个路径下
)

// NotificationService 服务结构
type NotificationService struct {
    // 可以添加更多成员变量，例如数据库连接等
}

// 实现 proto 文件中定义的接口
func (s *NotificationService) SendNotification(ctx context.Context, req *p.SendNotificationRequest) (*p.SendNotificationResponse, error) {
    if req == nil {
        return nil, status.Errorf(codes.InvalidArgument, "request cannot be nil")
    }

    // 这里可以添加更多的验证逻辑，例如检查通知内容等
    // ...

    // 发送通知的逻辑，例如存储到数据库，发送到消息队列等
    // 这里只是一个简单的示例，实际应用中需要根据具体需求实现
    fmt.Printf("Sending notification: %v
", req.GetMessage())

    // 返回成功响应
    return &p.SendNotificationResponse{Success: true}, nil
}

// 服务器端主函数
func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    fmt.Println("Server is running on port 50051")

    // 创建gRPC服务器
    s := grpc.NewServer()

    // 注册服务
    p.RegisterNotificationServiceServer(s, &NotificationService{})

    // 启动服务器
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
