package main

import (
	"gin-clean-architecture-api/pkg/infrastructure/persistence"
	"gin-clean-architecture-api/pkg/infrastructure/router"
	"gin-clean-architecture-api/pkg/interfaces/handler"
	"gin-clean-architecture-api/pkg/usecase"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// db起動待ち(マジックナンバーなのであまり良くない)
	time.Sleep(time.Second * 5)

	dsn := "root:password@tcp(CleanArchSampleDB)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Get the underlying SQL DB instance
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get SQL DB instance")
	}
	defer sqlDB.Close()

	shopRepository := persistence.NewShopRepository(db)
	shopUsecase := usecase.NewShopUsecase(shopRepository)
	shopHandler := handler.NewShopHandler(shopUsecase)

	r := router.SetupRouter(shopHandler)
	r.Run(":8080")
}
