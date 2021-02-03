package routes

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	r := gin.Default()
	r.Run()

	//r.POST("https://moneywave.herokuapp.com/v1/merchant/verify/", controllers)
}