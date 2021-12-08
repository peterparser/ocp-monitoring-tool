package main

import (
	"fmt"

	"./configuration"
	"./ocpAuth"
)

func main() {
	conf := configuration.ParseConfigurationFile("example.yaml")
	loginToken := ocpAuth.GetOcpToken(conf.OcpOauthUrl, conf.Username, conf.Password)
	fmt.Print(loginToken)
}
