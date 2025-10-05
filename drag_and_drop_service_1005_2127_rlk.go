// 代码生成时间: 2025-10-05 21:27:50
package main

import (
    "fmt"
    "context"
    "google.golang.org/grpc"
    "net"
    "log"
)

// Define the structure for our drag-and-drop items
type DragDropItem struct {
    Id   int    `json:"id"`
    Value string `json:"value"`
}

// Define the service
type DragAndDropService struct {}

// Define the gRPC server
type DragAndDropServer struct {
    DragAndDropService
}

// Define methods that our service will implement
func (s *DragAndDropServer) SortItems(ctx context.Context, req *SortItemsRequest) (*SortItemsResponse, error) {
    // Implement sorting logic here
    // For simplicity, we'll just return the items in the order they were received
    return &SortItemsResponse{
        Items: req.Items,
    }, nil
}

// Define request and response types for SortItems method
type SortItemsRequest struct {
    Items []*DragDropItem `json:"items"`
}

type SortItemsResponse struct {
    Items []*DragDropItem `json:"items"`
}

// Main function to start the gRPC server
func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Server is running on port 50051")

    // Create a new gRPC server
    s := grpc.NewServer()

    // Register our service with the gRPC server
    RegisterDragAndDropServiceServer(s, &DragAndDropServer{})

    // Start the server
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// RegisterDragAndDropServiceServer registers the service with the gRPC server
func RegisterDragAndDropServiceServer(s *grpc.Server, srv *DragAndDropServer) {
    pb.RegisterDragAndDropServiceServer(s, srv)
}
