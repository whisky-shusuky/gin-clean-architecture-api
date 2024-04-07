package usecase

import (
	"gin-clean-architecture-api/pkg/domain"
	"gin-clean-architecture-api/pkg/interfaces/repository"
)

type ShopUsecase interface {
	GetShops() ([]domain.Shop, error)
	GetShop(id uint64) (domain.Shop, error)
	CreateShop(shop domain.Shop) (domain.Shop, error)
	UpdateShop(shop domain.Shop) (domain.Shop, error)
	DeleteShop(id uint64) error
}

type shopUsecase struct {
	shopRepo repository.ShopRepository
}

func NewShopUsecase(shopRepo repository.ShopRepository) ShopUsecase {
	return &shopUsecase{shopRepo}
}

// 省略: GetShops, GetShop, CreateShop, UpdateShop, DeleteShop メソッドの実装
