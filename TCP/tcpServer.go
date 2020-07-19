package tcp

import (
        "bufio"
        "fmt"
        "net"
        "strings"
        "sync"
        "time"
)

func StartTCPServer(wg *sync.WaitGroup, address string, serverReady chan bool) {
        defer wg.Done()

        fmt.Println("Starting TCP Server...")
        
        l, err := net.Listen("tcp", address)
        if err != nil {
                fmt.Println(err)
                return
        }
        defer l.Close()


        fmt.Println("Server listening, sending signal to channel...")
        serverReady <- true

        runTCPServer(l)
}

func runTCPServer(l net.Listener) {
        c, err := l.Accept()
        if err != nil {
                fmt.Println(err)
                return
        }

        handleConnection(c)
}

func handleConnection(c net.Conn) {
        for {
                netData, err := bufio.NewReader(c).ReadString('\n')
                if err != nil {
                        fmt.Println(err)
                        return
                }

                fmt.Print("Server received -> ", string(netData))
                t := time.Now()
                myTime := t.Format(time.RFC3339) + "\n"
                c.Write([]byte(myTime))

                if strings.TrimSpace(string(netData)) == "STOP" {
                        fmt.Println("Exiting TCP server!")
                        return
                }
        }
}
    