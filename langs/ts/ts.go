package ts

import (
	"fmt"
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/mickeyreiss/firemodel"
	"github.com/mickeyreiss/firemodel/version"
	"github.com/pkg/errors"
)

func init() {
	firemodel.RegisterModeler("ts", &Modeler{})
}

type Modeler struct{}

func (m *Modeler) Model(schema *firemodel.Schema, sourceCoder firemodel.SourceCoder) error {
	f, err := sourceCoder.NewFile("firemodel.ts")
	if err != nil {
		return errors.Wrapf(err, "firemodel/ts: create typescript file")
	}
	defer f.Close()

	if err := tpl.Execute(f, schema); err != nil {
		return errors.Wrapf(err, "firemodel/ts: generating typescript")
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
			"getModelOption":   getModelOption,
			"getSchemaOption":  getSchemaOption,
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
			return fmt.Sprintf("DocumentReference<%s>", interfaceName(firetype.T.Name))
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
	case *firemodel.File:
		return "IFile"
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

func getSchemaOption(namespace string, key string, defaultValue string, options firemodel.SchemaOptions) string {
	ns, ok := options[namespace]
	if !ok {
		return defaultValue
	}
	opt, ok := ns[key]
	if !ok {
		return defaultValue
	}
	return opt
}

func getModelOption(namespace string, key string, required bool, options firemodel.SchemaModelOptions) string {
	ns, ok := options[namespace]
	if !ok {
		if required {
			err := errors.Errorf("option %s.%s is required but not set", namespace, key)
			panic(err)
		} else {
			return ""
		}
	}
	opt, ok := ns[key]
	if !ok {
		if required {
			err := errors.Errorf("option %s.%s is required but not set", namespace, key)
			panic(err)
		} else {
			return ""
		}
	}
	return opt
}

const (
	file = `// DO NOT EDIT - Code generated by firemodel {{firemodelVersion}}.
import { firestore } from 'firebase';

type Omit<T, K extends keyof T> = Pick<T, Exclude<keyof T, K>>;

export interface Query<DataType = firestore.DocumentData>
  extends firestore.Query {
  where(
    fieldPath: string | firestore.FieldPath,
    opStr: firestore.WhereFilterOp,
    value: any,
  ): Query<DataType>;
  orderBy(
    fieldPath: string | firestore.FieldPath,
    directionStr?: firestore.OrderByDirection,
  ): Query<DataType>;
  limit(limit: number): Query<DataType>;
  startAt(snapshot: DocumentSnapshot): Query<DataType>;
  startAt(...fieldValues: any[]): Query<DataType>;
  startAfter(snapshot: DocumentSnapshot): Query<DataType>;
  startAfter(...fieldValues: any[]): Query<DataType>;
  endBefore(snapshot: DocumentSnapshot): Query<DataType>;
  endBefore(...fieldValues: any[]): Query<DataType>;
  endAt(snapshot: DocumentSnapshot): Query<DataType>;
  endAt(...fieldValues: any[]): Query<DataType>;
  get(options?: firestore.GetOptions): Promise<QuerySnapshot<DataType>>;
  onSnapshot(observer: {
    next?: (snapshot: QuerySnapshot<DataType>) => void;
    error?: (error: Error) => void;
    complete?: () => void;
  }): () => void;
  onSnapshot(
    options: firestore.SnapshotListenOptions,
    observer: {
      next?: (snapshot: QuerySnapshot<DataType>) => void;
      error?: (error: Error) => void;
      complete?: () => void;
    },
  ): () => void;
  onSnapshot(
    onNext: (snapshot: QuerySnapshot<DataType>) => void,
    onError?: (error: Error) => void,
    onCompletion?: () => void,
  ): () => void;
  onSnapshot(
    options: firestore.SnapshotListenOptions,
    onNext: (snapshot: QuerySnapshot<DataType>) => void,
    onError?: (error: Error) => void,
    onCompletion?: () => void,
  ): () => void;
}


export interface DocumentSnapshot<DataType = firestore.DocumentData>
  extends firestore.DocumentSnapshot {
  data(options?: firestore.SnapshotOptions): DataType | undefined;
}
export interface QueryDocumentSnapshot<DataType = firestore.DocumentData>
  extends firestore.QueryDocumentSnapshot {
  data(options?: firestore.SnapshotOptions): DataType | undefined;
}
export interface QuerySnapshot<DataType = firestore.DocumentData>
  extends firestore.QuerySnapshot {
  readonly docs: QueryDocumentSnapshot<DataType>[];
}
export interface DocumentSnapshotExpanded<DataType = firestore.DocumentData> {
  exists: firestore.DocumentSnapshot['exists'];
  ref: firestore.DocumentSnapshot['ref'];
  id: firestore.DocumentSnapshot['id'];
  metadata: firestore.DocumentSnapshot['metadata'];
  data: DataType;
}
export interface QuerySnapshotExpanded<DataType = firestore.DocumentData> {
  metadata: {
    hasPendingWrites: firestore.QuerySnapshot['metadata']['hasPendingWrites'];
    fromCache: firestore.QuerySnapshot['metadata']['fromCache'];
  };
  size: firestore.QuerySnapshot['size'];
  empty: firestore.QuerySnapshot['empty'];
  docs: {
    [docId: string]: DocumentSnapshotExpanded<DataType>;
  };
}
export interface DocumentReference<DataType = firestore.DocumentData>
  extends firestore.DocumentReference {
  set(data: DataType, options?: firestore.SetOptions): Promise<void>;
  get(options?: firestore.GetOptions): Promise<DocumentSnapshot<DataType>>;
  onSnapshot(observer: {
    next?: (snapshot: DocumentSnapshot<DataType>) => void;
    error?: (error: firestore.FirestoreError) => void;
    complete?: () => void;
  }): () => void;
  onSnapshot(
    options: firestore.SnapshotListenOptions,
    observer: {
      next?: (snapshot: DocumentSnapshot<DataType>) => void;
      error?: (error: Error) => void;
      complete?: () => void;
    },
  ): () => void;
  onSnapshot(
    onNext: (snapshot: DocumentSnapshot<DataType>) => void,
    onError?: (error: Error) => void,
    onCompletion?: () => void,
  ): () => void;
  onSnapshot(
    options: firestore.SnapshotListenOptions,
    onNext: (snapshot: DocumentSnapshot<DataType>) => void,
    onError?: (error: Error) => void,
    onCompletion?: () => void,
  ): () => void;
}



export interface CollectionReference<DataType = firestore.DocumentData>
  extends Query<DataType>,
    Omit<firestore.CollectionReference, keyof Query> {
  add(data: DataType): Promise<DocumentReference>;
}
export interface Collection<DataType = firestore.DocumentData> {
  [id: string]: DocumentSnapshotExpanded<DataType>;
}


// tslint:disable-next-line:no-namespace
export namespace {{.Options | getSchemaOption "ts" "namespace" "firemodel"}} {
  type URL = string;

  export interface IFile {
    url: URL;
    mimeType: string;
    name: string;
  }

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
  {{- else}}

  /** TODO: Add documentation to {{.Name}} in firemodel schema. */
  {{- end}}
  export interface {{.Name | interfaceName | ToCamel}} {
    {{- range .Collections}}
    {{- if .Comment}}
    /** {{.Comment}} */
    {{- else }}
    /** TODO: Add documentation to {{.Name}} in firemodel schema. */
    {{- end}}
    {{.Name | ToLowerCamel}}: CollectionReference<{{.Type.Name | interfaceName | ToCamel}}>;
    {{- end}}

    {{- range .Fields}}
    {{- if .Comment}}
    /** {{.Comment}} */
    {{- else }}
    /** TODO: Add documentation to {{.Name}} in firemodel schema. */
    {{- end}}
    {{.Name | ToLowerCamel -}}?: {{toTypescriptType .Type}};
    {{- end}}
    {{- if .Options | getModelOption "firestore" "autotimestamp" false}}
    /** Record creation timestamp. */
    createdAt?: firestore.Timestamp;
    /** Record update timestamp. */
    updatedAt?: firestore.Timestamp;
		{{- end}}

    {{- if .Options | getModelOption "firestore" "autoversion" false}}
    /** Sync version */
    version?: number;
    /** Deletion tombstone */
    tombstone?: boolean;
    {{- end}}
  }`

	structTpl = `
  {{- if .Comment}}

  /** {{.Comment}} */
  {{- else}}

  /** TODO: Add documentation to {{.Name}} in firemodel schema. */
  {{- end}}
  export interface {{.Name | interfaceName | ToCamel}} {
    {{- range .Fields}}
    {{- if .Comment}}
    /** {{.Comment}} */
    {{- else }}
    /** TODO: Add documentation to {{.Name}} in firemodel schema. */
    {{- end}}
    {{.Name | ToLowerCamel -}}?: {{toTypescriptType .Type}};
    {{- end}}
  }`

	enum = `
  {{- if .Comment}}

  /** {{.Comment}} */
  {{- else}}

  /** TODO: Add documentation to {{.Name}} in firemodel schema. */
  {{- end}}
  export enum {{.Name | ToCamel}} {
    {{- range .Values}}
    {{- if .Comment}}
    /** {{.Comment}} */
    {{- else}}
    /** TODO: Add documentation to {{.Name}} in firemodel schema. */
    {{- end}}
    {{.Name}} = '{{.Name | ToScreamingSnake}}',
    {{- end}}
  }`
)
