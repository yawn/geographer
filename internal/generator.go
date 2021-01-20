package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"time"

	"github.com/dave/jennifer/jen"
	"github.com/yawn/geographer"
)

func main() {

	var buf bytes.Buffer

	if err := services(&buf); err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile("./generated.go", buf.Bytes(), 0644); err != nil {
		panic(err)
	}

}

func services(w io.Writer) error {

	r, err := geographer.GetServices(context.Background())

	if err != nil {
		return err
	}

	f := jen.NewFile("geographer")

	f.HeaderComment(fmt.Sprintf("Code generated on %s - DO NOT EDIT.", time.Now().Format(time.RFC3339)))

	m := r.Metadata

	f.HeaderComment(fmt.Sprintf("copyright %s", m.Copyright))
	f.HeaderComment(fmt.Sprintf("disclaimer %s", m.Disclaimer))
	f.HeaderComment(fmt.Sprintf("format-version %s", m.FormatVersion))
	f.HeaderComment(fmt.Sprintf("source-version %s", m.SourceVersion))

	f.Var().Id("Services").
		Op("=").
		Map(jen.String()).Id("Regions").Values(jen.DictFunc(func(m jen.Dict) {

		for service, regions := range r.Services() {

			m[jen.Lit(service)] = jen.Index().String().ValuesFunc(func(v *jen.Group) {

				for _, region := range regions {
					v.Lit(region)
				}

			})

		}

	}))

	return f.Render(w)

}
