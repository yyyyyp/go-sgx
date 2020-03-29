package main

// #cgo CFLAGS: -I${SRCDIR}/App -I/opt/intel/sgxsdk/include/
// #cgo LDFLAGS: /opt/intel/sgxsdk/lib64/libsgx_urts.so ${SRCDIR}/libapp.a
// #include<App.h>
import "C"

import (
        "fmt"
)

func main() {
    double := C.ff()
    fmt.Printf("%d \n",  double)
}

