package main

import (
	"flag"
	"fmt"
	"os"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

var (
	showVersion     = flag.Bool("version", false, "print the version and exit")
	omitempty       = flag.Bool("omitempty", true, "omit if google.api is empty")
	omitemptyPrefix = flag.String("omitempty_prefix", "", "omit if google.api is empty")
	externTemplate  = flag.String("extern_template", "", "use external template to generate file")
	transportPath   = flag.String("transport_path", "", "Custom HTTP transport import path")
)

var debug = (func() bool {
	return os.Getenv("DEBUG") == "1"
})()

func main() {
	flag.Parse()
	if *showVersion {
		dprintf("invoke show version\n")
		fmt.Printf("protoc-gen-go-http %v\n", release)
		return
	}

	protogen.Options{
		ParamFunc: flag.CommandLine.Set,
	}.Run(func(gen *protogen.Plugin) error {
		etpl := *externTemplate
		dprintf("Template Path: %s\n", etpl)

		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

		for _, f := range gen.Files {

			if !f.Generate {
				continue
			}
			generateFile(gen, f, *omitempty, *omitemptyPrefix, etpl)
		}
		return nil
	})
}

func dprintf(format string, args ...any) {
	if debug {
		fmt.Fprintf(os.Stderr, format, args...)
	}
}
