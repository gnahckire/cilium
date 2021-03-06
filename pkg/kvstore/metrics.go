// Copyright 2018 Authors of Cilium
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

package kvstore

import (
	"fmt"
	"strings"

	"github.com/cilium/cilium/pkg/metrics"
)

const (
	metricDelete = "delete"
	metricRead   = "read"
	metricSet    = "set"
)

func getScopeFromKey(key string) string {
	s := strings.SplitN(key, "/", 5)
	if len(s) != 5 {
		if len(key) >= 12 {
			return key[:12]
		}
		return key
	}
	return fmt.Sprintf("%s/%s", s[2], s[3])
}

func increaseMetric(key, kind, action string) {
	namespace := getScopeFromKey(key)
	metrics.KVStoreOperationsTotal.WithLabelValues(
		namespace, kind, action).Inc()
}
