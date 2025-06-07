package main

import "github.com/zguydev/openapi-fixer/pkg/fixup"

var Fixups = []fixup.OpenAPIFixup{
	&DiscriminatorFixup{},
	&NumberTypeFixup{},
	&EnumFixup{},
}
