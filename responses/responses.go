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

		for _, ValidTA := range ValidTransActions {
			areEqual, err := Matching(UserInput, ValidTA)

			if areEqual {fmt.Println("TransAction has been done successfully")}

			if err != nil{
				return 0, 0, fmt.Errorf(" Transaction is unvalid: %w", err)
			}

			if ValidTA == UserInput {
				ValidationFlag = 1
				break
			} else {
				ValidationFlag = 0
			}
		}
	return UserInput, ValidationFlag, nil
}

func Matching(a1, b1  int)  (areEq bool, err error){
	if a1 == b1{
		return true, nil
	}
	return true, nil
}
