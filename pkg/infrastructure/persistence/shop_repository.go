package persistence

import (
	"gin-clean-architecture-api/pkg/interfaces/repository"

	"gorm.io/gorm"
)

type shopRepository struct {
	db *gorm.DB
}

func NewShopRepository(db *gorm.DB) repository.ShopRepository {
	return &shopRepository{db}
}

// 省略: FindAll, FindByID, Create, Update, Delete メソッドの実装
