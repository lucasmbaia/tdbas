package operations

import (
	"testing"
)

func Test_CreateContainer(t *testing.T) {
	if _, err := CreateContainer("teste-sql-3"); err != nil {
		t.Fatal(err)
	}
}
