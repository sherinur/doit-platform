package producer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/sherinur/doit-platform/user-service/internal/adapter/nats/producer/dto"
	"github.com/sherinur/doit-platform/user-service/internal/domain/model"
	"github.com/sherinur/doit-platform/user-service/pkg/nats"
)

const PushTimeout = time.Second * 30

type User struct {
	natsClient *nats.Client
}

func NewUserProducer(
	natsClient *nats.Client,
) *User {
	return &User{
		natsClient: natsClient,
	}
}

func (c *User) Push(ctx context.Context, user model.User) error {
	ctx, cancel := context.WithTimeout(ctx, PushTimeout)
	defer cancel()

	pbCustomer := dto.FromUser(user)
	data, err := json.Marshal(pbCustomer)
	if err != nil {
		return err
	}

	err = c.natsClient.Conn.Publish("user.pused", data)
	if err != nil {
		return fmt.Errorf("c.natsClient.Conn.Publish: %w", err)
	}
	log.Println("user is pushed:", user)

	return nil
}
