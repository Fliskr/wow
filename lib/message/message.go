package message

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	SignalQuit int = iota + 1
	SignalRequestChallenge
	SignalResponseChallenge
	SignalRequestData
	SignalResponseData
	SignalError
)

type Message struct {
	Signal  int
	Message string
}

func (m *Message) String() string {
	return fmt.Sprintf("%d|%s\n", m.Signal, m.Message)
}

func (m *Message) Bytes() []byte {
	return []byte(m.String())
}

func (m *Message) Unmarshall(data []byte, message *Message) error {
	messageParts := strings.Split(string(data), "|")
	if len(messageParts) != 2 {
		return errors.New("Wrong message format")
	}

	signal, err := strconv.ParseInt(messageParts[0], 10, 32)
	if err != nil {
		return errors.New(fmt.Sprintf("Wrong message's signal format: %s", err))
	}
	message.Message = strings.Replace(messageParts[1], "\n", "", 1)
	message.Signal = int(signal)
	return nil
}

func (m *Message) Read(reader *bufio.Reader) error {
	data, err := reader.ReadBytes(byte(10))
	if err != nil {
		return err
	}
	err = m.Unmarshall(data, m)
	if err != nil {
		return err
	}

	return nil
}
