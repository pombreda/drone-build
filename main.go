package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/drone/drone-plugin-go/plugin"
)

func main() {

	conf := struct {
		Commands []string `json:"commands"`
	}{}
	clone := plugin.Clone{}
	plugin.Param("clone", &clone)
	plugin.Param("vargs", &conf)
	plugin.Parse()

	var buf bytes.Buffer

	// script should change working dir to the
	// repository root directory
	cd := fmt.Sprintf("cd %s", clone.Dir)
	buf.WriteString(newline(cd))

	for _, c := range conf.Commands {
		buf.WriteString(trace(c))
		buf.WriteString(newline(c))
	}

	os.MkdirAll("/drone/bin", 0777)
	ioutil.WriteFile("/drone/bin/build.sh", buf.Bytes(), 0777)
}

func trace(s string) string {
	cmd := fmt.Sprintf("$ %s\n", s)
	encoded := base64.StdEncoding.EncodeToString([]byte(cmd))
	return fmt.Sprintf("echo %s | base64 --decode\n", encoded)
}

func newline(s string) string {
	return fmt.Sprintf("%s\n", s)
}
