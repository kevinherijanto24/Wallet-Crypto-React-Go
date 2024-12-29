package routes

import (
	"github.com/gofiber/fiber/v2"
	"Wallet-Crypto-Crud/backend/database"
	"Wallet-Crypto-Crud/backend/models"
)

func RegisterWalletRoutes(app *fiber.App) {
	app.Post("/wallets", createWallet)
	app.Get("/wallets", getWallets)
	app.Get("/wallets/username/:username", getWalletsByUsername)
	app.Put("/wallets/:id", updateWallet) // Changed to use ID instead of username
	app.Delete("/wallets/:id", deleteWallet)
}

func createWallet(c *fiber.Ctx) error {
	wallet := new(models.Wallet)
	if err := c.BodyParser(wallet); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	database.DB.Create(&wallet)
	return c.Status(201).JSON(wallet)
}

func getWallets(c *fiber.Ctx) error {
	var wallets []models.Wallet
	database.DB.Find(&wallets)
	return c.JSON(wallets)
}

func getWalletsByUsername(c *fiber.Ctx) error {
	username := c.Params("username")
	var wallets []models.Wallet

	// Fetch wallets by username
	if err := database.DB.Where("username = ?", username).Find(&wallets).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "No wallets found for this username"})
	}

	return c.JSON(wallets)
}

func updateWallet(c *fiber.Ctx) error {
	id := c.Params("id") // Use the ID parameter for identifying the wallet
	var wallet models.Wallet

	// Check if the wallet exists
	if err := database.DB.First(&wallet, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Wallet not found"})
	}

	// Parse the request body
	var updatedData models.Wallet
	if err := c.BodyParser(&updatedData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Update wallet fields
	wallet.Address = updatedData.Address
	wallet.Balance = updatedData.Balance
	wallet.Currency = updatedData.Currency

	// Save the updated wallet
	if err := database.DB.Save(&wallet).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update wallet"})
	}

	return c.JSON(wallet)
}

func deleteWallet(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := database.DB.Delete(&models.Wallet{}, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Wallet not found"})
	}
	return c.SendStatus(204)
}
