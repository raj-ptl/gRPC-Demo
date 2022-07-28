package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/raj-ptl/gRPC-Demo/helper"
	"github.com/raj-ptl/gRPC-Demo/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.SumServiceServer
	pb.GetPrimesServiceServer
	pb.GetAverageServiceServer
}

func (s *server) GetSum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	res := &pb.SumResponse{
		SumResult: req.Num_1 + req.Num_2,
	}

	return res, nil
}

func (s *server) GetPrimes(req *pb.GetPrimesRequest, stream pb.GetPrimesService_GetPrimesServer) error {
	fmt.Println("Got request for primes till : ", req)
	primes := helper.Sieve(int(req.Num))
	for _, p := range primes {
		res := &pb.GetPrimesResponse{
			Num: int32(p),
		}
		stream.Send(res)
		time.Sleep(500 * time.Millisecond)
	}
	return nil
}

func (s *server) GetAverage(stream pb.GetAverageService_GetAverageServer) error {
	fmt.Println("Got request for calculating avg")

	var cumulativeSum int32
	var cumulativeCount int32

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GetAverageResponse{
				Avg: cumulativeSum / cumulativeCount,
			})
		}

		if err != nil {
			fmt.Println("Error while stream closure", err)
		}
		cumulativeSum += req.Num
		cumulativeCount++
	}

}

func main() {
	fmt.Println("Hello World from Server!")
	lis, err := net.Listen("tcp", "127.0.0.1:9091")
	if err != nil {
		fmt.Println("Failed to listen : ", err)
	}

	sv := grpc.NewServer()
	pb.RegisterSumServiceServer(sv, &server{})
	pb.RegisterGetPrimesServiceServer(sv, &server{})
	pb.RegisterGetAverageServiceServer(sv, &server{})
	reflection.Register(sv)
	if err := sv.Serve(lis); err != nil {
		fmt.Println("Failed to serve : ", err)
	}

}
