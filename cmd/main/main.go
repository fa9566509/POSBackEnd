package main

import (
	"fmt"

	"github.com/fa9566509/POSBackEnd/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	fmt.Print(`
	
    ____   ___  _____  __                   __    ______            __
   / __ \ / __\/ ___/ / /_    ____ _ _____ / /__ / ____/____   ____/ /
  / /_/ // / / \__ \ / __ \  / __ '// ___// //_// __/  / __ \ / __  / 
 / ____// /_/ /___/ // /_/ // /_/ // /__ / ,<  / /___ / / / // /_/ /  
/_/     \____//____//_.___/ \__,_/ \___//_/|_|/_____//_/ /_/ \__,_/   
                                                                      

`)

	utils.RegisterAllRoutes(app)
	app.Listen(":9010")
}
