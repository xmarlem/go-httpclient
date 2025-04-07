package examples

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	// initialization

	// execution
	endpoints, err := GetEndpoints()

	// valdiation

	fmt.Println(err)
	fmt.Println(endpoints)

}
