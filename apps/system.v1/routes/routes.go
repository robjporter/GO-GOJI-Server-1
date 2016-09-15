package routes

import (
	"net/http"

	"../controllers"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func Prefetch() {

}

func Routes(m *web.Mux) {
	goji.Handle("/api/v1/system/*", m)
	goji.Get("/api/v1/system", http.RedirectHandler("/api/v1/system/", 301))

	m.Get("/api/v1/system/tmp", controllers.TMP)
	m.Get("/api/v1/system/,tmp2", controllers.TMP2)

	// SYSTEM
	m.Get("/api/v1/system/", controllers.SysHome)
	m.Get("/api/v1/system/raw", controllers.SysCoreRaw)
	m.Get("/api/v1/system/cores", controllers.SysCoreCount)
	m.Get("/api/v1/system/ram", controllers.SysRamCount)
	m.Get("/api/v1/system/cpu/percent/sec", controllers.SysCPUPercent1Sec)
	m.Get("/api/v1/system/cpu/percent/sec/tot", controllers.SysCPUPercent1SecTot)
	m.Get("/api/v1/system/cpu/info/raw", controllers.SysCPUInfo)
	m.Get("/api/v1/system/cpu/info/vendor", controllers.SysCPUInfoVendor)
	m.Get("/api/v1/system/cpu/info/family", controllers.SysCPUInfoFamily)
	m.Get("/api/v1/system/cpu/info/model", controllers.SysCPUInfoModel)
	m.Get("/api/v1/system/cpu/info/cores", controllers.SysCPUInfoCores)
	m.Get("/api/v1/system/cpu/info/name", controllers.SysCPUInfoName)
	m.Get("/api/v1/system/cpu/info/speed", controllers.SysCPUInfoSpeed)
	m.Get("/api/v1/system/cpu/info/cache", controllers.SysCPUInfoCache)
	m.Get("/api/v1/system/cpu/info/flags", controllers.SysCPUInfoFlags)

	// HOST
	m.Get("/api/v1/system/host/boottime", controllers.HostBootTime)
	m.Get("/api/v1/system/host/os", controllers.HostOS)
	m.Get("/api/v1/system/host/version", controllers.HostVersion)
	m.Get("/api/v1/system/host/users", controllers.HostActiveUsers)
	m.Get("/api/v1/system/host/virtualisation/system", controllers.HostVirtualisationSystem)
	m.Get("/api/v1/system/host/virtualisation/role", controllers.HostVirtualisationRole)

	// NET
	m.Get("/api/v1/system/net/", controllers.NetInfoAdapterSummary)
	m.Get("/api/v1/system/net/bytes/sent", controllers.NetInfoAdapterSummaryBytesSent)
	m.Get("/api/v1/system/net/bytes/recv", controllers.NetInfoAdapterSummaryBytesRecv)
	m.Get("/api/v1/system/net/packets/sent", controllers.NetInfoAdapterSummaryPacketsSent)
	m.Get("/api/v1/system/net/packets/recv", controllers.NetInfoAdapterSummaryPacketsRecv)
	m.Get("/api/v1/system/net/errors/in", controllers.NetInfoAdapterSummaryErrorIn)
	m.Get("/api/v1/system/net/errors/out", controllers.NetInfoAdapterSummaryErrorOut)
	m.Get("/api/v1/system/net/adapters/summary", controllers.NetInfoAdaptersSummary)
	m.Get("/api/v1/system/net/adapters/:interface/summary", controllers.NetInfoAdaptersSummary)
	m.Get("/api/v1/system/net/adapters/summary/detail", controllers.NetInfoAdaptersSummaryDetail)
	m.Get("/api/v1/system/net/adapters/:interface/summary/detail", controllers.NetInfoAdaptersSummaryDetail)

	// MEM
	m.Get("/api/v1/system/mem", controllers.MemInfoSummary)
	m.Get("/api/v1/system/mem/swap/summary", controllers.MemInfoMemSwapSummary)
	m.Get("/api/v1/system/mem/swap/total", controllers.MemInfoMemSwapTotal)
	m.Get("/api/v1/system/mem/swap/used", controllers.MemInfoMemSwapUsed)
	m.Get("/api/v1/system/mem/swap/free", controllers.MemInfoMemSwapFree)
	m.Get("/api/v1/system/mem/swap/usedpercent", controllers.MemInfoMemSwapUsedPercent)
	m.Get("/api/v1/system/mem/swap/sin", controllers.MemInfoMemSwapSin)
	m.Get("/api/v1/system/mem/swap/sout", controllers.MemInfoMemSwapSout)
	m.Get("/api/v1/system/mem/virtual/summary", controllers.MemInfoMemVirtualSummary)
	m.Get("/api/v1/system/mem/virtual/total", controllers.MemInfoMemVirtualTotal)
	m.Get("/api/v1/system/mem/virtual/available", controllers.MemInfoMemVirtualAvailable)
	m.Get("/api/v1/system/mem/virtual/used", controllers.MemInfoMemVirtualUsed)
	m.Get("/api/v1/system/mem/virtual/usedpercent", controllers.MemInfoMemVirtualUsedPercent)
	m.Get("/api/v1/system/mem/virtual/free", controllers.MemInfoMemVirtualFree)
	m.Get("/api/v1/system/mem/virtual/active", controllers.MemInfoMemVirtualActive)
	m.Get("/api/v1/system/mem/virtual/inactive", controllers.MemInfoMemVirtualInactive)
	m.Get("/api/v1/system/mem/virtual/buffers", controllers.MemInfoMemVirtualBuffers)
	m.Get("/api/v1/system/mem/virtual/cached", controllers.MemInfoMemVirtualCached)
	m.Get("/api/v1/system/mem/virtual/wired", controllers.MemInfoMemVirtualWired)

	// LOAD
	m.Get("/api/v1/system/load", controllers.LoadInfoSummary)
}
