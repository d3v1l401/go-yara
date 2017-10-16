// Copyright © 2015-2017 Hilko Bengen <bengen@hilluzination.de>
// All rights reserved.
//
// Use of this source code is governed by the license that can be
// found in the LICENSE file.

package yara
/*
#cgo CFLAGS: -IC:/yara/win64/include
#cgo amd64 LDFLAGS: -LC:/yara/win64/lib -llibyara64
#cgo 386 LDFLAGS: -LC:/yara/32/lib -llibyara
*/
import "C"
