package models

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"reflect"
	"testing"
)

func setConfigDatabase() {
	var err error
	err = godotenv.Load(os.ExpandEnv("../../../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
}

func assertEqual(t testing.TB, got interface{}, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("Got %v want %v", got, want)
	}
}

func assertResult(t testing.TB, got interface{}, want interface{}) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v want %v", got, want)
	}
}
