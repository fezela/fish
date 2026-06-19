package models

type Board struct {
	Backlog []Card
	Doing []Card
	Done []Card
}

 func (b *Board) GetColumn(col string) []Card {
	 switch col {
	 case "Backlog":
		 return b.Backlog
	 case "Doing":
		 return b.Doing
	 case "Done":
		 return b.Done
	 }

	 return nil

 }

 func (b *Board) AddCard(c Card, destination string) {
	switch destination {
	case "Backlog":
		c.setStatus("Backlog")
		b.Backlog = append(b.Backlog, c)

	case "Doing":
		c.setStatus("Doing")
		b.Doing = append(b.Doing, c)
	
	case "Done":
		c.setStatus("Done")
		b.Done = append(b.Done, c)
	}
}
 
