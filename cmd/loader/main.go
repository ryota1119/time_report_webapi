package main

import (
	"fmt"
	"io"
	"os"

	"github.com/ryota1119/time_resport/internal/schema"

	"ariga.io/atlas-provider-gorm/gormschema"
)

func main() {
	stmts, err := gormschema.New("mysql").Load(
		&schema.Organization{},
		&schema.User{},
		&schema.Customer{},
		&schema.Project{},
		&schema.Timer{},
		&schema.Budget{},
	)
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "failed to load gorm schemaspy: %v\n", err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
	_, err = io.WriteString(os.Stdout, stmts)
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "failed to write output: %v\n", err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
}
