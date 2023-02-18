//go:build ignore

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/hedwigz/entviz"
)

func main() {

	/*

		spec := new(ogen.Spec)
		oas, err := entoas.NewExtension(entoas.Spec(spec))
		if err != nil {
			log.Fatalf("creating entoas extension: %v", err)
		}
		ogent, err := ogent.NewExtension(spec)
		if err != nil {
			log.Fatalf("creating ogent extension: %v", err)
		}
	*/

	opts := []entc.Option{
		entc.FeatureNames("schema/snapshot"), //"privacy", - will add down the line.
		entc.Extensions(entviz.Extension{}),  //ogent, oas
	}
	err := entc.Generate("./schema", &gen.Config{}, opts...)
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
