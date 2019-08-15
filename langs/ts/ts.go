package ts

import (
	"fmt"
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/pkg/errors"
	"github.com/visor-tax/firemodel"
	"github.com/visor-tax/firemodel/version"
)

func init() {
	firemodel.RegisterModeler("ts", &Modeler{})
}

type Modeler struct{}

func (m *Modeler) Model(schema *firemodel.Schema, options firemodel.GenOptions, sourceCoder firemodel.SourceCoder) error {
	f, err := sourceCoder.NewFile("firemodel.ts")
	if err != nil {
		return errors.Wrapf(err, "firemodel/ts: create typescript file")
	}
	defer f.Close()

	// d, err := os.Create("firebase.d.ts")
	d, err := sourceCoder.NewFile("firebase.d.ts")
	if err != nil {
		return errors.Wrapf(err, "firemodel/ts: create typescript definition file")
	}
	defer d.Close()

	if err := tpl.Execute(f, struct {
		Schema  *firemodel.Schema
		Options firemodel.GenOptions
	}{
		Schema:  schema,
		Options: options,
	}); err != nil {
		return errors.Wrapf(err, "firemodel/ts: generating typescript")
	}

	_, err = d.Write([]byte(definitions))
	if err != nil {
		return errors.Wrapf(err, "firemodel/ts: writing typescript definitions")
	}

	return nil
}

var (
	tpl = template.Must(template.
		New("file").
		Funcs(map[string]interface{}{
			"firemodelVersion": func() string { return version.Version },
			"toTypescriptType": toTypescriptType,
			"ToScreamingSnake": strcase.ToScreamingSnake,
			"ToLowerCamel":     strcase.ToLowerCamel,
			"ToCamel":          strcase.ToCamel,
			"interfaceName":    interfaceName,
		}).
		Parse(file),
	)
	_ = template.Must(tpl.New("model").Parse(model))
	_ = template.Must(tpl.New("enum").Parse(enum))
	_ = template.Must(tpl.New("struct").Parse(structTpl))
)

func interfaceName(sym string) string {
	return fmt.Sprintf("I%s", sym)
}

func toTypescriptType(firetype firemodel.SchemaFieldType) string {
	switch firetype := firetype.(type) {
	case *firemodel.Boolean:
		return "boolean"
	case *firemodel.Integer, *firemodel.Double:
		return "number"
	case *firemodel.Timestamp:
		return "firestore.Timestamp"
	case *firemodel.String:
		return "string"
	case *firemodel.Enum:
		return firetype.T.Name
	case *firemodel.URL:
		return "URL"
	case *firemodel.Bytes:
		return "firestore.Blob"
	case *firemodel.Reference:
		if firetype.T != nil {
			return fmt.Sprintf("firestore.DocumentReference<%s>", interfaceName(firetype.T.Name))
		} else {
			return "firestore.DocumentReference"
		}
	case *firemodel.GeoPoint:
		return "firestore.GeoPoint"
	case *firemodel.Array:
		if firetype.T != nil {
			return fmt.Sprintf("%s[]", toTypescriptType(firetype.T))
		} else {
			return "any[]"
		}
	case *firemodel.Struct:
		return interfaceName(firetype.T.Name)
	case *firemodel.Map:
		if firetype.T != nil {
			return fmt.Sprintf("{ [key: string]: %s; }", toTypescriptType(firetype.T))
		} else {
			return `{ [key: string]: any; }`
		}
	default:
		err := errors.Errorf("firemodel/ts: unknown type %s", firetype)
		panic(err)
	}
}

const (
	file = `// DO NOT EDIT - Code generated by firemodel {{firemodelVersion}}.
import { firestore } from 'firebase';

// tslint:disable-next-line:no-namespace
export namespace {{ .Options.Get "ts.namespace" }} {
  type URL = string;

  {{- range .Enums -}}
  {{- template "enum" .}}
  {{- end}}
  {{- range .Structs -}}
  {{- template "struct" .}}
  {{- end}}
  {{- range .Models -}}
  {{- template "model" .}}
  {{- end}}
}
`
	model = `
  {{- if .Comment}}

  /** {{.Comment}} */
  {{- end}}
  export interface {{.Name | interfaceName | ToCamel}} {
    {{- range .Collections}}
    {{- if .Comment}}
    /** {{.Comment}} */
    {{- end}}
    {{.Name | ToLowerCamel}}: firestore.CollectionReference<{{.Type.Name | interfaceName | ToCamel}}>;
    {{- end}}

    {{- range .Fields}}
    {{- if .Comment}}
    /** {{.Comment}} */
    {{- end}}
    {{.Name | ToLowerCamel -}}?: {{toTypescriptType .Type}};
    {{- end}}
  }`

	structTpl = `
  {{- if .Comment}}

  /** {{.Comment}} */
  {{- end}}
  export interface {{.Name | interfaceName | ToCamel}} {
    {{- range .Fields}}
    {{- if .Comment}}
    /** {{.Comment}} */
    {{- end}}
    {{.Name | ToLowerCamel -}}?: {{toTypescriptType .Type}};
    {{- end}}
  }`

	enum = `
  {{- if .Comment}}

  /** {{.Comment}} */
  {{- end}}
  export enum {{.Name | ToCamel}} {
    {{- range .Values}}
    {{- if .Comment}}
    /** {{.Comment}} */
    {{- end}}
    {{.Name}} = '{{.Name | ToScreamingSnake}}',
    {{- end}}
  }`

	definitions = `import * as FIREBASE from 'firebase';

declare module 'firebase' {
  namespace firestore {
    // Snapshots
    export interface DocumentSnapshot<T = DocumentData> {
      data(options?: SnapshotOptions): D | undefined;
    }
    export interface QueryDocumentSnapshot<T = DocumentData> extends DocumentSnapshot {
      data(options?: SnapshotOptions): T;
    }
    export interface QuerySnapshot<T = DocumentData> {
      readonly docs: QueryDocumentSnapshot<T>[];
      forEach(callback: (result: QueryDocumentSnapshot<T>) => void, thisArg?: any): void;
    }

    // References + Queries
    export interface DocumentReference<T = DocumentData> {
      onSnapshot(observer: {
        next?: (snapshot: DocumentSnapshot<T>) => void;
        error?: (error: FirestoreError) => void;
        complete?: () => void;
      }): () => void;
      onSnapshot(
        options: SnapshotListenOptions,
        observer: {
          next?: (snapshot: DocumentSnapshot<T>) => void;
          error?: (error: Error) => void;
          complete?: () => void;
        },
      ): () => void;
      onSnapshot(
        onNext: (snapshot: DocumentSnapshot<T>) => void,
        onError?: (error: Error) => void,
        onCompletion?: () => void,
      ): () => void;
      onSnapshot(
        options: SnapshotListenOptions,
        onNext: (snapshot: DocumentSnapshot<T>) => void,
        onError?: (error: Error) => void,
        onCompletion?: () => void,
      ): () => void;
    }
    export interface Query<T = DocumentData> {
      onSnapshot(observer: {
        next?: (snapshot: QuerySnapshot<T>) => void;
        error?: (error: Error) => void;
        complete?: () => void;
      }): () => void;
      onSnapshot(
        options: SnapshotListenOptions,
        observer: {
          next?: (snapshot: QuerySnapshot<T>) => void;
          error?: (error: Error) => void;
          complete?: () => void;
        },
      ): () => void;
      onSnapshot(
        onNext: (snapshot: QuerySnapshot<T>) => void,
        onError?: (error: Error) => void,
        onCompletion?: () => void,
      ): () => void;
      onSnapshot(
        options: SnapshotListenOptions,
        onNext: (snapshot: QuerySnapshot<T>) => void,
        onError?: (error: Error) => void,
        onCompletion?: () => void,
      ): () => void;
    }
    export interface CollectionReference<T = DocumentData> extends Query<T> {}
  }
}`
)
