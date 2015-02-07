package dali

//#cgo CFLAGS: -I/usr/local/include
//#cgo LDFLAGS: /usr/local/lib/libdali.a
//#include <libdali.h>
import "C"

import ()

// just use default values for now, go and c vargs don't play nicely ...

func Loginit(verbosity int) {
	C.dl_loginit((C.int)(verbosity), nil, nil, nil, nil)
}

func (d *DLCP) LoginitR(verbosity int) {
	C.dl_loginit_r((*C.struct_DLCP_s)(d), (C.int)(verbosity), nil, nil, nil, nil)
}
