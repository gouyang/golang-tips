package main

import (
	"fmt"
	"log"
	"reflect"
)

func UniqStr(col []string) []string {
	m := map[string]struct{}{}
	for _, v := range col {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
		}
	}
	list := make([]string, len(m))

	i := 0
	for v := range m {
		list[i] = v
		i++
	}
	return list
}

func main() {
	a := []struct{ in, out []string }{
		{[]string{"a", "c", "b", "a", "c"}, []string{"a", "c", "b"}},
		{[]string{"b", "c", "b", "c"}, []string{"b", "c"}},
	}

	for _, exp := range a {
		res := UniqStr(exp.in)
		if !reflect.DeepEqual(res, exp.out) {
			log.Fatal("%q did not match %q\n", res, exp.out)
		}
		for _, v := range res {
			fmt.Println(v)
		}
	}
}
