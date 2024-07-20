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

//go:embed libyolov5.so
var Libyolov5 []byte

//go:embed libyolov5_native.so
var Libyolov5_native []byte
