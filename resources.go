package resources

import _ "embed"

//go:embed libc++_shared.so
var Libcpp_shared []byte

//go:embed libopencv_java4.so
var Libopencv_java4 []byte

//go:embed libpaddle.so
var Libpaddle []byte

//go:embed libpaddle_native.so
var Libpaddle_native []byte

//go:embed libncnn.so
var Libncnn []byte

//go:embed libYoloV5Ncnn.so
var LibYoloV5Ncnn []byte

//go:embed ch_ppocr_mobile_v2.0_cls_opt.nb
var Ch_ppocr_mobile_cls_opt []byte

//go:embed ch_ppocr_mobile_v2.0_det_opt.nb
var Ch_ppocr_mobile_det_opt []byte

//go:embed ch_ppocr_mobile_v2.0_rec_opt.nb
var Ch_ppocr_mobile_rec_opt []byte

//go:embed ch_ppocr_keys.txt
var Ch_ppocr_keys []byte
