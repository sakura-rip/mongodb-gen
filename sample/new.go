package sample

type New struct {
	Id    string `bson:"_id"`
	Base1 Base1  `bson:"base_1"`
}
type Base1 struct {
	Base2 Base2 `bson:"base_2"`
}

type Base2 struct {
	Name string `bson:"name"`
}
