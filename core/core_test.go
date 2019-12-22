package core

import (
	"fmt"
	"testing"
	"context"
)

func Test_CreateDB(t *testing.T) {
	c, _ := NewCore(context.Background(), CoreConfig{PathVolume: "/data/SQLServer/"})

	ct, _ := c.CreateDB(Container{
		Organization: "lucas",
		Name:	      "teste-sql",
	})

	fmt.Println(ct)
}
