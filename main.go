package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-gst/go-gst/pkg/gst"
)

/*
# https://github.com/matthew1000/gstreamer-cheat-sheet/blob/master/rtmp.md
# Derived from the YouTube Example
# THIS IS WORKINGG !!!! SO I WILL TRY THIS IN GO
RTMP_DEST=rtmp://somewhere.any/live/name
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
*/

const RTMP_DEST = "rtmp://somewhere.any/live/name" // i'm unable to reveal the real url to the pubic

func createPipeline() (gst.Pipeline, error) {
	log.Printf("TryRTMP start creating pipeline")

	var (
		videotestsrc gst.Element
		videoconvert gst.Element
		capsfilter1  gst.Element // replaces "video/x-raw, width=1920, height=1080, framerate=25/1"
		queue        gst.Element
		x264enc      gst.Element
		capsfilter2  gst.Element // replaces "video/x-h264,profile=main"
		flvmux       gst.Element
		rtmpsink     gst.Element
		audiotestsrc gst.Element
		voaacenc     gst.Element
	)

	gst.Init()

	// Create a pipeline

	pipeline := gst.NewPipeline("").(gst.Pipeline)

	if videotestsrc = gst.ElementFactoryMake("videotestsrc", "videotestsrc"); videotestsrc == nil {
		log.Printf("TryRTMP videotestsrc is nil")
	}
	if videoconvert = gst.ElementFactoryMake("videoconvert", "videoconvert"); videoconvert == nil {
		log.Printf("TryRTMP videoconvert is nil")
	}
	if capsfilter1 = gst.ElementFactoryMake("capsfilter", "capsfilter1"); capsfilter1 == nil {
		log.Printf("TryRTMP capsfilter1 is nil")
	}
	if queue = gst.ElementFactoryMake("queue", "queue"); queue == nil {
		log.Printf("TryRTMP queue is nil")
	}
	if x264enc = gst.ElementFactoryMake("x264enc", "x264enc"); x264enc == nil {
		log.Printf("TryRTMP x264enc is nil")
	}
	if capsfilter2 = gst.ElementFactoryMake("capsfilter", "capsfilter2"); capsfilter2 == nil {
		log.Printf("TryRTMP capsfilter2 is nil")
	}
	if flvmux = gst.ElementFactoryMake("flvmux", "mux"); flvmux == nil {
		log.Printf("TryRTMP flvmux is nil")
	}
	if rtmpsink = gst.ElementFactoryMake("rtmpsink", "rtmpsink"); rtmpsink == nil {
		log.Printf("TryRTMP rtmpsink is nil")
	}
	if audiotestsrc = gst.ElementFactoryMake("audiotestsrc", "audiotestsrc"); audiotestsrc == nil {
		log.Printf("TryRTMP audiotestsrc is nil")
	}
	if voaacenc = gst.ElementFactoryMake("voaacenc", "voaacenc"); voaacenc == nil {
		log.Printf("TryRTMP voaacenc is nil")
	}

	log.Printf("TryRTMP end of ElementFactoryMake")

	// add capabilities

	// videotestsrc is-live=1
	videotestsrc.SetObjectProperty("is-live", true)

	// videoconvert
	// nothing needed here

	capsfilter1.SetObjectProperty("caps", gst.CapsFromString("video/x-raw, width=1920, height=1080, framerate=25/1"))

	// queue
	// nothing needed here

	// x264enc cabac=1 bframes=2 ref=1
	x264enc.SetObjectProperty("cabac", true)
	x264enc.SetObjectProperty("bframes", 2)
	x264enc.SetObjectProperty("ref", 1)

	// capsfilter2 caps=video/x-h264, profile=main
	capsfilter2.SetObjectProperty("caps", gst.CapsFromString("video/x-h264, profile=main"))

	// flvmux streamable=true name=mux
	flvmux.SetObjectProperty("streamable", true) // name not required

	// rtmpsink location="${RTMP_DEST} live=1"
	rtmpsink.SetObjectProperty("location", RTMP_DEST+" live=1")

	// audiotestsrc is-live=1 wave=ticks
	audiotestsrc.SetObjectProperty("is-live", true)
	// audiotestsrc.SetObjectProperty("wave", 8) // GstAudioTestSrcWave.ticks
	gst.UtilSetObjectArg(audiotestsrc, "wave", "tick") // alternative method

	// voaacenc bitrate=128000
	voaacenc.SetObjectProperty("bitrate", 128000)

	// Add the elements to the pipeline and link them

	if true {
		log.Printf("TryRTMP using pipeline.Add")
		if ok := pipeline.Add(videotestsrc); ok != true {
			log.Printf("TryRTMP failed to add videotestsrc")
		}
		if ok := pipeline.Add(videoconvert); ok != true {
			log.Printf("TryRTMP failed to add videoconvert")
		}
		if ok := pipeline.Add(capsfilter1); ok != true {
			log.Printf("TryRTMP failed to add capsfilter1")
		}
		if ok := pipeline.Add(queue); ok != true {
			log.Printf("TryRTMP failed to add queue")
		}
		if ok := pipeline.Add(x264enc); ok != true {
			log.Printf("TryRTMP failed to add x264enc")
		}
		if ok := pipeline.Add(capsfilter2); ok != true {
			log.Printf("TryRTMP failed to add capsfilter2")
		}
		if ok := pipeline.Add(flvmux); ok != true {
			log.Printf("TryRTMP failed to add flvmux")
		}
		if ok := pipeline.Add(rtmpsink); ok != true {
			log.Printf("TryRTMP failed to add rtmpsink")
		}
		if ok := pipeline.Add(capsfilter2); ok != true {
			log.Printf("TryRTMP failed to add capsfilter2")
		}
		if ok := pipeline.Add(audiotestsrc); ok != true {
			log.Printf("TryRTMP failed to add audiotestsrc")
		}
		if ok := pipeline.Add(voaacenc); ok != true {
			log.Printf("TryRTMP failed to add voaacenc")
		}
		if ok := pipeline.Add(flvmux); ok != true {
			log.Printf("TryRTMP failed to add flvmux")
		}
		// os.Exit(1)
	} else {
		log.Printf("TryRTMP using pipeline.AddMany")
		if ok := pipeline.AddMany(videotestsrc, videoconvert, capsfilter1,
			queue, x264enc, capsfilter2,
			flvmux, rtmpsink, audiotestsrc, voaacenc, flvmux); ok != true {
			log.Fatalf("TryRTMP fatal: pipeline.AddMany returned %v", ok)
		}
		// os.Exit(1)
	}

	// TODO: need to call Link (or LinkMany) for each “!” in the gst-launch call
	/*
		y@RSWilli sya I'm calling pipeline.LinkMany which will fail, because this call will link
		all elements in the order they were passed. My gst-launch pipeline is linked
		differently. I need to call Link (or LinkMany) for each “!” in your gst-launch call.

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
	*/

	if ok := gst.LinkMany(
		// videotestsrc,
		videoconvert,
		capsfilter1,
		queue,
		x264enc,
		capsfilter2,
		flvmux,
		rtmpsink,
		// audiotestsrc,
		voaacenc,
		flvmux); ok != true {
		log.Fatalf("TryRTMP fatal: gst.LinkMany returned %v", ok)
	}

	log.Printf("TryRTMP end of creating pipeline")
	os.Exit(1) // no point proceeding without a real destination url
	return pipeline, nil
}

func mainLoop(pipeline gst.Pipeline) error {
	// Start the pipeline

	pipeline.SetState(gst.StatePlaying)

	for msg := range pipeline.GetBus().Messages(context.Background()) {
		switch msg.Type() {
		case gst.MessageEos:
			return nil
		case gst.MessageError:
			debug, gerr := msg.ParseError()
			if debug != "" {
				fmt.Println(gerr.Error(), debug)
			}
			return gerr
		default:
			fmt.Println(msg)
		}

		pipeline.DebugBinToDotFileWithTs(gst.DebugGraphShowVerbose, "pipeline")
	}

	return fmt.Errorf("unexpected end of messages without EOS")
}

func StartGtreamer() error {
	pipeline, err := createPipeline()
	if err != nil {
		return fmt.Errorf("TryRTMP error creating pipeline:%s", err)
	}

	log.Printf("TryRTMP mainLoop(pipeline) WILL START")
	err = mainLoop(pipeline)
	if err != nil {
		return fmt.Errorf("TryRTMP error running pipeline:%s", err)
	}
	log.Printf("TryRTMP HAS ENDED")
	return nil
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("TryRTMP -> %s\n", RTMP_DEST)
	if err := StartGtreamer(); err != nil {
		log.Fatalln(err)
	}

}
