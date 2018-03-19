package main

import (
	"fmt"
	"os"
)

//string: ""
type _string  string
type _boolean  bool
type _true  _boolean
type _false  _boolean

type _func func(string, int) (int)


type _1 struct {
		_string _string
		_true   _boolean
		_false  _boolean
		_struct struct {
			x string
			b int
				sss struct {
					a bool
				    }
			}
}

func (r *_1) _unique(_s _string) _boolean {
		//init x = {}, declare x = {} string
		_map := make(map[interface{}]interface{})
		for _, _char := range _s {
				key := _char
				_, ok := _map[key]
				fmt.Printf("%v \n", _char)
				if ok {
						return r._false; }
				_map[key] = _char

		}
		return r._true;
}

func (r *_1) _main(_args []_string) {
		if len(_args) == 0 {
				_args = []_string{"abcd", "abcdcd "}
		}
		for _i, _v := range _args {
				fmt.Printf("str: %v", _v)
				fmt.Println(_i, ":", r._unique(_v))
		}
}

func main() {
	t := _1{_string: "", _true: true, _false: false}

	t._main(os.Args[1:])
}

/**/