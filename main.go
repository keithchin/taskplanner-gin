// main.go

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/keithchin/taskplanner-gin/common/database"
	"github.com/keithchin/taskplanner-gin/common/handlers/tasks"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	r := gin.Default()
	h := database.Init(dbUrl)

	tasks.RegisterRoutes(r, h)
	// register more routes here

	r.Run(port)
}
