package address

import (
	"log"
	"testing"
)

func TestAddress(t *testing.T) {
	log.Println(GetLocalAddress())
}
