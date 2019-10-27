package models

import (
	"time"
	"log"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/synergydesigns/stylesblitz-server/shared/utils"
)

type ServiceReview struct {
	ID        uint64    `gorm:"primary_key"`
	UserID    string    `json:"user_id"`
	VendorID  string    `json:"vendor_id"`
	ServiceID int       `json:"service_id"`
	Text      string
	Rating    string
	ParentID  int 		  `json:"parent_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ServiceReviewDBService struct {
	DB *gorm.DB
}

type ServiceReviewDB interface {
	CreateReview(userID string, vendorID string, serviceID int, text string, rating string, parentID int) (*ServiceReview, error)
	GetReviews(serviceID int) ([]*ServiceReview, error)
	// GetAllCarts(userID string) ([]*Cart, error)
	// UpdateCart(userID string, cartID string, quantity int, cartType string, typeID string) (*Cart, error)
	// DeleteCart(userID string, cartID string) (bool, error)
}

func (service *ServiceReviewDBService) CreateReview (userID string, vendorID string, serviceID int, text string, rating string, parentID int) (*ServiceReview, error) {
	review := ServiceReview{
		UserID: userID,
		VendorID: vendorID,
		ServiceID: serviceID,
		Text: text,
		Rating: rating,
		ParentID: parentID,
	}

	result := service.DB.Create(&review)

	if result.Error != nil {
		log.Printf("An error occurred creating review %v", result.Error.Error())
		// if utils.HasRecord(result.Error) {
		// 	return &review, fmt.Errorf("service with name %s already exit", serviceInput.Name)
		// }

		if utils.ForeignKeyNotExist(result.Error) {
			if utils.HasValue(result.Error.Error(), "vendor") {
				return &review, fmt.Errorf("vendor with id %s does not exit", vendorID)
			}

			// return &review, fmt.Errorf("category with id %d does not exit", serviceInput.CategoryID)
		}

		return &review, fmt.Errorf("an error occurred creating review %s", result.Error.Error())
	}

	return &review, nil
}

func (service *ServiceReviewDBService) GetReviews(serviceID int) ([]*ServiceReview, error) {
	var reviews []*ServiceReview

	result := service.DB.Where("service_id = ?", serviceID).Limit(50).Find(&reviews)
	if result.Error != nil {
		log.Printf("An error occurred getting reviews %v", result.Error.Error())
		return reviews, fmt.Errorf("An error occurred getting reviews %s", result.Error.Error())
	}

	return reviews, nil
}