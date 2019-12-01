package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type Role int

type Employee struct {
	JobRole Role   `json:"job_role"`
	Name    string `json:"name"`
}

const (
	_    Role = iota
	Cook Role = iota
	CTO  Role = iota
)

var roleMap = map[Role]string{
	Cook: "Cook",
	CTO:  "Chief technical officer",
}

func (e Role) MarshalJSON() ([]byte, error) {

	if v, ok := roleMap[e]; ok {

		b := bytes.NewBufferString(`"`)

		b.WriteString(v)
		b.WriteString(`"`)

		return b.Bytes(), nil
	}

	return nil, errors.New("Can't parse to json for " + string(e))
}

func main() {

	e := Employee{JobRole: CTO, Name: "tim"}

	if prettyJson, ok := json.MarshalIndent(e, "", "	"); ok == nil {
		fmt.Println(string(prettyJson))
	} else {
		fmt.Println(ok)
	}
}
