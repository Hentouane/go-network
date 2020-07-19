package udp

import (
        "bufio"
        "fmt"
        "net"
        "os"
		"strings"
		"sync"
)

func StartUDPClient(wg *sync.WaitGroup, address string) {
	defer wg.Done()

	s, err := net.ResolveUDPAddr("udp4", address)
	c, err := net.DialUDP("udp4", nil, s)
	if err != nil {
			fmt.Println(err)
			return
	}

	fmt.Printf("The UDP server is %s\n", c.RemoteAddr().String())
	defer c.Close()

	runUDPClient(c)
}

func runUDPClient(c *net.UDPConn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		data := []byte(text + "\n")
		_, err := c.Write(data)
		if strings.TrimSpace(string(data)) == "STOP" {
				fmt.Println("Exiting UDP client!")
				return
		}

		if err != nil {
				fmt.Println(err)
				return
		}

		buffer := make([]byte, 1024)
		n, _, err := c.ReadFromUDP(buffer)
		if err != nil {
				fmt.Println(err)
				return
		}
		fmt.Printf("Client received->: %s\n", string(buffer[0:n]))
}
}
      