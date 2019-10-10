package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

type ServiceCart struct {
	ID        uint64 `gorm:"primary_key"`
	VendorID  string `json:"vendor_id"`
	ServiceID int `json:"service_id"`
	CartID    string `json:"cart_id"`
	Quantity  int `gorm:"default:1"`
}

type ServiceCartDB interface {
	CreateServiceCart(userID string, vendorID string, serviceID int, quantity int) (*ServiceCart, error)
	UpdateServiceCart(userID string, cartID string, quantity int) (*ServiceCart, error)
	DeleteServiceCart(userID string, cartID string) (bool, error)
	GetServicesCart(userID string) ([]*ServiceCart, error)
}

type ServiceCartDBService struct {
	DB *gorm.DB
}

func (service *ServiceCartDBService) CreateServiceCart(userID string, vendorID string, serviceID int, quantity int) (*ServiceCart, error) {
	cart := Cart{
		UserID: userID,
	}

	cartResult := service.DB.Create(&cart)

	if cartResult.Error != nil {
		log.Printf("An error occurred creating cart %v", cartResult.Error.Error())

		return nil, fmt.Errorf("an error occurred creating cart %s", cartResult.Error.Error())
	}

	serviceCart := ServiceCart{
		VendorID:    vendorID,
		ServiceID:   serviceID,
		CartID:      cart.ID,
		Quantity:    quantity,
	}

	result := service.DB.Create(&serviceCart)

	if result.Error != nil {
		log.Printf("An error occurred creating service cart %v", result.Error.Error())
		return &serviceCart, fmt.Errorf("an error occurred creating service cart %s", result.Error.Error())
	}

	return &serviceCart, nil
}

func (service *ServiceCartDBService) UpdateServiceCart(userID string, cartID string, quantity int) (*ServiceCart, error) {
	serviceCart := ServiceCart{}
	value := make(map[string]interface{})

	value["quantity"] = quantity
	result := service.DB.Model(&serviceCart).Where("cart_id = ?", cartID).Updates(value)

	if result.Error != nil {
		log.Printf("An error occurred updating service cart %v", result.Error.Error())
		return &serviceCart, fmt.Errorf("an error occurred updating service cart %s", result.Error.Error())
	}

	result.First(&serviceCart, "cart_id = ?", cartID)

	return &serviceCart, nil
}

func (service *ServiceCartDBService) DeleteServiceCart(userID string, cartID string, )(bool, error) {
	tx := service.DB.Begin()
	defer func() {
    if r := recover(); r != nil {
      tx.Rollback()
    }
	}()
	if err := tx.Error; err != nil {
    return false, err
  }

	result := tx.Delete(&ServiceCart{}, "cart_id = ?", cartID)

	if result.Error != nil {
		tx.Rollback()

		log.Printf("An error occurred deleting service cart %v", result.Error.Error())
		return false, fmt.Errorf("An error occurred deleting service cart %s", result.Error.Error())
	}


	cartResult := tx.Delete(&Cart{}, "id = ?", cartID)
	if cartResult.Error != nil {
		tx.Rollback()

		log.Printf("An error occurred deleting cart %v", cartResult.Error.Error())
		return false, fmt.Errorf("An error occurred deleting cart %s", cartResult.Error.Error())
	}

	tx.Commit()

	return true, nil
}

func (service *ServiceCartDBService) GetServicesCart(userID string) ([]*ServiceCart, error) {
	var servicesCart []*ServiceCart
	var carts []*Cart

	service.DB.Where("user_id = ?", userID).Limit(20).Find(&carts)
	var cartIDsList []string

	for _, v := range carts {
		cartIDsList = append(cartIDsList, v.ID)
	}

	result := service.DB.Where("cart_id IN (?)", cartIDsList).Limit(20).Find(&servicesCart)

	if result.Error != nil {
		log.Printf("An error occurred getting all services cart %v", result.Error.Error())
		return nil, fmt.Errorf("An error occurred getting services cart %s", result.Error.Error())
	}

	return servicesCart, nil
}

func (sc *ServiceCart) Service() (*Service) {
	foundService := &Service{}
	database.Where("id = ?", uint64(sc.ServiceID)).Limit(1).First(&foundService)
	
	return foundService
}