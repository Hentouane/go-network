package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	U "github.com/Hentouane/go-network/UDP"
	T "github.com/Hentouane/go-network/TCP"
)

func main() {
	host := flag.String("host", "127.0.0.1", "The host address")
	port := flag.String("port", "555", "The host address")
	comm := flag.String("comm", "TCP", "The type of network protocol used.")
	flag.Parse()

	address := *host + ":" + *port
	fmt.Printf("Protocol used: %s\n", *comm)
	fmt.Printf("Address provided: %s\n", address)

	SetupCloseHandler()

	var wg sync.WaitGroup

	switch *comm {
		case "TCP":	UseTCP(&wg, address)
		case "UDP":	UseUDP(&wg, address)
		default:
			fmt.Printf("Communication protocol %s not recognized, exiting.\n", *comm)
			os.Exit(1)
	}
	
	wg.Wait()
}

func SetupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		os.Exit(0)
	}()
}

func UseTCP(wg *sync.WaitGroup, address string) {
	wg.Add(1)

	serverReady := make(chan bool,1)

	go T.StartTCPServer(wg, address, serverReady)

	<- serverReady

	wg.Add(1)

	go T.StartTCPClient(wg, address)
}

func UseUDP(wg *sync.WaitGroup, address string) {
	wg.Add(1)

	serverReady := make(chan bool,1)

	go U.StartUDPServer(wg, address, serverReady)

	<- serverReady

	wg.Add(1)

	go U.StartUDPClient(wg, address)
}