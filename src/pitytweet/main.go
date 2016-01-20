package main

import (
	"flag"
	"fmt"
)

func main() {
	var screen_name *string = flag.String("name", "", "twitter user name")
	flag.Parse()

	fmt.Printf("Pitiest gophers %s", *screen_name)
}
