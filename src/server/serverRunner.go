package server

import (
	"./config"
	"fmt"
)

func main()  {
	config.LoadConfig();
	fmt.Println("Load config success!")
}
