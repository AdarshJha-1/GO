package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type CarType int

const (
	_ CarType = iota
	Hatchback
	Sedan
	SUV
)

var (
	stringToCarType = map[string]CarType{
		"hatchback": Hatchback,
		"sedan":     Sedan,
		"suv":       SUV,
	}
	carTypeToString = map[CarType]string{
		Hatchback: "hatchback",
		Sedan:     "sedan",
		SUV:       "suv",
	}
)

func (c CarType) String() string {
	return carTypeToString[c]
}

func (c CarType) MarshalJSON() ([]byte, error) {
	t, ok := carTypeToString[c]
	if !ok {
		return nil, errors.New("invalid car type provided")
	}
	return []byte(`"` + t + `"`), nil
}

func (c *CarType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	t, ok := stringToCarType[s]
	if !ok {
		return fmt.Errorf("invalid car type: %s", s)
	}
	*c = t
	return nil
}

func main() {
	c := Hatchback
	n, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}

	fmt.Println(n)
	fmt.Println(string(n))

	var unmarshalCarType CarType
	err = json.Unmarshal([]byte(`"suv"`), &unmarshalCarType)
	if err != nil {
		panic(err)
	}

	fmt.Println(unmarshalCarType)
}
