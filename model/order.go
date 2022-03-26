package model

type Order struct {
	CustomerId string `bson:"customerid"`
	Amount     int32  `bson:"amount"`
	Product    []Product
}

type Product struct {
	ID          string `bson:"_id"`
	Name        string `bson:"name"`
	Description string `bson:"description"`
	Type        string `bson:"type"`
	Unit        int32  `bson:"unit"`
	Price       int32  `bson:"price"`
	Supplier    string `bson:"supplier"`
}

type RemoveOneOrder struct {
	OrderId   string `bson:"orderid"`
	ProductId string `bson:"productid"`
}
