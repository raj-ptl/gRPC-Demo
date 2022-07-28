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
	pb.ReturnIfMaxServiceServer
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

func (s *server) ReturnIfMax(stream pb.ReturnIfMaxService_ReturnIfMaxServer) error {
	fmt.Println("Got request for returning if Max")

	var max int32
	var isMaxInitialized bool = false

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			fmt.Println("Error while reading stream", err)
			return nil
		}

		newNum := req.GetNum()

		if !isMaxInitialized {
			max = newNum
			isMaxInitialized = true
			err := stream.Send(&pb.ReturnIfMaxResponse{
				MaxTillNow: newNum,
			})

			if err != nil {
				fmt.Println("Error while sending newnum to stream : ", err)
				return err
			}
		} else {
			if int32(newNum) > max {
				max = int32(newNum)
				err = stream.Send(&pb.ReturnIfMaxResponse{
					MaxTillNow: max,
				})
				if err != nil {
					fmt.Println("Error while sending newnum to stream : ", err)
				}
			}
		}
	}
}

func main() {

	lis, err := net.Listen("tcp", "127.0.0.1:9091")
	if err != nil {
		fmt.Println("Failed to listen : ", err)
	}
	fmt.Println("Server started on 127.0.0.1:9091")

	sv := grpc.NewServer()
	pb.RegisterSumServiceServer(sv, &server{})
	pb.RegisterGetPrimesServiceServer(sv, &server{})
	pb.RegisterGetAverageServiceServer(sv, &server{})
	pb.RegisterReturnIfMaxServiceServer(sv, &server{})
	reflection.Register(sv)
	if err := sv.Serve(lis); err != nil {
		fmt.Println("Failed to serve : ", err)
	}

}
