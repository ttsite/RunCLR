package core

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unicode/utf16"
	"unsafe"
)

//.NET Framework 版本:自4.0 之后可用

var (
	mscoree = syscall.NewLazyDLL("mscoree.dll")
	CLRCreateInstanceProc = mscoree.NewProc("CLRCreateInstance")
)

//https://docs.microsoft.com/zh-cn/dotnet/framework/unmanaged-api/hosting/clrcreateinstance-function

func CLRCreateInstance(clsid, riid *windows.GUID,ppInterface *uintptr) uintptr {
	//return three interface: ICLRMetaHost ICLRMetaHostPolicy ICLRDebugging
	Hresult,_, _ := CLRCreateInstanceProc.Call(
		uintptr(unsafe.Pointer(clsid)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(ppInterface)))
	return Hresult
}


func ReadUnicodeStr(ptr unsafe.Pointer) string {
	var byteVal uint16
	out := make([]uint16, 0)
	for i := 0; ; i++ {
		byteVal = *(*uint16)(unsafe.Pointer(ptr))
		if byteVal == 0x0000 {
			break
		}
		out = append(out, byteVal)
		ptr = unsafe.Pointer(uintptr(ptr) + 2)
	}
	return string(utf16.Decode(out))
}


func PrepareParameters(params []string) (*SafeArray, error) {
	sab := SafeArrayBound{
		cElements: uint32(len(params)),
		lLbound:   0,
	}
	listStrSafeArrayPtr, err := SafeArrayCreate(VT_BSTR, 1, &sab) // VT_BSTR
	if err != nil {
		return nil, err
	}
	for i, p := range params {
		bstr, err := SysAllocString(p)
		if err != nil {
			return nil, err
		}
		SafeArrayPutElement(listStrSafeArrayPtr, int32(i), bstr)
	}

	paramVariant := Variant{
		VT:  VT_BSTR | VT_ARRAY, // VT_BSTR | VT_ARRAY
		Val: uintptr(unsafe.Pointer(listStrSafeArrayPtr)),
	}

	sab2 := SafeArrayBound{
		cElements: uint32(1),
		lLbound:   0,
	}
	paramsSafeArrayPtr, err := SafeArrayCreate(VT_VARIANT, 1, &sab2) // VT_VARIANT
	if err != nil {
		return nil, err
	}
	err = SafeArrayPutElement(paramsSafeArrayPtr, int32(0), unsafe.Pointer(&paramVariant))
	if err != nil {
		return nil, err
	}
	return paramsSafeArrayPtr, nil
}