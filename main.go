package main

import (
	"flag"
	"fmt"

	"github.com/thzoid/ws-game-server/shared"
)

var (
	mapSize = shared.Coordinate{}
)

func main() {
	portPtr := flag.Uint("port", 8080, "port for the server to listen")
	widthPtr := flag.Uint("width", 10, "world map width")
	heightPtr := flag.Uint("height", 10, "world map height")
	flag.Parse()

	mapSize.X = int(*widthPtr)
	mapSize.Y = int(*heightPtr)

	fmt.Println("world map size: ("+fmt.Sprint(*widthPtr)+",", fmt.Sprint(*heightPtr)+")")
	listen(":" + fmt.Sprint(*portPtr))
}
