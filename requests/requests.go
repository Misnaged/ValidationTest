package requests

import (
	"errors"
	"fmt"
	"userTA/player"
	"userTA/responses"
)

//***** Building a transaction request if conditions are passed ****//
func RequestTA() error{
	TransactionAttempts := []int{1, 2, 3}

		for _, att := range TransactionAttempts {
			if player.AccessFlag == 1 {
					switch att {
					case 1:
						responses.ResponseToTA()

					case 2:
						responses.ResponseToTA()

					case 3:
						responses.ResponseToTA()

					}
				if responses.ValidationFlag == 1{
					break
				}
			} else if err, ok := IsAccessFlagValid(player.AccessFlag);
				err != nil && !ok {
				return fmt.Errorf("access flag check returned an error: %w", err)
			}
		}
return nil
}
//***additional checking for login/pass///
func IsAccessFlagValid(receivedFlag int) (error, bool){
	if receivedFlag != 1 {
		return errors.New("AccessFlag isn't valid"), false
	}
	return nil, true
}


