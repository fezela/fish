
package kanban 
import (
	
	"fish/kanban/models"
	"testing"


	
)


func TestAddCard(t *testing.T){
	u := models.User{
		Uname: "Snoopy",
		}	

	c := models.Card{
		Id:    1,
		Owner: u, //This should probably not be a string but the actual User object.
		Task:  "Build kanban card model",
		}
	
	b := models.Board{}
		

	b.AddCard(c, "Backlog")
	if len(b.Backlog) !=1 {
		t.Fatalf("expected 1 card, got %d", len(b.Backlog))
		}		
	}
	


