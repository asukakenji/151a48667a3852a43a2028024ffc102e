package main

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/kr/beanstalk"
)

const (
	// TimeForever represents a very large duration.
	TimeForever = math.MaxUint32 * time.Second
)

func main() {
	conn, err := beanstalk.Dial("tcp", "128.0.0.1:11300")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	id, err := conn.Put(([]byte)(os.Args[1]), 0, 0, 5*time.Second)
	if err != nil {
		if cerr, ok := err.(beanstalk.ConnError); !ok {
			panic(fmt.Errorf("Unknown error: %v", err))
		} else if cerr.Err == beanstalk.ErrBuried {
			fmt.Fprintf(os.Stderr, "Buried\n")
			panic(cerr)
		} else if cerr.Err == beanstalk.ErrNoCRLF {
			fmt.Fprintf(os.Stderr, "Expected CRLF\n")
			panic(cerr)
		} else if cerr.Err == beanstalk.ErrJobTooBig {
			fmt.Fprintf(os.Stderr, "Job too big\n")
			panic(cerr)
		} else if cerr.Err == beanstalk.ErrDraining {
			fmt.Fprintf(os.Stderr, "Draining\n")
			panic(cerr)
		} else {
			panic(fmt.Errorf("Unknown error: %v", err))
		}
	}
	fmt.Printf("id: %d\n", id)
}

func main1() {
	conn, err := beanstalk.Dial("tcp", "127.0.0.1:11300")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	for {
		id, body, err := conn.Reserve(5 * time.Second)
		if err != nil {
			if cerr, ok := err.(beanstalk.ConnError); !ok {
				panic(fmt.Errorf("Unknown error: %v", err))
			} else if cerr.Err == beanstalk.ErrTimeout {
				fmt.Fprintf(os.Stderr, "Timeout\n")
				continue
			} else if cerr.Err == beanstalk.ErrDeadline {
				fmt.Fprintf(os.Stderr, "Deadline soon\n")
				time.Sleep(1 * time.Second)
				continue
			} else {
				panic(fmt.Errorf("Unknown error: %v", err))
			}
		}
		fmt.Printf("id: %d\n", id)
		fmt.Printf("body: %s\n", body)
	}
}
