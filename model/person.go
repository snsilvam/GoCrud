package model

//Struct of Community
type Community struct {
	//Name the a community, example: Overwatch
	Name string
}

//Communities slice of community
type Communities []Community

//Struct of person
type Person struct {
	Name        string      //Name of a person, example: Nicolás
	Age         uint8       //Age of a person, example: 20
	Communities Communities // Comunidades a las que pertenece la persona
}
type Persons []Person //Slice of persons
