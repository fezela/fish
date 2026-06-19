package models 

import "time"
	
	

type Card struct {
	Id int //example: 1
	Owner User
	Task string
	Bounty int//this default to 0 allows the card to have various functions mapped to it by a numerical identifier.  It just lets it be a mutable object.
	Status string
	CreatedAt time.Time
	UpdatedAt time.Time

}

func (c Card) GetOwner() User {
	return c.Owner
}

func (c Card) GetTask() string {
	return c.Task
}

func (c Card) GetStatus() string {
	return c.Status
}

func (c Card) setStatus(status string) {
	c.Status = status
} 


func (c Card) SetBounty(bounty int) {
	c.Bounty = bounty
}

func (c Card) setUpdatedAt(){
	c.UpdatedAt = time.Now()
}

func (c Card) setCreatedAt(){
	c.CreatedAt = time.Now()
}



