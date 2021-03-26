package responses

import (
	"fmt"
	"log"
)

//****Defining UserInputs
var UserInput int

//******* Only 3 variant of transactions are allowed*****///
var ValidTransActions = []int {10000, 20000, 30000}

//*********** Flags used for managing the code ****//
var ValidationFlag int

func ResponseToTA() (int, int){

	if _, err := fmt.Scan(&UserInput);
		err != nil {
		log.Print(" Transaction has been failed due to ", err)
	}

	for _,ValidTA := range ValidTransActions{
		if UserInput == ValidTA{
			fmt.Println("Validation succeeded!")
		}else{
			log.Println("Validation rejected!")
		}
		if ValidTA == UserInput{
			ValidationFlag = 1
			break
		}else{
			ValidationFlag = 0
		}
	}

	return UserInput, ValidationFlag
}