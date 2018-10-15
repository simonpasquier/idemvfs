// Copyright 2018 Simon Pasquier
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package idemvfs

import (
	"net/http"
	"testing"
	"time"
)

var zeroTime time.Time = time.Time{}

type testIdentifier struct{}

func (t testIdentifier) Identify(_ string) (Identity, bool) { return &testIdentity{}, true }

type testIdentity struct{}

func (t testIdentity) Checksum() []byte   { return []byte{} }
func (t testIdentity) ModTime() time.Time { return zeroTime }
func (t testIdentity) Size() int64        { return 0 }

func TestOpen(t *testing.T) {
	fs := NewFileSystem(http.Dir("testdata/"), &testIdentifier{})

	f, err := fs.Open("file1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	s, err := f.Stat()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if zeroTime == s.ModTime() {
		t.Fatalf("unexpected modification time, got: %v", s.ModTime())
	}
}
