package avutil

/*
	#cgo pkg-config: libavutil
	#include <libavutil/frame.h>
	#include <stdlib.h>
*/
import "C"

func (f *Frame) GetNbSamples() int {
	return (int)(f.nb_samples)
}

func (f *Frame) SetNbSamples(nbSample int) {
	f.nb_samples = (C.int)(nbSample)
}

func (f *Frame) GetFormat() int {
	return (int)(f.format)
}

func (f *Frame) SetFormat(format int) {
	f.format = (C.int)(format)
}

func (f *Frame) GetChannelLayout() int {
	return (int)(f.channel_layout)
}

func (f *Frame) SetChannelLayout(layout int) {
	f.channel_layout = (C.ulong)(layout)
}
