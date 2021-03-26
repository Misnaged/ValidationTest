package main

import (
	"userTA/player"
	"userTA/requests"
)
var loginStruct player.LoginStruct

func main() {
	player.SetLogin(loginStruct)
	player.CheckLogin()
	requests.RequestTA()
}