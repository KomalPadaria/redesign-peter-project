package main

import (
	"fmt"

	"github.com/nurdsoft/redesign-grp-trust-portal-api/shared/jwt"
)

func main() {
	token, _ := jwt.GenerateWebhookToken()
	fmt.Println(token)
}
