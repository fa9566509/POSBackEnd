package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

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

func CreateOrder(ctx *fiber.Ctx) error {
	order := new(models.Orders)
	if err := ctx.BodyParser(order); err != nil {
		return err
	}
	order.CreateOrder()
	ctx.Response().Header.Add("Content-Type", "application/json")
	return nil
}

func GetOrderByID(ctx *fiber.Ctx) error {
	ID, _ := strconv.ParseInt(ctx.Params("ID"), 0, 0)
	order := models.GetOrderbyID(ID)
	o, err := json.Marshal(order)
	if err != nil {
		return err
	}
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write(o)
	return nil
}

func DeleteOrder(ctx *fiber.Ctx) error {
	ID, err := strconv.ParseInt(ctx.Params("ID"), 0, 0)
	if err != nil {
		fmt.Print("Error while parsing ID")
		return err
	}
	order := models.DeleteOrder(ID)
	o, err := json.Marshal(order)
	if err != nil {
		fmt.Println("Error while encoding order into json")
		return err
	}
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write(o)
	return nil
}
