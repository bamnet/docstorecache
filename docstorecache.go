package docstorecache

import (
	"context"
	"net/url"

	"gocloud.dev/docstore"
)

// DocstoreCache is an implementation of a Cache that stores responses in a
// Go Cloud Docstore (Google Cloud Firestore, Amazon DynamoDB, Azure Cosmos DB , etc).
type DocstoreCache struct {
	Collection *docstore.Collection
}

type doc struct {
	Key      string
	Response []byte
}

// New returns a new Cache that uses a docstore collection. When creating the
// collection you must indicate that "Key" is the primary field (partition key, name, id).
func New(collection *docstore.Collection) *DocstoreCache {
	return &DocstoreCache{Collection: collection}
}

// Get attempts to retrive a cached item from docstore.
func (c *DocstoreCache) Get(k string) (resp []byte, ok bool) {
	ctx := context.Background()
	e := doc{Key: key(k)}
	if err := c.Collection.Get(ctx, &e); err != nil || len(e.Response) == 0 {
		return []byte{}, false
	}

	return e.Response, true

}

// Set attempts to save an item in docstore. Errors are silently ignored.
func (c *DocstoreCache) Set(k string, resp []byte) {
	ctx := context.Background()
	e := &doc{
		Key:      key(k),
		Response: resp,
	}
	c.Collection.Put(ctx, e)
}

// Delete attempts to remove an item from docstore.
func (c *DocstoreCache) Delete(k string) {
	ctx := context.Background()
	e := doc{Key: key(k)}
	c.Collection.Delete(ctx, e)
}

func key(key string) string {
	// Encode the key so it looks less like a URL.
	// Some Docstore providers (at least Google Cloud Firestore) construct
	// URLs using this key and don't like it when the key contains raw "/" characters.
	return url.QueryEscape(key)
}
