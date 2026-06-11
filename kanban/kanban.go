package kanban

import (
	"fmt"
	"net/http"
	//"fish/kanban/models"

)

func createCard(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "creating a card with inputed data. FORM")
}

func getTasks(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "This is going to pull some cards from column specified in requests")

}

func Handler(w http.ResponseWriter, r *http.Request){

	switch r.URL.Path {

	case "/kanban":
	fmt.Fprintln(w, "hello from Kanban and a good stoping point.  Next up send a get request via postman! When you can do that do it from Android. No matter how annoying it is.  This is when you become a millionaire.")

	case "/kanban/createCard":
		fmt.Fprintln(w, "Feature not yet implemented")
	
}

//func main(){
//	http.Handlefunc("/kanban", home)
}
