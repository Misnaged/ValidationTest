package responses

import (
	"fmt"
)

//****Defining UserInputs
var UserInput int

//******* Only 3 variant of transactions are allowed*****///
var ValidTransActions = []int {10000, 20000, 30000}

//*********** Flags used for managing the code ****//
var ValidationFlag int


func ResponseToTA() (int, int, error){
	if _, err := fmt.Scan(&UserInput);
		err != nil {
		return 0, 0, fmt.Errorf(" Transaction has been failed due to: %w", err)
	}

	for _, ValidTA := range ValidTransActions{
		if ValidTA == UserInput {
			fmt.Println("TransAction has been done successfully")
			fmt.Printf("\n, %d", UserInput)
			ValidationFlag = 1
			break
		} else {
			ValidationFlag = 0
		}
	}
	return UserInput, ValidationFlag, nil
}

