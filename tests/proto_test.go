package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/emicklei/proto"
	"github.com/jhump/protoreflect/dynamic"
	"github.com/stretchr/testify/require"
)

func TestErr1(t *testing.T) {
	extgry := dynamic.NewExtensionRegistryWithDefaults()
	require.NotNil(t, extgry)
}

func TestProto_2(t *testing.T) {
	reader, _ := os.Open("./testdata/i18n.proto")
	defer reader.Close()

	parser := proto.NewParser(reader)
	definition, err := parser.Parse()
	require.NoError(t, err)

	proto.Walk(definition,
		proto.WithService(handleService),
		proto.WithMessage(handleMessage))
}

func handleService(s *proto.Service) {
	fmt.Println(s.Name)
}

func handleMessage(m *proto.Message) {
	lister := new(optionLister)
	for _, each := range m.Elements {
		each.Accept(lister)
	}
	fmt.Println(m.Name)
}

type optionLister struct {
	proto.NoopVisitor
}

func (l optionLister) VisitOption(o *proto.Option) {
	fmt.Println(o.Name)
}
