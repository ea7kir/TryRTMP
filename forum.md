I need help porting this command line to Go...
```
gst-launch-1.0 \
    videotestsrc is-live=1 \
    ! videoconvert \
    ! "video/x-raw, width=1920, height=1080, framerate=25/1" \
    ! queue \
    ! x264enc cabac=1 bframes=2 ref=1 \
    ! "video/x-h264,profile=main" \
    ! flvmux streamable=true name=mux \
    ! rtmpsink location="${RTMP_DEST} live=1" \
    audiotestsrc is-live=1 wave=ticks \
    ! voaacenc bitrate=128000 \
    ! mux.
```
All I have so far is this...
```
	gst.Init()
	pipeline := gst.NewPipeline("").(gst.Pipeline)

	videotestsrc := gst.ElementFactoryMake("videotestsrc", "")
	videoconvert := gst.ElementFactoryMake("videoconvert", "")
	video_x_raw := gst.ElementFactoryMake("video/x-raw", "")
	queue := gst.ElementFactoryMake("queue", "")
	x264enc := gst.ElementFactoryMake("x264enc", "")
	video_x_h264 := gst.ElementFactoryMake("video/x-h264", "")
	flvmux := gst.ElementFactoryMake("flvmux", "")
	rtmpsink := gst.ElementFactoryMake("rtmpsink", "")
	audiotestsrc := gst.ElementFactoryMake("audiotestsrc", "")
	voaacenc := gst.ElementFactoryMake("voaacenc", "")

	pipeline.AddMany(videotestsrc, videoconvert, video_x_raw, queue, x264enc,
		video_x_h264, flvmux, rtmpsink, audiotestsrc, voaacenc, flvmux)

	gst.LinkMany(videotestsrc, videoconvert, video_x_raw, queue, x264enc,
		video_x_h264, flvmux, rtmpsink, audiotestsrc, voaacenc, flvmux)

	// TODO: add arguments to elements
```
but I don't think this is the way to handle the ```mux.``` name, and I'm not confident my entire approach is correct.

Help and advice will be most welcome.
