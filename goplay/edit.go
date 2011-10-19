// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goplay

import (
	"appengine"
	"appengine/datastore"
	"http"
	"strings"
	"template"
)

func init() {
	http.HandleFunc("/", edit)
}

var editTemplate = template.Must(template.ParseFile("goplay/edit.html"))

func edit(w http.ResponseWriter, r *http.Request) {
	snip := &Snippet{Body: []byte(hello)}
	if strings.HasPrefix(r.URL.Path, "/p/") {
		c := appengine.NewContext(r)
		id := r.URL.Path[3:]
		key := datastore.NewKey(c, "Snippet", id, 0, nil)
		err := datastore.Get(c, key, snip)
		if err != nil {
			if err != datastore.ErrNoSuchEntity {
				c.Errorf("loading Snippet: %v", err)
			}
			http.Error(w, "Snippet not found", http.StatusNotFound)
			return
		}
	}
	editTemplate.Execute(w, snip)
}

const hello = `package main

import "fmt"

func main() {
	fmt.Println("Hello, playground")
}
`
