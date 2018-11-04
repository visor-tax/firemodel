// DO NOT EDIT - Code generated by firemodel (dev).

package firemodel

import (
	firestore "cloud.google.com/go/firestore"
	"context"
	"errors"
	"fmt"
	runtime "github.com/visor-tax/firemodel/runtime"
	latlng "google.golang.org/genproto/googleapis/type/latlng"
	"regexp"
	"time"
)

// A Test is a test model.
//
// Firestore document location: /users/{user_id}/test_models/{test_model_id}
type TestModel struct {
	// The name.
	Name string `firestore:"name"`
	// The age.
	Age int64 `firestore:"age"`
	// The number pi.
	Pi float64 `firestore:"pi"`
	// The birth date.
	Birthdate time.Time `firestore:"birthdate"`
	// True if it is good.
	IsGood bool `firestore:"isGood"`
	// TODO: Add comment to TestModel.data.
	Data []byte `firestore:"data"`
	// TODO: Add comment to TestModel.friend.
	Friend *firestore.DocumentRef `firestore:"friend"`
	// TODO: Add comment to TestModel.location.
	Location *latlng.LatLng `firestore:"location"`
	// TODO: Add comment to TestModel.colors.
	Colors []string `firestore:"colors"`
	// TODO: Add comment to TestModel.directions.
	Directions []TestEnum `firestore:"directions"`
	// TODO: Add comment to TestModel.models.
	Models []*TestStruct `firestore:"models"`
	// TODO: Add comment to TestModel.refs.
	Refs []*firestore.DocumentRef `firestore:"refs"`
	// TODO: Add comment to TestModel.model_refs.
	ModelRefs []*firestore.DocumentRef `firestore:"modelRefs"`
	// TODO: Add comment to TestModel.meta.
	Meta map[string]interface{} `firestore:"meta"`
	// TODO: Add comment to TestModel.meta_strs.
	MetaStrs map[string]string `firestore:"metaStrs"`
	// TODO: Add comment to TestModel.direction.
	Direction TestEnum `firestore:"direction"`
	// TODO: Add comment to TestModel.test_file.
	TestFile *runtime.File `firestore:"testFile"`
	// TODO: Add comment to TestModel.url.
	Url runtime.URL `firestore:"url"`
	// TODO: Add comment to TestModel.nested.
	Nested *TestStruct `firestore:"nested"`

	// Creation timestamp.
	CreatedAt time.Time `firestore:"createdAt"`
	// Update timestamp.
	UpdatedAt time.Time `firestore:"updatedAt"`
}

// TestModelPath returns the path to a particular TestModel in Firestore.
func TestModelPath(userId string, testModelId string) string {
	return fmt.Sprintf("users/%s/test_models/%s", userId, testModelId)
}

// TestModelRegexPath is a regex that can be use to filter out firestore events of TestModel
var TestModelRegexPath = regexp.MustCompile("^(?:projects/[^/]*/databases/[^/]*/documents/)?(?:/)?users/([a-zA-Z0-9]+)/test_models/([a-zA-Z0-9]+)$")

// TestModelRegexNamedPath is a named regex that can be use to filter out firestore events of TestModel
var TestModelRegexNamedPath = regexp.MustCompile("^(?:projects/[^/]*/databases/[^/]*/documents/)?(?:/)?users/(?P<user_id>[a-zA-Z0-9]+)/test_models/(?P<test_model_id>[a-zA-Z0-9]+)$")

// TestModelPathStruct is a struct that contains parts of a path of TestModel
type TestModelPathStruct struct {
	UserId      string
	TestModelId string
}

// TestModelPathToStruct is a function that turns a firestore path into a PathStruct of TestModel
func TestModelPathToStruct(path string) *TestModelPathStruct {
	parsed := TestModelRegexPath.FindStringSubmatch(path)
	result := &TestModelPathStruct{UserId: parsed[1], TestModelId: parsed[2]}
	return result
}

// TestModelStructToPath is a function that turns a PathStruct of TestModel into a firestore path
func TestModelStructToPath(path *TestModelPathStruct) string {
	built := fmt.Sprintf("users/%s/test_models/%s", path.UserId, path.TestModelId)
	return built
}

// TestModelWrapper is a struct wrapper that contains a reference to the firemodel instance and the path
type TestModelWrapper struct {
	Data    *TestModel
	Path    *TestModelPathStruct
	PathStr string
	// ---- Internal Stuffs ----
	client  *clientTestModel
	pathStr string
	ref     *firestore.DocumentRef
}

// TestModelFromSnapshot is a function that will create an instance of the model from a document snapshot
func TestModelFromSnapshot(snapshot *firestore.DocumentSnapshot) (*TestModelWrapper, error) {
	temp := &TestModel{}
	err := snapshot.DataTo(temp)
	if err != nil {
		return nil, err
	}
	path := TestModelPathToStruct(snapshot.Ref.Path)
	pathStr := TestModelStructToPath(path)
	wrapper := &TestModelWrapper{Path: path, PathStr: pathStr, pathStr: pathStr, ref: snapshot.Ref, Data: temp}
	return wrapper, nil
}

type clientTestModel struct {
	client *Client
}

func (c *clientTestModel) Set(ctx context.Context, path string, model *TestModel) (*TestModelWrapper, error) {
	ref := c.client.Client.Doc(path)
	snapshot, err := ref.Get(ctx)
	if snapshot.Exists() {
		temp, err := TestModelFromSnapshot(snapshot)
		if err != nil {
			// Don't do anything, just override
		} else {
			model.CreatedAt = temp.Data.CreatedAt
		}
	}
	wrapper := &TestModelWrapper{ref: ref, pathStr: path, PathStr: path, Path: TestModelPathToStruct(path), client: c, Data: model}
	wrapper.Data.UpdatedAt = time.Now()
	err = wrapper.Set(ctx)
	if err != nil {
		return nil, err
	}
	return wrapper, nil
}
func (c *clientTestModel) GetByPath(ctx context.Context, path string) (*TestModelWrapper, error) {
	reference := c.client.Client.Doc(path)
	snapshot, err := reference.Get(ctx)
	if err != nil {
		return nil, err
	}
	wrapper, err := TestModelFromSnapshot(snapshot)
	if err != nil {
		return nil, err
	}
	return wrapper, nil
}
func (m *TestModelWrapper) Set(ctx context.Context) error {
	if m.ref == nil {
		return errors.New("Cannot call set on a firemodel object that has no reference. Call `create` on the orm with this object instead")
	}
	_, err := m.ref.Set(ctx, m.Data)
	return err
}
