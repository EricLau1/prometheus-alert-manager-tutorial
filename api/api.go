package api

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"prometheus-alert-manager-tutorial/api/db"
	"prometheus-alert-manager-tutorial/api/handlers"
	"prometheus-alert-manager-tutorial/api/httpext"
	"prometheus-alert-manager-tutorial/api/middlewares"
	"prometheus-alert-manager-tutorial/api/services"
	"prometheus-alert-manager-tutorial/api/store"
	_ "prometheus-alert-manager-tutorial/docs"
)

var port int

func init() {
	gin.ForceConsoleColor()
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
	flag.IntVar(&port, "p", 8080, "api port")
	db.LoadConfigs(flag.CommandLine)
	flag.Parse()
}

func Run() {

	var (
		conn         = db.New()
		todosStore   = store.NewTodosStore(conn)
		todosService = services.NewTodosService(todosStore)
		router       = gin.New()
	)

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.Metrics())
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	handlers.RegisterIndex(router)
	handlers.RegisterMetrics(router)
	handlers.RegisterTodos(router, todosService)
	log.Fatal(router.Run(httpext.Port(port).Addr()))
}
