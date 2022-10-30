package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/fa9566509/POSBackEnd/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllOrderItems(ctx *fiber.Ctx) error {
	orderItems := models.GetAllOrdersItems()
	oI, err := json.Marshal(orderItems)
	if err != nil {
		fmt.Println("Error while encding json [Func: GetAllOrders()]")
		return err
	}
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write(oI)
	return nil
}

func CreateOrderItem(ctx *fiber.Ctx) error {
	orderItem := new(models.OrderItems)
	if err := ctx.BodyParser(orderItem); err != nil {
		return err
	}
	orderItem.CreateOrderItem()
	ctx.Response().Header.Add("Content-Type", "application/json")
	return nil
}

func GetOrderItemByID(ctx *fiber.Ctx) error {
	ID, _ := strconv.ParseInt(ctx.Params("ID"), 0, 0)
	orderItem := models.GetOrderbyID(ID)
	oI, err := json.Marshal(orderItem)
	if err != nil {
		return err
	}
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write(oI)
	return nil
}

func DeleteOrderItem(ctx *fiber.Ctx) error {
	ID, err := strconv.ParseInt(ctx.Params("ID"), 0, 0)
	if err != nil {
		fmt.Print("Error while parsing ID")
		return err
	}
	orderItem := models.DeleteOrder(ID)
	oI, err := json.Marshal(orderItem)
	if err != nil {
		fmt.Println("Error while encoding order into json")
		return err
	}
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write(oI)
	return nil
}

func UpdateOrderItem(ctx *fiber.Ctx) error {
	sid := ctx.Params("ID")
	id, err := strconv.ParseInt(sid, 0,0)
	if err != nil {
		fmt.Println("Error while parsing ID [Func: UpdateCustomer()]")
		return err
	}
	orderItem := models.Orders{}
	res := orderItem.UpdateOrder(id)
	resOrderItem, err := json.Marshal(res)
	if err != nil {
		fmt.Println("Error while encoding to json [Fucn: UpdateCustomer()]")
		return err
	}
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write(resOrderItem)
	return nil
}