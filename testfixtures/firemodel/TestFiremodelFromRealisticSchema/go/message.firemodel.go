// DO NOT EDIT - Code generated by firemodel (dev).

package firemodel

import firestore "cloud.google.com/go/firestore"

// Firestore document location: /users/{user_id}/messages/{message_id}
type Message struct {
	Content MessageContent         `firestore:"content,omitempty"`
	From    *firestore.DocumentRef `firestore:"from,omitempty"`
}
