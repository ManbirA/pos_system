package models

type Product struct{
	ID			string 		`json:"_id" bson:"_id"`
	Name		string 		`json:"name" bson:"name"`	
	Inventory 	int	   		`json:"inventory" bson:"inventory"`
	Price 		float64		`json:"price" bson:"price"`
}

type NewProduct struct{
	Name		string 		`json:"name" bson:"name"`	
	Inventory 	int	   		`json:"inventory" bson:"inventory"`
	Price 		float64		`json:"price" bson:"price"`
}