package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/fa9566509/POSBackEnd/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllPurchases(ctx *fiber.Ctx) error {
	purchases := models.GetAllProducts()
	p, err := json.Marshal(purchases)
	if err != nil {
		fmt.Println("Error while encding json [Func: GetAllOrders()]")
		return err
	}
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write(p)
	return nil
}

func CreatePurchase(ctx *fiber.Ctx) error {
	purchase := new(models.Purchases)
	if err := ctx.BodyParser(purchase); err != nil {
		return err
	}
	purchase.CreatePurchase()
	ctx.Response().Header.Add("Content-Type", "application/json")
	return nil
}

func GetPurchaseById(ctx *fiber.Ctx) error {
	ID, _ := strconv.ParseInt(ctx.Params("ID"), 0, 0)
	purchase := models.GetPurchaseByID(ID)
	p, err := json.Marshal(purchase)
	if err != nil {
		return err
	}
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write(p)
	return nil
}

func DeletePurchase(ctx *fiber.Ctx) error {
	ID, err := strconv.ParseInt(ctx.Params("ID"), 0, 0)
	if err != nil {
		fmt.Print("Error while parsing ID")
		return err
	}
	purchase := models.DeletePurchase(ID)
	p, err := json.Marshal(purchase)
	if err != nil {
		fmt.Println("Error while encoding order into json")
		return err
	}
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write(p)
	return nil
}

func UpdatePurchase(ctx *fiber.Ctx) error {
	sid := ctx.Params("ID")
	id, err := strconv.ParseInt(sid, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing ID [Func: UpdateCustomer()]")
		return err
	}
	purchase := models.Purchases{}
	res := purchase.UpdatePurchase(id)
	resPurchase, err := json.Marshal(res)
	if err != nil {
		fmt.Println("Error while encoding to json [Fucn: UpdateCustomer()]")
		return err
	}
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write(resPurchase)
	return nil
}
