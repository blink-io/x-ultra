package generator

import (
	"path"
	"sort"
	"strconv"

	"github.com/blink-io/x/nats/cmd/protoc-gen-go-nats/internal/typemap"

	descriptor "google.golang.org/protobuf/types/descriptorpb"
)

func (g *generator) generateImports(file *descriptor.FileDescriptorProto) {
	if len(file.Service) == 0 {
		return
	}

	// stdlib imports
	g.P(`import `, g.pkgs["context"], ` "context"`)
	g.P(`import `, g.pkgs["json"], ` "encoding/json"`)
	g.P(`import `, g.pkgs["fmt"], ` "fmt"`)
	g.P(`import `, g.pkgs["http"], ` "net/http"`)
	g.P()

	// dependency imports
	g.P(`import `, g.pkgs["nats"], ` "github.com/nats-io/nats.go"`)
	g.P(`import `, g.pkgs["micro"], ` "github.com/nats-io/nats.go/micro"`)

	// It's legal to import a message and use it as an input or output for a
	// method. Make sure to import the package of any such message. First, dedupe
	// them.
	deps := make(map[string]string) // Map of package name to quoted import path.
	ourImportPath := path.Dir(g.goFileName(file))
	for _, s := range file.Service {
		for _, m := range s.Method {
			defs := []*typemap.MessageDefinition{
				g.reg.MethodInputDefinition(m),
				g.reg.MethodOutputDefinition(m),
			}
			for _, def := range defs {
				importPath, _ := parseGoPackageOption(def.File.GetOptions().GetGoPackage())
				if importPath == "" { // no option go_package
					importPath := path.Dir(g.goFileName(def.File)) // use the dirname of the Go filename as import path
					if importPath == ourImportPath {
						continue
					}
				}

				if substitution, ok := g.importMap[def.File.GetName()]; ok {
					importPath = substitution
				}
				importPath = g.importPrefix + importPath

				pkg := g.goPackageName(def.File)
				if pkg != g.genPkgName {
					deps[pkg] = strconv.Quote(importPath)
				}
			}
		}
	}

	pkgs := make([]string, 0, len(deps))
	for pkg := range deps {
		pkgs = append(pkgs, pkg)
	}
	sort.Strings(pkgs)
	for _, pkg := range pkgs {
		g.P(`import `, pkg, ` `, deps[pkg])
	}
	if len(deps) > 0 {
		g.P()
	}

}
