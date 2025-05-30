package dto

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
)

type DeleteId struct {
	Id string `json:"id"`
}

func ToId(msg *nats.Msg) (string, error) {
	var id DeleteId
	err := json.Unmarshal(msg.Data, &id)
	if err != nil {
		return "", fmt.Errorf("json.Unmarshall: %w", err)
	}

	return id.Id, nil
}
