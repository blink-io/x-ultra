package generator

import (
	"fmt"
	"strings"

	"github.com/blink-io/x/nats/cmd/protoc-gen-go-nats/internal/stringutils"
	"github.com/blink-io/x/nats/cmd/protoc-gen-go-nats/internal/typemap"
	descriptor "google.golang.org/protobuf/types/descriptorpb"
)

func (g *generator) generateService(file *descriptor.FileDescriptorProto, service *descriptor.ServiceDescriptorProto, index int) {
	serviceName := stringutils.CamelCase(service.GetName())

	g.sectionComment(serviceName + ` Interface`)
	g.generateInterface(file, service)

	g.sectionComment(serviceName + ` NATS Client`)
	g.generateClient(file, service)

	g.sectionComment(serviceName + ` Micro Service Handlers`)
	g.generateServer(file, service)
}

// Big header comments to makes it easier to visually parse a generated file.
func (g *generator) sectionComment(sectionTitle string) {
	g.P()
	g.P(`// `, strings.Repeat("=", len(sectionTitle)))
	g.P(`// `, sectionTitle)
	g.P(`// `, strings.Repeat("=", len(sectionTitle)))
	g.P()
}

func (g *generator) generateInterface(file *descriptor.FileDescriptorProto, service *descriptor.ServiceDescriptorProto) {
	serviceName := stringutils.CamelCase(service.GetName())
	comments, err := g.reg.ServiceComments(file, service)
	if err == nil {
		g.printComments(comments)
		g.P(`// `)
	}
	g.P(`// Service will use the "`, file.GetPackage(), `.`, strings.ToLower(serviceName), `" service group`)

	g.P(`type `, serviceName, ` interface {`)
	for _, method := range service.Method {
		comments, err := g.reg.MethodComments(file, service, method)
		if err == nil {
			g.printComments(comments)
			g.P(`// `)
		}
		g.P(`// Method subject will be "`, file.GetPackage(), `.`, strings.ToLower(serviceName), `.`, strings.ToLower(method.GetName()), `"`)

		g.P(g.generateSignature(method))
		g.P()
	}
	g.P(`}`)
}

func (g *generator) generateClient(file *descriptor.FileDescriptorProto, service *descriptor.ServiceDescriptorProto) {
	pkgName := file.GetPackage()
	serviceName := stringutils.CamelCase(service.GetName())
	clientTypeName := stringutils.PascalCase(service.GetName())

	g.P(`	   	func New`, serviceName, `Client(nc *nats.Conn) `, serviceName, ` {`)
	g.P(`	   		return &`, clientTypeName, `Client{nc: nc}`)
	g.P(`	   	}`)
	g.P(``)
	g.P(`	   var _ `, serviceName, ` = (*`, clientTypeName, `Client)(nil)`)
	g.P(``)
	g.P(`	   	type `, clientTypeName, `Client struct {`)
	g.P(`	   		nc *nats.Conn`)
	g.P(`	   	}`)
	g.P(``)

	for _, method := range service.Method {
		inputType := g.goTypeName(method.GetInputType())
		outputType := g.goTypeName(method.GetOutputType())
		outputVarName := stringutils.PascalCase(outputType)
		endpointName := strings.ToLower(method.GetName())

		g.P(`	   	func (impl *`, clientTypeName, `Client) `, method.GetName(), `(ctx context.Context, req *`, inputType, `) (*`, outputType, `, error) {`)
		g.P(`	   		data, _ := json.Marshal(req)`)
		g.P(`	   		resp, err := impl.nc.RequestWithContext(ctx, "`, pkgName, `.`, strings.ToLower(serviceName), `.`, endpointName, `", data)`)
		g.P(`	   		if err != nil {`)
		g.P(`	   			return nil, err`)
		g.P(`	   		}`)
		g.P()
		g.P(`			errCode, errMessage := resp.Header.Get(micro.ErrorCodeHeader), resp.Header.Get(micro.ErrorHeader)`)
		g.P(`			if errCode != "" && errMessage != "" {`)
		g.P(`				return nil, fmt.Errorf("%s (%s)", errMessage, errCode)`)
		g.P(`			}`)
		g.P()
		g.P(`	   		var `, outputVarName, ` `, outputType)
		g.P(`	   		if err := json.Unmarshal(resp.Data, &`, outputVarName, `); err != nil {`)
		g.P(`	   			return nil, err`)
		g.P(`	   		}`)
		g.P()
		g.P(`	   		return &`, outputVarName, `, nil`)
		g.P(`	   	}`)
		g.P()
	}
	g.P()
}

func (g *generator) generateServer(file *descriptor.FileDescriptorProto, service *descriptor.ServiceDescriptorProto) {
	pkgName := file.GetPackage()
	serviceName := stringutils.CamelCase(service.GetName())
	microServiceName := strings.ToLower(strings.ReplaceAll(pkgName+"."+serviceName, ".", "-"))

	g.P(`// New`, serviceName, `Server builds a new micro.Service that will be registered with the instance provided`)
	g.P(`// Each RPC on the service will be mapped to a new endpoint within the micro service`)
	g.P(`func New`, serviceName, `Server(rootCtx `, g.pkgs["context"], `.Context, version string, nc *`, g.pkgs["nats"], `.Conn, impl `, serviceName, `) (`, g.pkgs["micro"], `.Service, error) {`)
	g.P(`	svc, err := micro.AddService(nc, micro.Config{`)
	g.P(`		Name:        `, fmt.Sprintf("%q", microServiceName), `,`) // escape the quotes if there are any
	g.P(`		Version:     version,`)
	comments, err := g.reg.ServiceComments(file, service)
	if err == nil {
		g.P(`		Description: "`, getComments(comments), `",`)
	}
	g.P(`		Metadata: map[string]string{`)
	g.P(`			"package": "`, pkgName, `",`)
	g.P(`			"name":    "`, serviceName, `",`)
	g.P(`		},`)
	g.P(`	})`)
	g.P()
	g.P(`	if err != nil {`)
	g.P(`		return nil, fmt.Errorf("failed to create nats service: %w", err)`)
	g.P(`	}`)
	g.P()
	g.P(`	group := svc.AddGroup("`, pkgName, `.`, strings.ToLower(serviceName), `")`)
	g.P()

	for _, method := range service.Method {
		inputType := g.goTypeName(method.GetInputType())
		outputType := g.goTypeName(method.GetOutputType())
		intpuVarName := stringutils.PascalCase(inputType)
		outputVarName := stringutils.PascalCase(outputType)
		endpointName := strings.ToLower(method.GetName())

		g.P(`	if err := group.AddEndpoint("`, endpointName, `", micro.ContextHandler(rootCtx, func(ctx context.Context, req micro.Request) {`)
		g.P(`		var `, intpuVarName, ` `, inputType, ``)
		g.P(`		if err := json.Unmarshal(req.Data(), &`, intpuVarName, `); err != nil {`)
		g.P(`			_ = req.Error(http.StatusText(http.StatusBadRequest), err.Error(), nil)`)
		g.P(`			return`)
		g.P(`		}`)
		g.P()
		g.P(`		`, outputVarName, `, err := impl.`, method.GetName(), `(ctx, &`, intpuVarName, `)`)
		g.P(`		if err != nil {`)
		g.P(`			_ = req.Error(http.StatusText(http.StatusInternalServerError), err.Error(), nil)`)
		g.P(`			return`)
		g.P(`		}`)
		g.P()
		g.P(`		_ = req.RespondJSON(`, outputVarName, `)`)
		g.P(`	})); err != nil {`)
		g.P(`		return nil, fmt.Errorf("failed to add endpoint %q: %w", "`, endpointName, `", err)`)
		g.P(`	}`)
		g.P()
	}

	g.P(`	return svc, nil`)
	g.P(`}`)
	g.P()
}

func (g *generator) generateSignature(method *descriptor.MethodDescriptorProto) string {
	methName := stringutils.CamelCase(method.GetName())
	inputType := g.goTypeName(method.GetInputType())
	outputType := g.goTypeName(method.GetOutputType())
	return fmt.Sprintf(`	%s(%s.Context, *%s) (*%s, error)`, methName, g.pkgs["context"], inputType, outputType)
}

func (g *generator) printComments(comments typemap.DefinitionComments) bool {
	text := strings.TrimSuffix(comments.Leading, "\n")
	if len(strings.TrimSpace(text)) == 0 {
		return false
	}
	split := strings.Split(text, "\n")
	for _, line := range split {
		g.P("// ", strings.TrimPrefix(line, " "))
	}
	return len(split) > 0
}

func getComments(comments typemap.DefinitionComments) string {
	text := strings.TrimSuffix(comments.Leading, "\n")
	if len(strings.TrimSpace(text)) == 0 {
		return ""
	}
	split := strings.Split(text, "\n")
	return strings.Join(split, " ")
}
