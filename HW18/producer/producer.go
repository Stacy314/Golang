package producer

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

type IDGeneratorService struct {
	lastId int
}

func NewIDGenerator() *IDGeneratorService {
	return &IDGeneratorService{}
}

func (ig *IDGeneratorService) GenerateID() int {
	ig.lastId++
	return ig.lastId
}

const orangesTopic = "oranges"
const brokerAddr = "localhost:9092"

func RunProducer() {
	writer := kafka.Writer{
		Addr:                   kafka.TCP(brokerAddr),
		Topic:                  orangesTopic,
		AllowAutoTopicCreation: true,
	}

	defer writer.Close()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	idService := NewIDGenerator()

	for {
		select {
		case <-ticker.C:
			id := idService.GenerateID()
			size := rand.Intn(300) + 1 // Розмір апельсина від 1 до 300 см
			randomMessage := fmt.Sprintf(`{"OrangeID": %d, "Size": %d}`, id, size)
			err := writer.WriteMessages(context.Background(),
				kafka.Message{
					Value: []byte(randomMessage),
				},
			)
			if err != nil {
				log.Fatal().Err(err).Msg("could not write message to kafka")
			} else {
				log.Info().Msgf("Produced message: %s", randomMessage)
			}
		}
	}
}
