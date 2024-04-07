package repository

import "gin-clean-architecture-api/pkg/domain"

type ShopRepository interface {
	FindAll() ([]domain.Shop, error)
	FindByID(id uint64) (domain.Shop, error)
	Create(shop domain.Shop) (domain.Shop, error)
	Update(shop domain.Shop) (domain.Shop, error)
	Delete(id uint64) error
}
