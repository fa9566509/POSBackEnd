package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/fa9566509/POSBackEnd/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func GetCustomers(ctx *fiber.Ctx) error {
	customers := models.GetAllCustomers()
	cs, err := json.Marshal(customers)
	if err != nil {
		fmt.Println("Error while encoding to json [Func: GetAllCustomers()]")
		return err
	}
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write(cs)
	return nil
}

func GetCustomerByID(ctx *fiber.Ctx) error {
	sID := ctx.Params("ID")
	id, err := strconv.ParseInt(sID,0,0)
	if err != nil {
		fmt.Println("Error while parsing id [Func: GetCustomerByID()]")
		return err
	}
	customer := models.GetCustomerByID(id)
	c, err := json.Marshal(customer)
	if err != nil {
		fmt.Println("Error while encoding to json [Func: GetCustomerByID()]")
		return err
	}
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write(c)
	return nil
}

func CreateCustomer(ctx *fiber.Ctx) error {
	customer := models.Customers{}
	if err := ctx.BodyParser(customer); err != nil {
		fmt.Println("Error while parsing request body [Func: CreateCustomer()]")
		return err
	}
	res := customer.CreateCustomer()
	resCustomer, err := json.Marshal(res)
	if err != nil {
		fmt.Println("Error while encoding to json [Func: CreateCustomer()]")
		return err
	}
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write(resCustomer)
	return nil
}

func DeleteCustomer(ctx *fiber.Ctx) error {
	sid := ctx.Params("ID")
	id, err := strconv.ParseInt(sid, 0,0)
	if err != nil {
		fmt.Println("Error while parsing ID [Func: DeleteCustomer()]")
		return err
	}
	res := models.DeleteCustomer(id)
	resCustomer, err := json.Marshal(res)
	if err != nil {
		fmt.Println("Error while encoding to json [Func: DeleteCustomer()]")
		return err
	}
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write(resCustomer)
	return nil
}


func UpdateCustomer(ctx *fiber.Ctx) error {
	sid := ctx.Params("ID")
	id, err := strconv.ParseInt(sid, 0,0)
	if err != nil {
		fmt.Println("Error while parsing ID [Func: UpdateCustomer()]")
		return err
	}
	customer := models.Customers{}
	res := customer.UpdateCustomer(id)
	resCustomer, err := json.Marshal(res)
	if err != nil {
		fmt.Println("Error while encoding to json [Fucn: UpdateCustomer()]")
		return err
	}
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write(resCustomer)
	return nil
}