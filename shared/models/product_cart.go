package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

type ProductCart struct {
	ID uint64 `gorm:"primary_key"`
	VendorID string
	ProductID string
	ShopID string
	CartID string
	Quantity int `gorm:"default:1"`
}

type ProductCartDBService struct {
	DB *gorm.DB
}

type ProductCartDB interface {
	CreateProductCart(userID string, vendorID string, productID string, quantity int) (*ProductCart, error)
	UpdateProductCart(userID string, cartID string, quantity int) (*ProductCart, error)
	DeleteProductCart(userID string, cartID string) (bool, error)
	GetProductsCart(userID string) ([]*ProductCart, error)
}

func (service *ProductCartDBService) CreateProductCart(userID string, vendorID string, productID string, quantity int) (*ProductCart, error) {
	cart := Cart{
		UserID: userID,
	}

	cartResult := service.DB.Create(&cart)

	if cartResult.Error != nil {
		log.Printf("An error occurred creating cart %v", cartResult.Error.Error())

		return nil, fmt.Errorf("an error occurred creating cart %s", cartResult.Error.Error())
	}

	productCart := ProductCart{
		VendorID:    vendorID,
		ProductID:   productID,
		CartID:      cart.ID,
		Quantity:    quantity,
	}

	result := service.DB.Create(&productCart)

	if result.Error != nil {
		log.Printf("An error occurred creating product cart %v", result.Error.Error())
		return &productCart, fmt.Errorf("an error occurred creating product cart %s", result.Error.Error())
	}

	return &productCart, nil
}

func (service *ProductCartDBService) UpdateProductCart(userID string, cartID string, quantity int) (*ProductCart, error) {
	productCart := ProductCart{}
	value := make(map[string]interface{})

	value["quantity"] = quantity

	result := service.DB.Model(&productCart).Where("cart_id = ?", cartID).Updates(value)

	if result.Error != nil {
		log.Printf("An error occurred updating product cart %v", result.Error.Error())
		return &productCart, fmt.Errorf("an error occurred updating product cart %s", result.Error.Error())
	}

	result.First(&productCart, "cart_id = ?", cartID)

	return &productCart, nil
}

func (service *ProductCartDBService) DeleteProductCart(userID string, cartID string)(bool, error) {
	tx := service.DB.Begin()
	defer func() {
    if r := recover(); r != nil {
      tx.Rollback()
    }
	}()
	if err := tx.Error; err != nil {
    return false, err
  }

	result := tx.Delete(&ProductCart{}, "cart_id = ?", cartID)

	if result.Error != nil {
		tx.Rollback()
		log.Printf("An error occurred deleting product cart %v", result.Error.Error())
		return false, fmt.Errorf("An error occurred deleting product cart %s", result.Error.Error())
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

func (service *ProductCartDBService) GetProductsCart(userID string) ([]*ProductCart, error) {
	var productsCart []*ProductCart
	var carts []*Cart

	service.DB.Where("user_id = ?", userID).Limit(20).Find(&carts)
	var cartIDsList []string

	for _, v := range carts {
		cartIDsList = append(cartIDsList, v.ID)
	}

	result := service.DB.Where("cart_id IN (?)", cartIDsList).Limit(20).Find(&productsCart)

	if result.Error != nil {
		log.Printf("An error occurred getting all products cart %v", result.Error.Error())
		return productsCart, fmt.Errorf("An error occurred getting products cart %s", result.Error.Error())
	}

	return productsCart, nil
}

func (pc *ProductCart) Product() (*Product) {
	foundProduct := &Product{}
	database.Where("id = ?", pc.ProductID).Limit(1).First(&foundProduct)
	
	return foundProduct
}
