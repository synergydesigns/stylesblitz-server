package models

import (
	"time"
	"log"
	"fmt"
	"strconv"

	"github.com/lucsky/cuid"
	"github.com/jinzhu/gorm"
	"github.com/synergydesigns/stylesblitz-server/shared/utils"
)

type Cart struct {
	ID        string    `gorm:"primary_key"`
	UserID    string    `json:"user_id"`
	Type      string    `json:"type"`
	TypeID    string    `json:"type_id"`
	VendorID  string    `json:"vendor_id"`
	Quantity  int       `gorm:"default:1"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CartType struct {
	Product
	Service
}

type CartDBService struct {
	DB *gorm.DB
}

func (cart *Cart) BeforeCreate(scope *gorm.Scope) error {
	if cart.ID == "" {
		scope.SetColumn("ID", cuid.New())
	}

	return nil
}

type CartDB interface {
	CreateCart(userID string, vendorID string, cartType string, typeID string, quantity int) (*Cart, error)
	GetAllCarts(userID string) ([]*Cart, error)
	UpdateCart(userID string, cartID string, quantity int, cartType string, typeID string) (*Cart, error)
	DeleteCart(userID string, cartID string) (bool, error)
}

func (service *CartDBService) CreateCart(userID string, vendorID string, cartType string, typeID string, quantity int) (*Cart, error) {
	cart := Cart{
		UserID: userID,
		VendorID: vendorID,
		Type: cartType,
		TypeID: typeID,
		Quantity: quantity,
	}

	if (cartType == "product") {
		var foundQuantity int
		foundProduct := Product{}
		result := service.DB.Where("id = ?", cart.TypeID).Limit(1).First(&foundProduct)
		if result.Error != nil {
			log.Printf("An error occurred creating cart.\n %v", result.Error.Error())

			return nil, fmt.Errorf("an error occurred creating cart %s", result.Error.Error())
		}
		foundQuantity = foundProduct.Available

		if (quantity > foundQuantity) {
			return nil, fmt.Errorf("Quantity is more than the available product. %v products is/are available", foundQuantity)
		}
	}

	cartResult := service.DB.Create(&cart)

	if cartResult.Error != nil {
		log.Printf("An error occurred creating cart %v", cartResult.Error.Error())

		if utils.ForeignKeyNotExist(cartResult.Error) {
			return nil, fmt.Errorf("Vendor with id %s does not exist", vendorID)
		}

		return nil, fmt.Errorf("an error occurred creating cart %s", cartResult.Error.Error())
	}

	return &cart, nil
}

func (service *CartDBService) GetAllCarts(userID string) ([]*Cart, error) {
	var carts []*Cart

	result := service.DB.Where("user_id = ?", userID).Limit(20).Find(&carts)
	if result.Error != nil {
		log.Printf("An error occurred getting all cart %v", result.Error.Error())
		return carts, fmt.Errorf("An error occurred getting all carts %s", result.Error.Error())
	}

	return carts, nil
}

func (cart *Cart) CartType() (CartType, error) {
	foundProduct := &Product{}
	foundService := &Service{}
	if cart.Type == "product" {
		database.Where("id = ?", cart.TypeID).Limit(1).First(&foundProduct)
		ct := CartType{}
		ct.Product = *foundProduct
		return ct, nil
	} else {
		val, err := strconv.Atoi(cart.TypeID)
		if err != nil {
			return CartType{}, err
		}
		database.Where("id = ?", val).Limit(1).First(&foundService)

		ct := CartType{}
		ct.Service = *foundService

		return ct, nil
	}
}

func (service *CartDBService) DeleteCart(userID string, cartID string)(bool, error) {
	cartResult := service.DB.Delete(&Cart{}, "id = ? AND user_id = ?", cartID, userID)
	if cartResult.Error != nil {
		log.Printf("An error occurred deleting cart %v", cartResult.Error.Error())
		return false, fmt.Errorf("An error occurred deleting cart. %s", cartResult.Error.Error())
	}
	if (cartResult.RowsAffected < 1) {
		return false, fmt.Errorf("An error occurred deleting cart")
	}

	return true, nil
}

func (service *CartDBService) UpdateCart(userID string, cartID string, quantity int, cartType string, typeID string) (*Cart, error) {
	cart := Cart{}
	value := make(map[string]interface{})

	if (cartType == "product") {
		var foundQuantity int
		foundProduct := Product{}
		result := service.DB.Where("id = ?", typeID).Limit(1).First(&foundProduct)
		if result.Error != nil {
			log.Printf("An error occurred updating cart %v", result.Error.Error())

			return nil, fmt.Errorf("an error occurred updating cart %s", result.Error.Error())
		}
		foundQuantity = foundProduct.Available

		if (quantity > foundQuantity) {
			return nil, fmt.Errorf("Quantity is more than the available product. %v products is/are available", foundQuantity)
		}
	}

	value["quantity"] = quantity
	result := service.DB.Model(&cart).Where("id = ?", cartID).Updates(value)

	if result.Error != nil {
		log.Printf("An error occurred updating cart %v", result.Error.Error())
		return &cart, fmt.Errorf("an error occurred updating cart %s", result.Error.Error())
	}

	result.First(&cart, "id = ?", cartID)

	return &cart, nil
}
