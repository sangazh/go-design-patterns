package template

import (
	"strings"
)

type MessageRetriever interface {
	Message() string
}

type Template interface {
	first() string
	third() string
	ExecuteAlgorithm(m MessageRetriever) string
}

type TemplateTmpl struct{}

func (t *TemplateTmpl) first() string {
	return "hello"
}

func (t *TemplateTmpl) third() string {
	return "template"
}

func (t *TemplateTmpl) ExecuteAlgorithm(m MessageRetriever) string {
	return strings.Join([]string{t.first(), m.Message(), t.third()}, " ")
}

type TestStruct struct {
	Template
}

func (m *TestStruct) Message() string {
	return "world"
}

type AnonymousTemplate struct{}

func (a *AnonymousTemplate) first() string {
	return "hello"
}

func (a *AnonymousTemplate) third() string {
	return "template"
}

func (a *AnonymousTemplate) ExecuteAlgorithm(f func() string) string {
	return strings.Join([]string{a.first(), f(), a.third()}, " ")
}

type TemplateAdapter struct {
	myFunc func() string
}

func (a *TemplateAdapter) Message() string {
	if a.myFunc != nil {
		return a.myFunc()
	}
	return ""
}

func MessageRetrieverAdapter(f func() string) MessageRetriever {
	return &TemplateAdapter{myFunc: f}
}
