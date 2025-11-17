package routes

import (
	handlers "VideoStreamingBackend/Handlers"
	mi "VideoStreamingBackend/Middleware"

	scalargo "github.com/bdpiprava/scalar-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/api.yaml", func(c *gin.Context) {
		c.File("./ApiDoc/api.yaml")
	})

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"set-cookie"},
	}))

	r.GET("/docs", func(c *gin.Context) {
		html, err := scalargo.NewV2(
			scalargo.WithSpecURL("/api.yaml"),
		)
		if err != nil {
			c.String(500, "Error generating Swagger UI: %v", err)
			return
		}
		c.Data(200, "text/html; charset=utf-8", []byte(html))
	})

	api := r.Group("/api")
	api.Static("images", "./Uploads/Images")

	api.POST("/users", handlers.CreateUserHandler)
	api.POST("/login", handlers.LoginHandler)

	protected := api.Group("")
	protected.Use(mi.AuthMiddleware())
	{
		protected.PUT("/users/:userId", handlers.UpdateUserHandler)
		protected.PATCH("/users/:userId", handlers.UpdatePasswordHandler)
		protected.DELETE("/users/:userId", handlers.DeleteUserHandler)
		protected.GET("/users/:userId", handlers.GetUserHandler)
		protected.POST("/logout", handlers.LogoutHandler)
		protected.GET("/videos/stream/:videoName", handlers.VideoStreamHandler)
		protected.POST("/videos", handlers.UploadVideoHandler)
		protected.GET("/videos/:videoId", handlers.GetVideoHandler)
		protected.GET("/videos/random30", handlers.Get30RandomVideosHandler)
		protected.POST("/videos/:videoId/like", handlers.LikeVideoHandler)
		protected.POST("/videos/:videoId/dislike", handlers.DislikedVideosHandler)
		protected.GET("/channels/:channelId", handlers.GetChannelHandler)
	}
}
