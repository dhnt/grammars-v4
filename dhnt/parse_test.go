package main

import (
	"testing"
	"path/filepath"
	"os"
	"fmt"
	"path"
	"sort"
)

func listFile(base string) []string {

	list := make([]string, 0, 10)

	err := filepath.Walk(base, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".jsn" || filepath.Ext(path) == ".json" {
			list = append(list, path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("walk error [%v]\n", err)
	}
	return list
}

//
func TestValues(t *testing.T) {
	cwd, _ := os.Getwd()
	base := path.Join(cwd, "/examples/values")

	list := listFile(base)

	sort.Strings(list)

	for _, f := range list {
		fmt.Println("Parsing file:", f)
		Parse(f)
	}
}

//json
func TestTestJson(t *testing.T) {
	_, err := ParseFile("test.json")

	if err != nil {
		t.Errorf("%v", err)
	}
}

//dhnt
func TestAssign(t *testing.T) {
	_, err := ParseFile("assign.jsn")

	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestBlock(t *testing.T) {
	_, err := ParseFile("block.jsn")

	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestChan(t *testing.T) {
	_, err := ParseFile("chan.jsn")

	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestDhnt(t *testing.T) {
	_, err := ParseFile("dhnt.jsn")

	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestError(t *testing.T) {
	_, err := ParseFile("error.jsn")

	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestExpr(t *testing.T) {
	_, err := ParseFile("expr.jsn")

	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestFunc(t *testing.T) {
	_, err := ParseFile("func.jsn")

	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestHelloworld(t *testing.T) {
	_, err := ParseFile("helloworld.jsn")

	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestJump(t *testing.T) {
	_, err := ParseFile("jump.jsn")

	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestKind(t *testing.T) {
	_, err := ParseFile("kind.jsn")

	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestLang(t *testing.T) {
	_, err := ParseFile("lang.jsn")

	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestMini(t *testing.T) {
	_, err := ParseFile("mini.jsn")

	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestObject(t *testing.T) {
		_, err := ParseFile("object.jsn")

		if err != nil {
				t.Errorf("%v", err)
		}
}

func TestRange(t *testing.T) {
	_, err := ParseFile("range.jsn")

	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestStmt(t *testing.T) {
	_, err := ParseFile("stmt.jsn")

	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestTimer(t *testing.T) {
	_, err := ParseFile("timer.jsn")

	if err != nil {
		t.Errorf("%v", err)
	}
}