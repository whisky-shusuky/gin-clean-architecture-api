package handler

import (
	"gin-clean-architecture-api/pkg/domain"
	"gin-clean-architecture-api/pkg/usecase"
	"net/http"
	"strconv"

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

func (h *shopHandler) GetShops(c *gin.Context) {
	shops, err := h.shopUsecase.GetShops()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, shops)
}

func (h *shopHandler) GetShop(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid shop ID"})
		return
	}

	shop, err := h.shopUsecase.GetShop(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, shop)
}

func (h *shopHandler) CreateShop(c *gin.Context) {
	var shop domain.Shop
	if err := c.ShouldBindJSON(&shop); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdShop, err := h.shopUsecase.CreateShop(shop)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdShop)
}

func (h *shopHandler) UpdateShop(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid shop ID"})
		return
	}

	var shop domain.Shop
	if err := c.ShouldBindJSON(&shop); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	shop.ID = id

	updatedShop, err := h.shopUsecase.UpdateShop(shop)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedShop)
}

func (h *shopHandler) DeleteShop(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid shop ID"})
		return
	}

	if err := h.shopUsecase.DeleteShop(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
