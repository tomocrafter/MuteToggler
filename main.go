package main

import (
	"github.com/toqueteos/ts3"
	"os"
	"strings"
)

func main() {
	conn, _ := ts3.Dial(":25639", true)
	defer conn.Close()

	fp, err := os.Open("apikey")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	buf := make([]byte, 29)
	fp.Read(buf)

	conn.Cmd("auth apikey=" + string(buf))
	r, _ := conn.Cmd("whoami")

	clid := Parse(r)["clid"]
	muted, _ := conn.Cmd("clientvariable clid=" + clid + " client_input_muted")
	isMuted := Parse(muted)["client_input_muted"] == "1"

	if isMuted {
		conn.Cmd("clientupdate client_input_muted=0")
	} else {
		conn.Cmd("clientupdate client_input_muted=1")
	}
}

func Parse(result string) (map[string]string) {
	m := map[string]string{}
	split := strings.Split(result, " ")
	for _, v := range split {
		aa := strings.Split(v, "=")
		m[aa[0]] = aa[1]
	}
	return m
}
