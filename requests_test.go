package main

import (
	"fmt"
	"testing"
	"userTA/player"
	"userTA/requests"
	"userTA/responses"
)


func TestRequestTA(t *testing.T){

	player.Authie["serega"] = "aye228"
player.CheckLogin()
if player.AccessFlag == 1{
	fmt.Println("AccessFlag has a true value, test has been passed")
}else{
	fmt.Println("AccessFlag test has been failed, Access flag ==", player.AccessFlag)
}
responses.UserInput = 10000
requests.RequestTA()

fmt.Println("Testing has been completed")
}

