package main

import (
	"testing"
    "reflect"
)
func TestEqualPlayer(t *testing.T) {
    expected := Player{
        name: "itsme",
        hp: 100,
    }
    have := Player {
        name: "itsme",
        hp: 100,
    }
    if !reflect.DeepEqual(expected, have){
        t.Errorf("expected %+v but got %+v",expected,have)
    }
}
