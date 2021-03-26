package requests

import (
	"fmt"
	"userTA/player"
	"userTA/responses"

	//"userTA/EventHandler"
)

//***** Building a transaction request if conditions are passed ****//
func RequestTA(){
	TransactionAttempts := []int{1, 2, 3}
	for _, att := range TransactionAttempts{
		switch att {
		case 1:
			CallTransAct()
			fmt.Println(responses.ValidationFlag)
		case 2:
			CallTransAct()
			fmt.Println(responses.ValidationFlag)

		case 3:
			CallTransAct()
			fmt.Println(responses.ValidationFlag)
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

/*
func VerifyAuth() {
	h1 := func(w http.ResponseWriter, r *http.Request){
		http.ServeFile(w, r, "user.html")
	}
	h2 := func(w http.ResponseWriter, req *http.Request){
		 l := req.FormValue("username")
		 p := req.FormValue("passphrase")


		loginname := l
		password := p

		for key, value := range player.Authie {
			if (loginname == key && password == value){
				fmt.Println("login and password are accepted")
				fmt.Fprintf(w, "%s", "\n Accepted")
			}else{
				fmt.Println("wrong login/password")
			}
		}
	}
	http.HandleFunc("/", h1)
	http.HandleFunc("/postform", h2)
	fmt.Println("Server is listening..., \n open localhost:8181 on your web browser and fill form")
	http.ListenAndServe(":8181", nil)
}


 */


/*
func FromRequest(req *http.Request) (net.IP, error) {

}

		err:= r.ParseForm() //parse in middleware, data will be contained in r.PostForm
		if err != nil {
			log.Fatalf("Failed to decode postFormByteSlice: %v", err)
		}
		loginname := r.PostFormValue(l)
		password := r.PostFormValue(p)

		for key, value := range player.Authie {
			if (loginname == key && password == value){
				fmt.Println("zaebok")
			}else{
			fmt.Println("xyeta")
			}
		}
*/