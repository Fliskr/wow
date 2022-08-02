package message

import (
	"bufio"
	"bytes"
	"testing"
)

func TestMessage(t *testing.T) {
	t.Run("New message String()", func(t *testing.T) {
		msg := Message{}

		if msg.String() != "0|\n" {
			t.Errorf("Wrong initial values: result %s, want 0|", msg.String())
		}
	})

	t.Run("Unmarshal message", func(t *testing.T) {
		var msg Message

		err := msg.Unmarshall([]byte("1|test"), &msg)
		if err != nil {
			t.Errorf("Error unmarshalling message: %s", err)
		}

		if msg.Signal != 1 {
			t.Errorf("Wrong signal: result %d, want %d", msg.Signal, 1)
		}

		if msg.Message != "test" {
			t.Errorf("Wrong message: result %s, want %s", msg.Message, "test")
		}

		err = msg.Unmarshall([]byte("2|test2"), &msg)
		if err != nil {
			t.Errorf("Error unmarshalling message: %s", err)
		}

		if msg.Signal != 2 {
			t.Errorf("Wrong signal: result %d, want %d", msg.Signal, 2)
		}

		if msg.Message != "test2" {
			t.Errorf("Wrong message: result %s, want %s", msg.Message, "test2")
		}
	})

	t.Run("Read message", func(t *testing.T) {
		var msg Message
		buf := bytes.NewBuffer([]byte("0|test\n"))
		reader := bufio.NewReader(buf)

		err := msg.Read(reader)
		if err != nil {
			t.Errorf("Error reading message: %s", err)
		}
		if msg.String() != "0|test\n" {
			t.Errorf("Wrong message: result '%s', want '0|test'", msg.String())
		}
		if msg.Signal != 0 {
			t.Errorf("Wrong signal: result %d, want 0", msg.Signal)
		}

		reader.Reset(buf)
		buf.WriteString("1|test2\n")

		err = msg.Read(reader)
		if err != nil {
			t.Errorf("Error reading message: %s", err)
		}
		if msg.String() != "1|test2\n" {
			t.Errorf("Wrong message: result '%s', want '1|test2\n'", msg.String())
		}
		if msg.Signal != 1 {
			t.Errorf("Wrong signal: result %d, want 1", msg.Signal)
		}
	})

}
