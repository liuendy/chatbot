package repository

import (
	"chatbot/internal/domain/model"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoMessageTemplateRepository struct {
	collection *mongo.Collection
}

// NewMongoMessageTemplateRepository 创建一个新的消息模板仓库实例
func NewMongoMessageTemplateRepository(db *mongo.Database, collectionName string) *MongoMessageTemplateRepository {
	return &MongoMessageTemplateRepository{
		collection: db.Collection(collectionName),
	}
}

// Save 保存或更新消息模板
func (repo *MongoMessageTemplateRepository) Save(ctx context.Context, template *model.MessageTemplate) error {
	_, err := repo.collection.ReplaceOne(
		ctx,
		bson.M{"templateId": template.TemplateID},
		template,
		options.Replace().SetUpsert(true),
	)
	return err
}

// FindByID 根据模板ID查找消息模板
func (repo *MongoMessageTemplateRepository) FindByID(ctx context.Context, templateID string) (*model.MessageTemplate, error) {
	var template model.MessageTemplate
	err := repo.collection.FindOne(ctx, bson.M{"templateId": templateID}).Decode(&template)
	if err != nil {
		return nil, err
	}
	return &template, nil
}

// Delete 根据模板ID删除消息模板
func (repo *MongoMessageTemplateRepository) Delete(ctx context.Context, templateID string) error {
	_, err := repo.collection.DeleteOne(ctx, bson.M{"templateId": templateID})
	return err
}

// FindAll 获取所有消息模板
func (repo *MongoMessageTemplateRepository) FindAll(ctx context.Context) ([]*model.MessageTemplate, error) {
	var templates []*model.MessageTemplate
	cursor, err := repo.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &templates); err != nil {
		return nil, err
	}
	return templates, nil
}
