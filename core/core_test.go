package core

import (
	"fmt"
	"testing"
	"context"
)

func Test_CreateDB(t *testing.T) {
	c, _ := NewCore(context.Background())

	ct, _ := c.CreateDB(Container{
		Name:	  "teste-sql",
		Replicas: 2,
	})

	fmt.Println(ct)
}
