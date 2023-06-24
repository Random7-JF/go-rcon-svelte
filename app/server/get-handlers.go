package server

import (
	"fmt"

	"github.com/Random7-JF/go-rcon/app/model"
	"github.com/Random7-JF/go-rcon/app/rcon"
	"github.com/gofiber/fiber/v2"
)

func IndexHandler(c *fiber.Ctx) error {
	return c.Render("pages/index", model.TempalteData{
		Title: "Home",
	}, "layouts/main")
}

func DashboardHandler(c *fiber.Ctx) error {

	players, err := rcon.GetPlayers()
	if err != nil {
		return err
	}
	whitelist, err := rcon.GetWhitelist()
	if err != nil {
		return err
	}
	commands, err := model.GetCommandLog(5)
	if err != nil {
		return err
	}

	data := make(map[string]interface{})
	data["Players"] = players
	data["Rcon"] = AppConfig.RconSettings.Connection
	data["Whitelist"] = whitelist
	data["Commands"] = commands

	return c.Render("pages/dashboard", model.TempalteData{
		Title: "Dashboard",
		Data:  data,
	}, "layouts/main")
}

func PlayersPageHandler(c *fiber.Ctx) error {
	players, err := rcon.GetPlayers()
	if err != nil {
		return err
	}

	data := make(map[string]interface{})
	data["Players"] = players

	return c.Render("pages/players", model.TempalteData{
		Title: "Players",
		Data:  data,
	}, "layouts/main")
}

func CommandsHandler(c *fiber.Ctx) error {
	return c.Render("pages/commands", model.TempalteData{
		Title: "Commands",
	}, "layouts/main")
}

func WhitelistHandler(c *fiber.Ctx) error {
	whitelist, err := rcon.GetWhitelist()
	if err != nil {
		return err
	}

	data := make(map[string]interface{})
	data["Whitelist"] = whitelist

	return c.Render("pages/whitelist", model.TempalteData{
		Title: "Whitelist",
		Data:  data,
	}, "layouts/main")
}

func LoginHandler(c *fiber.Ctx) error {

	session, err := AppConfig.Store.Get(c)
	if err != nil {
		fmt.Println(err)
	}
	auth := session.Get("auth")
	fmt.Println("Auth: ", auth)

	keys := session.Keys()
	fmt.Println("Keys: ", keys)

	return c.Render("pages/login", model.TempalteData{
		Title: "Login",
	}, "layouts/main")
}
