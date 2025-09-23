// 代码生成时间: 2025-09-23 09:24:24
package main

import (
    "context"
    "fmt"
    "log"
    "net"
    "os"
    "runtime"
    "time"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "google.golang.org/grpc/status"
    "google.golang.org/grpc/health"
    "google.golang.org/grpc/health/grpc_health_v1"
)

// MemoryUsage defines the memory usage details
type MemoryUsage struct {
    // Alloc is bytes of allocated heap objects
    Alloc uint64 `protobuf:"varint,1,opt,name=alloc,proto3" json:"alloc,omitempty"`
    // TotalAlloc is bytes of total allocated space
    TotalAlloc uint64 `protobuf:"varint,2,opt,name=total_alloc,json=totalAlloc,proto3" json:"total_alloc,omitempty"`
    // Sys is the total bytes of memory obtained from the OS
    Sys uint64 `protobuf:"varint,3,opt,name=sys,proto3" json:"sys,omitempty"`
    // Mallocs is the count of heap objects
    Mallocs uint64 `protobuf:"varint,4,opt,name=mallocs,proto3" json:"mallocs,omitempty"`
    // Frees is the count of freed heap objects
    Frees uint64 `protobuf:"varint,5,opt,name=frees,proto3" json:"frees,omitempty"`
    // LiveObjects is the number of live heap objects
    LiveObjects uint64 `protobuf:"varint,6,opt,name=live_objects,json=liveObjects,proto3" json:"live_objects,omitempty"`
    // HeapAlloc is the bytes of allocated heap
    HeapAlloc uint64 `protobuf:"varint,7,opt,name=heap_alloc,json=heapAlloc,proto3" json:"heap_alloc,omitempty"`
    // HeapInuse is the bytes of in-use heap
    HeapInuse uint64 `protobuf:"varint,8,opt,name=heap_inuse,json=heapInuse,proto3" json:"heap_inuse,omitempty"`
    // HeapReleased is the bytes of released heap
    HeapReleased uint64 `protobuf:"varint,9,opt,name=heap_released,json=heapReleased,proto3" json:"heap_released,omitempty"`
}

// MemoryUsageServiceServer is the server API for MemoryUsageService service
type MemoryUsageServiceServer struct {
    // UnimplementedMemoryUsageServiceServer must be embedded to have forward compatible methods
    grpc_health_v1.UnimplementedHealthServer
}

// GetMemoryUsage returns the current memory usage details
func (s *MemoryUsageServiceServer) GetMemoryUsage(ctx context.Context, in *grpc_health_v1.HealthCheckRequest) (*MemoryUsage, error) {
    m := new(runtime.MemStats)
    runtime.ReadMemStats(m)

    mu := &MemoryUsage{
        Alloc:        m.Alloc,
        TotalAlloc:   m.TotalAlloc,
        Sys:          m.Sys,
        Mallocs:      m.Mallocs,
        Frees:        m.Frees,
        LiveObjects:  m.LiveObjects,
        HeapAlloc:    m.HeapAlloc,
        HeapInuse:    m.HeapInuse,
        HeapReleased: m.HeapReleased,
    }

    return mu, nil
}

// Check checks the service's health
func (s *MemoryUsageServiceServer) Check(ctx context.Context, in *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
    return &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}, nil
}

// Watch is not implemented
func (s *MemoryUsageServiceServer) Watch(in *grpc_health_v1.HealthCheckRequest, srv grpc_health_v1.Health_WatchServer) error {
    return status.Errorf(codes.Unimplemented, "Watch not implemented")
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Server listening on port 50051")

    grpcServer := grpc.NewServer()
    memoryUsageServiceServer := &MemoryUsageServiceServer{}

    grpc_health_v1.RegisterHealthServer(grpcServer, memoryUsageServiceServer)
    grpcServer.RegisterService(&_MemUsageService_serviceDesc, memoryUsageServiceServer)
    reflection.Register(grpcServer)

    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
