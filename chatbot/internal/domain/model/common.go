package model

import "github.com/google/uuid"

// generateID 生成一个新的全局唯一标识符 (UUID)
func GenerateID() string {
	return uuid.New().String()
}
