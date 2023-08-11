package rcon

import (
	"fmt"

	"github.com/Random7-JF/go-rcon/app/config"

	mcrcon "github.com/Kelwing/mc-rcon"
)

type Connection struct {
	Rcon      *mcrcon.MCConn
	Connected bool
}

var RconSession *Connection

func SetupRconSession(a *config.App) *Connection {
	return &Connection{
		Rcon: a.Rcon.Session,
	}
}

func NewRconSession(c *Connection) {
	RconSession = c
}

func SetupConnection(App *config.App) error {
	App.Rcon.Session = new(mcrcon.MCConn)

	ip := App.Rcon.Ip + ":" + App.Rcon.Port
	err := App.Rcon.Session.Open(ip, App.Rcon.Password)
	if err != nil {
		fmt.Println("Error opening rcon connection:", err)
		return err
	}
	err = App.Rcon.Session.Authenticate()
	if err != nil {
		fmt.Println("Error authenticating rcon connection:", err)
		return err
	}

	test, err := App.Rcon.Session.SendCommand("list")
	if err != nil {
		fmt.Println("Error sending command:", err)
		return err
	}

	fmt.Println(test)
	App.Rcon.Connection = true
	return nil
}
