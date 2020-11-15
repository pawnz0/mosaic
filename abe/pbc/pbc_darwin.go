// +build darwin,cgo,pbc,!wasm

package pbc

/*
#cgo CFLAGS: -std=gnu99
#cgo LDFLAGS: -lpbc -lgmp
#include <pbc/pbc.h>
#include "gmp.h"

#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <stdint.h>
#include <unistd.h>
#include <errno.h>

typedef struct memstream_s {
	char*  buffer;
	size_t cap;
	size_t len;
	fpos_t cursor;
	FILE*  fd;
	int    closed;
} memstream_t;

memstream_t* pbc_open_memstream();
FILE* pbc_memstream_to_fd(memstream_t* m);
int pbc_close_memstream(memstream_t* m, char** bufp, size_t* sizep);

static void memstream_realloc(memstream_t* m, size_t size) {
	m->cap = size;
	m->buffer = realloc(m->buffer, size);
	if (size < m->len) {
		m->len = size;
		m->cursor = m->len;
	}
}

static int memstream_read(void* cookie, char* buf, int nbytes) {
	memstream_t* m = (memstream_t*)cookie;
	if (m->closed) { errno = EBADF; return -1; }
	if (m->cursor >= m->len) return 0;
	size_t toRead = (m->len - m->cursor);
	if (toRead > (size_t)nbytes) toRead = (size_t)nbytes;
	memcpy(buf, &m->buffer[m->cursor], toRead);
	m->cursor += toRead;
	return toRead;
}

static int memstream_write(void* cookie, const char* buf, int nbytes) {
	memstream_t* m = (memstream_t*)cookie;
	if (m->closed) { errno = EBADF; return -1; }
	size_t zeros = m->cursor - m->len;
	size_t dataLen = (size_t)nbytes;
	size_t neededSpace = zeros + dataLen;
	if (neededSpace < dataLen) {
		errno = EFBIG;
		return -1;
	}
	size_t newLen = m->len + neededSpace;
	if (newLen < m->len) {
		errno = EFBIG;
		return -1;
	}
	if (newLen > m->cap) {
		size_t newCap = m->cap;
		do {
			if (SIZE_MAX - newCap < newCap) {
				newCap = SIZE_MAX;
			} else {
				newCap <<= 1;
			}
		} while (newLen > newCap);
		memstream_realloc(m, newCap);
	}
	if (zeros > 0) {
		memset(&m->buffer[m->len], 0, zeros);
	}
	memcpy(&m->buffer[m->cursor], buf, dataLen);
	m->cursor += dataLen;
	m->len = newLen;
	return neededSpace;
}

static fpos_t memstream_seek(void* cookie, fpos_t offset, int whence) {
	memstream_t* m = (memstream_t*)cookie;
	if (m->closed) { errno = EBADF; return -1; }
	fpos_t base = 0;
	switch (whence) {
		case SEEK_SET: base = 0; break;
		case SEEK_CUR: base = m->cursor; break;
		case SEEK_END: base = m->len; break;
		// SEEK_HOLE and SEEK_DATA are not supported on darwin
		default: errno = EINVAL; return -1;
	}
	fpos_t desired = base + offset;
	if (desired < 0) {
		errno = EINVAL;
		return -1;
	}
	if (offset > 0 && desired < base) {
		errno = EOVERFLOW;
		return -1;
	}
	return (m->cursor = desired);
}

static int memstream_close(void* cookie) {
	memstream_t* m = (memstream_t*)cookie;
	m->closed = 1;
	return 0;
}

memstream_t* pbc_open_memstream() {
	memstream_t* m = malloc(sizeof(memstream_t));
	m->buffer = NULL;
	memstream_realloc(m, 1024);
	m->len = 0;
	m->cursor = 0;
	m->closed = 0;
	m->fd = funopen(m, memstream_read, memstream_write, memstream_seek, memstream_close);
	return m;
}

FILE* pbc_memstream_to_fd(memstream_t* m) { return m->fd; }
int pbc_close_memstream(memstream_t* m, char** bufp, size_t* sizep) {
	fclose(m->fd);
	*bufp = m->buffer;
	*sizep = m->len;
	free(m);
	return 1;
}

mpz_t* newMpzT() {
	mpz_t* x = malloc(sizeof(mpz_t));
	mpz_init(*x);
	return x;
}

void freeMpzT(mpz_t* x) {
	mpz_clear(*x);
	free(x);
}

struct pbc_param_s* newParamStruct() {
  return malloc(sizeof(struct pbc_param_s));
}

void freeParamStruct(struct pbc_param_s* x) {
	pbc_param_clear(x);
	free(x);
}

int param_out_str_wrapper(char** bufp, size_t* sizep, pbc_param_t p) {
	memstream_t* stream = pbc_open_memstream();
	if (stream == NULL) return 0;
	pbc_param_out_str(pbc_memstream_to_fd(stream), p);
	return pbc_close_memstream(stream, bufp, sizep);
}

struct pairing_s* newPairingStruct() {
  return malloc(sizeof(struct pairing_s));
}

void freePairingStruct(struct pairing_s* x) {
	pairing_clear(x);
	free(x);
}

struct element_s* newElementStruct() {
  return malloc(sizeof(struct element_s));
}

void freeElementStruct(struct element_s* x) {
	element_clear(x);
	free(x);
}

size_t wordsize() {
  return sizeof(size_t);
}
*/
import "C"

import (
	"math/big"
	"runtime"
	"unsafe"
)

type mpz struct {
	p *C.mpz_t
}

func freeMpz(x *mpz) {
	C.freeMpzT(x.p)
}

func newMpz() *mpz {
	out := &mpz{p: C.newMpzT()}
	runtime.SetFinalizer(out, freeMpz)
	return out
}

func big2mpz(num *big.Int) *mpz {
	out := newMpz()
	words := num.Bits()
	if len(words) > 0 {
		C.mpz_import(&out.p[0], C.size_t(len(words)), -1, C.wordsize(), 0, 0, unsafe.Pointer(&words[0]))
	}
	return out
}

type Params struct {
	cptr *C.struct_pbc_param_s
}

func freeParams(params *Params) {
	C.freeParamStruct(params.cptr)
}

func makeParams() *Params {
	params := &Params{cptr: C.newParamStruct()}
	runtime.SetFinalizer(params, freeParams)
	return params
}

func NewParamsFromString(s string) *Params {
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	params := makeParams()
	C.pbc_param_init_set_str(params.cptr, cstr)
	return params
}

func GenerateA(rbits uint32, qbits uint32) *Params {
	params := makeParams()
	C.pbc_param_init_a_gen(params.cptr, C.int(rbits), C.int(qbits))
	return params
}

func GenerateF(bits uint32) *Params {
	params := makeParams()
	C.pbc_param_init_f_gen(params.cptr, C.int(bits))
	return params
}

func (params *Params) String() string {
	var buf *C.char
	var bufLen C.size_t
	if C.param_out_str_wrapper(&buf, &bufLen, params.cptr) == 0 {
		return ""
	}
	str := C.GoStringN(buf, C.int(bufLen))
	C.free(unsafe.Pointer(buf))
	return str
}

type Pairing struct {
	params *Params
	cptr   *C.struct_pairing_s
}

func freePairing(pairing *Pairing) {
	C.freePairingStruct(pairing.cptr)
}

func makePairing(params *Params) *Pairing {
	pairing := &Pairing{
		params: params, cptr: C.newPairingStruct()}
	runtime.SetFinalizer(pairing, freePairing)
	return pairing
}

func NewPairing(params *Params) *Pairing {
	pairing := makePairing(params)
	C.pairing_init_pbc_param(pairing.cptr, params.cptr)
	pairing.params = nil
	return pairing
}

func (pairing *Pairing) NewUncheckedElement(group string) *Element {
	return makeUncheckedElement(pairing, group)
}

type Element struct {
	pairing   *Pairing
	cptr      *C.struct_element_s
	fieldPtr  *C.struct_field_s
	isInteger bool
}

func freeElement(element *Element) {
	C.freeElementStruct(element.cptr)
}

func makeUncheckedElement(pairing *Pairing, group string) *Element {
	element := &Element{
		cptr: C.newElementStruct(), pairing: pairing}

	switch group {
	case "G1":
		C.element_init_G1(element.cptr, pairing.cptr)
	case "G2":
		C.element_init_G2(element.cptr, pairing.cptr)
	case "GT":
		C.element_init_GT(element.cptr, pairing.cptr)
	case "Zr":
		C.element_init_Zr(element.cptr, pairing.cptr)
	}

	runtime.SetFinalizer(element, freeElement)
	return element
}

func (e *Element) BytesLen() int {
	return int(C.element_length_in_bytes(e.cptr))
}

func (e *Element) Bytes() []byte {
	buf := make([]byte, e.BytesLen())
	C.element_to_bytes((*C.uchar)(unsafe.Pointer(&buf[0])), e.cptr)
	return buf
}

func (e *Element) SetBytes(buf []byte) *Element {
	C.element_from_bytes(e.cptr, (*C.uchar)(unsafe.Pointer(&buf[0])))
	return e
}

func (e *Element) SetFromHash(hash []byte) *Element {
	C.element_from_hash(e.cptr, unsafe.Pointer(&hash[0]), C.int(len(hash)))
	return e
}

func (e *Element) Rand() *Element {
	C.element_random(e.cptr)
	return e
}

func (e *Element) Set1() *Element {
	C.element_set1(e.cptr)
	return e
}

func (e *Element) Mul(x *Element, y *Element) *Element {
	C.element_mul(e.cptr, x.cptr, y.cptr)
	return e
}

func (e *Element) Div(x *Element, y *Element) *Element {
	C.element_div(e.cptr, x.cptr, y.cptr)
	return e
}

func (e *Element) Invert(x *Element) *Element {
	C.element_invert(e.cptr, x.cptr)
	return e
}

func (e *Element) ThenInvert() *Element {
	return e.Invert(e)
}

func (e *Element) PowBig(x *Element, i *big.Int) *Element {

	C.element_pow_mpz(e.cptr, x.cptr, &big2mpz(i).p[0])
	return e
}

func (e *Element) Pair(x *Element, y *Element) *Element {
	C.pairing_apply(e.cptr, x.cptr, y.cptr, e.pairing.cptr)
	return e
}
