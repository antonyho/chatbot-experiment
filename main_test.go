package main

import (
	"fmt"
	"github.com/FlashBoys/go-finance"
	"github.com/davecgh/go-spew/spew"
	"strings"
	"testing"
)

func TestQuote(t *testing.T) {
	q, _ := finance.GetQuote(fmt.Sprintf("%s.HK", strings.TrimSpace("0005")))
	t.Logf("0005.HK: %v", spew.Sdump(q))
}
