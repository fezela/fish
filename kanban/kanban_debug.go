
package main 
import (
	"fmt"
	"kanban/models"

	
)

func main() {
	u := models.User{
		Uname: "Snoopy",
	}

	c := models.Card{
		Id:    1,
		Owner: u.Uname,
		Task:  "Build kanban card model",
	}
	b := models.Board{
		Backlog:  []models.Card{c},
	}
	
	fmt.Println("Kanban Debug")
	fmt.Println("Still needs to show what it should be that or break or through something if something isn't right")
	fmt.Println("Card.GetOwner: ",c.GetOwner())
	fmt.Println("Card.GetTask: ",c.GetTask())
	fmt.Println("Card.GetBounty: ",c.Bounty) // should print 0
	fmt.Println("Board.Backlog: ", b.Backlog)

}
