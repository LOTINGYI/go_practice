package main


import (
	"fmt"
	"encoding/json"
)

type Monster struct{
	Name string
	Age int
	Birthday string
	Sal float64
	Skill string
}

func testStruct()  {
	monster := Monster{
		Name:"牛魔王",
		Age: 500,
		Birthday: "2018-01-01",
		Sal: 8000.0,
		Skill:"牛摩拳",
	}
	data,err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("序列化失敗\n",err)
		return
	}

	fmt.Printf("monster序列化後=%v",string(data))

}

func testMap()  {
	var a map[string]interface{}
	a = make(map[string]interface{})

	a["name"] = "紅孩兒"
	a["age"] = 30
	a["address"] = "taipei"
	data,err := json.Marshal(a)
	if err != nil {
		fmt.Println("序列化失敗\n",err)
		return
	}

	fmt.Printf("map序列化後=%v",string(data))

}


func testSlice()  {
	var slice []map[string]interface{}
	
	var m1 map[string]interface{}

	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = "7"
	m1["address"] = "北京"

	slice = append(slice,m1)

	var m2 map[string]interface{}

	m2 = make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = "20"
	m2["address"] = "墨西哥"
	slice = append(slice,m2)

	data,err := json.Marshal(slice)
	if err != nil {
		fmt.Println("序列化失敗\n",err)
		return
	}

	fmt.Printf("slice 序列化後=%v",string(data))
}
// 將json反序列化成struct
func unMarshalStruct()  {
	str := "{\"Name\":\"牛魔王\",\"Age\":500,\"Birthday\":\"2018-01-01\",\"Sal\":8000,\"Skill\":\"牛摩拳\"}"

	var monster Monster
	err := json.Unmarshal([]byte(str),&monster)

	if err != nil {
		fmt.Printf("unmarshal err=%v\n",err)
	}
	fmt.Printf("反序列化 monster=%v",monster)
}

func unMarshalMap()  {
	str := "{\"address\":\"山東\",\"age\":30,\"name\":\"紅孩兒\"}"

	// 定義一個map
	var a map[string]interface{}

	// 反序列化
	err := json.Unmarshal([]byte(str),&a)

	if err != nil {
		fmt.Printf("unmarshal err=%v\n",err)
	}
	fmt.Printf("反序列化 monster=%v",a)
}

func unMarshalSlice()  {
	str := "[{\"address\":\"山東\",\"age\":\"30\",\"name\":\"紅孩兒\"},"+
			"{\"address\":[\"美國\",\"墨西哥\"],\"age\":\"21\",\"name\":\"jack\"}]"

	// 定義一個map
	var slice []map[string]interface{}

	// 反序列化
	err := json.Unmarshal([]byte(str),&slice)

	if err != nil {
		fmt.Printf("unmarshal err=%v\n",err)
	}
	fmt.Printf("反序列化 monster=%v",slice)
}

func main() {
	// testStruct()
	// testSlice()
	// testMap()
	// unMarshalStruct()
	// unMarshalMap()
	unMarshalSlice()
}