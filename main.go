package main

import (
	"gorm.io/gorm"
	"medici.vn/commission-serivce/config"
	api "medici.vn/commission-serivce/routes"
	"os"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()
)

func main() {
	defer config.CloseDatabaseConnection(db)

	r := api.Router()

	var port = os.Getenv("PORT")

	err := r.Run("0.0.0.0:" + port)

	if err != nil {
		return
	}
}
