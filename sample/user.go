package sample

type User struct {
	Id    string `bson:"_id"`
	Email string `bson:"email"`

	Profile  Profile   `bson:"profile"`
	Profiles []Profile `bson:"profiles"`
	Items    []string  `bson:"items"`

	Hoge interface{} `bson:"hoge"`

	Fuck int            `bson:"fuck"`
	Tag  map[string]Tag `bson:"tag"`
}

type Profile struct {
	Name string `bson:"name"`
}

type Tag struct {
	Description string `bson:"description"`
}
