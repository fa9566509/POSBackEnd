package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/fa9566509/POSBackEnd/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllCategories(ctx *fiber.Ctx) error {
	categories := models.GetAllCateories()
	c, err := json.Marshal(categories)
	if err != nil {
		fmt.Println("Error while encoding json [Func: GetAllCategories()]")
		return err
	}
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write(c)
	return nil
}

func GetCategoryByID(ctx *fiber.Ctx) error {
	ID, err := strconv.ParseInt(ctx.Params("ID"), 0, 0)
	if err != nil {
		fmt.Println("Error while parsing ID [Func: GetCategoryByID()]")
		return err
	}
	category := models.GetCategoryByID(ID)
	c, err := json.Marshal(category)
	if err != nil {
		fmt.Println("Error while encoding json [Func: GetUserByID()]")
		return err
	}
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write(c)
	return nil
}

func CreateCategory(ctx *fiber.Ctx) error {
	category := new(models.Categories)
	if err := ctx.BodyParser(category); err != nil {
		fmt.Println("Error while decoding json [Func: CreateCategory()]")
		return err
	}
	c := category.CreateCategory()
	cn, err := json.Marshal(c)
	if err != nil {
		fmt.Println("Error while encoding json [Func: CreateCategory()]")
		return err
	}
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write(cn)
	return nil
}

func DeleteCategory(ctx *fiber.Ctx) error {
	sID := ctx.Params("ID")
	ID, err := strconv.ParseInt(sID, 0,0)
	if err != nil {
		fmt.Println("Error while parsing ID [Func: DeleteCategory()]")
		return err
	}
	category := models.DeleteCategory(ID)
	c, err := json.Marshal(category)
	if err != nil {
		fmt.Println("Error while encodig to json [Func: DeleteCategory()]")
		return err
	}
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write(c)
	return nil
}

func UpdateCategory(ctx *fiber.Ctx) error {
	sID := ctx.Params("ID")
	ID, err := strconv.ParseInt(sID, 0,0)
	if err != nil {
		fmt.Println("Error while parsing ID [Func: UpdateCategory()]")
		return err
	}
	category := new(models.Categories)
	if err := ctx.BodyParser(category); err != nil {
		fmt.Println("Error while parsing body [Func: UpdateCategory()]")
		return err
	}
	c := category.UpdateCategory(ID)
	cn, err := json.Marshal(c)
	if err != nil {
		fmt.Println("Error while encoding json [Func: UpdateCategory()]")
		return err
	}
	ctx.Response().Header.Add("Content-Type", "application/json")
	ctx.Write(cn)
	return nil
}