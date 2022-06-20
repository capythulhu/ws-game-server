package main

import (
	"flag"
	"fmt"
)

func main() {
	portPtr := flag.Uint("port", 8080, "port for the server to listen")
	widthPtr := flag.Uint("width", 20, "world map width")
	heightPtr := flag.Uint("height", 10, "world map height")
	hbIntervalPtr := flag.Uint("interval", 100, "heartbeat interval in ms")
	flag.Parse()

	matchMap.Size.X = int(*widthPtr)
	matchMap.Size.Y = int(*heightPtr)

	hbInterval = int(*hbIntervalPtr)

	fmt.Println("world map size: ("+fmt.Sprint(*widthPtr)+",", fmt.Sprint(*heightPtr)+")")
	listen(":" + fmt.Sprint(*portPtr))
}
