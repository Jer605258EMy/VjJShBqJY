// 代码生成时间: 2025-09-22 10:59:19
package main

import (
    "context"
    "fmt"
    "log"
    "net"

    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "google.golang.org/protobuf/types/known/emptypb"
)

// CartService provides a service for managing a shopping cart
type CartService struct {
    // Carts stores the shopping carts, where the key is the user ID
    Carts map[string]map[string]int
}

// NewCartService creates a new instance of CartService
func NewCartService() *CartService {
    return &CartService{
        Carts: make(map[string]map[string]int),
    }
}

// AddItem adds an item to the cart
func (s *CartService) AddItem(ctx context.Context, in *AddItemRequest) (*emptypb.Empty, error) {
    if _, exists := s.Carts[in.UserId]; !exists {
        s.Carts[in.UserId] = make(map[string]int)
    }
    s.Carts[in.UserId][in.ProductId] += in.Quantity
    return &emptypb.Empty{}, nil
}

// RemoveItem removes an item from the cart
func (s *CartService) RemoveItem(ctx context.Context, in *RemoveItemRequest) (*emptypb.Empty, error) {
    if cart, exists := s.Carts[in.UserId]; exists {
        if quantity, exists := cart[in.ProductId]; exists {
            if quantity > 0 {
                cart[in.ProductId] -= 1
            } else {
                return nil, status.Errorf(codes.OutOfRange, "Item quantity is zero")
            }
        } else {
            return nil, status.Errorf(codes.NotFound, "Item not found in cart")
        }
    } else {
        return nil, status.Errorf(codes.NotFound, "Cart not found for user")
    }
    return &emptypb.Empty{}, nil
}

// EmptyCart empties the cart
func (s *CartService) EmptyCart(ctx context.Context, in *EmptyCartRequest) (*emptypb.Empty, error) {
    if _, exists := s.Carts[in.UserId]; exists {
        delete(s.Carts, in.UserId)
    } else {
        return nil, status.Errorf(codes.NotFound, "Cart not found for user\)
    }
    return &emptypb.Empty{}, nil
}

// GetCart returns the cart items for a user
func (s *CartService) GetCart(ctx context.Context, in *GetCartRequest) (*GetCartResponse, error) {
    if cart, exists := s.Carts[in.UserId]; exists {
        items := make([]*CartItem, 0, len(cart))
        for productID, quantity := range cart {
            items = append(items, &CartItem{ProductId: productID, Quantity: quantity})
        }
        return &GetCartResponse{Items: items}, nil
    } else {
        return nil, status.Errorf(codes.NotFound, "Cart not found for user")
    }
}

// The main function starts the gRPC server
func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Server is running on port 50051")
    s := grpc.NewServer()
    RegisterCartServiceServer(s, NewCartService())
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// Define the request and response messages for the service methods
type AddItemRequest struct {
    UserId    string
    ProductId string
    Quantity  int32
}

type RemoveItemRequest struct {
    UserId    string
    ProductId string
}

type EmptyCartRequest struct {
    UserId string
}

type GetCartRequest struct {
    UserId string
}

type GetCartResponse struct {
    Items []*CartItem
}

type CartItem struct {
    ProductId string
    Quantity  int32
}

// RegisterCartServiceServer registers the CartService with the gRPC server
func RegisterCartServiceServer(s *grpc.Server, srv *CartService) {
    // Register the CartService with the gRPC server
    RegisterCartServiceServer(s, srv)
}
