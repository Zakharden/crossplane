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
	"time"
)

func TestNewManagerCacheOptions(t *testing.T) {
	syncInterval := 5 * time.Minute
	namespace := "crossplane-system"

	o := newManagerCacheOptions(syncInterval, namespace, true)

	if o.SyncPeriod == nil || *o.SyncPeriod != syncInterval {
		t.Fatalf("expected sync period %s, got %v", syncInterval, o.SyncPeriod)
	}

	if _, ok := o.DefaultNamespaces[namespace]; !ok {
		t.Fatalf("expected manager cache to be restricted to namespace %q, got %v", namespace, o.DefaultNamespaces)
	}
}

func TestNewManagerCacheOptionsClusterWide(t *testing.T) {
	o := newManagerCacheOptions(time.Minute, "crossplane-system", false)

	if len(o.DefaultNamespaces) != 0 {
		t.Fatalf("expected manager cache to be cluster-wide, got %v", o.DefaultNamespaces)
	}
}

func TestNewAPIExtensionsCacheOptionsClusterWide(t *testing.T) {
	syncInterval := 5 * time.Minute

	o := newAPIExtensionsCacheOptions(syncInterval)

	if o.SyncPeriod == nil || *o.SyncPeriod != syncInterval {
		t.Fatalf("expected sync period %s, got %v", syncInterval, o.SyncPeriod)
	}

	if len(o.DefaultNamespaces) != 0 {
		t.Fatalf("expected API extensions cache to stay cluster-wide, got %v", o.DefaultNamespaces)
	}
}
