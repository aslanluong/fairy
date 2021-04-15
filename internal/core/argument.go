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

func (al ArgumentList) Get(index int) Argument {
	if index < 0 || index >= len(al) {
		return Argument("")
	}

	return Argument(al[index])
}

func (al ArgumentList) IndexOf(value string) int {
	for i, arg := range al {
		if Argument(value) == arg {
			return i
		}
	}

	return -1
}

func (al ArgumentList) Contains(value string) bool {
	return al.IndexOf(value) > -1
}

func (al ArgumentList) Splice(index, deleteCount int) ArgumentList {
	l := len(al)
	if index >= l {
		return al
	}

	if index+deleteCount >= l {
		return al[:index]
	}

	return append(al[:index], al[index+deleteCount:]...)
}
