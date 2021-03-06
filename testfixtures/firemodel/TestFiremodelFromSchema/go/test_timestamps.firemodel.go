// DO NOT EDIT - Code generated by firemodel (dev).

package firemodel

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"fmt"
	"regexp"
	"time"
)

//
// Firestore document location: /timestamps/{test_timestamps_id}
type TestTimestamps struct {

	// Creation timestamp.
	CreatedAt time.Time `firestore:"createdAt"`
	// Update timestamp.
	UpdatedAt time.Time `firestore:"updatedAt"`
}

// TestTimestampsPath returns the path to a particular TestTimestamps in Firestore.
func TestTimestampsPath(testTimestampsId string) string {
	return fmt.Sprintf("timestamps/%s", testTimestampsId)
}

// TestTimestampsRegexPath is a regex that can be use to filter out firestore events of TestTimestamps
var TestTimestampsRegexPath = regexp.MustCompile("^(?:projects/[^/]*/databases/[^/]*/documents/)?(?:/)?timestamps/([a-zA-Z0-9]+)$")

// TestTimestampsRegexNamedPath is a named regex that can be use to filter out firestore events of TestTimestamps
var TestTimestampsRegexNamedPath = regexp.MustCompile("^(?:projects/[^/]*/databases/[^/]*/documents/)?(?:/)?timestamps/(?P<test_timestamps_id>[a-zA-Z0-9]+)$")

// TestTimestampsPathStruct is a struct that contains parts of a path of TestTimestamps
type TestTimestampsPathStruct struct {
	TestTimestampsId string
}

// TestTimestampsPathToStruct is a function that turns a firestore path into a PathStruct of TestTimestamps
func TestTimestampsPathToStruct(path string) *TestTimestampsPathStruct {
	parsed := TestTimestampsRegexPath.FindStringSubmatch(path)
	result := &TestTimestampsPathStruct{TestTimestampsId: parsed[1]}
	return result
}

// TestTimestampsStructToPath is a function that turns a PathStruct of TestTimestamps into a firestore path
func TestTimestampsStructToPath(path *TestTimestampsPathStruct) string {
	built := fmt.Sprintf("timestamps/%s", path.TestTimestampsId)
	return built
}

// TestTimestampsWrapper is a struct wrapper that contains a reference to the firemodel instance and the path
type TestTimestampsWrapper struct {
	Data    *TestTimestamps
	Path    *TestTimestampsPathStruct
	PathStr string
	// ---- Internal Stuffs ----
	client  *clientTestTimestamps
	pathStr string
	ref     *firestore.DocumentRef
}

// TestTimestampsFromSnapshot is a function that will create an instance of the model from a document snapshot
func TestTimestampsFromSnapshot(snapshot *firestore.DocumentSnapshot) (*TestTimestampsWrapper, error) {
	temp := &TestTimestamps{}
	err := snapshot.DataTo(temp)
	if err != nil {
		return nil, err
	}
	path := TestTimestampsPathToStruct(snapshot.Ref.Path)
	pathStr := TestTimestampsStructToPath(path)
	wrapper := &TestTimestampsWrapper{Path: path, PathStr: pathStr, pathStr: pathStr, ref: snapshot.Ref, Data: temp}
	return wrapper, nil
}

type clientTestTimestamps struct {
	client *Client
}

func (c *clientTestTimestamps) Set(ctx context.Context, path string, model *TestTimestamps) (*TestTimestampsWrapper, error) {
	ref := c.client.Client.Doc(path)
	snapshot, err := ref.Get(ctx)
	if snapshot.Exists() {
		temp, err := TestTimestampsFromSnapshot(snapshot)
		if err != nil {
			// Don't do anything, just override
		} else {
			model.CreatedAt = temp.Data.CreatedAt
		}
	}
	wrapper := &TestTimestampsWrapper{ref: ref, pathStr: path, PathStr: path, Path: TestTimestampsPathToStruct(path), client: c, Data: model}
	wrapper.Data.UpdatedAt = time.Now()
	err = wrapper.Set(ctx)
	if err != nil {
		return nil, err
	}
	return wrapper, nil
}
func (c *clientTestTimestamps) GetByPath(ctx context.Context, path string) (*TestTimestampsWrapper, error) {
	reference := c.client.Client.Doc(path)
	snapshot, err := reference.Get(ctx)
	if err != nil {
		return nil, err
	}
	wrapper, err := TestTimestampsFromSnapshot(snapshot)
	if err != nil {
		return nil, err
	}
	return wrapper, nil
}
func (c *clientTestTimestamps) GetByPathTx(ctx context.Context, tx *firestore.Transaction, path string) (*TestTimestampsWrapper, error) {
	reference := c.client.Client.Doc(path)
	snapshot, err := tx.Get(reference)
	if err != nil {
		return nil, err
	}
	wrapper, err := TestTimestampsFromSnapshot(snapshot)
	if err != nil {
		return nil, err
	}
	return wrapper, nil
}
func (m *TestTimestampsWrapper) Set(ctx context.Context) error {
	if m.ref == nil {
		return errors.New("Cannot call set on a firemodel object that has no reference. Call `create` on the orm with this object instead")
	}
	_, err := m.ref.Set(ctx, m.Data)
	return err
}
func (m *TestTimestampsWrapper) SetTx(ctx context.Context, tx *firestore.Transaction) error {
	if m.ref == nil {
		return errors.New("Cannot call set on a firemodel object that has no reference. Call `create` on the orm with this object instead")
	}
	err := tx.Set(m.ref, m.Data)
	return err
}
