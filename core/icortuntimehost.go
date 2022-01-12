package core

import (
	"syscall"
	"unsafe"
)

type ICORRuntimeHost struct {
	lpVtbl *ICORRuntimeHostVtbl
}

type ICORRuntimeHostVtbl struct {
	QueryInterface                uintptr
	AddRef                        uintptr
	Release                       uintptr
	CreateLogicalThreadState      uintptr
	DeleteLogicalThreadState      uintptr
	SwitchInLogicalThreadState    uintptr
	SwitchOutLogicalThreadState   uintptr
	LocksHeldByLogicalThreadState uintptr
	MapFile                       uintptr
	GetConfiguration              uintptr
	Start                         uintptr
	Stop                          uintptr
	CreateDomain                  uintptr
	GetDefaultDomain              uintptr
	EnumDomains                   uintptr
	NextDomain                    uintptr
	CloseEnum                     uintptr
	CreateDomainEx                uintptr
	CreateDomainSetup             uintptr
	CreateEvidence                uintptr
	UnloadDomain                  uintptr
	CurrentDomain                 uintptr
}

func (obj *ICORRuntimeHost)QueryInterface(){}
func (obj *ICORRuntimeHost)CreateLogicalThreadState(){}
func (obj *ICORRuntimeHost)DeleteLogicalThreadState(){}
func (obj *ICORRuntimeHost)SwitchInLogicalThreadState(){}
func (obj *ICORRuntimeHost)SwitchOutLogicalThreadState(){}
func (obj *ICORRuntimeHost)LocksHeldByLogicalThreadState(){}
func (obj *ICORRuntimeHost)MapFile(){}
func (obj *ICORRuntimeHost)GetConfiguration(){}
func (obj *ICORRuntimeHost)Stop(){}
func (obj *ICORRuntimeHost)CreateDomain(){}
func (obj *ICORRuntimeHost)EnumDomains(){}
func (obj *ICORRuntimeHost)NextDomain(){}
func (obj *ICORRuntimeHost)CloseEnum(){}
func (obj *ICORRuntimeHost)CreateDomainEx(){}
func (obj *ICORRuntimeHost)CreateDomainSetup(){}
func (obj *ICORRuntimeHost)CreateEvidence(){}
func (obj *ICORRuntimeHost)UnloadDomain(){}
func (obj *ICORRuntimeHost)CurrentDomain(){}


func NewICORRuntimeHost(ppv uintptr) *ICORRuntimeHost {
	return (*ICORRuntimeHost)(unsafe.Pointer(ppv))
}

func (obj *ICORRuntimeHost) AddRef() uintptr {
	ret, _, _ := syscall.Syscall(
		obj.lpVtbl.AddRef,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	return ret
}

func (obj *ICORRuntimeHost) Release() uintptr {
	ret, _, _ := syscall.Syscall(
		obj.lpVtbl.Release,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	return ret
}

func (obj *ICORRuntimeHost) Start() uintptr {
	ret, _, _ := syscall.Syscall(
		obj.lpVtbl.Start,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	return ret
}

func (obj *ICORRuntimeHost) GetDefaultDomain(pAppDomain *uintptr) uintptr {
	ret, _, _ := syscall.Syscall(
		obj.lpVtbl.GetDefaultDomain,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(pAppDomain)),
		0)
	return ret
}