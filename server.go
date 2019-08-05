// Copyright 2019 kdevb0x Ltd. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause license
// The full license text can be found in the LICENSE file.

package coschedule

import (
	"sync"
)

type Client struct {
	Addr          string
	localDocState *DocumentState
}

type DocumentState struct {
	sync.Locker
}
