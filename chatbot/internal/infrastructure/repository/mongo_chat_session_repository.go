package repository

import (
	"chatbot/internal/domain/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoChatSessionRepository 是 ChatSessionRepository 的 MongoDB 实现
type MongoChatSessionRepository struct {
	collection *mongo.Collection
}

// NewMongoChatSessionRepository 创建新的 MongoDB 聊天会话仓库实例
func NewMongoChatSessionRepository(db *mongo.Database, collectionName string) *MongoChatSessionRepository {
	return &MongoChatSessionRepository{
		collection: db.Collection(collectionName),
	}
}

func (repo *MongoChatSessionRepository) Save(ctx context.Context, session *model.ChatSession) error {
	_, err := repo.collection.ReplaceOne(
		ctx,
		bson.M{"id": session.ID},
		session,
		options.Replace().SetUpsert(true),
	)
	return err
}

func (repo *MongoChatSessionRepository) FindByID(ctx context.Context, sessionID string) (*model.ChatSession, error) {
	var session model.ChatSession
	err := repo.collection.FindOne(ctx, bson.M{"id": sessionID}).Decode(&session)
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (repo *MongoChatSessionRepository) AddMessage(ctx context.Context, sessionID string, message model.Message) error {
	filter := bson.M{"id": sessionID}
	update := bson.M{"$push": bson.M{"messages": message}}
	_, err := repo.collection.UpdateOne(ctx, filter, update)
	return err
}

func (repo *MongoChatSessionRepository) GetMessageHistory(ctx context.Context, sessionID string) ([]model.Message, error) {
	var session model.ChatSession
	err := repo.collection.FindOne(ctx, bson.M{"id": sessionID}).Decode(&session)
	if err != nil {
		return nil, err
	}
	return session.Messages, nil
}

func (repo *MongoChatSessionRepository) RemoveMessage(ctx context.Context, sessionID string, messageID string) error {
	filter := bson.M{"id": sessionID}
	update := bson.M{"$pull": bson.M{"messages": bson.M{"id": messageID}}}
	_, err := repo.collection.UpdateOne(ctx, filter, update)
	return err
}
