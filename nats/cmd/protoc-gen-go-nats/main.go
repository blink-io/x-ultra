package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/blink-io/x/nats/cmd/protoc-gen-go-nats/generator"

	"google.golang.org/protobuf/proto"
	plugin "google.golang.org/protobuf/types/pluginpb"
)

func main() {
	versionFlag := flag.Bool("version", false, "print version and exit")
	flag.Parse()

	if *versionFlag {
		fmt.Println(generator.GEN_VERSION)
		os.Exit(0)
	}

	req := readGeneratorRequest(os.Stdin)
	resp, err := generator.Generate(req)
	if err != nil {
		logErr(err, "generating response")
	}

	writeGeneratorResponse(os.Stdout, resp)
}

func readGeneratorRequest(r io.Reader) *plugin.CodeGeneratorRequest {
	data, err := io.ReadAll(r)
	if err != nil {
		logErr(err, "reading input")
	}

	var req plugin.CodeGeneratorRequest
	if err := proto.Unmarshal(data, &req); err != nil {
		logErr(err, "parsing input proto")
	}

	if len(req.FileToGenerate) == 0 {
		logFail("no files to generate")
	}

	return &req
}

func writeGeneratorResponse(w io.Writer, resp *plugin.CodeGeneratorResponse) {
	data, err := proto.Marshal(resp)
	if err != nil {
		logErr(err, "marshaling response")
	}

	_, err = w.Write(data)
	if err != nil {
		logErr(err, "writing response")
	}
}

func logFail(msgs ...string) {
	s := strings.Join(msgs, " ")
	log.Print("error:", s)
	os.Exit(1)
}

func logErr(err error, msgs ...string) {
	s := strings.Join(msgs, " ") + ":" + err.Error()
	log.Print("error:", s)
	os.Exit(1)
}
