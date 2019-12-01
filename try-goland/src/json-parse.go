package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type Role int

const (
	// _                    Role = iota
	cook                 Role = iota
	softwareEngineer     Role = iota
	manager              Role = iota
	chiefTechnicalOffice Role = iota
)

type Employee struct {
	Role Role   `json:"role"`
	// name string // `json:"name"`
}

func (r Role) string() (string, error) {

	roles := []string{
		"Cook",
		"Software engineer",
		"Manger",
		"Chief technical officer",
	}

	i := int(r) - 1

	if i >= 0 && i < len(roles) {
		return roles[i], nil
	}

	return "", errors.New("unknow role")
}

func (e Employee) MarshalJSON() ([]byte, error) {

	if rStr, rOk := e.Role.string(); rOk == nil {

		b := bytes.NewBufferString(`"`)

		fmt.Println(rStr)
		b.WriteString(rStr)

		b.WriteString(`"`)
		return b.Bytes(), nil
	} else {
		// fmt.Println(rOk)
	}

	return nil, errors.New("Couldn't marshal employee role")
}

func (g *Role) UnmarshalJSON(b []byte) error {
	var j string
	if err := json.Unmarshal(b, &j); err != nil {
		return err
	}
	*g = chiefTechnicalOffice
	return nil
}

func main() {

	fmt.Println(manager)
	fmt.Println(chiefTechnicalOffice)
	fmt.Println(softwareEngineer)
	fmt.Println(cook)

	e := Employee{name: "toufiq", Role: manager}

	if employeeJson, mOk := json.MarshalIndent(e, "", "	"); mOk == nil {
		fmt.Println(employeeJson)
		fmt.Println(string(employeeJson))
	} else {
		fmt.Println(mOk)
	}
}
