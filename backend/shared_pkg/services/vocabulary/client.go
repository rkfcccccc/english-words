package vocabulary

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/rkfcccccc/english_words/proto/vocabulary"
	"github.com/segmentio/kafka-go"
)

type Client struct {
	conn  *grpc.ClientConn
	grpc  pb.VocabularyServiceClient
	kafka *kafka.Writer
}

func NewClient(grpcAddr string, kafkaAddr string) *Client {
	conn, err := grpc.Dial(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	producer := kafka.Writer{
		Addr:  kafka.TCP(kafkaAddr),
		Topic: "vocabulary",
	}

	client := pb.NewVocabularyServiceClient(conn)
	return &Client{conn, client, &producer}
}

func (c *Client) GetChallenge(ctx context.Context, userId int) (*Challenge, error) {
	response, err := c.grpc.GetChallenge(ctx, &pb.GetChallengeRequest{UserId: int32(userId)})
	return &Challenge{WordId: response.GetWordId(), LearningStep: int(response.GetLearningStep())}, err
}

func (c *Client) PromoteWord(ctx context.Context, userId int, wordId string) error {
	_, err := c.grpc.PromoteWord(ctx, &pb.PromoteWordRequest{UserId: int32(userId), WordId: wordId})
	return err
}

func (c *Client) ResistWord(ctx context.Context, userId int, wordId string) error {
	_, err := c.grpc.ResistWord(ctx, &pb.ResistWordRequest{UserId: int32(userId), WordId: wordId})
	return err
}

func (c *Client) SetAlreadyLearned(ctx context.Context, userId int, wordId string, isAlreadyLearned bool) error {
	_, err := c.grpc.SetAlreadyLearned(ctx, &pb.SetAlreadyLearnedRequest{UserId: int32(userId), WordId: wordId, IsAlreadyLearned: isAlreadyLearned})
	return err
}

func (c *Client) writeWordActionsToKafka(ctx context.Context, userId int, add bool, wordIds ...string) error {
	messages := make([]kafka.Message, len(wordIds))
	for i, wordId := range wordIds {
		action := WordAction{UserId: userId, WordId: wordId, Add: add}

		bytes, err := json.Marshal(action)
		if err != nil {
			return fmt.Errorf("json.Marshal: %v", err)
		}

		messages[i].Value = bytes
	}

	err := c.kafka.WriteMessages(ctx, messages...)
	if err != nil {
		return fmt.Errorf("kafka.WriteMessages: %v", err)
	}

	return nil
}

func (c *Client) AddWords(ctx context.Context, userId int, wordIds ...string) error {
	return c.writeWordActionsToKafka(ctx, userId, true, wordIds...)
}

func (c *Client) DeleteWords(ctx context.Context, userId int, wordIds ...string) error {
	return c.writeWordActionsToKafka(ctx, userId, true, wordIds...)
}
