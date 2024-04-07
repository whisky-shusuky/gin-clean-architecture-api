package handler

import (
	"gin-clean-architecture-api/pkg/usecase"

	"github.com/gin-gonic/gin"
)

type ShopHandler interface {
	GetShops(c *gin.Context)
	GetShop(c *gin.Context)
	CreateShop(c *gin.Context)
	UpdateShop(c *gin.Context)
	DeleteShop(c *gin.Context)
}

type shopHandler struct {
	shopUsecase usecase.ShopUsecase
}

func NewShopHandler(shopUsecase usecase.ShopUsecase) ShopHandler {
	return &shopHandler{shopUsecase}
}

// 省略: GetShops, GetShop, CreateShop, UpdateShop, DeleteShop メソッドの実装
