//
//package main
//
//import (
//	"chatbot/internal/application/service"
//	"chatbot/internal/domain/model"
//	"chatbot/internal/infrastructure/repository"
//	"chatbot/internal/interfaces/web"
//	"context"
//"go.mongodb.org/mongo-driver/mongo"
//"go.mongodb.org/mongo-driver/mongo/options"
//"log"
//	"os"
//)
//
//func main() {
//	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongodb:27017"))
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer client.Disconnect(context.TODO())
//
//	db := client.Database("chatbot")
//	chatSessionRepo := repository.NewMongoChatSessionRepository(db, "chat_sessions")
//
//	// 示例：使用仓库实例
//	ctx := context.Background()
//	session := model.NewChatSession("user123")
//	err = chatSessionRepo.Save(ctx, session)
//	if err != nil {
//		log.Printf("Failed to save session: %v", err)
//	}
//
//	retrievedSession, err := chatSessionRepo.FindByID(ctx, session.ID)
//	if err != nil {
//		log.Printf("Failed to retrieve session: %v", err)
//	}
//	log.Printf("Retrieved session: %+v", retrievedSession)
//}

package main

import (
	"chatbot/internal/application/service"
	"chatbot/internal/infrastructure/repository"
	"chatbot/internal/interfaces/web"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func main() {
	// 获取环境变量
	mongoHost := "localhost:27017"

	//mongoHost := "mongo:27017"

	port := "8080" // 默认端口

	// 构建连接字符串
	mongoURI := fmt.Sprintf("mongodb://admin:admin@%s/chatbot?authSource=admin", mongoHost)

	// 创建 MongoDB 客户端
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {

		}
	}(client, context.TODO())

	// 选择数据库和集合
	db := client.Database("chatbot")

	chatRepo := repository.NewMongoChatSessionRepository(db, "chat_sessions")
	feedbackRepo := repository.NewMongoFeedbackSessionRepository(db, "feedback_sessions")
	templateRepo := repository.NewMongoMessageTemplateRepository(db, "message_templates")

	// 初始化服务
	chatService := service.NewChatService(chatRepo)
	feedbackService := service.NewFeedbackService(feedbackRepo)
	templateService := service.NewMessageTemplateService(templateRepo)

	// 初始化路由器
	router := web.NewRouter(chatService, feedbackService, templateService)

	// 启动服务器
	log.Printf("Starting server on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
