package controllers

import (
	db "api_fiber/src/config"
	"api_fiber/src/models"
	"fmt"
	"strconv"

	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateCashier(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Invalid data",
			},
		)
	}

	if data["name"] == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Cashier name is required",
			},
		)
	}
	if data["passcode"] == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Passcode is required",
			},
		)
	}
	// Можно вот так объеденить условия в одно
	// if data["name"] == "" || data["passcode"] == "" {
	// 	return c.Status(400).JSON(fiber.Map{
	// 		"success": false,
	// 		"message": "Cashier Name is required",
	// 		"error":   map[string]interface{}{},
	// 	})
	// }
	cashier := models.Cashier{
		Name:      data["name"],
		Passcode:  data["passcode"],
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	db.DB.Create(&cashier)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Cashier created successfully",
		"data":    &cashier,
	})
}

func CashiersList(c *fiber.Ctx) error {
	var cashier []models.Cashier
	limit, _ := strconv.Atoi(c.Query("limit"))
	fmt.Println(limit)
	skip, _ := strconv.Atoi(c.Query("skip"))
	fmt.Println(skip)

	var count int64

	//вот тут порядок важен при запросе, если Count(который я не понял пока зачем нужен),
	// то будет ошибка, если Find в начало например поставить, то Limit and Offset тупо не сработают
	db.DB.Select("*").Limit(limit).Offset(skip).Find(&cashier).Count(&count)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Cashires list",
		"data":    cashier,
	})
}

func GetCashierDetails(c *fiber.Ctx) error {
	cashierId := c.Params("cashierId")
	var cashier models.Cashier

	db.DB.Select("id, name, created_at, updated_at").Where("id =?", cashierId).First(&cashier)

	cashierData := make(map[string]interface{})
	cashierData["cashierId"] = cashier.Id
	cashierData["name"] = cashier.Name
	cashierData["created_at"] = cashier.CreatedAt
	cashierData["update_at"] = cashier.UpdatedAt

	if cashier.Id == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Cashier Not Found",
			"error":   map[string]interface{}{},
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    cashierData,
	})
}

func UpdateCashier(c *fiber.Ctx) error {
	cashierId := c.Params("cashierId")
	var cashier models.Cashier

	err := c.BodyParser(&cashier)

	if err != nil {
		return err
	}

	db.DB.Find(&cashier, "id=?", cashierId)

	if cashier.Name == "" || cashier.Id == 0 {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Cashier Not Found or Cashier Name is required",
		})
	}
	var updateCashier models.Cashier

	err1 := c.BodyParser(&updateCashier)
	if err1 != nil {
		return err1
	}

	if updateCashier.Name == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Cashier Name is required",
		})
	}

	cashier.Name = updateCashier.Name
	db.DB.Save(&cashier)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Cashier updated successfully",
		"data":    &cashier,
	})
}

func DeleteCashier(c *fiber.Ctx) error {
	cashierId := c.Params("cashierId")
	var cashier models.Cashier

	db.DB.Where("id = ?", cashierId).First(&cashier)
	if cashier.Id == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Cashier Not Found",
		})
	}
	db.DB.Where("id = ?", cashierId).Delete(&cashier)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Cashier deleted successfully",
	})
}
