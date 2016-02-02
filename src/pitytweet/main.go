package main

import (
	"flag"
	"fmt"
	"github.com/mrjones/oauth"
	"io/ioutil"
	// "golang.org/x/net/context"
	// "golang.org/x/oauth2"
	"log"
	"os"
	"os/exec"
)

func main() {
	var screenName *string = flag.String("name", "", "twitter user name")
	var consumerKey *string = flag.String("key", "", "")
	var consumerSecret *string = flag.String("secret", "", "")

	flag.Parse()

	fmt.Printf("fetching for: %s \n", *screenName)
	if *consumerKey == "" || *consumerSecret == "" {
		// panic("needed key and secret, exiting")
		log.Fatal("needed key and secret, exiting")
		os.Exit(1)
	}

	c := oauth.NewConsumer(
		*consumerKey,
		*consumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
			AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		})

	c.Debug(true)

	requestToken, u, err := c.GetRequestTokenAndUrl("oob")
	if err != nil {
		log.Fatal(err)
	}

	exec.Command("google-chrome-stable", u).Run()

	fmt.Println("(1) Grant access, you should get back a verification code.")
	fmt.Println("(2) Enter that verification code here: ")

	verificationCode := ""
	fmt.Scanln(&verificationCode)

	accessToken, err := c.AuthorizeToken(requestToken, verificationCode)
	if err != nil {
		log.Fatal(err)
	}

	client, err := c.MakeHttpClient(accessToken)
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Get(
		"https://api.twitter.com/1.1/statuses/home_timeline.json?count=1")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	bits, err := ioutil.ReadAll(response.Body)
	fmt.Println("The newest item in your home timeline is: " + string(bits))

}
