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

        stopServer := make(chan bool)

        fmt.Println("Starting TCP Server...")
        
        l, err := net.Listen("tcp", address)
        if err != nil {
                fmt.Println(err)
                return
        }
        defer l.Close()


        fmt.Println("Server listening, sending signal to channel...")
        serverReady <- true

        go runTCPServer(l, stopServer)

        <- stopServer
        
        return
}

func runTCPServer(l net.Listener, stopServer chan bool) {
        for {
                c, err := l.Accept()
                if err != nil {
                        fmt.Println(err)
                        return
                }

                go handleConnection(c, stopServer)
        }
}

func handleConnection(c net.Conn, stopServer chan bool) {
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
                        stopServer <- true
                        return
                }
        }
}
    