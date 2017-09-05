package main

import (
	"github.com/FlashBoys/go-finance"
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestQuote(t *testing.T) {
	q, _ := finance.GetQuote("0005.HK")
	t.Log("0005.HK: %v", spew.Sdump(q))
}
