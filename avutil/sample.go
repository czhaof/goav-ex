// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

// Package avutil is a utility library to aid portable multimedia programming.
// It contains safe portable string functions, random number generators, data structures,
// additional mathematics functions, cryptography and multimedia related functionality.
// Some generic features and utilities provided by the libavutil library
package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/avutil.h>
//#include <libavutil/samplefmt.h>
//#include <stdlib.h>
import "C"
import "unsafe"

func AvSamplesAllocArrayAndSamples(audioData ***uint8, lineSize *int, nbChannels int, nbSamples int, sampleFmt AvSampleFormat, align int) int {
	return int(C.av_samples_alloc_array_and_samples((***C.uint8_t)(unsafe.Pointer(audioData)), (*C.int)(lineSize), (C.int)(nbChannels), (C.int)(nbSamples), (C.enum_AVSampleFormat)(sampleFmt), (C.int)(align)))
}
