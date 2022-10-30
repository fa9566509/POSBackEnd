package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/fa9566509/POSBackEnd/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllStore(ctx *fiber.Ctx) error {
	store := models.GetAllStores()
	s, err := json.Marshal(store)
	if err != nil {
		fmt.Println("Error while encding json [Func: GetAllOrders()]")
		return err
	}
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write(s)
	return nil
}

func CreateStore(ctx *fiber.Ctx) error {
	store := new(models.Stores)
	if err := ctx.BodyParser(store); err != nil {
		return err
	}
	store.CreateStore()
	ctx.Response().Header.Add("Content-Type", "application/json")
	return nil
}

func GetStoreById(ctx *fiber.Ctx) error {
	ID, _ := strconv.ParseInt(ctx.Params("ID"), 0, 0)
	store := models.GetStoreByID(ID)
	s, err := json.Marshal(store)
	if err != nil {
		return err
	}
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write(s)
	return nil
}

func DeleteStore(ctx *fiber.Ctx) error {
	ID, err := strconv.ParseInt(ctx.Params("ID"), 0, 0)
	if err != nil {
		fmt.Print("Error while parsing ID")
		return err
	}
	store := models.DeleteStore(ID)
	s, err := json.Marshal(store)
	if err != nil {
		fmt.Println("Error while encoding order into json")
		return err
	}
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write(s)
	return nil
}

func UpdateStore(ctx *fiber.Ctx) error {
	sid := ctx.Params("ID")
	id, err := strconv.ParseInt(sid, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing ID [Func: UpdateCustomer()]")
		return err
	}
	store := models.Stores{}
	res := store.UpdateStore(id)
	resStore, err := json.Marshal(res)
	if err != nil {
		fmt.Println("Error while encoding to json [Fucn: UpdateCustomer()]")
		return err
	}
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write(resStore)
	return nil
}
