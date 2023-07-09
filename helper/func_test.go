package helper

import (
	"fmt"
	"testing"
)

func TestWorkDir(t *testing.T) {
	got, err := WorkDir()
	fmt.Println(got)
	fmt.Println(err)
}
