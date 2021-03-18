// auto generated by https://github.com/sakura-rip/mongodb-gen
// Version : develop@v1
// DO NOT EDIT

package dao

import (
	"github.com/sakura-rip/mongodb-gen/sample"
	"go.mongodb.org/mongo-driver/bson"
)

func (cl *NewDaoClient) DeleteBase1(id string) error {
	return cl.DeleteAttributes(id, bson.M{"base_1": 1})
}

func (cl *NewDaoClient) InsertOrUpdateBase1(id string, object interface{}) error {
	return cl.UpdateAttribute(id, bson.M{"base_1": object})
}

func (cl *NewDaoClient) GetNewBase1(id string) (sample.Base1, error) {
	attribute, err := cl.GetAttribute(id, bson.D{{"base_1", 1}})
	if err != nil {
		return sample.Base1{}, err
	}
	return attribute.Base1, err
}

func (cl *NewDaoClient) GetNewByBase1(profile sample.Base1) ([]sample.New, error) {
	return cl.Get(bson.M{"base_1": profile})
}

func (cl *NewDaoClient) UpdateOrInsertNewBase1Base2(id string, base2 Base2) error {
	return cl.UpdateAttribute(id, bson.M{"base_1.base_2": base2})
}

func (cl *NewDaoClient) DeleteNewBase1Base2(id string) error {
	return cl.DeleteAttributes(id, bson.M{"base_1.base_2": 1})
}

func (cl *NewDaoClient) GetNewBase1Base2(id string) (val Base2, err error) {
	new, err := cl.GetAttribute(id, bson.D{{"base_1.base_2", 1}})
	if err != nil {
		return val, err
	}
	return new.Base1.Base2, err
}

func (cl *NewDaoClient) GetNewByBase1Base2(base2 Base2) ([]sample.New, error) {
	return cl.Get(bson.M{"base_1.base_2": base2})
}
func (cl *NewDaoClient) UpdateOrInsertNewBase1Name(id string, name string) error {
	return cl.UpdateAttribute(id, bson.M{"base_1.base_2.name": name})
}

func (cl *NewDaoClient) DeleteNewBase1Name(id string) error {
	return cl.DeleteAttributes(id, bson.M{"base_1.base_2.name": 1})
}

func (cl *NewDaoClient) GetNewBase1Name(id string) (val string, err error) {
	new, err := cl.GetAttribute(id, bson.D{{"base_1.base_2.name", 1}})
	if err != nil {
		return val, err
	}
	return new.Base1.Base2.Name, err
}

func (cl *NewDaoClient) GetNewByBase1Name(name string) ([]sample.New, error) {
	return cl.Get(bson.M{"base_1.base_2.name": name})
}
