package dali

//#cgo LDFLAGS: -ldali
//#include <libdali.h>
//#include <portable.h>
import "C"

import (
	"errors"
	"unsafe"
)

const (
	DLERROR    int = -1
	DLENDED    int = 0
	DLPACKET   int = 1
	DLNOPACKET int = 2
)

type DLCP _Ctype_DLCP
type DLPacket _Ctype_DLPacket
type DLTime int64

func (d *DLCP) Pktid() int64 {
	return (int64)(d.pktid)
}
func (d *DLCP) Pkttime() DLTime {
	return (DLTime)(d.pkttime)
}

func NewDLCP(address string, progname string) (*DLCP, error) {
	caddress := C.CString(address)
	defer C.free(unsafe.Pointer(caddress))
	cprogname := C.CString(progname)
	defer C.free(unsafe.Pointer(cprogname))
	ptr := (*DLCP)(C.dl_newdlcp(caddress, cprogname))
	if ptr == nil {
		return nil, errors.New("Unable to create new DLCP struct")
	}
	return ptr, nil
}

func FreeDLCP(d *DLCP) {
	C.dl_freedlcp((*C.struct_DLCP_s)(d))
}

func (d *DLCP) ExchangeIDs(parseresp int) (int, error) {
	rc := (int)(C.dl_exchangeIDs((*C.struct_DLCP_s)(d), (C.int)(parseresp)))
	if rc == -1 {
		return 0, errors.New("an error has occurred")
	}
	return rc, nil
}

func (d *DLCP) Position(pktid int64, pkttime DLTime) (int64, error) {
	rc := (int64)(C.dl_position((*C.struct_DLCP_s)(d), (C.int64_t)(pktid), (C.dltime_t)(pkttime)))
	if rc == -1 {
		return 0, errors.New("an error has occurred")
	}
	return rc, nil
}

func (d *DLCP) PositionAfter(datatime DLTime) (int64, error) {
	rc := (int64)(C.dl_position_after((*C.struct_DLCP_s)(d), (C.dltime_t)(datatime)))
	if rc == -1 {
		return 0, errors.New("an error has occurred")
	}
	return rc, nil
}

func (d *DLCP) Match(matchpattern string) (int64, error) {
	cmatchpattern := C.CString(matchpattern)
	defer C.free(unsafe.Pointer(cmatchpattern))
	rc := (int64)(C.dl_match((*C.struct_DLCP_s)(d), cmatchpattern))
	if rc == -1 {
		return 0, errors.New("an error has occurred")
	}
	return rc, nil
}

func (d *DLCP) Reject(rejectpattern string) (int64, error) {
	crejectpattern := C.CString(rejectpattern)
	defer C.free(unsafe.Pointer(crejectpattern))
	rc := (int64)(C.dl_reject((*C.struct_DLCP_s)(d), crejectpattern))
	if rc == -1 {
		return 0, errors.New("an error has occurred")
	}
	return rc, nil
}

func (d *DLCP) Write(packet []byte, packetlen int, streamid string, datastart DLTime, dataend DLTime, ack bool) (int64, error) {
	var cack int

	if ack {
		cack = 1
	} else {
		cack = 0
	}

	cstreamid := C.CString(streamid)
	defer C.free(unsafe.Pointer(cstreamid))
	rc := (int64)(C.dl_write((*C.struct_DLCP_s)(d), (unsafe.Pointer)(&packet), (C.int)(packetlen), cstreamid, (C.dltime_t)(datastart), (C.dltime_t)(dataend), (C.int)(cack)))
	if rc == -1 {
		return 0, errors.New("an error has occurred")
	}
	if ack && rc == 0 {
		return 0, errors.New("an error has occurred")
	}
	return rc, nil
}

func (d *DLCP) Read(pktid int64, packet *DLPacket, packetdata []byte, maxdatasize uint) (int, error) {
	rc := (int)(C.dl_read((*C.struct_DLCP_s)(d), (C.int64_t)(pktid), (*C.struct_DLPacket_s)(packet), (unsafe.Pointer)(&packetdata), (C.size_t)(maxdatasize)))
	if rc == -1 {
		return 0, errors.New("an error has occurred")
	}
	return rc, nil
}

func (d *DLCP) Getinfo(infotype string, infomatch string, maxinfosize uint32) ([]string, error) {
	cinfotype := C.CString(infotype)
	defer C.free(unsafe.Pointer(cinfotype))
	cinfomatch := C.CString(infomatch)
	defer C.free(unsafe.Pointer(cinfomatch))
	var cinfodata **C.char

	var infodata []string

	rc := (int)(C.dl_getinfo((*C.struct_DLCP_s)(d), cinfotype, cinfomatch, (cinfodata), (C.size_t)(maxinfosize)))
	if rc == DLERROR {
		return infodata, errors.New("an error has occurred")
	}
	infodata = make([]string, rc, rc)
	for i := 0; i < rc; i++ {
		infodata[i] = ""
	}
	return infodata, nil
}

func (d *DLCP) Collect(packet *DLPacket, packetdata []byte, maxdatasize uint32, endflag int8) (int, error) {
	rc := (int)(C.dl_collect((*C.struct_DLCP_s)(d), (*C.struct_DLPacket_s)(packet), (unsafe.Pointer)(&packetdata), (C.size_t)(maxdatasize), (C.int8_t)(endflag)))
	if rc == DLERROR {
		return 0, errors.New("an error has occurred")
	}
	return rc, nil
}

func (d *DLCP) CollectNB(packet *DLPacket, packetdata []byte, maxdatasize uint32, endflag int8) (int, error) {
	rc := (int)(C.dl_collect_nb((*C.struct_DLCP_s)(d), (*C.struct_DLPacket_s)(packet), (unsafe.Pointer)(&packetdata), (C.size_t)(maxdatasize), (C.int8_t)(endflag)))
	if rc == DLERROR {
		return 0, errors.New("an error has occurred")
	}
	return rc, nil
}

func (d *DLCP) Handlereply(buffer []byte, buflen int) (int64, error) {
	var value C.int64_t
	rc := (int)(C.dl_handlereply((*C.struct_DLCP_s)(d), (unsafe.Pointer)(&buffer), (C.int)(buflen), &value))
	if rc == -1 {
		return 0, errors.New("unable to handle reply")
	}
	if rc == 1 {
		return 0, errors.New("ERROR reply received")
	}
	return (int64)(value), nil
}

func (d *DLCP) Terminate() {
	C.dl_terminate((*C.struct_DLCP_s)(d))
}

func (d *DLCP) IsTerminated() bool {
	if d.terminate != 0 {
		return true
	}
	return false
}
