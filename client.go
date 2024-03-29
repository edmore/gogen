package main

import (
	"flag"
	"fmt"
	"github.com/edmore/gogen/service"
	"log"
	"net/rpc"
	"os"
)

var (
	ip   = flag.String("ip", "", "server IP address; must be set.")
	port = flag.Int("port", 9999, "default server port; can be reset.")
)

func main() {
	flag.Parse()
	if *ip == "" {
		flag.Usage()
		os.Exit(2)
	}

	client, err := rpc.DialHTTP("tcp", fmt.Sprintf(*ip+":%d", *port))
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var pong, ping string

	// Arith Service
	args := &service.Args{7, 8}
	var reply int
	err = client.Call("Arith.Mul", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)

	// Ping Service
	for ping != "exit" {
		fmt.Printf("\nPlease enter your ping string :\n")
		fmt.Scanf("%s", &ping)

		err = client.Call("Ping.Pong", ping, &pong)
		if err != nil {
			log.Fatal("ping error:", err)
		}
		fmt.Printf(pong)
	}
}
