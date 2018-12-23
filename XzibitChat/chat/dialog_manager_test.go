package chat

import (
	"reflect"
	"testing"
)

func TestHubManager_Add(t *testing.T) {
	hubm := NewHubManager()
	hubm.Add(1, &Dialog{})
	if count := hubm.Count(); count != 1 {
		t.Error("Expected 1 got ", count)
	}
	hubm.Add(2, &Dialog{})
	if count := hubm.Count(); count != 2 {
		t.Error("Expected 2 got ", count)
	}
}

func TestHubManager_AddExisted(t *testing.T) {
	hubm := NewHubManager()
	hubm.Add(1, &Dialog{})
	if count := hubm.Count(); count != 1 {
		t.Error("Expected 1 got ", count)
	}
	err := hubm.Add(1, &Dialog{})
	if err == nil {
		t.Error("Added existed hub")
	}
}

func TestHubManager_Contains(t *testing.T) {
	hubm := NewHubManager()
	hubm.Add(1, &Dialog{})
	if !hubm.Contains(1) {
		t.Error("Not found hub with id: 1")
	}
}


func TestHubManager_Remove(t *testing.T) {
	hubm := NewHubManager()
	hubm.Add(1, &Dialog{})
	if !hubm.Remove(1) {
		t.Error("Error with removing hub")
	}
	if hubm.Remove(1) {
		t.Error("Not existed hub is deleted")
	}
}

func TestHubManager_Get(t *testing.T) {
	hubm := NewHubManager()
	hub := &Dialog{}
	hubm.Add(1, hub)
	h, ok := hubm.Get(1)
	if !ok {
		t.Error("Can't get hub")
	}
	if !reflect.DeepEqual(hub, h) {
		t.Error("Hubs are not equal")
	}
}
