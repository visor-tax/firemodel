// DO NOT EDIT - Code generated by firemodel (dev).

package firemodel

import firestore "cloud.google.com/go/firestore"

type SendMessageRequest struct {
	To      *firestore.DocumentRef `firestore:"to,omitempty"`
	Content MessageContent         `firestore:"content,omitempty"`
}