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

func (u *shopUsecase) GetShops() ([]domain.Shop, error) {
	return u.shopRepo.FindAll()
}

func (u *shopUsecase) GetShop(id uint64) (domain.Shop, error) {
	return u.shopRepo.FindByID(id)
}

func (u *shopUsecase) CreateShop(shop domain.Shop) (domain.Shop, error) {
	return u.shopRepo.Create(shop)
}

func (u *shopUsecase) UpdateShop(shop domain.Shop) (domain.Shop, error) {
	return u.shopRepo.Update(shop)
}

func (u *shopUsecase) DeleteShop(id uint64) error {
	return u.shopRepo.Delete(id)
}
