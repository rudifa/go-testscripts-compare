package goeval

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
)

// MiniEval3J compiles a schema and data and returns compact JSON selected by expr
func MiniEval3J(schema2, data, expr string) (string, error) {

	ctx := cuecontext.New()

	v2 := ctx.CompileString(schema2)
	if err := v2.Err(); err != nil {
		return "", err
	}

	v3 := ctx.CompileString(data, cue.Scope(v2))
	if err := v3.Err(); err != nil {
		return "", err
	}

	res := v2.Unify(v3).LookupPath(cue.ParsePath(expr))
	if err := res.Err(); err != nil {
		return "", err
	}

	bytes, err := res.MarshalJSON()
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func TestMiniEval3J() {
	schema, _ := ReadString("data/combo.cue")
	data, _ := ReadString("data/raw.json")
	expr := "#toApp"
	res, err := MiniEval3J(schema, data, expr)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	ppres, _ := Prettyfmt(res)
	fmt.Println(ppres)
}

func ReadString(filepath string) (string, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func Prettyfmt(input string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(input), "", "\t"); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}
