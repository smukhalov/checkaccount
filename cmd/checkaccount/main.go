// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package main

import (
	"checkaccount/app"
)

// This is the name you will use for the NET START command
const svcName = "checkaccount"

// This is the name that will appear in the Services control panel
const svcNameLong = "Check account MO"

// Площадка развертывания
var environment string

func svcLauncher() error {

	err := app.Run(elog, svcName, environment)
	if err != nil {
		return err
	}

	return nil
}
