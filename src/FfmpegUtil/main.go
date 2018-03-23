package main

//#cgo pkg-config: libavformat libavcodec libavutil libavdevice libavfilter libswresample libswscale
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavformat/avformat.h>
//#include <libavcodec/avcodec.h>
//#include <libavutil/avutil.h>
//#include <libavutil/opt.h>
//#include <libavdevice/avdevice.h>
import "C"
import (
	"fmt"
	"unsafe"
	"FfmpegUtil/util"
)

type (
	Context C.struct_AVFormatContext
	InputFormat C.struct_AVInputFormat
	Dictionary C.struct_AVDictionary
	CodecContext C.struct_AVCodecContext
	MediaType C.enum_AVMediaType
	Codec C.struct_AVCodec
	Stream C.struct_AVStream
	CodecId C.enum_AVCodecID
)

//Initialize libavformat and register all the muxers, demuxers and protocols.
func AvRegisterAll() {
	C.av_register_all()
}

//Allocate an Context.
func AvformatAllocContext() *Context {
	return (*Context)(C.avformat_alloc_context())
}

//Open an input stream and read the header.
func AvformatOpenInput(ps **Context, fi string, fmt *InputFormat, d **Dictionary) int {
	return int(C.avformat_open_input((**C.struct_AVFormatContext)(unsafe.Pointer(ps)), C.CString(fi), (*C.struct_AVInputFormat)(fmt), (**C.struct_AVDictionary)(unsafe.Pointer(d))))
}

//Read packets of a media file to get stream information.
func (s *Context) AvformatFindStreamInfo(d **Dictionary) int {
	return int(C.avformat_find_stream_info((*C.struct_AVFormatContext)(s), (**C.struct_AVDictionary)(unsafe.Pointer(d))))
}

func (ctxt *Context) Streams() *Stream {
	return (*Stream)(unsafe.Pointer(ctxt.streams))
}

func (ctxt *Context) NbStreams() uint {
	return uint(ctxt.nb_streams)
}

func (avs *Stream) CodecContext() *CodecContext {
	return (*CodecContext)(unsafe.Pointer(avs.codec))
}

func (ctxt *CodecContext) CodecId() CodecId {
	return (CodecId)(ctxt.codec_id)
}

func (ctxt *CodecContext) Profile() int {
	return int(ctxt.profile)
}

func (ctxt *CodecContext) Codec() *Codec {
	return (*Codec)(ctxt.codec)
}

func (ctxt *CodecContext) CodecType() MediaType {
	return (MediaType)(ctxt.codec_type)
}

/*func (ctxt *CodecContext) CodecName() string {
	return C.GoString(ctxt.codec_name)
}*/

func main() {

	source := "rtsp://localhost:9990/test.mp4"

	fmt.Println("Valid source: ", isValidStream(source))

}

func isValidStream(source string) bool {

	// Register all formats and codecs
	AvRegisterAll()
	util.WriteFrames("/home/user/tools/ffserver/drone_video.mp4", "/home/user/tools/ffserver/drone_video%d.ppm")

	ctxtFormat := AvformatAllocContext()

	// Open source
	if AvformatOpenInput(&ctxtFormat, source, nil, nil) != 0 {
		fmt.Println("Error: Couldn't open file.")
		return false;
	}

	// Find stream info
	// Retrieve stream information
	if ctxtFormat.AvformatFindStreamInfo(nil) < 0 {
		fmt.Println("Error: Couldn't find stream information.")
		return false
	}

	//ctxtFormat->nb_streams
	n := ctxtFormat.NbStreams()

	//ctxtFormat->streams[]
	s := ctxtFormat.Streams()
	//s2 := avformat.StreamsOne(ctxtFormat, 1)

	fmt.Println("Number of Streams:", n)

	for i := 0; i < int(n); i++ {
		fmt.Println("Stream Number:", i)

		if (*CodecContext)(s.CodecContext()) != nil {
			codec := s.CodecContext()
			codecContext := (*CodecContext)(unsafe.Pointer(&codec))
			fmt.Println("Codec ID: ", codecContext.CodecId())
			fmt.Println("Profile:", codecContext.Profile())
			//codecActual := codecContext.Codec()
			//fmt.Println("Codec Name:", codecContext.CodecName())
			fmt.Println("Media Type Name:", codecContext.CodecType())
		}
	}

	return true
}
