package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/config"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/handlers"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/mapper"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/middleware"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/repositories"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/services"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB, cfg *config.Config) *gin.Engine {
	r := gin.New()

	// Global Middleware
	r.Use(middleware.LoggerMiddleware())
	r.Use(middleware.RecoveryMiddleware())
	r.Use(middleware.ContextMiddleware(30 * time.Second))

	// Repositories
	userRepo := repositories.NewUserRepository(db)
	storeRepo := repositories.NewStoreRepository(db)
	faqRepo := repositories.NewFaqRepository(db)
	faqCategoryRepo := repositories.NewFaqCategoryRepository(db)
	faqTranslationRepo := repositories.NewFaqTranslationRepository(db)

	// Mappers
	userMapper := &mapper.UserMapper{}
	storeMapper := &mapper.StoreMapper{}
	faqMapper := &mapper.FAQMapper{}
	faqCategoryMapper := &mapper.FAQCategoryMapper{}
	faqTranslationMapper := &mapper.FAQTranslationMapper{}

	// Services
	authService := services.NewAuthService(userRepo, storeRepo, cfg, userMapper, storeMapper)
	userService := services.NewUserService(userRepo, userMapper)
	storeService := services.NewStoreService(storeRepo, storeMapper)
	faqService := services.NewFAQService(faqRepo, faqMapper)
	faqCategoryService := services.NewFAQCategoryService(faqCategoryRepo, faqCategoryMapper)
	faqTranslationService := services.NewFAQTranslationService(faqTranslationRepo, faqTranslationMapper)

	// Handlers
	authHandler := handlers.NewAuthHandler(authService, cfg)
	userHandler := handlers.NewUserHandler(userService)
	storeHandler := handlers.NewStoreHandler(storeService)
	faqHandler := handlers.NewFAQHandler(faqService, storeService)
	faqCategoryHandler := handlers.NewFAQCategoryHandler(faqCategoryService)
	faqTranslationHandler := handlers.NewFAQTranslationHandler(faqTranslationService)

	// Routes
	api := r.Group("/api")
	{
		// Auth
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		// Public FAQ Routes
		api.GET("/faqs", faqHandler.GetAllFAQs)
		api.GET("/faqs/:id", faqHandler.GetFAQ)
		api.GET("/categories", faqCategoryHandler.GetAllCategories)
		api.GET("/categories/:id", faqCategoryHandler.GetCategory)
		api.GET("/translations/:id", faqTranslationHandler.GetTranslation)
		api.GET("/faqs/:id/translations", faqTranslationHandler.GetTranslationsByFAQID)

		// Protected Routes
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware(cfg))
		{
			// User
			protected.GET("/users/me", userHandler.GetMe)
			protected.DELETE("/users/:id", userHandler.DeleteUser) // Admin or Self

			// Merchant Stores
			stores := protected.Group("/stores")
			stores.Use(middleware.RoleMiddleware("merchant"))
			{
				stores.GET("/me", storeHandler.GetMyStore)
				stores.PUT("/me", storeHandler.UpdateStore)
			}

			// Merchant FAQs
			merchantFAQs := protected.Group("/merchant/faqs")
			merchantFAQs.Use(middleware.RoleMiddleware("merchant"))
			{
				merchantFAQs.POST("", faqHandler.CreateFAQ)
				merchantFAQs.PUT("/:id", faqHandler.UpdateFAQ)
				merchantFAQs.DELETE("/:id", faqHandler.DeleteFAQ)

				// Translations for Merchant FAQs
				merchantFAQs.POST("/translations", faqTranslationHandler.CreateTranslation)
				merchantFAQs.PUT("/translations/:id", faqTranslationHandler.UpdateTranslation)
				merchantFAQs.DELETE("/translations/:id", faqTranslationHandler.DeleteTranslation)
			}

			// Admin
			admin := protected.Group("/admin")
			admin.Use(middleware.RoleMiddleware("admin"))
			{
				// Categories
				admin.POST("/categories", faqCategoryHandler.CreateCategory)
				admin.PUT("/categories/:id", faqCategoryHandler.UpdateCategory)
				admin.DELETE("/categories/:id", faqCategoryHandler.DeleteCategory)

				// Global FAQs (Admin can manage all, but specifically global ones)
				admin.POST("/faqs", faqHandler.CreateFAQ) // Can create global
				admin.PUT("/faqs/:id", faqHandler.UpdateFAQ)
				admin.DELETE("/faqs/:id", faqHandler.DeleteFAQ)

				// Translations
				admin.POST("/translations", faqTranslationHandler.CreateTranslation)
				admin.PUT("/translations/:id", faqTranslationHandler.UpdateTranslation)
				admin.DELETE("/translations/:id", faqTranslationHandler.DeleteTranslation)
			}
		}
	}

	return r
}
