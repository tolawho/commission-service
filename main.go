package main

import (
	"gorm.io/gorm"
	"medici.vn/commission-serivce/config"
	api "medici.vn/commission-serivce/routes"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()
)

func main() {
	defer config.CloseDatabaseConnection(db)

	r := api.Router()

	err := r.Run("0.0.0.0:8081")

	if err != nil {
		return
	}
}
