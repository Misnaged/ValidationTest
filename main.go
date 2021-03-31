package main

import (
	"log"
	"os"
	"userTA/player"
	"userTA/requests"
	"userTA/responses"

	//"userTA/responses"
)
var loginStruct player.LoginStruct


func main() {

	if err:= player.SetLogin(loginStruct); err != nil{
		log.Println(err)
		os.Exit(1)
		return
	}

	if _, err := player.CheckLogin(loginStruct); err != nil{
		log.Println(err)
		os.Exit(1)
		return
	}

	if err := requests.RequestTA(); err != nil{
		log.Println(err)
		os.Exit(1)
		return
	}


	if _,_,err := responses.ResponseToTA(); err != nil{
		log.Println(err)
		os.Exit(1)
		return
	}

}
