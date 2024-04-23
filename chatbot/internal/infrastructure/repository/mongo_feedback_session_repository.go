package repository

import (
	"chatbot/internal/domain/model"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoFeedbackSessionRepository 是 FeedbackSessionRepository 的具体 MongoDB 实现
type MongoFeedbackSessionRepository struct {
	collection *mongo.Collection
}

// NewMongoFeedbackSessionRepository 创建一个新的 MongoDB 反馈会话仓库实例
func NewMongoFeedbackSessionRepository(db *mongo.Database, collectionName string) *MongoFeedbackSessionRepository {
	return &MongoFeedbackSessionRepository{
		collection: db.Collection(collectionName),
	}
}

// Save 存储反馈会话到 MongoDB
func (repo *MongoFeedbackSessionRepository) Save(ctx context.Context, session *model.FeedbackSession) error {
	_, err := repo.collection.ReplaceOne(
		ctx,
		bson.M{"sessionId": session.SessionID},
		session,
		options.Replace().SetUpsert(true),
	)
	return err
}

// FindByID 根据 ID 从 MongoDB 查找反馈会话
func (repo *MongoFeedbackSessionRepository) FindByID(ctx context.Context, sessionID string) (*model.FeedbackSession, error) {
	var session model.FeedbackSession
	err := repo.collection.FindOne(ctx, bson.M{"sessionId": sessionID}).Decode(&session)
	if err != nil {
		return nil, err
	}
	return &session, nil
}
