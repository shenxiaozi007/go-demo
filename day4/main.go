package day4

import "unsafe"

type waitq struct {

}

type hchan struct {
    qcount uint
    dataqsiz uint
    buf unsafe.Pointer
    elemsize uint16
    closed uint32
    sendx uint
    recvx uint
    recvq waitq
    sendq waitq
}