package controllers

import (
	"encoding/json"
	"github.com/fa9566509/POSBackEnd/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllOrders(ctx *fiber.Ctx) error {
	orders := models.GetAllOrders()
	o, err := json.Marshal(orders)
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write(o)
	return err
}