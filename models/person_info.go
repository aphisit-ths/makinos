package models

type Telephone struct {
	Id        string `bson:"_id" `
	IdCard    string `bson:"idCard"`
	Telephone string `bson:"telephone"`
}

type PersonInfo struct {
	Id        string `bson:"_id" `
	IdCard    string `bson:"idCard"`
	FirstName string `bson:"firstName"`
	LastName  string `bson:"lastName"`
	PreName   string `bson:"preName"`
}
