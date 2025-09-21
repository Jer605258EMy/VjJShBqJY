// 代码生成时间: 2025-09-22 07:22:47
package main

import (
    "fmt"
    "io"
    "math/rand"
    "time"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
)

// RandomNumberService 定义了随机数生成服务
type RandomNumberService struct{}

// GenerateRandomNumber 响应请求并生成一个随机数
func (s *RandomNumberService) GenerateRandomNumber(ctx context.Context, in *GenerateRandomNumberRequest) (*GenerateRandomNumberResponse, error) {
    if err := validateRequest(in); err != nil {
        return nil, err
    }

    // 生成随机数
    randomNumber := rand.Intn(in.GetMaximum())

    // 返回随机数
    return &GenerateRandomNumberResponse{
        Value: int32(randomNumber),
    }, nil
}

// validateRequest 验证请求参数
func validateRequest(in *GenerateRandomNumberRequest) error {
    if in.GetMaximum() <= 0 {
        return fmt.Errorf("maximum must be greater than 0")
    }
    return nil
}

// Define the gRPC server we'll use to handle requests
type server struct{
    RandomNumberServiceServer
}

// Implement the methods of the RandomNumberServiceServer interface
func (s *server) GenerateRandomNumber(ctx context.Context, in *GenerateRandomNumberRequest) (*GenerateRandomNumberResponse, error) {
    return (&RandomNumberService{}).GenerateRandomNumber(ctx, in)
}

// main 函数设置并启动 gRPC 服务
func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("listening on port 50051")
    s := grpc.NewServer()
    RegisterRandomNumberServiceServer(s, &server{})
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// GenerateRandomNumberRequest is the request for the GenerateRandomNumber method
type GenerateRandomNumberRequest struct{
    Maximum int32 `protobuf:"varint,1,opt,name=maximum" json:"maximum,omitempty"`
}

// GenerateRandomNumberResponse is the response for the GenerateRandomNumber method
type GenerateRandomNumberResponse struct{
    Value int32 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

// RandomNumberServiceServer is the server API for RandomNumberService service
type RandomNumberServiceServer interface {
    // GenerateRandomNumber generates a random number
    GenerateRandomNumber(context.Context, *GenerateRandomNumberRequest) (*GenerateRandomNumberResponse, error)
}

// RegisterRandomNumberServiceServer registers the server with the gRPC server
func RegisterRandomNumberServiceServer(s *grpc.Server, srv *RandomNumberService) {
    RegisterRandomNumberServiceHandler(s, srv)
}

// RandomNumberServiceHandler is the implementation of the RandomNumberService service
func (m *RandomNumberService) ServeRandomNumberService(server *grpc.Server, srv RandomNumberServiceServer) {
    p := grpc.NewServer(server, grpc.UnknownStreamHandler(srv), grpc.UnknownUnaryHandler(srv))
    p.RegisterService(&RandomNumberService_ServiceDesc, srv)
}

// RandomNumberService_ServiceDesc is the service descriptor for RandomNumberService
var RandomNumberService_ServiceDesc = grpc.ServiceDesc{
    ServiceName: "RandomNumberService",
    HandlerType: (*RandomNumberServiceServer)(nil),
    Methods: []grpc.MethodDesc{
        {
            MethodName: "GenerateRandomNumber",
            Handler: grpc.HandlerFunc(func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
                in := new(GenerateRandomNumberRequest)
                if err := dec(in); err != nil {
                    return nil, err
                }
                return srv.(RandomNumberServiceServer).GenerateRandomNumber(ctx, in)
            })},
    },
    Streams: []grpc.StreamDesc{},
   Metadata: "random_number_generator.proto",
}

// Initialize the random seed
func init() {
    rand.Seed(time.Now().UnixNano())
}