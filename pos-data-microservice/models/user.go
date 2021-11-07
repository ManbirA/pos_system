package models

type NewUser struct{
	Name	string		  `json:"name" bson:"name"`	
	Gender	string		  `json:"gender" bson:"gender"`
	Age		int			  `json:"age" bson:"age"`
}

type User struct{
	ID		string `json:"_id" bson:"_id"`
	Name	string		  `json:"name" bson:"name"`	
	Gender	string		  `json:"gender" bson:"gender"`
	Age		int			  `json:"age" bson:"age"`
}