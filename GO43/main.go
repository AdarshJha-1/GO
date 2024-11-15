package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
	{
		"user_id": 12345,
		"name": 'User A",
		"age": 22,
		"password": "my-password",
		"roles": ["admin", "collaborator"]
	}
*/

type User struct {
	ID          int      `json:"user_id"`
	Name        string   `json:"name,omitempty"`
	Age         int      `json:"age"`
	Password    string   `json:"-"`     // ignore password
	Permissions []string `json:"roles"` // rename "Permissions" to roles
}

type Person struct {
	Name       string   `json:"full_name"`           // Customize field name
	Age        int      `json:"years_old,omitempty"` // Exclude it zero
	occupation string   `json:"-"`                   // ignore this field
	Languages  []string `json:"spoken_languages"`    // rename "Permissions" to roles
}

type PersoN struct {
	Name   string `json:"username"`
	Age    int    `json:"age"`
	Active bool   `json:"is_active"`
}

func main() {
	u1 := User{
		ID:          1,
		Name:        "Luffy",
		Age:         19,
		Password:    "meat",
		Permissions: []string{"Captain", "Samurai", "Ninja"},
	}

	b, err := json.Marshal(u1)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
	// for _, v := range b {
	// 	fmt.Println(v, ":", string(v))
	// }

	jsonData := `{
  "full_name": "Satoru Gojo",
  "years_old": 28,
  "spoken_languages": ["Japanese", "English"]
}`

	var p1 Person
	if err := json.Unmarshal([]byte(jsonData), &p1); err != nil {
		panic(err)
	}
	fmt.Println(p1)

	p2 := PersoN{
		Name:   "Zoro",
		Age:    21,
		Active: true,
	}

	f, err := os.Create("output.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err = json.NewEncoder(f).Encode(p2); err != nil {
		panic(err)
	}

	var p3 PersoN
	file, err := os.Open("output.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if err = json.NewDecoder(file).Decode(&p3); err != nil {
		panic(err)
	}

	fmt.Println(p3)
}
