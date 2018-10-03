package ios

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/mickeyreiss/firemodel"
	"github.com/mickeyreiss/firemodel/version"
	"github.com/pkg/errors"
)

func init() {
	firemodel.RegisterModeler("ios", &Modeler{})
}

type Modeler struct{}

func (m *Modeler) Model(schema *firemodel.Schema, sourceCoder firemodel.SourceCoder) error {
	f, err := sourceCoder.NewFile("Firemodel.swift")
	if err != nil {
		return errors.Wrapf(err, "firemodel/ios: create swift file")
	}
	defer f.Close()

	if err := tpl.Execute(f, schema); err != nil {
		return errors.Wrapf(err, "firemodel/ios: generating swift")
	}
	return nil
}

var (
	tpl = template.Must(template.
		New("file").
		Funcs(map[string]interface{}{
			"firemodelVersion":         func() string { return version.Version },
			"toSwiftType":              toSwiftType,
			"toScreamingSnake":         strcase.ToScreamingSnake,
			"toCamel":                  strcase.ToCamel,
			"toLowerCamel":             strcase.ToLowerCamel,
			"filterFieldsEnumsOnly":    filterFieldsEnumsOnly,
			"filterFieldsNonEnumsOnly": filterFieldsNonEnumsOnly,
			"filterFieldsStructsOnly":  filterFieldsStructsOnly,
			"hasFieldsOrStructs":       hasFieldsOrStructs,
			"firestorePath":            firestorePath,
		}).
		Parse(file),
	)
	_ = template.Must(tpl.New("model").Parse(model))
	_ = template.Must(tpl.New("enum").Parse(enum))
	_ = template.Must(tpl.New("struct").Parse(structTpl))
)

func hasFieldsOrStructs(in []*firemodel.SchemaField) bool {
	if len(filterFieldsStructsOnly(in)) > 0 {
		return true
	}
	if len(filterFieldsStructsOnly(in)) > 0 {
		return true
	}
	return false
}

func filterFieldsEnumsOnly(in []*firemodel.SchemaField) []*firemodel.SchemaField {
	var out []*firemodel.SchemaField
	for _, i := range in {
		if _, ok := i.Type.(*firemodel.Enum); !ok {
			continue
		}
		out = append(out, i)
	}
	return out
}

func filterFieldsNonEnumsOnly(in []*firemodel.SchemaField) []*firemodel.SchemaField {
	var out []*firemodel.SchemaField
	for _, i := range in {
		if _, ok := i.Type.(*firemodel.Enum); ok {
			continue
		}
		out = append(out, i)
	}
	return out
}

func filterFieldsStructsOnly(in []*firemodel.SchemaField) []*firemodel.SchemaField {
	var out []*firemodel.SchemaField
	for _, i := range in {
		if _, ok := i.Type.(*firemodel.Struct); !ok {
			continue
		}
		out = append(out, i)
	}
	return out
}

func toSwiftType(root bool, firetype firemodel.SchemaFieldType) string {
	switch firetype := firetype.(type) {
	case *firemodel.Boolean:
		return "Bool = false"
	case *firemodel.Integer:
		return "Int = 0"
	case *firemodel.Double:
		return "Float = 0.0"
	case *firemodel.Timestamp:
		if root {
			return "Date?"
		} else {
			return "Date"
		}
	case *firemodel.URL:
		if root {
			return "URL?"
		} else {
			return "URL"
		}
	case *firemodel.String:
		if root {
			return "String?"
		} else {
			return "String"
		}
	case *firemodel.Bytes:
		if root {
			return "Data?"
		} else {
			return "Data"
		}
	case *firemodel.Reference:
		if firetype.T != nil {
			if root {
				return fmt.Sprintf("Pring.Reference<%s> = .init()", strcase.ToCamel(firetype.T.Name))
			} else {
				return fmt.Sprintf("Pring.Reference<%s>", strcase.ToCamel(firetype.T.Name))
			}
		} else {
			return "Pring.AnyReference"
		}
	case *firemodel.GeoPoint:
		if root {
			return "Pring.GeoPoint?"
		} else {
			return "Pring.GeoPoint"
		}
	case *firemodel.Array:
		if firetype.T != nil {
			return fmt.Sprintf("[%s]?", toSwiftType(false, firetype.T))
		}
		return "[Any]"
	case *firemodel.File:
		if root {
			return "Pring.File?"
		} else {
			return "Pring.File"
		}
	case *firemodel.Struct:
		if root {
			return fmt.Sprintf("%s?", firetype.T.Name)
		} else {
			return fmt.Sprintf("%s", firetype.T.Name)
		}
	case *firemodel.Enum:
		if root {
			return fmt.Sprintf("%s?", strcase.ToCamel(firetype.T.Name))
		} else {
			return fmt.Sprintf("%s", strcase.ToCamel(firetype.T.Name))
		}
	case *firemodel.Map:
		if firetype.T != nil {
			return fmt.Sprintf("[String: %s] = [:]", toSwiftType(false, firetype.T))
		} else {
			return "[String: Any] = [:]"
		}
	default:
		err := errors.Errorf("firemodel/ios: unknown type %s", firetype)
		panic(err)
	}
}

func firestorePath(model firemodel.SchemaModel) string {
	format, args, err := model.Options.GetFirestorePath()
	if err != nil {
		panic(err)
	}

	if len(args) == 0 {
		fmt.Printf("warning: no firestore path for %s", model.Name)
		return ""
	}

	var out strings.Builder

	for _, arg := range args {
		fmt.Fprintf(&out, "static var %s: String = \"\"\n", strcase.ToLowerCamel(arg))
	}

	argsWrappedInInterpolationParens := make([]interface{}, len(args))
	for idx, arg := range args {
		argsWrappedInInterpolationParens[idx] = fmt.Sprintf(`\(%s)`, strcase.ToLowerCamel(arg))
	}
	path := fmt.Sprintf(format, argsWrappedInInterpolationParens...)

	fmt.Fprintf(&out, "    override class var path: String { return \"%s\" }", path)

	return out.String()
}

const (
	file = `// DO NOT EDIT - Code generated by firemodel {{firemodelVersion}}.

import Foundation
import Pring
{{range .Enums -}}
{{template "enum" .}}
{{- end}}
{{range .Structs -}}
{{template "struct" .}}
{{- end}}
{{- range .Models -}}
{{- template "model" .}}
{{- end -}}`

	model = `
{{- if .Comment}}
// {{.Comment}}
{{- else}}
// TODO: Add documentation to {{.Name | toCamel}} in firemodel schema.
{{- end}}
@objcMembers class {{.Name | toCamel}}: Pring.Object {
    {{. | firestorePath}}
    {{- range .Fields}}
    {{- if .Comment}}
    // {{.Comment}}
    {{- else }}
    // TODO: Add documentation to {{.Name | toLowerCamel}} in firemodel schema.
    {{- end}}
    dynamic var {{.Name | toLowerCamel -}}: {{.Type | toSwiftType true}}
    {{- end}}
    {{- range .Collections}}
    {{- if .Comment}}
    // {{.Comment}}
    {{- else }}
    // TODO: Add documentation to {{.Name}} in firemodel schema.
    {{- end}}
    dynamic var {{.Name | toLowerCamel}}: Pring.NestedCollection<{{.Type.Name}}> = []
    {{- end}}
    {{- if .Fields | hasFieldsOrStructs }}

    override func encode(_ key: String, value: Any?) -> Any? {
        switch key {
        {{- range .Fields | filterFieldsEnumsOnly}}
        case "{{.Name | toLowerCamel}}":
            return self.{{.Name | toLowerCamel}}?.firestoreValue
        {{- end}}
        {{- range .Fields | filterFieldsStructsOnly}}
        case "{{.Name | toLowerCamel}}":
            return self.{{.Name | toLowerCamel}}?.rawValue
        {{- end}}
        default:
            break
        }
        return nil
    }

    override func decode(_ key: String, value: Any?) -> Bool {
        switch key {
        {{- range .Fields | filterFieldsEnumsOnly}}
        case "{{.Name | toLowerCamel}}":
            self.{{.Name | toLowerCamel}} = {{.Type | toSwiftType false }}(firestoreValue: value)
        {{- end}}
        {{- range .Fields | filterFieldsStructsOnly}}
        case "{{.Name | toLowerCamel}}":
          if let value = value as? [String: Any] {
            self.{{.Name | toLowerCamel}} = {{.Type | toSwiftType false}}(id: self.id, value: value)
            return true
          }
          {{- end}}
        default:
            break
        }
        return false
    }
    {{- end}}
}
`
	enum = `
{{- if .Comment}}
// {{.Comment}}
{{- else}}
// TODO: Add documentation to {{.Name | toCamel}} in firemodel schema.
{{- end}}
@objc enum {{.Name | toCamel }}: Int {
    {{- range .Values}}
    {{- if .Comment}}
    // {{.Comment}}
    {{- else}}
    // TODO: Add documentation to {{.Name | toCamel}} in firemodel schema.
    {{- end}}
    case {{.Name | toLowerCamel}}
    {{- end}}
}

extension {{.Name | toCamel}}: CustomDebugStringConvertible {
    init?(firestoreValue value: Any?) {
        guard let value = value as? String else {
            return nil
        }
        switch value {
        {{- range $v := .Values}}
        case "{{$v.Name | toScreamingSnake}}":
            self = .{{$v.Name | toLowerCamel }}
        {{- end}}
        default:
            return nil
        }
    }

    var firestoreValue: String? {
        switch self {
        {{- range .Values}}
        case .{{.Name | toLowerCamel}}:
            return "{{.Name | toScreamingSnake}}"
        {{- end}}
        }
    }

    var debugDescription: String { return firestoreValue ?? "<INVALID>" }
}`

	structTpl = `
{{- if .Comment}}
// {{.Comment}}
{{- else}}
// TODO: Add documentation to {{.Name}} in firemodel schema.
{{- end}}
@objcMembers class {{.Name | toCamel }}: Pring.Object {
  {{- range .Fields}}
  {{- if .Comment}}
  // {{.Comment}}
  {{- else}}
  // TODO: Add documentation to {{.Name}} in firemodel schema.
  {{- end}}
  var {{.Name | toLowerCamel -}}: {{.Type | toSwiftType true}}
  {{- end}}
}
`
)
