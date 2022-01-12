package main

import (
	"Assembly/core"
	"fmt"
	"io/ioutil"
	"log"
	"runtime"
	"strings"
	"syscall"
)

func CheckH(hr uintptr,caller string) {
	if hr != 0x0 {
		fmt.Println(fmt.Sprintf("%s return 0x%08x",caller,hr))
		return
	}
}


func main(){
	var params = []string{"123"}
	exebyte ,_ := ioutil.ReadFile("SharpSQLTools.exe")

	var MetaHost uintptr
	hr := core.CLRCreateInstance(&core.CLSID_CLRMetaHost,&core.IID_ICLRMetaHost,&MetaHost)
	CheckH(hr,"CLRCreateInstance")
	metaHost := core.NewICLRMetaHost(MetaHost)
	pwzVersion, _ := syscall.UTF16PtrFromString("v4.0.30319")
	var RuntimeInfo uintptr
	hr = metaHost.GetRuntime(pwzVersion,&core.IID_ICLRRuntimeInfo,&RuntimeInfo)
	CheckH(hr,"metaHost.Runtime")
	runtimeinfo := core.NewICLRRuntimeInfo(RuntimeInfo)
	var isLoadable bool
	hr = runtimeinfo.IsLoadable(&isLoadable)
	CheckH(hr,"runtimeinfo.IsLoadable")
	if !isLoadable {
		fmt.Println("[!] IsLoadable return false.")
		return
	}
	hr = runtimeinfo.BindAsLegacyV2Runtime()
	CheckH(hr,"BindAsLegacyV2Runtime")
	var RuntimeHost uintptr
	hr = runtimeinfo.GetInterface(&core.CLSID_CorRuntimeHost,&core.IID_ICorRuntimeHost,&RuntimeHost)
	runtimeHost := core.NewICORRuntimeHost(RuntimeHost)
	hr = runtimeHost.Start()
	CheckH(hr,"runtimehost.Start")
	fmt.Println("[+] Loaded CLR into this process")
	var AppDomain uintptr
	var IUnknown uintptr
	hr = runtimeHost.GetDefaultDomain(&IUnknown)
	CheckH(hr,"runtimehost.GetDefaultDomain")
	iu := core.NewIUnknown(IUnknown)
	hr = iu.QueryInterface(&core.IID_AppDomain,&AppDomain)
	CheckH(hr,"iu.QueryInterface")
	appDomain := core.NewAppDomain(AppDomain)
	fmt.Println("[+] Get Default AppDomain")

	safeArray, err := core.CreateSafeArray(exebyte)
	if err != nil {
		fmt.Println("[!]",err)
		return
	}
	runtime.KeepAlive(safeArray)
	fmt.Println("[+] Crated SafeArray from byte array")
	var pAssembly uintptr
	assembly,_ := appDomain.Load_3((*core.SafeArray)(safeArray))
	CheckH(hr, "appDomain.Load_3")
	//assembly := core.NewAssembly(pAssembly)
	fmt.Printf("[+] Executable loaded into memory at 0x%08x\n", pAssembly)

	//var pEntryPointInfo uintptr
	methodInfo, _ := assembly.GetEntryPoint()
	CheckH(hr, "assembly.GetEntryPoint")
	//fmt.Printf("[+] Executable entrypoint found at 0x%08x. Calling...\n", pEntryPointInfo)
	fmt.Println("-------")
	//methodInfo := core.NewMethodInfo(pEntryPointInfo)

	var paramSafeArray *core.SafeArray
	methodSignature, err := methodInfo.GetString()
	if err != nil {
		return
	}

	fmt.Println("[+] Checking if the assembly requires arguments...")
	if !strings.Contains(methodSignature, "Void Main()") {
		if len(params) < 1 {
			log.Fatal("the assembly requires arguments but none were provided\nUsage: EXEfromMemory.exe <exe_file> <exe_args>")
		}
		if paramSafeArray, err = core.PrepareParameters(params); err != nil {
			log.Fatal(fmt.Sprintf("there was an error preparing the assembly arguments:\r\n%s", err))
		}
	}

	nullVariant := core.Variant{
		VT:  1,
		Val: uintptr(0),
	}
	fmt.Println("[+] Invoking...")
	err = methodInfo.Invoke_3(nullVariant, paramSafeArray)
	if err != nil {
		fmt.Println("[!]",err)
	}

	fmt.Println("-------")

	CheckH(hr, "methodInfo.Invoke_3")
	fmt.Printf("[+] Executable returned code")

	appDomain.Release()
	runtimeHost.Release()
	runtimeinfo.Release()
	metaHost.Release()
}
