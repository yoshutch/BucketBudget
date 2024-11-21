package models

import (
	"errors"
	"strconv"
	"strings"
)

type Amount struct {
	Cents int64
}

var (
	InvalidAmountParseError = errors.New("Invalid string to parse into an amount")
)

func ParseAmountFromString(value string) (Amount, error) {
	if !strings.Contains(value, ".") {
		return Amount{}, InvalidAmountParseError
	}
	value = strings.ReplaceAll(value, ".", "")
	num, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return Amount{}, InvalidAmountParseError
	}
	return Amount{Cents: num}, nil
}

func NewAmount(dollars, cents int64) (Amount, error) {
	// validate stuff
	return Amount{
		Cents: cents,
	}, nil
}

func (a Amount) Add(other Amount) Amount {
	return Amount{
		Cents: a.Cents + other.Cents,
	}
}
