// Copyright (c) 2020, Mohlmann Solutions SRL. All rights reserved.
// Use of this source code is governed by a License that can be found in the LICENSE file.
// SPDX-License-Identifier: BSD-3-Clause

package pqdsn

import "testing"

func BenchmarkBuilder_addString(t *testing.B) {
	const (
		k = "key"
		s = "foo's bar is a long string that's nice"
	)

	b := &builder{escape: true}

	for i := 0; i < t.N; i++ {
		b.addString(k, s)
	}
}
