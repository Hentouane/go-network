package tcp

import (
        "bufio"
        "fmt"
        "net"
        "os"
        "strings"
        "sync"
)

func StartTCPClient(wg *sync.WaitGroup, address string) {
        defer wg.Done()

        fmt.Println("Starting TCP Client...")

        c, err := net.Dial("tcp", address)
        if err != nil {
                fmt.Println(err)
                return
        }

        runTCPClient(c)
}

func runTCPClient(c net.Conn) {
        for {
                reader := bufio.NewReader(os.Stdin)
                fmt.Print(">> ")
                text, _ := reader.ReadString('\n')
                fmt.Fprintf(c, text+"\n")

                if strings.TrimSpace(string(text)) == "STOP" {
                        fmt.Println("TCP client exiting...")
                        return
                }


                message, _ := bufio.NewReader(c).ReadString('\n')
                fmt.Print("Client received ->: " + message)
        }
}
    