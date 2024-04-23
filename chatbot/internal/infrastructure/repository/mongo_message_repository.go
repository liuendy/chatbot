package repository

import (
	"chatbot/internal/domain/model"
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoMessageRepository struct {
	collection *mongo.Collection
}

// NewMongoMessageRepository 创建一个新的消息仓库实例
func NewMongoMessageRepository(db *mongo.Database, collectionName string) *MongoMessageRepository {
	return &MongoMessageRepository{
		collection: db.Collection(collectionName),
	}
}

// Save 保存或更新消息
func (repo *MongoMessageRepository) Save(ctx context.Context, message *model.Message) error {
	_, err := repo.collection.ReplaceOne(
		ctx,
		bson.M{"messageID": message.ID},
		message,
		options.Replace().SetUpsert(true),
	)
	return err
}

// FindBySessionID 根据会话ID查找所有相关消息
func (repo *MongoMessageRepository) FindBySessionID(ctx context.Context, sessionID string) ([]*model.Message, error) {
	var messages []*model.Message
	cursor, err := repo.collection.Find(ctx, bson.M{"sessionID": sessionID})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &messages); err != nil {
		return nil, err
	}
	return messages, nil
}

// Delete 根据消息ID删除消息
func (repo *MongoMessageRepository) Delete(ctx context.Context, messageID string) error {
	_, err := repo.collection.DeleteOne(ctx, bson.M{"messageID": messageID})
	return err
}
