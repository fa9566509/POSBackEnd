package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/fa9566509/POSBackEnd/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(ctx *fiber.Ctx) error {
	users := models.GetAllUsers
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
	ID, err := strconv.ParseInt(ctx.Params("ID"), 0,0)
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