package main

import (
	"fmt"
	// "os"

	"github.com/gin-gonic/gin"
	"github.com/nexentra/inteligpt/pkg/auth"
	"github.com/nexentra/inteligpt/pkg/comments"
	"github.com/nexentra/inteligpt/pkg/common/db"
	"github.com/nexentra/inteligpt/middlware/jsight"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           CloudNua: Comment Service
// @version         1.0
// @description     This is a simple CRUD HTTP service.
// @termsOfService  http://swagger.io/terms/

// @contact.name   CloudNua API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@cloudnua.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /v1/comment

// @securityDefinitions.basic  BasicAuth

// Run - is going to be responsible for
// the instantiation and startup of the
// go application
func Run() error {
	fmt.Println("starting up api service")

	return nil
}

func setupRouter() *gin.Engine {

	fmt.Println("Go-Gin REST API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}

	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	// port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	// dsn := fmt.Sprintf(
	// 	"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
	// 	os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_PORT"),
	// 	os.Getenv("DB_USERNAME"),
	// 	os.Getenv("DB_TABLE"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("SSL_MODE"),
	// )

	// fmt.Println(dsn)

	r := gin.Default()
	
	r.Use(jsight.Validator())

	h := db.InitDatabase(dbUrl)

	comments.RegisterRoutes(r, h)
	auth.RegisterRoutes(r,h)

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "CloudNua Comment Service",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}

func main() {
	r := setupRouter()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
