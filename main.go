package main

import (
	"context"
	"errors"
	"fmt"
	"go_blog/controller"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func setRouter() *gin.Engine {
	router := gin.New()
	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC3339Nano),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome Go Blog Server")
	})
	router.POST("/login", controller.Login)
	router.POST("/register", controller.Register)

	loginCheckRouter := router.Group("/", controller.AuthCheck())
	{
		loginCheckRouter.POST("/logout", controller.Logout)
		userRouter := loginCheckRouter.Group("/user")
		{
			userRouter.GET("/", controller.GetUser)
			userRouter.POST("/", controller.UpdateUser)
			userRouter.DELETE("/", controller.DeleteUser)
		}
		articleRouter := loginCheckRouter.Group("/article")
		{
			articleRouter.GET("/", controller.GetArticles)
			articleRouter.POST("/", controller.CreateArticle)
			articleRouter.GET("/:id", controller.GetArticle)
			articleRouter.POST("/:id", controller.UpdateArticle)
			articleRouter.DELETE("/:id", controller.DeleteArticle)
		}
	}
	return router
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	gin_mode := os.Getenv("GIN_MODE")
	gin.SetMode(gin_mode)
	// Disable log's color
	// gin.DisableConsoleColor()

	router := setRouter()

	port := os.Getenv("PORT")

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	// need size 1 to get a signal from buffer.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
