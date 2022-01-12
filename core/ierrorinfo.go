package core

import (
	"fmt"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type IErrorInfo struct {
	lpVtbl *IErrorInfoVtbl
}

// https://docs.microsoft.com/en-us/previous-versions/windows/desktop/ms723041(v=vs.85)

type IErrorInfoVtbl struct {
	// QueryInterface Retrieves pointers to the supported interfaces on an object.
	QueryInterface uintptr
	// AddRef Increments the reference count for an interface pointer to a COM object.
	// You should call this method whenever you make a copy of an interface pointer.
	AddRef uintptr
	// Release Decrements the reference count for an interface on a COM object.
	Release uintptr
	// GetDescription Returns a text description of the error
	GetDescription uintptr
	// GetGUID Returns the GUID of the interface that defined the error.
	GetGUID uintptr
	// GetHelpContext Returns the Help context ID for the error.
	GetHelpContext uintptr
	// GetHelpFile Returns the path of the Help file that describes the error.
	GetHelpFile uintptr
	// GetSource Returns the name of the component that generated the error, such as "ODBC driver-name".
	GetSource uintptr
}

// https://docs.microsoft.com/en-us/previous-versions/windows/desktop/ms714318(v=vs.85)
func (obj *IErrorInfo) GetDescription() (pbstrDescription *string, err error) {
	hr, _, err := syscall.Syscall(
		obj.lpVtbl.GetDescription,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&pbstrDescription)),
		0,
	)

	if err != syscall.Errno(0) {
		err = fmt.Errorf("the IErrorInfo::GetDescription method returned an error:%s", err)
		return
	}
	if hr != 0x00 {
		err = fmt.Errorf("the IErrorInfo::GetDescription method method returned a non-zero HRESULT: 0x%x", hr)
		return
	}
	err = nil
	return
}


// https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-ierrorinfo-getguid

func (obj *IErrorInfo) GetGUID() (pGUID *windows.GUID, err error) {
	hr, _, err := syscall.Syscall(
		obj.lpVtbl.GetGUID,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(pGUID)),
		0,
	)

	if err != syscall.Errno(0) {
		err = fmt.Errorf("the IErrorInfo::GetGUID method returned an error:%s", err)
		return
	}
	if hr != 0x00 {
		err = fmt.Errorf("the IErrorInfo::GetGUID method method returned a non-zero HRESULT: 0x%x", hr)
		return
	}
	err = nil
	return
}


// https://docs.microsoft.com/en-us/windows/win32/api/oleauto/nf-oleauto-geterrorinfo

func GetErrorInfo() (pperrinfo *IErrorInfo, err error) {
	modOleAut32 := syscall.MustLoadDLL("OleAut32.dll")
	procGetErrorInfo := modOleAut32.MustFindProc("GetErrorInfo")
	hr, _, err := procGetErrorInfo.Call(0, uintptr(unsafe.Pointer(&pperrinfo)))
	if err != syscall.Errno(0) {
		err = fmt.Errorf("the OleAu32.GetErrorInfo procedure call returned an error:%s", err)
		return
	}
	if hr != 0x00 {
		err = fmt.Errorf("the OleAu32.GetErrorInfo procedure call returned a non-zero HRESULT code: 0x%x", hr)
		return
	}
	err = nil
	return
}
