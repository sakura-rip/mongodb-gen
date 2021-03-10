package sample

type User struct {
	Id    string `bson:"_id"`
	Email string `bson:"email"`

	Profile Profile `bson:"profile"`

	Items []string `bson:"items"`

	Tag map[string]Tag `bson:"tag"`
}

type Profile struct {
	Name string `bson:"name"`
}

type Tag struct {
	Description string `bson:"description"`
}
