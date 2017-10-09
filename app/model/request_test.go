package model

import (
	"fmt"
	"testing"
)

func TestGetAllTokenRequests(t *testing.T) {
	reqs, err := GetAllTokenRequests("5-Q9IV8pp")
	if err != nil {
		t.Error(err)
	}
	for i := 0; i < len(reqs); i++ {
		fmt.Printf("Headers: %s", reqs[i].ReqHeaders)
	}
}
