package main

import (
	"os"
	"path"

	//_ "github.com/maxim-shestakov/final-yandex-project/docs"

	// h "github.com/maxim-shestakov/final-yandex-project/pkg/handlers"
	// db "github.com/maxim-shestakov/final-yandex-project/internal/db"
	// repo "github.com/maxim-shestakov/final-yandex-project/pkg/repo"
	// "github.com/maxim-shestakov/final-yandex-project/pkg/service"

	gin "github.com/gin-gonic/gin"
	//swaggerFiles "github.com/swaggo/files"     // swagger embed files
	//ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title Swagger Example API
// @version 1.0

// @contact.name Max Shestakov
// @contact.url https://github.com/maxim-shestakov

func main() {
	// config := db.DBConfig{}
	// config.User =os.Getenv("POSTGRES_USER")
	// config.Password =os.Getenv("POSTGRES_PASSWORD")
	// config.Host = os.Getenv("POSTGRES_HOST")
	// config.Port = os.Getenv("POSTGRES_PORT")
	// config.Database = os.Getenv("POSTGRES_DB")
	// if len(config.User) == 0 || len(config.Password) == 0 || len(config.Host) == 0 || len(config.Port) == 0 || len(config.Database) == 0 {
	// 	log.Fatal("DB environment variables are not set")
	// }

	// DB, err := db.New(&config)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer DB.Close()

	// Repository:=repo.New(DB)
	// Service:=service.New(Repository)
	// Handlers:=h.New(Service)

	appPort := ":" + os.Getenv("TODO_PORT")
	if appPort == ":" {
		appPort = ":7540"
	}

	r := gin.Default()
	r.Use(gin.Logger(), gin.Recovery())
	webDir := "../web/"
	r.StaticFile("/", webDir)
	r.StaticFile("/js/scripts.min.js", path.Join(webDir, "js"))
	r.StaticFile("/css/style.css", path.Join(webDir, "css"))
	r.StaticFile("/favicon.ico", path.Join(webDir, "favicon.ico"))

	//r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := r.Run(appPort)
	if err != nil {
		panic(err)
	}
}
