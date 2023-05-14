package test

import (
	"chatgpt-web/internal/app/utils"

	"chatgpt-web/internal/app/dao"
	"fmt"
	"reflect"
	"testing"
	"time"
)

type Class struct {
	Id         string
	ClassId    int
	Master     Teacher
	CreateTime time.Time
}
type Teacher struct {
	Name string
	Sex  bool
	Age  int
}

func TestMapMarshal(t *testing.T) {
	data := map[string]interface{}{"Id": "hello", "ClassId": 1, "Master": map[string]interface{}{"Name": "xiyang", "Sex": true, "Age": 16}, "CreateTime": time.Now()}
	value, err := utils.Marshal(data, reflect.TypeOf(Class{}))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%v", value)
}

func TestJWTMarshal(t *testing.T) {
	user := dao.NewUser("xiyang")
	fmt.Println(utils.MarshalJWT(user))
}
