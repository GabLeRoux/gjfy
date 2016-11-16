package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"
)

// In-memory representation of a secret.
type StoreEntry struct {
	Secret    string    `json:"secret"`
	MaxClicks int       `json:"max_clicks"`
	Clicks    int       `json:"clicks"`
	DateAdded time.Time `json:"date_added"`
}

// Secret augmented with computed fields.
type StoreEntryInfo struct {
	StoreEntry
	Id        string `json:"id"`
	PathQuery string `json:"path_query"`
	Url       string `json:"url"`
}

type secretStore map[string]StoreEntry

// hashStruct returns a hash from an arbitrary structure, usable in a URL.
func hashStruct(data interface{}) (hash string) {
	hashBytes := sha256.Sum256([]byte(fmt.Sprintf("%#v", data)))
	hash = base64.URLEncoding.EncodeToString(hashBytes[:])
	return
}

// AddEntry adds a secret object to the store.
func (st secretStore) AddEntry(e StoreEntry) string {
	e.DateAdded = time.Now()
	id := hashStruct(e)
	st[id] = e
	return id
}

// NewEntry adds a new secret to the store.
func (st secretStore) NewEntry(secret string, maxclicks int) string {
	return st.AddEntry(StoreEntry{secret, maxclicks, 0, time.Time{}})
}

// GetEntry retrieves a secret from the store.
func (st secretStore) GetEntry(id string) (se StoreEntry, ok bool) {
	se, ok = st[id]
	return
}

// GetEntryInfo wraps GetEntry and adds some computed fields.
func (st secretStore) GetEntryInfo(id string) (si StoreEntryInfo, ok bool) {
	entry, ok := st.GetEntry(id)
	pathQuery := uGet + "?" + id
	url := schemeHost + listen + pathQuery
	return StoreEntryInfo{entry, id, pathQuery, url}, ok
}

// Click increases the click counter for an entry.
func (st secretStore) Click(id string) {
	entry, ok := st.GetEntry(id)
	if ok {
		if entry.Clicks < entry.MaxClicks-1 {
			entry.Clicks += 1
			st[id] = entry
		} else {
			delete(st, id)
		}
	}
	return
}