package core

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

//https://stackoverflow.com/questions/37781676/how-to-use-com-component-object-model-in-golang
//https://docs.microsoft.com/zh-cn/dotnet/framework/unmanaged-api/hosting/iclrmetahost-interface

type ICLRMetaHost struct {
	lpVtbl *ICLRMetaHostVtbl
}

func NewICLRMetaHost(ptr uintptr) *ICLRMetaHost {
	return (*ICLRMetaHost)(unsafe.Pointer(ptr))
}

//metahost.h

type ICLRMetaHostVtbl struct {
	QueryInterface uintptr
	AddRef uintptr
	Release uintptr
	GetRuntime uintptr
	GetVersionFromFile uintptr
	EnumerateInstalledRuntimes uintptr
	EnumerateLoadedRuntimes uintptr
	RequestRuntimeLoadedNotification uintptr
	QueryLegacyV2RuntimeBinding uintptr
	ExitProcess uintptr
}



func (obj *ICLRMetaHost)EnumerateLoadedRuntimes(){}
func (obj *ICLRMetaHost)ExitProcess(){}
func (obj *ICLRMetaHost)QueryInterface(){}
func (obj *ICLRMetaHost)GetVersionFromFile(){}
func (obj *ICLRMetaHost)RequestRuntimeLoadedNotification(){}
func (obj *ICLRMetaHost)QueryLegacyV2RuntimeBinding(){}


//https://docs.microsoft.com/zh-cn/dotnet/framework/unmanaged-api/hosting/iclrmetahost-getruntime-method

func (obj *ICLRMetaHost)GetRuntime(pwzVersion *uint16, riid *windows.GUID,ppRuntime *uintptr) uintptr {
	ret,_,_ := syscall.Syscall6(
		obj.lpVtbl.GetRuntime,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(pwzVersion)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(ppRuntime)),
		0,
		0,
		)
	return ret
}

func (obj *ICLRMetaHost)AddRef() uintptr {
	ret,_,_ := syscall.Syscall(
		obj.lpVtbl.AddRef,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	return ret
}

func (obj *ICLRMetaHost)Release() uintptr {
	ret,_,_ := syscall.Syscall(
		obj.lpVtbl.Release,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	return ret
}

//https://docs.microsoft.com/zh-cn/dotnet/framework/unmanaged-api/hosting/iclrmetahost-enumerateinstalledruntimes-method

func (obj *ICLRMetaHost)EnumerateInstalledRuntimes(ppEnumerator *uintptr) uintptr {
	ret, _,_ := syscall.Syscall(
		obj.lpVtbl.EnumerateInstalledRuntimes,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(ppEnumerator)),
		0)
	return ret
}