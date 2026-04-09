package movistarapi

import (
	"fmt"
	"testing"
)

func TestEnsureSyntax(t *testing.T) {
	hgu, err := HGULogin("password123")
	if err != nil {
		t.Fatal(err)
	}

	localMap, err := hgu.LocalMap()
	if err != nil {
		return
	}

	fmt.Println(localMap)
}
