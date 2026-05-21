package models 

type Card struct {
	Id int //example: 1
	Owner string 
	Task string
	Bounty int//this default to 0 allows the card to have various functions mapped to it by a numerical identifier.  It just lets it be a mutable object.

}

func (c Card) GetOwner() string {
	return c.Owner
}

func (c Card) GetTask() string {
	return c.Task
}

func (c Card) SetBounty(bounty int) {
	c.Bounty = bounty
}


