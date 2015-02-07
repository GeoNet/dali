## go wrapper for libdali ##

ftp://ftp.iris.washington.edu/pub/programs/ringserver/

An initial minimal wrapper for the libdali library used
to create clients for connecting to datalink servers.

The library needs to be compiled and placed, together with
the libdali.h and portable.h files, somewhere where
the go build (cgo) routines can find them.

The logging aspect of the library hasn't been wrapped.

To test the build an operational ringserver should be
running locally on port 16000.

Mark Chadwick
