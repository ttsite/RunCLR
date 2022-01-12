package core

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type IUnknown struct {
	lpVtbl *IUnknownVtbl
}

type IUnknownVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
}

func NewIUnknown(ppv uintptr) *IUnknown {
	return (*IUnknown)(unsafe.Pointer(ppv))
}

func (obj *IUnknown) QueryInterface(riid *windows.GUID, ppvObject *uintptr) uintptr {
	ret, _, _ := syscall.Syscall(
		obj.lpVtbl.QueryInterface,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(ppvObject)))
	return ret
}

func (obj *IUnknown) AddRef() uintptr {
	ret, _, _ := syscall.Syscall(
		obj.lpVtbl.AddRef,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	return ret
}

func (obj *IUnknown) Release() uintptr {
	ret, _, _ := syscall.Syscall(
		obj.lpVtbl.Release,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	return ret
}