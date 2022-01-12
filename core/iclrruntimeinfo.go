package core

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICLRRuntimeInfo struct {
	lpVtbl *ICLRRuntimeInfoVtbl
}

type ICLRRuntimeInfoVtbl struct {
	QueryInterface uintptr
	AddRef uintptr
	Release uintptr
	GetVersionString uintptr
	GetRuntimeDirectory uintptr
	IsLoaded uintptr
	LoadErrorString uintptr
	LoadLibrary uintptr
	GetProcAddress uintptr
	GetInterface uintptr
	IsLoadable uintptr
	SetDefaultStartupFlags uintptr
	GetDefaultStartupFlags uintptr
	BindAsLegacyV2Runtime uintptr
	IsStarted uintptr
}

func (obj *ICLRRuntimeInfo)GetRuntimeDirectory(){}
func (obj *ICLRRuntimeInfo)LoadErrorString(){}
func (obj *ICLRRuntimeInfo)LoadLibrary(){}
func (obj *ICLRRuntimeInfo)GetProcAddress(){}
func (obj *ICLRRuntimeInfo)QueryInterface(){}
func (obj *ICLRRuntimeInfo)SetDefaultStartupFlags(){}
func (obj *ICLRRuntimeInfo)GetDefaultStartupFlags(){}
func (obj *ICLRRuntimeInfo)IsStarted(){}
func (obj *ICLRRuntimeInfo)GetVersionString(){}
func (obj *ICLRRuntimeInfo)IsLoaded(){}


func NewICLRRuntimeInfo(ppv uintptr) *ICLRRuntimeInfo {
	return (*ICLRRuntimeInfo)(unsafe.Pointer(ppv))
}

func (obj *ICLRRuntimeInfo)AddRef()uintptr {
	ret,_,_ := syscall.Syscall(
		obj.lpVtbl.AddRef,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	return ret
}

func (obj *ICLRRuntimeInfo)Release()uintptr{
	ret,_,_ := syscall.Syscall(
		obj.lpVtbl.Release,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	return ret
}

func (obj *ICLRRuntimeInfo) BindAsLegacyV2Runtime() uintptr {
	ret, _, _ := syscall.Syscall(
		obj.lpVtbl.BindAsLegacyV2Runtime,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return ret
}

func (obj *ICLRRuntimeInfo) IsLoadable(pbLoadable *bool) uintptr {
	ret, _, _ := syscall.Syscall(
		obj.lpVtbl.IsLoadable,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(pbLoadable)),
		0)
	return ret
}

func (obj *ICLRRuntimeInfo) GetInterface(rclsid *windows.GUID, riid *windows.GUID, ppUnk *uintptr) uintptr {
	ret, _, _ := syscall.Syscall6(
		obj.lpVtbl.GetInterface,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(rclsid)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(ppUnk)),
		0,
		0)
	return ret
}