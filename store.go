package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	hiddenString = "#HIDDEN#"
)

var (
	expFactor = realExpFactor
)

// In-memory representation of a secret.
type StoreEntry struct {
	Secret    string    `json:"secret"`
	MaxClicks int       `json:"max_clicks"`
	Clicks    int       `json:"clicks"`
	DateAdded time.Time `json:"date_added"`
	ValidFor  int       `json:"valid_for"`
	AuthToken string    `json:"auth_token"`
}

// Secret augmented with computed fields.
type StoreEntryInfo struct {
	StoreEntry
	Id        string `json:"id"`
	PathQuery string `json:"path_query"`
	Url       string `json:"url"`
	ApiUrl    string `json:"api_url"`
}

type secretStore map[string]StoreEntry

// hashStruct returns a hash from an arbitrary structure, usable in a URL.
func hashStruct(data interface{}) (hash string) {
	hashBytes := sha256.Sum256([]byte(fmt.Sprintf("%v", data)))
	hash = base64.RawURLEncoding.EncodeToString(hashBytes[:])
	return
}

// AddEntry adds a secret to the store.
func (st secretStore) AddEntry(e StoreEntry, id string) string {
	e.DateAdded = time.Now()
	if id == "" {
		id = hashStruct(e)
	}
	if e.ValidFor == 0 {
		e.ValidFor = defaultValidity
	}
	st[id] = e
	return id
}

// NewEntry adds a new secret to the store. Set id to ""
// to have it auto-generated by hashing the entry.
func (st secretStore) NewEntry(secret string, maxclicks int, validfor int, at string, id string) string {
	return st.AddEntry(StoreEntry{secret, maxclicks, 0, time.Time{}, validfor, at}, id)
}

// GetEntry retrieves a secret from the store.
func (st secretStore) GetEntry(id string) (se StoreEntry, ok bool) {
	se, ok = st[id]
	return
}

// GetEntryInfo wraps GetEntry and adds some computed fields.
func (st secretStore) GetEntryInfo(id string) (si StoreEntryInfo, ok bool) {
	entry, ok := st.GetEntry(id)
	pathQuery := uGet + "?id=" + id
	url := getURLBase() + pathQuery
	apiurl := getURLBase() + uApiGet + id
	return StoreEntryInfo{entry, id, pathQuery, url, apiurl}, ok
}

// GetEntryInfo wraps GetEntry and adds some computed fields. In addition it
// hides the "secret" value.
func (st secretStore) GetEntryInfoHidden(id string) (si StoreEntryInfo, ok bool) {
	si, ok = st.GetEntryInfo(id)
	si.Secret = hiddenString
	return
}

// Click increases the click counter for an entry and sends a notification
func (st secretStore) Click(id string, r *http.Request) {
	var msg string
	entry, ok := st.GetEntry(id)
	if ok {
		if entry.Clicks < entry.MaxClicks-1 {
			entry.Clicks += 1
			st[id] = entry
		} else {
			delete(st, id)
		}
		msg = fmt.Sprintf(`
Id:          %s
Clicked:     %d time(s)
Clicks left: %d
Request:     %s %s %s %s
User-Agent:  %s
`,
			id, entry.Clicks, entry.MaxClicks-entry.Clicks,
			r.RemoteAddr, r.Method, r.URL.Path, r.Proto, r.Header.Get("User-Agent"))
		NotifyMail(entry.AuthToken, msg)
	}
	return
}

// realExpFactor will scale the given int value to a certain amount of time
func realExpFactor(v int) time.Duration {
	return time.Hour * 24 * time.Duration(v)
}

// Expiry checks for expired entries at regular intervals
func (st secretStore) Expiry(interval time.Duration) {
	tck := time.NewTicker(interval)
	log.Printf("checking for expiration every %s\n", interval)
	for {
		now := time.Now()
		for id, e := range st {
			expDate := e.DateAdded
			expDate = expDate.Add(expFactor(e.ValidFor))
			if now.After(expDate) {
				log.Printf("%s expired\n", id)
				delete(st, id)
			}
		}
		<-tck.C
	}
}
