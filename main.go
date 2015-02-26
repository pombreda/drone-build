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
	conf := plugin.Config{}
	clone := plugin.Clone{}
	plugin.Param("clone", &clone)
	plugin.Param("config", &conf)
	plugin.Parse()

	var buf bytes.Buffer

	// nativate to the repository root dir
	// TODO: should we just set WORKDIR in the build container?
	cd := fmt.Sprintf("cd %s", clone.Dir)
	buf.WriteString(newline(cd))

	// TODO: should we just pass the into the container with -e?
	for _, c := range conf.Env {
		exp := fmt.Sprintf("export %s", c)
		buf.WriteString(newline(exp))
	}

	for _, c := range conf.Script {
		buf.WriteString(trace(c))
		buf.WriteString(newline(c))
	}

	os.MkdirAll("/tmp", 0777)
	ioutil.WriteFile("/tmp/build.sh", buf.Bytes(), 0777)
}

func trace(s string) string {
	cmd := fmt.Sprintf("$ %s\n", s)
	encoded := base64.StdEncoding.EncodeToString([]byte(cmd))
	return fmt.Sprintf("echo %s | base64 --decode\n", encoded)
}

func newline(s string) string {
	return fmt.Sprintf("%s\n", s)
}
