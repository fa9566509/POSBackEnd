package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/fa9566509/POSBackEnd/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllSuppliers(ctx *fiber.Ctx) error {
	suppliers := models.GetAllSuppliers()
	s, err := json.Marshal(suppliers)
	if err != nil {
		fmt.Println("Error while encding json [Func: GetAllOrders()]")
		return err
	}
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write(s)
	return nil
}

func CreateSupplier(ctx *fiber.Ctx) error {
	supplier := new(models.Suppliers)
	if err := ctx.BodyParser(supplier); err != nil {
		return err
	}
	supplier.CreateSupplier()
	ctx.Response().Header.Add("Content-Type", "application/json")
	return nil
}

func GetSupplierById(ctx *fiber.Ctx) error {
	ID, _ := strconv.ParseInt(ctx.Params("ID"), 0, 0)
	supplier := models.GetSupplierByID(ID)
	s, err := json.Marshal(supplier)
	if err != nil {
		return err
	}
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write(s)
	return nil
}

func DeleteSupplier(ctx *fiber.Ctx) error {
	ID, err := strconv.ParseInt(ctx.Params("ID"), 0, 0)
	if err != nil {
		fmt.Print("Error while parsing ID")
		return err
	}
	supplier := models.DeleteSupplier(ID)
	s, err := json.Marshal(supplier)
	if err != nil {
		fmt.Println("Error while encoding order into json")
		return err
	}
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write(s)
	return nil
}

func UpdateSupplier(ctx *fiber.Ctx) error {
	sid := ctx.Params("ID")
	id, err := strconv.ParseInt(sid, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing ID [Func: UpdateCustomer()]")
		return err
	}
	supplier := models.Suppliers{}
	res := supplier.UpdateSupplier(id)
	resSupplier, err := json.Marshal(res)
	if err != nil {
		fmt.Println("Error while encoding to json [Fucn: UpdateCustomer()]")
		return err
	}
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write(resSupplier)
	return nil
}
