package main

import (
	"github.com/joho/godotenv"
	"github.com/joseph-beck/chit-chat/api/internal/app"
)

//	@title			Chit Chat
//	@version		0.1
//	@description	a realtime chatting and calling app

//	@contact.name	Joseph Beck

//	@license.name	MIT
//	@license.url	https://opensource.org/license/mit/

// @host		localhost:8080
// @BasePath	/
func main() {
	a := app.New()
	a.Run()
	a.Close()
}

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}
