package player

import (
	"fmt"
	"log"
)
//players database path
//const playerDBpath  = "playersDB.json"
var Authie map[string]string

// interface used for abstract defining
type LogPass interface {
	BuildLogin(login, password string)
	//PlayersFlagData(blocked, transacting bool)
}

//*********Login data building structure****///
type LoginStruct struct {
	Login, Password string
	//IsBlocked, IsInTransactionSession bool
}


var (
	login = "serega"
	passphrase = "aye228"
)

var AccessFlag int


//************************************//

/*
//  defines entire Players DataBase
type AllPlayers struct {
	Players []LoginStruct
}
// global variable for getting an access to DB
var playersDB AllPlayers

//   AN IMPLEMENT METHODS FOR INTERFACE
//   ConfigFlag-builder method
func (login LoginStruct) PlayersFlagData(blocked, transacting bool) {
	getPlayersConfigFlags, _ := ioutil.ReadFile(playerDBpath)
	_ = json.Unmarshal(getPlayersConfigFlags, &playersDB)
	PlayerConfigFlags := LoginStruct{IsBlocked: blocked, IsInTransactionSession: transacting}
	playersDB.Players = append(playersDB.Players, PlayerConfigFlags)
	setPlayerConfigFlags, _ := json.MarshalIndent(&playersDB, "", "  ")

	_ = ioutil.WriteFile(playerDBpath, setPlayerConfigFlags, 0644)
	//   primitive CF implementation is using true/false booleans to "switch" needed properties
}

//    Login-builder method
/*
}
func (login LoginStruct) BuildLogin(loginname, password string, playerID int) {
	if    _, err := fmt.Scan(&loginname, &password);    err != nil {
		log.Print("  Scan for i, j & k failed, due to ", err)
		return
	}
	playersinfo, _ := ioutil.ReadFile(playerDBpath)

	_ = json.Unmarshal(playersinfo, &playersDB)

	instanciatePlayer := LoginStruct{
		Login: loginname,
		Password: password,
		PlayerID: playerID,
	}

	playersDB.Players = append(playersDB.Players, instanciatePlayer)

	newplayerinfo, _ := json.MarshalIndent(&playersDB, "", "  ")


	_ = ioutil.WriteFile(playerDBpath, newplayerinfo, 0644)
}

func SetupPlayerFlag(loginIF LogPass){
	var blockflag, transflag bool
	loginIF.PlayersFlagData(blockflag, transflag)
}

*/


func (login LoginStruct) BuildLogin(loginname, password string){
	fmt.Println("Insert Login and Passphrase")
	Authie = make(map[string]string)
	if _, err := fmt.Scan(&loginname, &password);
	err != nil {
		log.Print("  Scan for i, j & k failed, due to ", err)
		return
	}
	Authie[loginname] = password
}

//-----------------------------------------------------------------------------------------//

//************FUNCTIONS FOR EXTERNAL CALLING*********//

func CheckLogin() int{
	for key, value := range Authie {
		if (login == key && passphrase == value) {
			fmt.Println("login and password are accepted")
			AccessFlag = 1
			//fmt.Fprintf(w, "%s", "\n Accepted")
		} else {
			AccessFlag = 0
			fmt.Println("wrong login/password")
		}
	}
	return AccessFlag
}

func SetLogin(loginIF LogPass){
	//PlayerID := 228
	loginIF.BuildLogin("","")
}

//**********************************************//


