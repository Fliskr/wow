package main

import (
	"bufio"
	"context"
	"log"
	"net"
	"time"

	"wow/lib/message"
	"wow/lib/pow"
)

const timeout = time.Second * 10

func main() {
	connect()
}

func connect() {
	var d net.Dialer
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	conn, err := d.DialContext(ctx, "tcp", "localhost:9001")
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}

	defer conn.Close()

	msg := message.Message{
		Signal: message.SignalRequestChallenge,
	}
	conn.Write(msg.Bytes())

	reader := bufio.NewReader(conn)

	var rMsg message.Message
	err = rMsg.Read(reader)
	if err != nil {
		log.Fatal(err)
		return
	}

	p := pow.Pow{
		Hash: rMsg.Message,
	}
	if err := p.Solve(); err != nil {
		log.Fatal(err)
		return
	}

	msg = message.Message{
		Signal:  message.SignalRequestData,
		Message: p.Hash,
	}
	conn.Write(msg.Bytes())

	err = rMsg.Read(reader)
	if err != nil {
		log.Fatal(err)
		return
	}

	println(rMsg.Message)
}
