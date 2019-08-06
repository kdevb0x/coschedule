// Copyright 2019 kdevb0x Ltd. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause license
// The full license text can be found in the LICENSE file.

package coschedule

type Client struct {
	Addr          string
	localDocState *DocumentState
}

func (d *DocumentState) AddToRenderQueue(c *Client) chan DocumentState {
	if d.Viewers == nil {
		d.Viewers = make(map[*Client]chan DocumentState)
		d.Viewers[c] = make(chan DocumentState)
	}
	if q, ok := d.Viewers[c]; ok {
		return q
	}
	cn := make(chan DocumentState)
	d.Viewers[c] = cn
	return cn
}
