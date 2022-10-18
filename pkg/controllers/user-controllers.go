package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/fa9566509/POSBackEnd/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(ctx *fiber.Ctx) error {
	users := models.GetAllUsers()
	u, err := json.Marshal(users)
	if err != nil {
		fmt.Println("Error while encoding json [Func: GetAllUsers()]")
		return err
	}
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write(u)
	return nil
}

func GetUserByID(ctx *fiber.Ctx) error {
	ID, err := strconv.ParseInt(ctx.Params("ID"), 0, 0)
	if err != nil {
		fmt.Println("Error while parsing ID [Func: GetUserByID()]")
		return err
	}
	user := models.GetUserByID(ID)
	u, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error while encoding json [Func: GetUserByID()]")
		return err
	}
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write(u)
	return nil
}

func CreateUser(ctx *fiber.Ctx) error {
	user := new(models.Users)
	if err := ctx.BodyParser(user); err != nil {
		fmt.Println("Error while decoding json [Func: CreateUSer()]")
		return err
	}
	u := user.CreateUser()
	un, err := json.Marshal(u)
	if err != nil {
		fmt.Println("Error while encoding json [Func: CreateUser()]")
		return err
	}
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write(un)
	return nil
}

func DeleteUser(ctx *fiber.Ctx) error {
	sID := ctx.Params("ID")
	ID, err := strconv.ParseInt(sID, 0,0)
	if err != nil {
		fmt.Println("Error while parsing ID [Func: DeleteUser()]")
		return err
	}
	user := models.DeleteUser(ID)
	u, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error while encodig to json [Func: DeleteUser()]")
		return err
	}
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write(u)
	return nil
}

func UpdateUser(ctx *fiber.Ctx) error {
	sID := ctx.Params("ID")
	ID, err := strconv.ParseInt(sID, 0,0)
	if err != nil {
		fmt.Println("Error while parsing ID [Func: UpdateUser()]")
		return err
	}
	user := new(models.Users)
	if err := ctx.BodyParser(user); err != nil {
		fmt.Println("Error while parsing body [Func: UpdateUser()]")
		return err
	}
	u := user.UpdateUser(ID)
	un, err := json.Marshal(u)
	if err != nil {
		fmt.Println("Error while encoding json [Func: UpdateUser()]")
		return err
	}
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write(un)
	return nil
}