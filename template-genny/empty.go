package template_text

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "TypeT=string,int"

import "github.com/cheekybits/genny/generic"

type TypeT generic.Type

type EmptyTypeT struct {}

func (* EmptyTypeT) Filter(in map[string]TypeT) map[string]TypeT {
	return in
}

