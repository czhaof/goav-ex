// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avformat

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"
import (
	"unsafe"

	"github.com/czhaof/goav-ex/avcodec"
	"github.com/czhaof/goav-ex/avutil"
)

func (avs *Stream) CodecParameters() *avcodec.AvCodecParameters {
	return (*avcodec.AvCodecParameters)(unsafe.Pointer(avs.codecpar))
}

func (avs *Stream) Metadata() *avutil.Dictionary {
	return (*avutil.Dictionary)(unsafe.Pointer(avs.metadata))
}

func (avs *Stream) AttachedPic() avcodec.Packet {
	return *fromCPacket(&avs.attached_pic)
}

func (avs *Stream) AvgFrameRate() avcodec.Rational {
	return newRational(avs.avg_frame_rate)
}

// func (avs *Stream) DisplayAspectRatio() *Rational {
// 	return (*Rational)(unsafe.Pointer(avs.display_aspect_ratio))
// }

func (avs *Stream) RFrameRate() avcodec.Rational {
	return newRational(avs.r_frame_rate)
}

func (avs *Stream) SampleAspectRatio() avcodec.Rational {
	return newRational(avs.sample_aspect_ratio)
}

func (avs *Stream) TimeBase() avcodec.Rational {
	return newRational(avs.time_base)
}

// func (avs *Stream) RecommendedEncoderConfiguration() string {
// 	return C.GoString(avs.recommended_encoder_configuration)
// }

func (avs *Stream) Discard() AvDiscard {
	return AvDiscard(avs.discard)
}

func (avs *Stream) Disposition() int {
	return int(avs.disposition)
}

func (avs *Stream) EventFlags() int {
	return int(avs.event_flags)
}

func (avs *Stream) Id() int {
	return int(avs.id)
}

func (avs *Stream) Index() int {
	return int(avs.index)
}

func (avs *Stream) NbSideData() int {
	return int(avs.nb_side_data)
}

func (avs *Stream) Duration() int64 {
	return int64(avs.duration)
}

// func (avs *Stream) FirstDiscardSample() int64 {
// 	return int64(avs.first_discard_sample)
// }

// func (avs *Stream) LastDiscardSample() int64 {
// 	return int64(avs.last_discard_sample)
// }

func (avs *Stream) NbFrames() int64 {
	return int64(avs.nb_frames)
}

// func (avs *Stream) StartSkipSamples() int64 {
// 	return int64(avs.start_skip_samples)
// }

func (avs *Stream) StartTime() int64 {
	return int64(avs.start_time)
}

// func (avs *Stream) PrivPts() *FFFrac {
// 	return (*FFFrac)(unsafe.Pointer(avs.priv_pts))
// }

func (avs *Stream) Free() {
	C.av_freep(unsafe.Pointer(avs))
}
