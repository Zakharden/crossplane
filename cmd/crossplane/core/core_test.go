/*
Copyright 2026 The Crossplane Authors.

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

package core

import (
	"testing"

	"github.com/crossplane/crossplane-runtime/v2/pkg/event"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

func TestRestrictNamespacedEventsFilter(t *testing.T) {
	type args struct {
		obj runtime.Object
	}

	tests := map[string]struct {
		args args
		want bool
	}{
		"FiltersClusterScopedObject": {
			args: args{
				obj: &corev1.Namespace{},
			},
			want: true,
		},
		"FiltersDefaultNamespaceObject": {
			args: args{
				obj: &corev1.ConfigMap{},
			},
			want: true,
		},
		"AllowsNonDefaultNamespaceObject": {
			args: args{
				obj: &corev1.ConfigMap{},
			},
			want: false,
		},
		"FiltersObjectsWithoutMetadata": {
			args: args{
				obj: &runtime.Unknown{},
			},
			want: true,
		},
	}

	tests["FiltersDefaultNamespaceObject"].args.obj.(*corev1.ConfigMap).SetNamespace("default")
	tests["AllowsNonDefaultNamespaceObject"].args.obj.(*corev1.ConfigMap).SetNamespace("tenant-a")

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := restrictNamespacedEventsFilter(tt.args.obj, event.Event{})
			if got != tt.want {
				t.Fatalf("restrictNamespacedEventsFilter(...) = %t, want %t", got, tt.want)
			}
		})
	}
}
