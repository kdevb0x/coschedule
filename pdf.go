// Copyright 2019 kdevb0x Ltd. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause license
// The full license text can be found in the LICENSE file.

package coschedule

import (
	"sync"

	"github.com/jung-kurt/gofpdf"
	pdf "github.com/jung-kurt/gofpdf/pdf"
)


type DocumentState struct {
	sync.Locker
	Viewers map[*Client]chan DocumentState
	Doc *gofpdf.Fpdf
}

var _ pdf.
