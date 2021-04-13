package core

import "strconv"

type Argument string

func (a Argument) AsString() string {
	return string(a)
}

func (a Argument) AsInt() (int, error) {
	return strconv.Atoi(a.AsString())
}

func (a Argument) AsBool() (bool, error) {
	return strconv.ParseBool(a.AsString())
}

type ArgumentList []Argument
