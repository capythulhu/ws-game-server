package main

import (
	"flag"
	"fmt"
)

func main() {
	portPtr := flag.Uint("port", 8080, "port for the server to listen")
	widthPtr := flag.Uint("width", 10, "world map width")
	heightPtr := flag.Uint("height", 10, "world map height")
	flag.Parse()
	fmt.Println("world map size:", fmt.Sprint(*widthPtr)+",", *heightPtr)
	listen(":" + fmt.Sprint(*portPtr))
}
