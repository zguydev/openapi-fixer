package main

import "github.com/zguydev/openapi-fixer/internal/fixer"

var Fixups = []fixer.OpenAPIFixup{
	&DiscriminatorFixup{},
	&NumberTypeFixup{},
	&EnumFix{},
}
