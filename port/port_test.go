package port

import (
	"fmt"
	"testing"
)

func Test_Port(t *testing.T) {
	pis, err := PortInfo("27017")
	if err != nil {
		t.Fail()
	}

	fmt.Println(pis)
}
