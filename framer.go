package tcp

import "fmt"
import "bufio"
import "net"
import "encoding/binary"

func FrameConnections(tcpListener net.Listener) chan *Responder {
    fmt.Println("Now attempting to listen")

    outChannel := make(chan *Responder, 100)
    go handleConnections(tcpListener, outChannel)
    return outChannel
}

// This should spin forever
func handleConnections(tcpListener net.Listener, outChannel chan *Responder) {
    for {
        fmt.Println("Awaiting connection...")
        conn, acceptErr := tcpListener.Accept()
        if acceptErr != nil {
            fmt.Printf("The was an accept error %v/n", acceptErr)
            continue
        }

        go handleMessages(conn, outChannel)
    }
}

// this should also spin until the connection closes.
func handleMessages(conn net.Conn, outChannel chan *Responder) {
    reader := bufio.NewReader(conn)
    messageLength := int(0)

    for {
        fmt.Printf("About to start reading from new connection\n")
        binary.Read(reader, binary.LittleEndian, &messageLength)
        fmt.Printf("Message recieved with length %v\n", messageLength)
        buffer := make([]byte, messageLength)

        reader.Read(buffer)

        fmt.Printf("Message recieved %v\n", buffer)
        outChannel <- &Responder{outChannel, buffer, messageLength}
    }
}
