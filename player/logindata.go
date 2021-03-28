package player

import (
	"fmt"
)


// interface used for abstract defining
type LogPass interface {
	BuildLogin(login, password string) error
}

//*********Login data building structure****///
type LoginStruct struct {
	Login, Password string
}


var (
	AuLogin = "serega"
	AuPassphrase = "aye228"
)

var AccessFlag int

var Authie = make(map[string]string)


//************************************//

func (login LoginStruct) BuildLogin(loginname, password string) error{
	fmt.Println("Insert Login and Passphrase")
	if _, err := fmt.Scan(&loginname, &password);
	err != nil {
		return fmt.Errorf("invalid login/pass: %w", err)
	}
	Authie[loginname] = password
	return nil
}

//-----------------------------------------------------------------------------------------//

//************FUNCTIONS FOR EXTERNAL CALLING*********//

func CheckLogin() int{
	for key, value := range Authie {
		if (AuLogin == key && AuPassphrase == value) {
			fmt.Println("login and password are accepted")
			AccessFlag = 1

		} else {
			AccessFlag = 0
			fmt.Println("wrong login/password")
		}
	}
	return AccessFlag
}

func SetLogin(loginIF LogPass){
	loginIF.BuildLogin("","")
}

//**********************************************//


