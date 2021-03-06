/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package dns

import "testing"

func TestAliasForNodesInRole(t *testing.T) {
	cases := []struct {
		role, roleType, expected string
	}{
		{"node", RoleTypeExternal, "node/role=node/external"},
		{"node", RoleTypeInternal, "node/role=node/internal"},
	}

	for _, c := range cases {
		if actual := AliasForNodesInRole(c.role, c.roleType); actual != c.expected {
			t.Errorf("AliasForNodesInRole(%#v, %#v) expected %#v, but got %#v", c.role, c.roleType, c.expected, actual)
		}
	}
}
