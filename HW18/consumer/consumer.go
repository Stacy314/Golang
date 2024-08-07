package consumer

import (
	"context"
	"encoding/json"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

const (
	orangesTopic = "oranges"
	brokerAddr   = "localhost:9092"
)

type Orange struct {
	OrangeID int `json:"OrangeID"`
	Size     int `json:"Size"`
}

func RunConsumer() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{brokerAddr},
		Topic:     orangesTopic,
		Partition: 0,
		MinBytes:  10e3, 
		MaxBytes:  10e6, 
	})
	defer reader.Close()

	small, medium, large := 0, 0, 0

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			log.Info().Msgf("Oranges: small=%d, medium=%d, large=%d", small, medium, large)
		default:
			m, err := reader.ReadMessage(context.Background())
			if err != nil {
				log.Fatal().Err(err).Msg("could not read message from kafka")
				continue
			}

			var orange Orange
			err = json.Unmarshal(m.Value, &orange)
			if err != nil {
				log.Error().Err(err).Msg("could not unmarshal message")
				continue
			}

			if orange.Size <= 100 {
				small++
			} else if orange.Size <= 200 {
				medium++
			} else {
				large++
			}

			log.Info().Msgf("Consumed message: %s", string(m.Value))
		}
	}
}
