package dictionary

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	. "github.com/rkfcccccc/english_words/shared_pkg/services/dictionary/models"
)

type repository struct {
	collection *mongo.Collection
}

func NewMongoRepository(collection *mongo.Collection) Repository {
	return &repository{collection}
}

func (repo *repository) CreateWordIndex(ctx context.Context) error {
	model := mongo.IndexModel{
		Keys: bson.M{
			"word": 1,
		}, Options: options.Index().SetUnique(true),
	}

	_, err := repo.collection.Indexes().CreateOne(ctx, model)
	return err
}

func (repo *repository) Create(ctx context.Context, entry *WordEntry) (string, error) {
	result, err := repo.collection.InsertOne(ctx, entry)
	if err != nil {
		return "", fmt.Errorf("collection.InsertOne: %v", err)
	}

	insertedId, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("error conversion %v to objectId", result.InsertedID)
	}

	return insertedId.Hex(), nil
}

func (repo *repository) GetById(ctx context.Context, wordId string) (*WordEntry, error) {
	objectId, err := primitive.ObjectIDFromHex(wordId)
	if err != nil {
		return nil, fmt.Errorf("primitive.ObjectIDFromHex: %v", err)
	}

	var result WordEntry
	err = repo.collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&result)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("collection.FindOne: %v", err)
	}

	return &result, nil
}

func (repo *repository) GetByWord(ctx context.Context, word string) (*WordEntry, error) {
	var result WordEntry
	err := repo.collection.FindOne(ctx, bson.M{"word": word}).Decode(&result)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (repo *repository) Delete(ctx context.Context, wordId string) error {
	objectId, err := primitive.ObjectIDFromHex(wordId)
	if err != nil {
		return fmt.Errorf("primitive.ObjectIDFromHex: %v", err)
	}

	_, err = repo.collection.DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		return fmt.Errorf("collection.DeleteOne: %v", err)
	}

	return nil
}

func (repo *repository) SetPictures(ctx context.Context, wordId string, pictures []SourcedPicture) error {
	objectId, err := primitive.ObjectIDFromHex(wordId)
	if err != nil {
		return fmt.Errorf("primitive.ObjectIDFromHex: %v", err)
	}

	_, err = repo.collection.UpdateByID(ctx, objectId, bson.D{{Key: "$set", Value: bson.M{"pictures": pictures}}})
	if err != nil {
		return fmt.Errorf("collection.DeleteOne: %v", err)
	}

	return nil
}

// TODO: manage limit and offsets with parameters
func (repo *repository) Search(ctx context.Context, query string) ([]*WordEntry, error) {
	filter := bson.M{"word": bson.M{"$regex": primitive.Regex{
		Pattern: "^" + query,
		Options: "i",
	}}}

	cursor, err := repo.collection.Find(ctx, filter, options.Find().SetLimit(20))
	if err != nil {
		return nil, fmt.Errorf("collection.Find: %v", err)
	}

	var result []*WordEntry
	if err := cursor.All(ctx, &result); err != nil {
		return nil, fmt.Errorf("cursor.All: %v", err)
	}

	if result == nil {
		result = []*WordEntry{}
	}

	return result, nil
}
