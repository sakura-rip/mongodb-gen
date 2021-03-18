# mongodb-gen

The mongodb-gen tool produces a mongodb dao(data access object) from given go struct.

# Installation

Download from release [page](https://github.com/sakura-rip/mongodb-gen/releases)

or using shell

```shell
go get -u github.com/sakura-rip/mongodb-gen
go install github.com/sakura-rip/mongodb-gen
```

# Generate example

Here is sample dir

```shell
$ tree
D:.
├─sample  
│      group.go  
│      new.go  
│      user.go  
```

then

```shell
mongodb-gen sample
```

out:

```shell
$ tree /f
D:.
├─sample
│      group.go
│      new.go
│      user.go
│
├─sample_dao
│      client_gen.go
│      GroupCollectionBase_gen.go
│      GroupMembersField_gen.go
│      GroupNameField_gen.go
│      query.go
      ...etx
```

# Usage of generated file

```go
package main

import (
	"fmt"
	"github.com/sakura-rip/mongodb-gen/sample"
	dao "github.com/sakura-rip/mongodb-gen/sample_dao"
	"log"
)

func main() {
	uuid := "39f436034ff548748229433587dce645"
	userCl := dao.NewUserDaoClient()
	_, err := userCl.InsertOne(&sample.User{
		Id:      uuid,
		Email:   "sakura@example.com",
		Profile: sample.Profile{Name: "sakura"},
	})
	if err != nil {
		log.Fatal("failed to insert user")
	}
	user, err := userCl.GetUserByProfileName("sakura")
	if err != nil {
		fmt.Printf("%#v\n", user)
	}
}
```

read generated file to get more detail