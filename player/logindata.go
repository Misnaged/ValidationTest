package player

import (
	"errors"
	"fmt"
	"strings"
)


// interface used for abstract defining
type LogPass interface {
	BuildLogin(login, password string) error
	LogPassEq(key, value string) (error, bool)
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

func (login LoginStruct) LogPassEq(key, value string) (error, bool){
	if strings.EqualFold(AuLogin, key) && strings.EqualFold(AuPassphrase, value) {
		return nil, true
	} else{
		return errors.New("error in Login/Passphrase"), false
	}
	}
//-----------------------------------------------------------------------------------------//

//************FUNCTIONS FOR EXTERNAL CALLING*********//



func SetLogin(loginIF LogPass) error{

	if err := loginIF.BuildLogin("",""); err != nil{
		return fmt.Errorf("error: %w", err)
	}
	return nil
}


func CheckLogin(loginIF LogPass) (int, error){
	for key, value := range Authie {
		err, ok := loginIF.LogPassEq(key, value)
		if ok {
			fmt.Println("login and password are accepted")
			AccessFlag = 1
		} else if err != nil && !ok{
			AccessFlag = 0
			return 0, fmt.Errorf("login/pass failed due to: %w", err)
		}
	}
	return AccessFlag, nil
}
//**********************************************//

