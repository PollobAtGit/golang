package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Gender int

const (
	GenderNotSet Gender = iota
	GenderMale   Gender = iota
	GenderFemale Gender = iota
	GenderOther  Gender = iota
)

var toString = map[Gender]string{
	GenderNotSet: "Not Set",
	GenderMale:   "Male",
	GenderFemale: "Female",
	GenderOther:  "Other",
}

func (g Gender) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)

	fmt.Println(toString[g])

	buffer.WriteString(toString[g])
	buffer.WriteString(`"`)

	t := buffer.Bytes()

	fmt.Println(string(t))

	return t, nil
}

type Human struct {
	Gender Gender `json:"gender"`
}

func main() {
	me := Human{Gender: GenderMale}
	prettyJSON, _ := json.MarshalIndent(me, "", "    ")
	fmt.Println(string(prettyJSON))
}
