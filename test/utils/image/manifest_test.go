/*
Copyright 2019 The Kubernetes Authors.

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

package image

import (
	"testing"
)

var registryTests = []struct {
	in  string
	out string
}{
	{"docker.io/library/test:123", "test.io/library/test:123"},
	{"docker.io/library/test", "test.io/library/test"},
	{"test", "test.io/library/test"},
	{"gcr.io/kubernetes-e2e-test-images/test:123", "test.io/kubernetes-e2e-test-images/test:123"},
	{"k8s.gcr.io/test:123", "test.io/test:123"},
	{"gcr.io/k8s-authenticated-test/test:123", "test.io/k8s-authenticated-test/test:123"},
	{"gcr.io/google-samples/test:latest", "test.io/google-samples/test:latest"},
	{"unknwon.io/google-samples/test:latest", "unknwon.io/google-samples/test:latest"},
}

// go test k8s.io/kubernetes/test/utils/image
func TestReplaceRegistryInImageURL(t *testing.T) {
	// Set custom registries
	dockerLibraryRegistry = "test.io/library"
	e2eRegistry = "test.io/kubernetes-e2e-test-images"
	gcRegistry = "test.io"
	PrivateRegistry = "test.io/k8s-authenticated-test"
	sampleRegistry = "test.io/google-samples"

	for _, tt := range registryTests {
		t.Run(tt.in, func(t *testing.T) {
			s := ReplaceRegistryInImageURL(tt.in)
			if s != tt.out {
				t.Errorf("got %q, want %q", s, tt.out)
			}
		})
	}
}
