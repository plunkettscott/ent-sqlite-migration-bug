package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	err := entc.Generate("./schema", &gen.Config{
		Schema:  "./schema",
		Target:  "../internal/ent",
		Package: "entgo.io/bug/internal/ent",
		Features: []gen.Feature{
			gen.FeatureEntQL,
			gen.FeaturePrivacy,
			gen.FeatureLock,
			gen.FeatureUpsert,
			gen.FeatureModifier,
			gen.FeatureSnapshot,
		},
	})

	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
