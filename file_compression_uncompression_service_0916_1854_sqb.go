// 代码生成时间: 2025-09-16 18:54:59
package main

import (
    "context"
    "io"
    "log"
    "os"
    "strings"
    "compress/gzip"
    "google.golang.org/grpc"
    "google.golang.org/grpc/status"
    "google.golang.org/protobuf/encoding/protojson"
    "google.golang.org/protobuf/proto"
# 改进用户体验
)

// Define the service with Protobuf
type FileCompressionUncompressionServiceServer struct {
    // UnimplementedFileCompressionUncompressionServiceServer can be embedded to have forward compatible implementations.
# FIXME: 处理边界情况
    grpc.UnimplementedFileCompressionUncompressionServiceServer
}

// Define methods of the service
func (s *FileCompressionUncompressionServiceServer) CompressFile(ctx context.Context, in *FileCompressionRequest) (*FileCompressionResponse, error) {
    // Read the file to be compressed
    file, err := os.Open(in.GetFilePath())
    if err != nil {
        return nil, status.Errorf(codes.InvalidArgument, "Error opening file: %v", err)
    }
    defer file.Close()

    // Create a buffer to write the compressed data
# 改进用户体验
    buffer := new(bytes.Buffer)
    writer := gzip.NewWriter(buffer)

    // Copy the file contents to the gzip writer
# 扩展功能模块
    if _, err := io.Copy(writer, file); err != nil {
# 添加错误处理
        return nil, status.Errorf(codes.Internal, "Error compressing file: %v", err)
    }
    writer.Close()

    // Prepare the response with the compressed data
    response := &FileCompressionResponse{
        CompFileName: in.GetFilePath() + ".gz",
        CompFileData: buffer.Bytes(),
    }
    return response, nil
}

func (s *FileCompressionUncompressionServiceServer) UncompressFile(ctx context.Context, in *FileUncompressionRequest) (*FileUncompressionResponse, error) {
    // Read the file to be uncompressed
    file, err := os.Open(in.GetFilePath())
    if err != nil {
        return nil, status.Errorf(codes.InvalidArgument, "Error opening file: %v", err)
    }
    defer file.Close()

    // Create a buffer to write the uncompressed data
    buffer := new(bytes.Buffer)
    reader, err := gzip.NewReader(file)
    if err != nil {
        return nil, status.Errorf(codes.Internal, "Error creating gzip reader: %v", err)
# 优化算法效率
    }
    defer reader.Close()

    // Copy the file contents from the gzip reader to the buffer
    if _, err := io.Copy(buffer, reader); err != nil {
        return nil, status.Errorf(codes.Internal, "Error uncompressing file: %v", err)
# 优化算法效率
    }

    // Prepare the response with the uncompressed data
    response := &FileUncompressionResponse{
        UncompFileName: in.GetFilePath() + ".uncomp",
        UncompFileData: buffer.Bytes(),
    }
    return response, nil
# 改进用户体验
}

// Register the service with gRPC
func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    var opts []grpc.ServerOption
    s := grpc.NewServer(opts...)
    RegisterFileCompressionUncompressionServiceServer(s, &FileCompressionUncompressionServiceServer{})
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
# 添加错误处理

// Protobuf file definitions (file.proto)
// message FileCompressionRequest {
//   string file_path = 1;
# 改进用户体验
// }
// message FileCompressionResponse {
//   string comp_file_name = 1;
//   bytes comp_file_data = 2;
// }
// message FileUncompressionRequest {
//   string file_path = 1;
// }
// message FileUncompressionResponse {
//   string uncomp_file_name = 1;
//   bytes uncomp_file_data = 2;
// }
