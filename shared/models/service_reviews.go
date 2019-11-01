package models

import (
	"time"
	"log"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/synergydesigns/stylesblitz-server/shared/utils"
)

type ServiceReview struct {
	ID        uint64          `gorm:"primary_key"`
	UserID    string          `json:"user_id"`
	VendorID  string          `json:"vendor_id"`
	ServiceID int             `json:"service_id"`
	Text      string
	Rating    int
	Replies   []ServiceReview `gorm:"foreignkey:ParentID;association_foreignkey:ID"`
	ParentID  int 		        `json:"parent_id"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt *time.Time      `json:"deleted_at"`
}

type ServiceReviewWithAverageRating struct {
	Reviews         []*ServiceReview
	AverageRatings  float64
}

type ServiceReviewDBService struct {
	DB *gorm.DB
}

type ServiceReviewDB interface {
	CreateReview(userID string, vendorID string, serviceID int, text string, rating int, parentID int) (*ServiceReview, error)
	GetReviews(serviceID int) ([]*ServiceReview, error)
	UpdateReview(userID string, text string, rating int, id int) (*ServiceReview, error)
}

func (service *ServiceReviewDBService) CreateReview (userID string, vendorID string, serviceID int, text string, rating int, parentID int) (*ServiceReview, error) {
	review := ServiceReview{
		UserID: userID,
		VendorID: vendorID,
		ServiceID: serviceID,
		Text: text,
		Rating: rating,
		ParentID: parentID,
	}

	// For creating replies. Might need to be moved to a seperate method
	if (parentID != 0) {
		var foundReview ServiceReview
		service.DB.Where("service_id = ? AND id = ?", serviceID, parentID).First(&foundReview)

		if (foundReview.ServiceID == 0) {
			return nil, fmt.Errorf("an error occurred creating review")
		}

		if (foundReview.ParentID != 0) {
			return nil, fmt.Errorf("An error occurred. You cannot reply a reply :)")
		}

		review.Rating = 0
		
	}

	result := service.DB.Create(&review)

	if result.Error != nil {
		log.Printf("An error occurred creating review %v", result.Error.Error())

		if utils.ForeignKeyNotExist(result.Error) {
			if utils.HasValue(result.Error.Error(), "vendor") {
				return &review, fmt.Errorf("vendor with id %s does not exit", vendorID)
			}

		}

		return &review, fmt.Errorf("an error occurred creating review %s", result.Error.Error())
	}

	return &review, nil
}

func (service *ServiceReviewDBService) GetReviews(serviceID int) ([]*ServiceReview, error) {
	var reviews []*ServiceReview
	result := service.DB.Where("service_id = ? AND parent_id = ?", serviceID, 0).Order("created_at desc").Preload("Replies").Limit(50).Find(&reviews)

	if result.Error != nil {
		log.Printf("An error occurred getting reviews %v", result.Error.Error())
		return reviews, fmt.Errorf("An error occurred getting reviews %s", result.Error.Error())
	}

	return reviews, nil
}

func (service *ServiceReviewDBService) UpdateReview (userID string, text string, rating int, id int) (*ServiceReview, error) {
	review := ServiceReview{}
	value := make(map[string]interface{})

	value["rating"] = rating
	value["text"] = text
	result := service.DB.Model(&review).Where("id = ? AND user_id = ?", id, userID).Updates(value)

	if result.Error != nil {
		log.Printf("An error occurred updating review %v", result.Error.Error())

		return &review, fmt.Errorf("an error occurred updating review %s", result.Error.Error())
	}

	result.First(&review, "id = ?", id)

	return &review, nil
}
