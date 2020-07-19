package udp

import (
	"fmt"
	"math/rand"
	"net"
	"strings"
	"sync"
	"time"
)

func StartUDPServer(wg *sync.WaitGroup, address string, serverReady chan bool) {
	defer wg.Done()

	s, err := net.ResolveUDPAddr("udp4", address)
	if err != nil {
			fmt.Println(err)
			return
	}

	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
			fmt.Println(err)
			return
	}

	serverReady <- true

	defer connection.Close()
	buffer := make([]byte, 1024)
	rand.Seed(time.Now().Unix())

	for {
			n, addr, err := connection.ReadFromUDP(buffer)
			fmt.Println("Server received -> ", string(buffer[0:n-1]))

			if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
					fmt.Println("Exiting UDP server!")
					return
			}

			t := time.Now()
			myTime := t.Format(time.RFC3339) + "\n"
			data := []byte(myTime)
			
			_, err = connection.WriteToUDP(data, addr)
			if err != nil {
					fmt.Println(err)
					return
			}
	}
}