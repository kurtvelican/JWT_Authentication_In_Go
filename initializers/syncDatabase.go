package initializers

import "JWT_Authentication_In_Go/models"

func SyncDatabase() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		return
	}
}
