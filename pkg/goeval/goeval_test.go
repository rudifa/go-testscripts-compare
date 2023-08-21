package goeval_test

import (
	"testing"

	"github.com/rudifa/testscripts-compare/pkg/goeval"
)

func TestMiniEval3J(t *testing.T) {

	schema, err := goeval.ReadString("testdata/combo.cue")
	if err != nil {
		t.Error(err)
	}
	data, err := goeval.ReadString("testdata/raw.json")
	if err != nil {
		t.Error(err)
	}
	expr := "#toApp"

	res, err := goeval.MiniEval3J(schema, data, expr)
	if err != nil {
		t.Error(err)
	}
	ppres, err := goeval.Prettyfmt(res)
	if err != nil {
		t.Error(err)
	}

	expect, err := goeval.ReadString("testdata/stdout.json")

	if err != nil {
		t.Error(err)
	}

	if ppres != expect {
		t.Errorf("got\n%s, want\n%s", ppres, expect)
	}
}
