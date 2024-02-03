package main

import (
	"fmt"
	"net"

	pb "github.com/hayk2377/Grpc-go/protobuf/phew"
	"google.golang.org/grpc"
)

type LoadBalancer struct{}

func (l *LoadBalancer) NewServer(req *pb.ServerRequest, res *pb.ServerResponse) error {
    fmt.Println("Received newServer request from IP:", req.Ip)
    res.ServerId = "1234"
    return nil
}

func (l *LoadBalancer) HeartBeat(req *pb.ServerRequest, res *pb.ServerResponse) error {
    fmt.Println("Received heartBeat request from IP:", req.Ip)
    res.ServerId = "ok"
    return nil
}

func (l *LoadBalancer) Notify(req *pb.NotifyRequest, res *pb.ServerResponse) error {
    fmt.Printf("Received notify request for GameID: %s, ServerIP: %s\n", req.GameId, req.ServerIp)
    res.ServerId = "ok"
    return nil
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
