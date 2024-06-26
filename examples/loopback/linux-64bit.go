// Copyright 2017 Keybase Inc. All rights reserved.
// Use of this source code is governed by a BSD
// license that can be found in the LICENSE file.

//go:build linux && !386
// +build linux,!386

package main

import (
	"syscall"
	"time"
)

func tToTv(t time.Time) (tv syscall.Timeval) {
	tv.Sec = int64(t.Unix())
	tv.Usec = int64(t.UnixNano() % time.Second.Nanoseconds() /
		time.Microsecond.Nanoseconds())
	return tv
}
