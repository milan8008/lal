// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package hevc_test

import (
	"testing"

	"github.com/q191201771/lal/pkg/hevc"
	"github.com/q191201771/naza/pkg/assert"
)

// https://github.com/ksvc/FFmpeg/blob/release/3.3/libavformat/hevc.c#L936
var goldenSH = []byte{
	0x1c, 0x00, 0x00, 0x00, 0x00,
	0x01, // configurationVersion
	0x01,
	0x60, 0x00, 0x00, 0x00,
	0x90, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x5d,
	0xf0, 0x00,
	0xfc,
	0xfd,
	0xf8,
	0xf8,
	0x00, 0x00, // favgFrameRate
	0x0f,
	0x03, // numOfArrarys
	0x20, 0x00, 0x01, 0x00, 0x17,
	0x40, 0x01, 0x0c, 0x01, 0xff, 0xff, 0x01, 0x60, 0x00, 0x00, 0x03, 0x00, 0xb0, 0x00, 0x00, 0x03, 0x00, 0x00, 0x03, 0x00, 0x5d, 0xac, 0x09,
	0x21, 0x00, 0x01, 0x00, 0x25,
	0x42, 0x01, 0x01, 0x01, 0x60, 0x00, 0x00, 0x03, 0x00, 0xb0, 0x00, 0x00, 0x03, 0x00, 0x00, 0x03, 0x00, 0x5d, 0xa0, 0x04, 0x82, 0x00, 0x40, 0x16, 0x5a, 0xee, 0x4c, 0x92, 0xea, 0x52, 0x0a, 0x0c, 0x0c, 0x05, 0xda, 0x14, 0x25,
	0x22, 0x00, 0x01, 0x00, 0x08,
	0x44, 0x01, 0xc0, 0xe3, 0x0f, 0x03, 0xb0, 0x84,
}

var goldenVPS = []byte{
	0x40, 0x01, 0x0c, 0x01, 0xff, 0xff, 0x01, 0x60, 0x00, 0x00, 0x03, 0x00, 0xb0, 0x00, 0x00, 0x03, 0x00, 0x00, 0x03, 0x00, 0x5d, 0xac, 0x09,
}

var goldenSPS = []byte{
	0x42, 0x01, 0x01, 0x01, 0x60, 0x00, 0x00, 0x03, 0x00, 0xb0, 0x00, 0x00, 0x03, 0x00, 0x00, 0x03, 0x00, 0x5d, 0xa0, 0x04, 0x82, 0x00, 0x40, 0x16, 0x5a, 0xee, 0x4c, 0x92, 0xea, 0x52, 0x0a, 0x0c, 0x0c, 0x05, 0xda, 0x14, 0x25,
}

var goldenPPS = []byte{
	0x44, 0x01, 0xc0, 0xe3, 0x0f, 0x03, 0xb0, 0x84,
}

var goldenVPSSPSPPSAnnexB = []byte{
	0x00, 0x00, 0x00, 0x01,
	0x40, 0x01, 0x0c, 0x01, 0xff, 0xff, 0x01, 0x60, 0x00, 0x00, 0x03, 0x00, 0xb0, 0x00, 0x00, 0x03, 0x00, 0x00, 0x03, 0x00, 0x5d, 0xac, 0x09,
	0x00, 0x00, 0x00, 0x01,
	0x42, 0x01, 0x01, 0x01, 0x60, 0x00, 0x00, 0x03, 0x00, 0xb0, 0x00, 0x00, 0x03, 0x00, 0x00, 0x03, 0x00, 0x5d, 0xa0, 0x04, 0x82, 0x00, 0x40, 0x16, 0x5a, 0xee, 0x4c, 0x92, 0xea, 0x52, 0x0a, 0x0c, 0x0c, 0x05, 0xda, 0x14, 0x25,
	0x00, 0x00, 0x00, 0x01,
	0x44, 0x01, 0xc0, 0xe3, 0x0f, 0x03, 0xb0, 0x84,
}

func TestParseVPSSPSPPSFromSeqHeader(t *testing.T) {
	vps, sps, pps, err := hevc.ParseVPSSPSPPSFromSeqHeader(goldenSH)
	assert.Equal(t, nil, err)
	assert.Equal(t, goldenVPS, vps)
	assert.Equal(t, goldenSPS, sps)
	assert.Equal(t, goldenPPS, pps)
}

func TestVPSSPSPPSSeqHeader2AnnexB(t *testing.T) {
	out, err := hevc.VPSSPSPPSSeqHeader2AnnexB(goldenSH)
	assert.Equal(t, nil, err)
	assert.Equal(t, goldenVPSSPSPPSAnnexB, out)
}
