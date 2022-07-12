package methods

import (
	"bot/models"
	"bot/views"
)

func Read(data views.Data) interface{} {
	var db = models.DB
	var user models.User
	db.Where("user_id = ?", data.UserId).First(&user)
	if user.Id == 0 {
		user.UserId = data.UserId
		user.Integral = 100
		user.Permission = 0
		db.Create(&user)
	}
	return user
}
