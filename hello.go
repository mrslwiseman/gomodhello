package hello

import (
	"strings"
)

var version string = "2.0.0"

func SayHello(name string) string {
	return strings.Join([]string{"Hello", name, version}, " ")
}
