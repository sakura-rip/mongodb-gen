package sample

type Group struct {
	Id      string   `bson:"_id"`
	Name    string   `bson:"name"`
	Members []string `bson:"members"`
	Setting Setting  `bson:"setting"`
}

type Setting struct {
	AllowSearch bool `bson:"allow_search"`
}
