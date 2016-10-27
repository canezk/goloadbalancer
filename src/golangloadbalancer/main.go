package golangloadbalancer

import (
	"fmt"
	"golangloadbalancer/config"
)

func main() {
	config.LoadConfig()
	fmt.Println("Load config success!")
}
