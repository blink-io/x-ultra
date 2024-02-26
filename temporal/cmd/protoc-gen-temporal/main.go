package main

import (
	_ "embed"
	"html/template"
	"log"
	"os"

	"google.golang.org/protobuf/compiler/protogen"
)

//go:embed go.tpl
var serviceData []byte
var serviceTemplate = template.Must(template.New("").Parse(string(serviceData)))

func GenService(file *protogen.GeneratedFile, service *protogen.Service) error {
	log.Print("----- BEGIN SERVICE ", service.GoName, " -----")
	if err := serviceTemplate.Execute(file, service); err != nil {
		return err
	}
	log.Print("----- END SERVICE ", service.GoName, " -----")
	return nil
}

type TemplateInput struct {
	File    *protogen.File
	GenFile *protogen.GeneratedFile
}

func GenFile(plu *protogen.Plugin, file *protogen.File) error {
	fileName := file.GeneratedFilenamePrefix + ".temporal.go"
	genFile := plu.NewGeneratedFile(fileName, file.GoImportPath)
	log.Print("----- BEGIN FILE ", file.Desc.Path(), " -----")
	defer log.Print("----- END FILE ", file.Desc.Path(), " -----")
	if err := serviceTemplate.Execute(genFile, TemplateInput{
		File:    file,
		GenFile: genFile,
	}); err != nil {
		return err
	}

	// genFile.P("// Code generated by protoc-gen-temporal. DO NOT EDIT.")
	// genFile.P("// source: ", *file.Proto.Name)
	// genFile.P("package ", file.GoPackageName)
	// for _, service := range file.Services {
	// 	if err := GenService(genFile, service); err != nil {
	// 		return err
	// 	}
	// }
	return nil
}

func Generate(plu *protogen.Plugin) error {
	log.Print("----- BEGIN PLUGIN -----")
	defer log.Print("----- END PLUGIN -----")
	for _, file := range plu.Files {
		if !file.Generate {
			continue
		}
		if err := GenFile(plu, file); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	log.SetOutput(os.Stderr)
	// log.SetOutput(ioutil.Discard)
	protogen.Options{}.Run(Generate)
}