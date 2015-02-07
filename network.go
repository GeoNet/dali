package dali

//#cgo LDFLAGS: -ldali
//#include <libdali.h>
import "C"

import (
	"errors"
	"unsafe"
)

func (d *DLCP) Connect() (int, error) {
	rc := (int)(C.dl_connect((*C.struct_DLCP_s)(d)))
	if rc == -1 {
		return 0, errors.New("unable to connect to datalink server")
	}
	return rc, nil
}

func (d *DLCP) Disonnect() {
	C.dl_disconnect((*C.struct_DLCP_s)(d))
}

func (d *DLCP) Senddata(buffer []byte, sendlen uint) (int, error) {

	rc := (int)(C.dl_senddata((*C.struct_DLCP_s)(d), (unsafe.Pointer)(&buffer), (C.size_t)(sendlen)))
	if rc == -1 {
		return 0, errors.New("unable to send data")
	}
	return rc, nil
}

func (d *DLCP) Sendpacket(headerbuf []byte, headerlen uint, databuf []byte, datalen uint, respbuf []byte, resplen int) (int, error) {

	rc := (int)(C.dl_sendpacket((*C.struct_DLCP_s)(d), (unsafe.Pointer)(&headerbuf), (C.size_t)(headerlen), (unsafe.Pointer)(&databuf), (C.size_t)(datalen), (unsafe.Pointer)(&respbuf), (C.int)(resplen)))
	if rc == -1 {
		return 0, errors.New("unable to send packet")
	}
	return rc, nil
}

func (d *DLCP) Recvdata(buffer []byte, buflen uint, blockflag bool) (int, error) {

	var cblockflag uint8 = 0
	if blockflag {
		cblockflag = 1
	}

	rc := (int)(C.dl_recvdata((*C.struct_DLCP_s)(d), (unsafe.Pointer)(&buffer), (C.size_t)(buflen), (C.uint8_t)(cblockflag)))
	if rc == -1 {
		return 0, errors.New("unable to receive header, connection shutdown")
	} else if rc == -2 {
		return 0, errors.New("unable to receive header")
	}
	return 0, nil
}

func (d *DLCP) Recvheader(buffer []byte, buflen uint, blockflag bool) (int, error) {

	var cblockflag uint8 = 0
	if blockflag {
		cblockflag = 1
	}

	rc := (int)(C.dl_recvheader((*C.struct_DLCP_s)(d), (unsafe.Pointer)(&buffer), (C.size_t)(buflen), (C.uint8_t)(cblockflag)))
	if rc == -1 {
		return 0, errors.New("unable to receive header, connection shutdown")
	} else if rc == -2 {
		return 0, errors.New("unable to receive header")
	}
	return 0, nil
}
