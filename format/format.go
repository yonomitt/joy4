package format

import (
	"github.com/yonomitt/joy4/format/mp4"
	"github.com/yonomitt/joy4/format/ts"
	"github.com/yonomitt/joy4/format/rtmp"
	"github.com/yonomitt/joy4/format/rtsp"
	"github.com/yonomitt/joy4/format/flv"
	"github.com/yonomitt/joy4/format/aac"
	"github.com/yonomitt/joy4/av/avutil"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}

