package wshub_test

import (
	"github.com/user/2019_1_newTeam2/pkg/wshub"
	"reflect"
	"testing"
)

func TestNewClientsMap(t *testing.T) {
	newMap := wshub.NewClientsMap()
	if newMap == nil{
		t.Errorf("map wasn't created")
	}
}

func TestClientsMap_Delete(t *testing.T) {
	newMap := wshub.NewClientsMap()
	newMap.Store(&wshub.Client{ID:1})
	paniced := false
	defer func() {
		paniced = true
		recover()
	}()
	newMap.Delete(1)
	if paniced {
		t.Error("your map paniced")
	}
}

func TestClientsMap_Store(t *testing.T) {
	newMap := wshub.NewClientsMap()
	newMap.Store(&wshub.Client{ID:1})
	cl, ok := newMap.Load(1)
	if !ok {
		t.Error("ur map failed to load")
	}
	if !reflect.DeepEqual(cl, &wshub.Client{ID:1}) {
		t.Errorf("don't match %v and %v", cl, &wshub.Client{ID:1})
	}
}


