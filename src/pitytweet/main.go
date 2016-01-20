package main

import (
	"flag"
	"fmt"
)

func main() {
	var screen_name *string = flag.String("name", "", "twitter user name")
	var consumer_key *string = flag.String("key", "", "")
	var consumer_secret *string = flag.String("secret", "", "")

	flag.Parse()

	fmt.Printf("public flags: name %s \n", *screen_name)
	if *consumer_key != "" && *consumer_secret != "" {
		fmt.Println("Credentials recieved")
	}

}
