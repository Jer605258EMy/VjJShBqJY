// 代码生成时间: 2025-09-24 07:22:32
package main

import (
    "context"
    "log"
    "net"
    "time"

    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"

    "golang.org/x/time/rate"
)

// CacheItem defines the structure for cache items
type CacheItem struct {
    Value    string    `json:"value"`
    Expiry   time.Time `json:"expiry"`
}

// CacheService is the service that implements the caching logic
type CacheService struct {
    // Stores the cached items with their expiry times
    cache map[string]CacheItem
    // Mutex to protect concurrent access to the cache
    mu    sync.Mutex
    // Limiter to control the rate of cache writes
    limiter *rate.Limiter
}

// NewCacheService creates a new instance of the CacheService
func NewCacheService() *CacheService {
    // Initialize the cache and limiter
    return &CacheService{
        cache:  make(map[string]CacheItem),
        limiter: rate.NewLimiter(1, 1),
    }
}

// Set caches a new item with an expiry time
func (c *CacheService) Set(ctx context.Context, req *SetRequest) (*SetResponse, error) {
    if !c.limiter.Allow() {
        return nil, status.Errorf(codes.ResourceExhausted, "cache set rate limit exceeded")
    }

    c.mu.Lock()
    defer c.mu.Unlock()

    item := CacheItem{
        Value:    req.Value,
        Expiry:   time.Now().Add(req.ExpiryDuration),
    }
    c.cache[req.Key] = item
    return &SetResponse{Success: true}, nil
}

// Get retrieves an item from the cache if it has not expired
func (c *CacheService) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
    c.mu.Lock()
    defer c.mu.Unlock()

    item, exists := c.cache[req.Key]
    if !exists {
        return nil, status.Errorf(codes.NotFound, "cache item not found")
    }

    if time.Now().After(item.Expiry) {
        delete(c.cache, req.Key) // Remove expired item from cache
        return nil, status.Errorf(codes.Unavailable, "cache item has expired")
    }

    return &GetResponse{Value: item.Value}, nil
}

// StartGRPCServer starts the GRPC server with the cache service
func StartGRPCServer(port string, service *CacheService) error {
    lis, err := net.Listen("tcp", port)
    if err != nil {
        return err
    }

    grpcServer := grpc.NewServer()
    // Register the cache service with the GRPC server
    RegisterCacheServiceServer(grpcServer, service)
    return grpcServer.Serve(lis)
}

func main() {
    // Create a new cache service instance
    cacheService := NewCacheService()
    // Start the GRPC server on port 50051
    if err := StartGRPCServer(":50051", cacheService); err != nil {
        log.Fatalf("failed to start GRPC server: %v", err)
    }
}

// SetRequest defines the request for setting a cache item
type SetRequest struct {
    Key          string    `json:"key"`
    Value        string    `json:"value"`
    ExpiryDuration time.Duration `json:"expiryDuration"`
}

// SetResponse defines the response for setting a cache item
type SetResponse struct {
    Success bool `json:"success"`
}

// GetRequest defines the request for getting a cache item
type GetRequest struct {
    Key string `json:"key"`
}

// GetResponse defines the response for getting a cache item
type GetResponse struct {
    Value string `json:"value"`
}

// CacheServiceServer is the server API for CacheService service
type CacheServiceServer struct{
    // UnimplementedCacheServiceServer can be embedded to have forward compatible implementations.
    UnimplementedCacheServiceServer 
}

// RegisterCacheServiceServer registers the server with the GRPC framework
func RegisterCacheServiceServer(s *grpc.Server, srv *CacheService) {
    s.RegisterService(&_CacheService_serviceDesc, &CacheServiceServer{srv})
}

// _CacheService_serviceDesc is the descriptor for the service
var _CacheService_serviceDesc = grpc.ServiceDesc{
    ServiceName: "CacheService",
    HandlerType: (*CacheServiceServer)(nil),
    Methods: []grpc.MethodDesc{
        {
            MethodName: "Set",
            Handler: _CacheService_Set_Handler,
        },
        {
            MethodName: "Get",
            Handler: _CacheService_Get_Handler,
        },
    },
    Streams:  []grpc.StreamDesc{},
    Metadata: "cache_service.proto",
}

// _CacheService_Set_Handler is the handler for the Set method
func _CacheService_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
    in := new(SetRequest)
    if err := dec(in); err != nil {
        return nil, err
    }
    if interceptor == nil { return srv.(CacheServiceServer).Set(ctx, in) }
    info := &grpc.UnaryServerInfo{
        Server:     srv,
         FullMethod: "CacheService/Set",
    }
    handler := func(ctx context.Context, req interface{}) (interface{}, error) {
        return srv.(CacheServiceServer).Set(ctx, req.(*SetRequest))
    }
    return interceptor(ctx, in, info, handler)
}

// _CacheService_Get_Handler is the handler for the Get method
func _CacheService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
    in := new(GetRequest)
    if err := dec(in); err != nil {
        return nil, err
    }
    if interceptor == nil { return srv.(CacheServiceServer).Get(ctx, in) }
    info := &grpc.UnaryServerInfo{
        Server:     srv,
         FullMethod: "CacheService/Get",
    }
    handler := func(ctx context.Context, req interface{}) (interface{}, error) {
        return srv.(CacheServiceServer).Get(ctx, req.(*GetRequest))
    }
    return interceptor(ctx, in, info, handler)
}