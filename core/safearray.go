package core

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"syscall"
	"unsafe"
)

type SafeArray struct {
	cDims      uint16
	fFeatures  uint16
	cbElements uint32
	cLocks     uint32
	pvData     uintptr
	rgsabound  [1]SafeArrayBound
}

type SafeArrayBound struct {
	cElements uint32
	lLbound   int32
}

func CreateSafeArray(rawBytes []byte) (unsafe.Pointer, error) {
	saPtr, err := CreateEmptySafeArray(0x11, len(rawBytes)) // VT_UI1
	if err != nil {
		return nil, err
	}
	// now we need to use RtlCopyMemory to copy our bytes to the SafeArray
	modNtDll := syscall.MustLoadDLL("ntdll.dll")
	procRtlCopyMemory := modNtDll.MustFindProc("RtlCopyMemory")
	sa := (*SafeArray)(saPtr)
	_, _, err = procRtlCopyMemory.Call(
		sa.pvData,
		uintptr(unsafe.Pointer(&rawBytes[0])),
		uintptr(len(rawBytes)))
	if err != syscall.Errno(0) {
		return nil, err
	}
	return saPtr, nil

}


func CreateEmptySafeArray(arrayType int, size int) (unsafe.Pointer, error) {
	modOleAuto := syscall.MustLoadDLL("OleAut32.dll")
	procSafeArrayCreate := modOleAuto.MustFindProc("SafeArrayCreate")

	sab := SafeArrayBound{
		cElements: uint32(size),
		lLbound:   0,
	}
	vt := uint16(arrayType)
	ret, _, err := procSafeArrayCreate.Call(
		uintptr(vt),
		uintptr(1),
		uintptr(unsafe.Pointer(&sab)))

	if err != syscall.Errno(0) {
		return nil, err
	}

	return unsafe.Pointer(ret), nil
}

func SafeArrayCreate(vt uint16, cDims uint32, rgsabound *SafeArrayBound) (safeArray *SafeArray, err error) {
	modOleAuto := syscall.MustLoadDLL("OleAut32.dll")
	procSafeArrayCreate := modOleAuto.MustFindProc("SafeArrayCreate")

	ret, _, err := procSafeArrayCreate.Call(
		uintptr(vt),
		uintptr(cDims),
		uintptr(unsafe.Pointer(rgsabound)),
	)

	if err != syscall.Errno(0) {
		return
	}
	err = nil

	if ret == 0 {
		err = fmt.Errorf("the OleAut32!SafeArrayCreate function return 0x%x and the SafeArray was not created", ret)
		return
	}

	// Unable to avoid misuse of unsafe.Pointer because the Windows API call returns the safeArray pointer in the "ret" value. This is a go vet false positive
	safeArray = (*SafeArray)(unsafe.Pointer(ret))
	return
}

func SysAllocString(str string) (unsafe.Pointer, error) {

	modOleAuto := syscall.MustLoadDLL("OleAut32.dll")
	sysAllocString := modOleAuto.MustFindProc("SysAllocString")

	input := utf16Le(str)
	ret, _, err := sysAllocString.Call(
		uintptr(unsafe.Pointer(&input[0])),
	)

	if err != syscall.Errno(0) {
		return nil, err
	}
	// TODO Return a pointer to a BSTR instead of an unsafe.Pointer
	// Unable to avoid misuse of unsafe.Pointer because the Windows API call returns the safeArray pointer in the "ret" value. This is a go vet false positive
	return unsafe.Pointer(ret), nil
}


func utf16Le(s string) []byte {
	enc := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewEncoder()
	var buf bytes.Buffer
	t := transform.NewWriter(&buf, enc)
	_, err := t.Write([]byte(s))
	if err != nil {
		return nil
	}
	return buf.Bytes()
}

func SafeArrayPutElement(psa *SafeArray, rgIndices int32, pv unsafe.Pointer) error {

	modOleAuto := syscall.MustLoadDLL("OleAut32.dll")
	safeArrayPutElement := modOleAuto.MustFindProc("SafeArrayPutElement")

	hr, _, err := safeArrayPutElement.Call(
		uintptr(unsafe.Pointer(psa)),
		uintptr(unsafe.Pointer(&rgIndices)),
		uintptr(pv),
	)
	if err != syscall.Errno(0) {
		return err
	}
	if hr != 0x0 {
		return fmt.Errorf("the OleAut32!SafeArrayPutElement call returned a non-zero HRESULT: 0x%x", hr)
	}
	return nil
}