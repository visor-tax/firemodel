// DO NOT EDIT - Code generated by firemodel (dev).

package firemodel

import (
	"fmt"
	"time"
)

// TODO: Add comment to TestTimestamps in firemodel schema.
//
// Firestore document location: /timestamps/{test_timestamps_id}
type TestTimestamps struct {

	// Creation timestamp.
	CreatedAt time.Time `firestore:"createdAt,serverTimestamp"`
	// Update timestamp.
	UpdatedAt time.Time `firestore:"updatedAt,serverTimestamp"`
}

// TestTimestampsPath returns the path to a particular TestTimestamps in Firestore.
func TestTimestampsPath(testTimestampsId string) string {
	return fmt.Sprintf("timestamps/%s", testTimestampsId)
}