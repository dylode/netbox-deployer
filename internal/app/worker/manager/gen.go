//go:build ignore

package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"reflect"
	"text/template"

	"dylaan.nl/netbox-deployer/internal/pkg/netbox"
	"github.com/gookit/goutil/dump"
)

const outputName = "detector_gen.go"

func main() {
	var buf bytes.Buffer

	fmt.Fprint(&buf, "// Code generated by \"gen.go\"; DO NOT EDIT.\n\n")
	fmt.Fprint(&buf, "package manager\n\n")
	fmt.Fprint(&buf, "import \"dylaan.nl/netbox-deployer/internal/pkg/netbox\"\n")

	genAllModelNames(&buf)
	fmt.Fprint(&buf, "\n\n")
	genHasComponent(&buf)

	err := os.WriteFile(outputName, buf.Bytes(), 0644)
	if err != nil {
		log.Fatalf("writing output: %s", err)
	}
}

func getUniqueModelTags(v any, uniqueTags map[string]struct{}) {
	val := reflect.ValueOf(v)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return
	}

	t := val.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("model")

		if tag != "" {
			uniqueTags[tag] = struct{}{}
		}

		fieldValue := val.Field(i)

		switch fieldValue.Kind() {
		case reflect.Struct:
			getUniqueModelTags(fieldValue.Interface(), uniqueTags)
		case reflect.Slice:
			for j := 0; j < fieldValue.Len(); j++ {
				getUniqueModelTags(fieldValue.Index(j).Interface(), uniqueTags)
			}
		}
	}
}

func allModelNames() []string {
	uniqueTags := make(map[string]struct{})
	getUniqueModelTags(netbox.VirtualMachine{}, uniqueTags)

	result := make([]string, 0, len(uniqueTags))
	for tag := range uniqueTags {
		result = append(result, tag)
	}
	return result
}

var allModelNamesTemplate = `
var allNetboxModelNames []netbox.ModelName

func init() {
	allNetboxModelNames = []netbox.ModelName {
		{{- range . }}
		netbox.ModelName("{{ . }}"),
		{{- end }}
	}
}`

func genAllModelNames(buf *bytes.Buffer) {
	tmpl, err := template.New("modelNames").Parse(allModelNamesTemplate)
	if err != nil {
		panic(err)
	}

	var tmplBuf bytes.Buffer
	err = tmpl.Execute(&tmplBuf, allModelNames())
	if err != nil {
		panic(err)
	}

	buf.Write(tmplBuf.Bytes())
}

var hasComponentTemplate = `
{{- define "node" -}}
{{- if .Children -}}
{{- range $child := .Children }}
{{- if $child.Slice }}
for _, {{ $child.Name }} := range {{ .Path }} {
	{{- range $subChild := $child.Children }}
		{{- template "node" $subChild }}
	{{- end }}
}
{{- end -}}
{{- if $child.Struct }}
	{{- range $subChild := $child.Children }}
		{{- template "node" $subChild }}
	{{- end }}
{{- end -}}
{{- end -}}
{{- else if eq .Name "ID" }}
if event.ModelName == "{{ .Model }}" && event.ModelID == {{ .Path }} {
	return true
}
{{- end -}}
{{- end -}}
func hasComponent(vm netbox.VirtualMachine, event netbox.WebhookEvent) bool {
	{{- template "node" . }}
	return false
}
`

func genHasComponent(buf *bytes.Buffer) {
	tmpl, err := template.
		New("hasComponent").
		Parse(hasComponentTemplate)
	if err != nil {
		panic(err)
	}

	var tmplBuf bytes.Buffer
	err = tmpl.Execute(&tmplBuf, getVirtualMachineComponents())
	if err != nil {
		panic(err)
	}

	buf.Write(tmplBuf.Bytes())
}

type node struct {
	Root     bool
	Slice    bool
	Struct   bool
	Name     string
	Model    string
	Parent   *node
	Path     string
	Children []*node
}

func getVirtualMachineComponents() *node {
	var walk func(t reflect.Type, parent *node)
	walk = func(t reflect.Type, parent *node) {
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)

			switch field.Type.Kind() {
			case reflect.Slice:
				child := &node{
					Name:   field.Name,
					Slice:  true,
					Parent: parent,
					Model:  field.Tag.Get("model"),
					Path:   fmt.Sprintf("%s.%s", parent.Path, field.Name),
				}

				walk(field.Type.Elem(), child)
				parent.Children = append(parent.Children, child)
			case reflect.Struct:
				child := &node{
					Name:   field.Name,
					Struct: true,
					Parent: parent,
					Model:  field.Tag.Get("model"),
					Path:   fmt.Sprintf("%s.%s", parent.Path, field.Name),
				}

				if parent.Slice {
					child.Path = fmt.Sprintf("%s.%s", parent.Name, field.Name)
				}
				walk(field.Type, child)
				parent.Children = append(parent.Children, child)
			default:
				if field.Name != "ID" {
					continue
				}

				child := &node{
					Name:   field.Name,
					Parent: parent,
					Model:  parent.Model,
					Path:   fmt.Sprintf("%s.%s", parent.Path, field.Name),
				}
				if parent.Slice {
					child.Path = fmt.Sprintf("%s.%s", parent.Name, field.Name)
				}

				parent.Children = append(parent.Children, child)
			}
		}
	}

	root := &node{Root: true, Name: "vm", Path: "vm"}
	walk(reflect.ValueOf(netbox.VirtualMachine{}).Type(), root)

	dump.P(root)
	return root
}
