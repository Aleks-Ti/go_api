package controllers

import "github.com/gofiber/fiber/v2"

func CreateCashier(c *fiber.Ctx) error {
	return nil
}

func UpdateCashier(c *fiber.Ctx) error {
	return nil
}

func CashiersList(c *fiber.Ctx) error {
	return c.SendString("Cashiers List")
}

func GetCashierDetails(c *fiber.Ctx) error {
	return nil
}

func DeleteCashier(c *fiber.Ctx) error {
	return nil
}
