package requests

import (
	"fmt"
	"userTA/player"
	"userTA/responses"
)

//***** Building a transaction request if conditions are passed ****//
func RequestTA(){
	TransactionAttempts := []int{1, 2, 3}
	for _, att := range TransactionAttempts{
		switch att {
		case 1:
			CallTransAct()
			fmt.Println("validation flag is", responses.ValidationFlag)
		case 2:
			CallTransAct()
			fmt.Println("validation flag is", responses.ValidationFlag)

		case 3:
			CallTransAct()
			fmt.Println("validation flag is", responses.ValidationFlag)
		}
		if responses.ValidationFlag == 1{
			break
		}else{
			continue
		}
	}

}

//**** checking if Login/Password are passed succesfully ****////
func CallTransAct(){
	if player.AccessFlag == 1{
		responses.ResponseToTA()
	}else{
		return
	}
}

