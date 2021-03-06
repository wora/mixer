// Copyright 2016 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package denyChecker

import (
	"testing"

	"istio.io/mixer/adapters"
	"istio.io/mixer/adapters/testutil"
)

func TestBuilderInvariants(t *testing.T) {
	b := NewBuilder()
	testutil.TestBuilderInvariants(b, t)
}

func TestAll(t *testing.T) {
	b := NewBuilder()
	b.Configure(b.DefaultBuilderConfig())

	a, err := b.NewAdapter(b.DefaultAdapterConfig())
	if err != nil {
		t.Errorf("Unable to create adapter: %v", err)
	}
	listChecker := a.(adapters.ListChecker)

	cases := []string{"", "ABC", "ABC/DEF", "AAA", "123"}
	for _, c := range cases {
		ok, err := listChecker.CheckList(c)
		if err != nil {
			t.Errorf("CheckList(%s) failed with error %v", c, err)
		}

		if ok {
			t.Errorf("CheckList(%s) => got 'true'; want 'false'", c)
		}
	}

	listChecker.Close()
}
