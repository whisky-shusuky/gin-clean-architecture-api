package persistence

import (
	"gin-clean-architecture-api/pkg/domain"
	"gin-clean-architecture-api/pkg/interfaces/repository"

	"gorm.io/gorm"
)

type shopRepository struct {
	db *gorm.DB
}

func NewShopRepository(db *gorm.DB) repository.ShopRepository {
	return &shopRepository{db}
}

func (r *shopRepository) Create(shop domain.Shop) (domain.Shop, error) {
	if err := r.db.Create(&shop).Error; err != nil {
		return domain.Shop{}, err
	}
	return shop, nil
}

func (r *shopRepository) FindAll() ([]domain.Shop, error) {
	var shops []domain.Shop
	if err := r.db.Find(&shops).Error; err != nil {
		return nil, err
	}
	return shops, nil
}

func (r *shopRepository) FindByID(id uint64) (domain.Shop, error) {
	var shop domain.Shop
	if err := r.db.First(&shop, id).Error; err != nil {
		return domain.Shop{}, err
	}
	return shop, nil
}

func (r *shopRepository) Update(shop domain.Shop) (domain.Shop, error) {
	if err := r.db.Save(&shop).Error; err != nil {
		return domain.Shop{}, err
	}
	return shop, nil
}

func (r *shopRepository) Delete(id uint64) error {
	if err := r.db.Delete(&domain.Shop{}, id).Error; err != nil {
		return err
	}
	return nil
}
