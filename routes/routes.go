package routes

import (
	"fmt"

	_ "github.com/angelicalombas/merge-list-api/docs"

	"github.com/angelicalombas/merge-list-api/controllers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/SaveLists", controllers.SaveLists)
	r.GET("/Merge", controllers.Merge)

	fmt.Println("Server started at port 8080")
	r.Run(":8080")
}
