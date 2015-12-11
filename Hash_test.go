package Password

import (
	"log"
	"testing"
)

func TestHash(t *testing.T) {
	var hash, err = Hash("abc", PASSWORD_DEFAULT)

	if err != nil {
		log.Print(err)
		t.Fail()
	}
	if hash[:3] != "$2y" {
		t.Fail()
	}
}
