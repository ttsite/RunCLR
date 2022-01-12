package core

import (
	"fmt"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)


// https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nn-oaidl-isupporterrorinfo

type ISupportErrorInfo struct {
	lpVtbl *ISupportErrorInfoVtbl
}

type ISupportErrorInfoVtbl struct {
	// QueryInterface Retrieves pointers to the supported interfaces on an object.
	QueryInterface uintptr
	// AddRef Increments the reference count for an interface pointer to a COM object.
	// You should call this method whenever you make a copy of an interface pointer.
	AddRef uintptr
	// Release Decrements the reference count for an interface on a COM object.
	Release uintptr
	// InterfaceSupportsErrorInfo Indicates whether an interface supports the IErrorInfo interface.
	// https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-isupporterrorinfo-interfacesupportserrorinfo
	InterfaceSupportsErrorInfo uintptr
}

// https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-queryinterface(refiid_void)

func (obj *ISupportErrorInfo) QueryInterface(riid windows.GUID, ppvObject unsafe.Pointer) error {

	hr, _, err := syscall.Syscall(
		obj.lpVtbl.QueryInterface,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&riid)), // A reference to the interface identifier (IID) of the interface being queried for.
		uintptr(ppvObject),
	)
	if err != syscall.Errno(0) {
		return fmt.Errorf("the IUknown::QueryInterface method returned an error:\r\n%s", err)
	}
	if hr != 0x0 {
		return fmt.Errorf("the IUknown::QueryInterface method method returned a non-zero HRESULT: 0x%x", hr)
	}
	return nil
}


// https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref

func (obj *ISupportErrorInfo) AddRef() (count uint32, err error) {
	ret, _, err := syscall.Syscall(
		obj.lpVtbl.AddRef,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	if err != syscall.Errno(0) {
		return 0, fmt.Errorf("the IUnknown::AddRef method returned an error:\r\n%s", err)
	}
	err = nil
	// Unable to avoid misuse of unsafe.Pointer because the Windows API call returns the safeArray pointer in the "ret" value. This is a go vet false positive
	count = *(*uint32)(unsafe.Pointer(ret))
	return
}


// https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-release

func (obj *ISupportErrorInfo) Release() (count uint32, err error) {

	ret, _, err := syscall.Syscall(
		obj.lpVtbl.Release,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	if err != syscall.Errno(0) {
		return 0, fmt.Errorf("the IUnknown::Release method returned an error:\r\n%s", err)
	}
	err = nil
	// Unable to avoid misuse of unsafe.Pointer because the Windows API call returns the safeArray pointer in the "ret" value. This is a go vet false positive
	count = *(*uint32)(unsafe.Pointer(ret))
	return
}


// https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-isupporterrorinfo-interfacesupportserrorinfo

func (obj *ISupportErrorInfo) InterfaceSupportsErrorInfo(riid windows.GUID) error {

	hr, _, err := syscall.Syscall(
		obj.lpVtbl.InterfaceSupportsErrorInfo,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&riid)),
		0,
	)
	if err != syscall.Errno(0) {
		return fmt.Errorf("the ISupportErrorInfo::InterfaceSupportsErrorInfo method returned an error:\r\n%s", err)
	}
	if hr != 0x0 {
		return fmt.Errorf("the ISupportErrorInfo::InterfaceSupportsErrorInfo method method returned a non-zero HRESULT: 0x%x", hr)
	}
	return nil
}
