package main

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func main() {
	mapWorld := emoji.Sprint("Hello :world_map:")
	fmt.Println(mapWorld)
}
