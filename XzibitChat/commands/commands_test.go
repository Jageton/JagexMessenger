package commands

import (
	"github.com/mitchellh/mapstructure"
	"reflect"
	"testing"
)

type TestCommand struct {
	IsOK bool
	Msg  string
}

func TestNotFullParseWithData(t *testing.T) {
	actual := &TestCommand{}
	mp := map[string]interface{}{
		"IsOK": true,
	}
	expected := &TestCommand{
		IsOK: true,
		Msg: "Hello",
	}
	err := mapstructure.Decode(mp, &actual)
	if err != nil {
		t.Error(err)
	}
	if reflect.DeepEqual(expected, actual) {
		t.Error("Expected ", expected, "got ", actual)
	}
}

func TestNotFullParse(t *testing.T) {
	actual := &TestCommand{}
	mp := map[string]interface{}{
		"IsOK": true,
	}
	expected := &TestCommand{
		IsOK: true,
	}
	err := mapstructure.Decode(mp, &actual)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Error("Expected ", expected, "got ", actual)
	}
}

func TestParse(t *testing.T) {
	actual := &TestCommand{}
	mp := map[string]interface{}{
		"IsOK": true,
		"Msg": "OK",
	}
	expected := &TestCommand{
		IsOK: true,
		Msg:  "OK",
	}
	err := mapstructure.Decode(mp, &actual)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Error("Expected ", expected, "got ", actual)
	}
}
