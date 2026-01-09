package scripts

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/config"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/dto"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/mapper"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/models"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/repositories"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestLogic() {
	cfg := config.LoadConfig()
	dbUser := "postgres"
	dbPort := "5432"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, dbUser, cfg.DBPassword, cfg.DBName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(&models.User{}, &models.Store{}, &models.FAQCategory{}, &models.FAQ{}, &models.FAQTranslation{})

	userRepo := repositories.NewUserRepository(db)
	storeRepo := repositories.NewStoreRepository(db)
	faqRepo := repositories.NewFaqRepository(db)
	faqTransRepo := repositories.NewFaqTranslationRepository(db)
	faqCategoryRepo := repositories.NewFaqCategoryRepository(db)

	userMapper := &mapper.UserMapper{}
	storeMapper := &mapper.StoreMapper{}
	faqMapper := &mapper.FAQMapper{}
	faqTransMapper := &mapper.FAQTranslationMapper{}

	authService := services.NewAuthService(userRepo, storeRepo, cfg, userMapper, storeMapper)
	faqService := services.NewFAQService(faqRepo, faqMapper)
	faqTransService := services.NewFAQTranslationService(faqTransRepo, faqTransMapper)

	ctx := context.Background()

	merchantEmail := fmt.Sprintf("merchant_%d@test.com", time.Now().Unix())
	merchantReq := dto.SignupRequest{
		Email:    merchantEmail,
		Password: "password",
		Role:     "merchant",
	}
	merchantRes, err := authService.Register(ctx, merchantReq)
	if err != nil {
		log.Fatalf("Failed to register merchant: %v", err)
	}
	fmt.Printf("Registered Merchant: %s\n", merchantRes.ID)

	// Verify Store Created
	fmt.Println("find store by merchant id ...")
	store, err := storeRepo.FindStoreByMerchantID(ctx, merchantRes.ID)
	if err != nil {
		log.Fatalf("Store NOT created for merchant: %v", err)
	}
	fmt.Printf("Store created for merchant: %s\n", store.ID)

	userEmail := fmt.Sprintf("user_%d@test.com", time.Now().Unix())
	userReq := dto.SignupRequest{
		Email:    userEmail,
		Password: "password",
		Role:     "customer",
	}
	userRes, err := authService.Register(ctx, userReq)
	if err != nil {
		log.Fatalf("Failed to register user: %v", err)
	}
	fmt.Printf("Registered User: %s\n", userRes.ID)

	_, err = storeRepo.FindStoreByMerchantID(ctx, userRes.ID)
	if err == nil {
		log.Fatalf("Store SHOULD NOT be created for user, but was found")
	}
	fmt.Println("Correctly confirmed no store for regular user.")

	fmt.Println("create category ...")
	category := models.FAQCategory{ID: uuid.New(), Name: "General"}
	if err := faqCategoryRepo.CreateFaqCategory(ctx, category); err != nil {
		log.Fatal(err)
	}
	fmt.Println("create faq ...")
	faqID := uuid.New()
	faq := models.FAQ{
		ID:         faqID,
		CategoryID: category.ID,
		StoreID:    &store.ID,
		IsGlobal:   false,
		CreatedBy:  merchantRes.ID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	if err := faqRepo.CreateFaq(ctx, faq); err != nil {
		log.Fatalf("Failed to create FAQ: %v", err)
	}

	trans1 := models.FAQTranslation{ID: uuid.New(), FAQID: faqID, LanguageCode: "en", Question: "EN Q", Answer: "EN A"}
	trans2 := models.FAQTranslation{ID: uuid.New(), FAQID: faqID, LanguageCode: "ar", Question: "AR Q", Answer: "AR A"}
	if err := faqTransRepo.CreateFaqTranslation(ctx, trans1); err != nil {
		log.Fatal(err)
	}
	if err := faqTransRepo.CreateFaqTranslation(ctx, trans2); err != nil {
		log.Fatal(err)
	}
	translations, err := faqTransService.GetTranslationsByFAQID(ctx, faqID.String())
	if err != nil {
		log.Fatalf("Failed to get translations: %v", err)
	}
	if len(translations) != 2 {
		log.Fatalf("Expected 2 translations, got %d", len(translations))
	}
	fmt.Printf("Successfully retrieved %d translations for FAQ.\n", len(translations))

	globalFAQ := models.FAQ{
		ID:         uuid.New(),
		CategoryID: category.ID,
		IsGlobal:   true,
		CreatedBy:  merchantRes.ID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	if err := faqRepo.CreateFaq(ctx, globalFAQ); err != nil {
		log.Fatal(err)
	}

	faqs, err := faqService.GetAllFAQs(ctx, store.ID.String())
	if err != nil {
		log.Fatalf("Failed to get FAQs: %v", err)
	}

	foundGlobal := false
	foundStore := false
	for _, f := range faqs {
		if f.ID == globalFAQ.ID {
			foundGlobal = true
		}
		if f.ID == faq.ID {
			foundStore = true
		}
	}

	if !foundGlobal {
		log.Fatal("Global FAQ not found in store view")
	}
	if !foundStore {
		log.Fatal("Store FAQ not found in store view")
	}
	fmt.Println("Successfully filtered FAQs (Global + Store).")
	fmt.Println("ALL TESTS PASSED!")
}
