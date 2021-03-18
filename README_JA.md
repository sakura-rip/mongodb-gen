# mongodb-gen

このツールは、MongodbのDAO(data access object)を与えられたディレクトリの中にある構造体から生成します

# Installation

[リリースページ](https://github.com/sakura-rip/mongodb-gen/releases) からダウンロード 

またはシェルを利用してインストール

```shell
go get -u github.com/sakura-rip/mongodb-gen
go install github.com/sakura-rip/mongodb-gen
```

# Generate example

サンプルのディレクトリがあるとします

```shell
$ tree /f
D:.
├─sample  
│      group.go  
│      new.go  
│      user.go  
```

そして下のコマンドを実行します(`windowsの場合はmongo-gen.txt sample`)

```shell
mongodb-gen sample
```

生成されたファイルは下のようになります

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

# 生成されたDAOの利用方法

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

生成されたファイルを読むことで更なる情報を手に入れることができます。