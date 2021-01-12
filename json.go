package main

import (
	"fmt"
	"os"
	"encoding/json"
)

type response1 struct {
	Page int
	Fruits []string
}

type response2 struct {
	Page int `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	// Marshalling
	strs := []string{"apple", "peach", "pear"}
	jstrs, _ := json.Marshal(strs)
	fmt.Println(string(jstrs))

	// Doesn't use the tagged properties
	res1 := &response1{
		Page: 1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1j, _ := json.Marshal(res1)
	fmt.Println(string(res1j))

	res2 := &response2{
		Page: 1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2j, _ := json.Marshal(res2)
	fmt.Println(string(res2j))

	// Decode when we don't know the format
	byt := []byte(`{"num":6.13, "strs":["a", "b"]}`)

	// Not sure what this is, interface{}
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println("dat", dat)
	fmt.Println(dat["num"].(float64))

	// Decode using tags
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println("parsed response:", res)
	fmt.Println("first fruit:", res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
