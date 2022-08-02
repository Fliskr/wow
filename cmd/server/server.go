package main

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"

	"wow/config"
	"wow/lib/message"
	"wow/res/store"
	"wow/res/store/wowstore"
)

func main() {

	listen, err := net.Listen(config.TYPE, config.HOST+":"+config.PORT)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer listen.Close()

	store := wowstore.NewStore()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		go handleIncomingRequest(conn, store)
	}
}

func handleIncomingRequest(conn net.Conn, s store.Store) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	nonce := rand.Intn(config.DIFFICULTY/2) + config.DIFFICULTY/2

	hashedSecret := sha1.New()
	hashedSecret.Write([]byte(fmt.Sprintf("%d", nonce)))
	hashedSecretHex := fmt.Sprintf("%x", hashedSecret.Sum(nil))
mainLoop:
	for {
		var rMsg message.Message
		err := rMsg.Read(reader)
		if err != nil {
			break
		}

		switch rMsg.Signal {
		case message.SignalRequestChallenge:
			msg := message.Message{
				Signal:  message.SignalResponseChallenge,
				Message: hashedSecretHex,
			}
			conn.Write(msg.Bytes())
		case message.SignalRequestData:
			if rMsg.Message != hashedSecretHex {
				break mainLoop
			}
			msg := message.Message{
				Signal:  message.SignalResponseData,
				Message: s.GetRandomQuote(),
			}
			conn.Write(msg.Bytes())
		case message.SignalQuit:
			break mainLoop
		default:
			break mainLoop
		}
	}
}
