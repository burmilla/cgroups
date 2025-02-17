// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package cmplx provides basic constants and mathematical functions for
// complex numbers.
package cmplx

import "github.com/containerd/cgroups/math"

// Abs returns the absolute value (also called the modulus) of x.
func Abs(x complex128) float64 { return math.Hypot(real(x), imag(x)) }
