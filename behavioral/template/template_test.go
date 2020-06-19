package template

import (
	"strings"
	"testing"
)

func TestTemplate_ExecuteAlgorithm(t *testing.T) {
	t.Run("use interfaces", func(t *testing.T) {
		s := &TestStruct{&TemplateTmpl{}}
		res := ExecuteAlgorithm(s)
		expected := "world"
		if !strings.Contains(res, expected) {
			t.Error("not expected")
		}
		t.Log(res)
	})

	t.Run("using anonymous", func(t *testing.T) {
		m := new(AnonymousTemplate)
		res := m.ExecuteAlgorithm(func() string {
			return "world"
		})
		expectOrError(res, " world ", t)
	})

	t.Run("using anonymous adapted to the interface", func(t *testing.T) {
		mr := MessageRetrieverAdapter(func() string {
			return "world"
		})

		if mr == nil {
			t.Fatal("message retriever is nil")
		}
		s := &TemplateTmpl{}
		res := s.ExecuteAlgorithm(mr)

		expectOrError(res, " world ", t)
	})

}

func expectOrError(got, expect string, t *testing.T) {
	if !strings.Contains(got, expect) {
		t.Error("not expected")
	}
}
