package shortener

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"
)

type Data struct {
	content     string
	lastUpdated time.Time
}

type shortenedData map[string]Data

var Short = New()

func New() shortenedData {
	return make(shortenedData)
}

func (s shortenedData) Add(content string) (string, error) {
	d := Data{
		content:     content,
		lastUpdated: time.Now(),
	}
	jsonData, err := json.Marshal(d)
	if err != nil {
		return "", fmt.Errorf("failed to marshal data for shortening. Err: %v", err)
	}
	hasher := sha256.New()
	hasher.Write(jsonData)
	hash := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	shortHash := hash[:6]
	s[shortHash] = d
	return shortHash, nil
}

func (s shortenedData) Remove(hash string) {
	delete(s, hash)
}

func (d Data) GetContent() string {
	return d.content
}

func (d Data) GetLastUpdatedTime() time.Time {
	return d.lastUpdated
}

// StartGarbageCollector is used to expire links that are older than the given age
func (s shortenedData) StartGarbageCollector(runsEvery time.Duration, dataAge time.Duration) {
	for range time.Tick(runsEvery) {
		for hash, data := range s {
			if time.Now().Sub(data.lastUpdated) >= dataAge {
				s.Remove(hash)
			}
		}
	}
}
