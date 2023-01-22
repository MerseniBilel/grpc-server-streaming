package counter

import (
	"errors"
	"fmt"

	pb "github.com/MerseniBilel/streamrpc/grpc/counter"
)

type counterServiceImplement struct {
	pb.UnimplementedCounterServiceServer
}

// returns new instance of counterServiceimplement
func NewCounterService() *counterServiceImplement {
	return &counterServiceImplement{}
}

// implement the function initCounter that takes an integer n and returns a stream from n -> 0
func (s *counterServiceImplement) InitCounter(in *pb.CounterRequest, srv pb.CounterService_InitCounterServer) error {
	fmt.Printf("returning counter response from number %d \n", in.InitCounter)
	quit := make(chan bool)

	go func() {
		for i := 0; i < int(in.InitCounter); i++ {
			SendStream(i, srv, quit)

		}
	}()

	for result := range quit {
		if result {
			return errors.New("error")
		}
	}

	defer close(quit)
	return nil
}

func SendStream(i int, srv pb.CounterService_InitCounterServer, c chan bool) { // sleep
	resp := pb.CounterResponse{Counter: int32(i)}

	err := srv.Send(&resp)
	if err != nil {
		c <- true
	} else {
		fmt.Printf("sending response number %d \n", i)
	}

	// if err := srv.Send(&resp); err != nil {
	// 	c <- true
	// }

}
