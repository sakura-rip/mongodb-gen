package templates

import "go.mongodb.org/mongo-driver/bson"

func And(attrs ...bson.M) bson.M {
	return bson.M{"$and": attrs}
}

func Or(attrs ...bson.M) bson.M {
	return bson.M{"$or": attrs}
}

func Nor(attrs ...bson.M) bson.M {
	return bson.M{"$nor": attrs}
}

func Not(attrs ...bson.M) bson.M {
	return bson.M{"$not": attrs}
}

func Unset(attrs bson.M) bson.M {
	return bson.M{"$unset": attrs}
}

func Set(attrs bson.M) bson.M {
	return bson.M{"$set": attrs}
}

func AddToSet(attrs bson.M) bson.M {
	return bson.M{"$addToSet": attrs}
}

func Pull(attrs bson.M) bson.M {
	return bson.M{"$pull": attrs}
}
