package messaging

import (
	"encoding/binary"
	"encoding/json"
	"io"
)

type Message struct {
	Type   string `json:"type"`
	Action string `json:"action,omitempty"`
}

func ReadMessage(r io.Reader) (*Message, error) {
	var length uint32
	if err := binary.Read(r, binary.LittleEndian, &length); err != nil {
		return nil, err
	}

	payload := make([]byte, length)
	if _, err := io.ReadFull(r, payload); err != nil {
		return nil, err
	}

	var msg Message
	if err := json.Unmarshal(payload, &msg); err != nil {
		return nil, err
	}

	return &msg, nil
}

func WriteMessage(w io.Writer, msg Message) error {
	payload, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	if err := binary.Write(w, binary.LittleEndian, uint32(len(payload))); err != nil {
		return err
	}

	_, err = w.Write(payload)
	return err
}
