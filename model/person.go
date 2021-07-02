package model

//Struct of Community
type Community struct {
	//Name the a community, example: Overwatch
	Name string `json:"name"`
}

//Communities slice of community
type Communities []Community

//Struct of person
type Person struct {
	Name        string      `json:"name"`        //Name of a person, example: Nicolás
	Age         uint8       `json:"age"`         //Age of a person, example: 20
	Communities Communities `json:"communities"` // Comunidades a las que pertenece la persona
}
type Persons []Person //Slice of persons
