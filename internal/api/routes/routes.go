package routes

import (
	"github.com/Alperen10/Cms/internal/api/handlers"
	"github.com/Alperen10/Cms/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all the routes for our application
func SetupRoutes(router *gin.Engine) {
	// API v1 group
	v1 := router.Group("/api/v1")

	// Auth routes
	auth := v1.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
	}

	// Protected routes
	protected := v1.Group("")
	protected.Use(middleware.AuthMiddleware())
	{
		// Users routes (to be implemented)
		users := protected.Group("/users")
		{
			users.GET("", middleware.RoleMiddleware("admin"))
			users.GET("/:id")
			users.PUT("/:id")
			users.DELETE("/:id", middleware.RoleMiddleware("admin"))
		}

		// Posts routes
		posts := protected.Group("/posts")
		{
			posts.GET("", handlers.GetPosts)
			posts.POST("", middleware.RoleMiddleware("admin", "editor"), handlers.CreatePost)
			posts.GET("/:id", handlers.GetPost)
			posts.PUT("/:id", middleware.RoleMiddleware("admin", "editor"), handlers.UpdatePost)
			posts.DELETE("/:id", middleware.RoleMiddleware("admin"), handlers.DeletePost)
		}

	}
}
