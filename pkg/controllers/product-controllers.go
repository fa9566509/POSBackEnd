package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/fa9566509/POSBackEnd/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllProducts(ctx *fiber.Ctx) error {
	products := models.GetAllProducts()
	p, err := json.Marshal(products)
	if err != nil {
		fmt.Println("Error while encding json [Func: GetAllOrders()]")
		return err
	}
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write(p)
	return nil
}

func CreateProduct(ctx *fiber.Ctx) error {
	product := new(models.Products)
	if err := ctx.BodyParser(product); err != nil {
		return err
	}
	product.CreateProduct()
	ctx.Response().Header.Add("Content-Type", "application/json")
	return nil
}

func GetProductById(ctx *fiber.Ctx) error {
	ID, _ := strconv.ParseInt(ctx.Params("ID"), 0, 0)
	product := models.GetProductByID(ID)
	p, err := json.Marshal(product)
	if err != nil {
		return err
	}
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write(p)
	return nil
}

func DeleteProduct(ctx *fiber.Ctx) error {
	ID, err := strconv.ParseInt(ctx.Params("ID"), 0, 0)
	if err != nil {
		fmt.Print("Error while parsing ID")
		return err
	}
	product := models.DeleteProduct(ID)
	p, err := json.Marshal(product)
	if err != nil {
		fmt.Println("Error while encoding order into json")
		return err
	}
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write(p)
	return nil
}

func UpdateProduct(ctx *fiber.Ctx) error {
	sid := ctx.Params("ID")
	id, err := strconv.ParseInt(sid, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing ID [Func: UpdateCustomer()]")
		return err
	}
	Product := models.Products{}
	res := Product.UpdateProduct(id)
	resProduct, err := json.Marshal(res)
	if err != nil {
		fmt.Println("Error while encoding to json [Fucn: UpdateCustomer()]")
		return err
	}
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write(resProduct)
	return nil
}
