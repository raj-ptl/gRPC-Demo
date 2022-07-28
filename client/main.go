package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/raj-ptl/gRPC-Demo/pb"
	"google.golang.org/grpc"
)

func GetSum(clt pb.SumServiceClient) {
	fmt.Println("Sending a GetSum request")
	req := &pb.SumRequest{
		Num_1: 15,
		Num_2: 4,
	}
	res, err := clt.GetSum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while invoking Get Sum Call : %v\n", err)
		return
	}

	fmt.Println("Response from Get Sum Call : ", res.SumResult)
}

func GetPrimes(clt pb.GetPrimesServiceClient) {
	fmt.Println("Sending a get primes request")
	req := &pb.GetPrimesRequest{
		Num: 30,
	}

	res, err := clt.GetPrimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while invoking Get Primes Call :%v\n ", err)
	}

	for {
		stream, err := res.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream : %v\n", err)
			return
		}

		fmt.Println("resp froms svr : ", stream.GetNum())
	}
}

func GetAverage(clt pb.GetAverageServiceClient) {
	fmt.Println("Sending a GetAverage request")

	stream, err := clt.GetAverage(context.Background())

	if err != nil {
		log.Fatalf("Error while invoking Get Average Call : %v\n", err)
	}

	sampleNumbers := []*pb.GetAverageRequest{
		{
			Num: 10,
		},
		{
			Num: 20,
		},
		{
			Num: 30,
		},
		{
			Num: 40,
		},
	}

	for _, val := range sampleNumbers {
		fmt.Printf("Sending num : %v\n", val)
		err := stream.Send(val)
		if err != nil {
			log.Fatalf("Error while sending value : %v\n", err)
		}
		time.Sleep(500 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while CloseAndRecv() : %v\n", err)
	}

	fmt.Printf("Response from Server : %v\n", res.Avg)
}

func ReturnIfMax(clt pb.ReturnIfMaxServiceClient) {
	fmt.Println("Sending a ReturnIfMax request")

	sampleNumbers := []int32{1, 3, 7, 9, 2, 5, 22, 22, 15, 21, 19}

	stream, err := clt.ReturnIfMax(context.Background())

	if err != nil {
		log.Fatalf("Error while invoking ReturnIfMax Call : %v\n", err)
	}

	ch := make(chan int)

	go func() {
		for _, val := range sampleNumbers {
			fmt.Printf("Sending num : %v\n", val)
			err := stream.Send(&pb.ReturnIfMaxRequest{
				Num: val,
			})

			if err != nil {
				log.Fatalf("Error while sending value : %v\n", err)
			}
			time.Sleep(500 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while Recv() : %v\n", err)
			}

			fmt.Printf("Response from Server : %v\n", res.MaxTillNow)
		}
		close(ch)
	}()

	<-ch
}

func main() {
	fmt.Printf("Client Started ...\n")
	//fmt.Println("CLI arg passed : ", os.Args[1])
	newConn, err := grpc.Dial("127.0.0.1:9001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to establish connection to 127.0.0.1:9001 : %v\n", err)
	}
	defer newConn.Close()

	switch os.Args[1] {
	case "sum":
		clt := pb.NewSumServiceClient(newConn)
		GetSum(clt)

	case "primes":
		clt := pb.NewGetPrimesServiceClient(newConn)
		GetPrimes(clt)

	case "max":
		clt := pb.NewReturnIfMaxServiceClient(newConn)
		ReturnIfMax(clt)

	case "avg":
		clt := pb.NewGetAverageServiceClient(newConn)
		GetAverage(clt)

	default:
		log.Printf("Pass a valid cli arg. Exiting ...")
	}
}
