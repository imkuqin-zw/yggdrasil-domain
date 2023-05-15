package gendomain

import (
	"bytes"
	"strings"
	"text/template"
)

var domainTemplate = `
{{range .Events}}

const Topic{{.Topic}} = "{{.Topic}}"

func (e *{{.MessageName}}) ID() string {
	return {{$.FmtPkg}}Sprintf("{{.ID.Format}}", {{.ID.Args}})
}

func (*{{.MessageName}}) Topic() string {
	return Topic{{.Topic}}
}

func (e *{{.MessageName}}) Content() []byte {
	content, _ := {{$.ProtoPkg}}Marshal(e)
	return content
}

func (e *{{.MessageName}}) Unmarshal(content []byte) error {
	return {{$.ProtoPkg}}Unmarshal(content, e)
}
{{end}}
`

type EventIDDesc struct {
	Format string
	Args   string
}

type EventDesc struct {
	MessageName string
	ID          EventIDDesc
	Topic       string
}

type domainDesc struct {
	FmtPkg   string
	ProtoPkg string
	Events   []*EventDesc
}

func (s *domainDesc) execute() string {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("domain").Parse(strings.TrimSpace(domainTemplate))
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(buf, s); err != nil {
		panic(err)
	}
	return string(buf.Bytes())
}
