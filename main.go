package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/hayk2377/Grpc-go/phew"
	"google.golang.org/grpc"
)


type LoadBalancer struct {
    pb.UnimplementedLoadBalancerServer
}

func (l *LoadBalancer) NewServer(ctx context.Context, req *pb.ServerRequest) (*pb.ServerResponse, error) {
    fmt.Println("Received newServer request from IP:", req.Ip)
    return &pb.ServerResponse{ServerId: "1234"}, nil
}


func (l *LoadBalancer) HeartBeat(ctx context.Context, req *pb.Heartreq) (*pb.ServerResponse, error) {
    fmt.Println("Received heartBeat request from IP:", req.Status)
    return &pb.ServerResponse{ServerId: "ok"}, nil
}


func (l *LoadBalancer) Notify(ctx context.Context, req *pb.NotifyRequest) (*pb.ServerResponse, error) {
    fmt.Printf("Received notify request for GameID: %s, ServerIP: %s\n", req.GameId, req.ServerIp)
    return &pb.ServerResponse{ServerId: "ok"}, nil
}


func startRPCServer() {
    listener, err := net.Listen("tcp", ":1234")
    if err != nil {
        fmt.Println("Error starting RPC server:", err)
        return
    }

    grpcServer := grpc.NewServer()
    pb.RegisterLoadBalancerServer(grpcServer, &LoadBalancer{})
    fmt.Println("gRPC server started on :1234")
    grpcServer.Serve(listener)
}

func main() {
    go startRPCServer()
    select {}
}
