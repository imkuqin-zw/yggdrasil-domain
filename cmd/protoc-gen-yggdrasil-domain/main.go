package main

import (
	"flag"

	"github.com/imkuqin-zw/yggdrasil-domain/internal/gendomain"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	flag.Parse()
	var flags flag.FlagSet
	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			if err := gendomain.GenerateFile(gen, f); err != nil {
				return err
			}
		}
		return nil
	})
}
