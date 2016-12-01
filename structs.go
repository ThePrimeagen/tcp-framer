package tcp

import "fmt"

type Responder struct {
    resp chan *Responder
    message []byte
    length int
}

func (resp *Responder) Respond(out []byte) {
    fmt.Println("This should really work, but it does not currently")
}

func (resp Responder) String() {
    return fmt.Sprintf(
        "I am a responder with message length %v\nmessage: %v\n")
}

