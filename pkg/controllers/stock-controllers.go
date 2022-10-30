package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/fa9566509/POSBackEnd/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllStocks(ctx *fiber.Ctx) error {
	stocks := models.GetAllStocks()
	s, err := json.Marshal(stocks)
	if err != nil {
		fmt.Println("Error while encding json [Func: GetAllOrders()]")
		return err
	}
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write(s)
	return nil
}

func CreateStock(ctx *fiber.Ctx) error {
	stock := new(models.Stock)
	if err := ctx.BodyParser(stock); err != nil {
		return err
	}
	stock.CreateSock()
	ctx.Response().Header.Add("Content-Type", "application/json")
	return nil
}

func GetStockById(ctx *fiber.Ctx) error {
	ID, _ := strconv.ParseInt(ctx.Params("ID"), 0, 0)
	stock := models.GetStockByID(ID)
	s, err := json.Marshal(stock)
	if err != nil {
		return err
	}
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write(s)
	return nil
}

func DeleteStock(ctx *fiber.Ctx) error {
	ID, err := strconv.ParseInt(ctx.Params("ID"), 0, 0)
	if err != nil {
		fmt.Print("Error while parsing ID")
		return err
	}
	stock := models.DeleteStock(ID)
	s, err := json.Marshal(stock)
	if err != nil {
		fmt.Println("Error while encoding order into json")
		return err
	}
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write(s)
	return nil
}

func UpdateStock(ctx *fiber.Ctx) error {
	sid := ctx.Params("ID")
	id, err := strconv.ParseInt(sid, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing ID [Func: UpdateCustomer()]")
		return err
	}
	stock := models.Stock{}
	res := stock.UpdateStock(id)
	resStock, err := json.Marshal(res)
	if err != nil {
		fmt.Println("Error while encoding to json [Fucn: UpdateCustomer()]")
		return err
	}
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write(resStock)
	return nil
}
