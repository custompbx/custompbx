package metrics

import (
	"custompbx/altData"
	"custompbx/cache"
	"custompbx/dataSourceAdapter"
	"custompbx/mainStruct"
	"custompbx/pbxcache"
	"custompbx/webcache"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"time"
)

func dealwithErr(err error) {
	if err != nil {
		fmt.Println(err)
		//os.Exit(-1)
	}
}

func GetHardwareData() {
	//runtimeOS := runtime.GOOS
	vmStat, err := mem.VirtualMemory()
	dealwithErr(err)
	diskStat, err := disk.Usage("/")
	dealwithErr(err)
	cpuStat, err := cpu.Info()
	dealwithErr(err)
	percentage, err := cpu.Percent(0, true)
	dealwithErr(err)
	hostStat, err := host.Info()
	dealwithErr(err)
	//interfStat, err := net.Interfaces()
	//dealwithErr(err)
	/*	log.Println("OS : " + runtimeOS)
		log.Println("Total memory: " + strconv.FormatUint(vmStat.Total, 10) + " bytes")
		log.Println("Free memory: " + strconv.FormatUint(vmStat.Free, 10) + " bytes")
		log.Println("Percentage used memory: " + strconv.FormatFloat(vmStat.UsedPercent, 'f', 2, 64) + "%")
		log.Println("Total disk space: " + strconv.FormatUint(diskStat.Total, 10) + " bytes")
		log.Println("Used disk space: " + strconv.FormatUint(diskStat.Used, 10) + " bytes")
		log.Println("Free disk space: " + strconv.FormatUint(diskStat.Free, 10) + " bytes")
		log.Println("Percentage disk space usage: " + strconv.FormatFloat(diskStat.UsedPercent, 'f', 2, 64) + "%")
		log.Println("CPU index number: " + strconv.FormatInt(int64(cpuStat[0].CPU), 10))
		log.Println("VendorID: " + cpuStat[0].VendorID)
		log.Println("Family: " + cpuStat[0].Family)
		log.Println("Number of cores: " + strconv.FormatInt(int64(cpuStat[0].Cores), 10))
		log.Println("Model Name: " + cpuStat[0].ModelName)
		log.Println("Speed: " + strconv.FormatFloat(cpuStat[0].Mhz, 'f', 2, 64) + " MHz")

		for idx, cpupercent := range percentage {
			log.Println("Current CPU utilization: [" + strconv.Itoa(idx) + "] " + strconv.FormatFloat(cpupercent, 'f', 2, 64) + "%")
		}

		log.Println("Hostname: " + hostStat.Hostname)
		log.Println("Uptime: " + strconv.FormatUint(hostStat.Uptime, 10))
		log.Println("Number of processes running: " + strconv.FormatUint(hostStat.Procs, 10))
		log.Println("OS: " + hostStat.OS)
		log.Println("Platform: " + hostStat.Platform)
		log.Println("Host ID(uuid): " + hostStat.HostID)
		for _, interf := range interfStat {
			log.Println("------------------------------------------------------")
			log.Println("Interface Name: " + interf.Name)

			if interf.HardwareAddr != "" {
				log.Println("Hardware(MAC) Address: " + interf.HardwareAddr)
			}

			for _, flag := range interf.Flags {
				log.Println("Interface behavior or flags: " + flag)
			}
			log.Println(interf.String())
			for _, addr := range interf.Addrs {
				log.Println("IPv6 or IPv4 addresses: " + addr.String())

			}


		}

	*/
	data := mainStruct.DashboardData{
		Timestamp:    time.Now(),
		Hostname:     hostStat.Hostname,
		OS:           hostStat.OS,
		Platform:     hostStat.Platform,
		CPUModel:     cpuStat[0].ModelName,
		CPUFrequency: cpuStat[0].Mhz,

		DynamicMetrics: mainStruct.DynamicMetrics{
			PercentageDiskUsage:  diskStat.UsedPercent,
			TotalDiscSpace:       diskStat.Total,
			FreeDiskSpace:        diskStat.Free,
			TotalMemory:          vmStat.Total,
			FreeMemory:           vmStat.Free,
			PercentageUsedMemory: vmStat.UsedPercent,
			CoreUtilization:      percentage,
		},
	}
	webcache.DashBoardSetStaticData(data)

}

func UpdateMetrics() {
	tick := time.Tick(1 * time.Second)
	webcache.DashBoardSetSipRegs(cache.GetDomainSipRegsCounter())
	profiles, gateways := altData.GetSofiaProfilesAndGateways()
	profiles = dataSourceAdapter.UpdateSofiaProfileStatuses(profiles)
	gateways = dataSourceAdapter.UpdateSofiaGatewayStatuses(gateways)
	webcache.DashBoardSetSofiaData(profiles, gateways)
	webcache.DashBoardSetCallsCounter(pbxcache.GetChannelsCounter())
	GetHardwareData()

	for {
		select {
		case <-tick:
			GetHardwareData()
			//save to db then
		}
	}
}

/*
2019/08/30 15:00:49 Hostname: debian-01
2019/08/30 15:00:49 OS: linux
2019/08/30 15:00:49 Platform: debian

2019/08/30 15:00:49 Total memory: 1032429568 bytes
2019/08/30 15:00:49 Free memory: 87228416 bytes
2019/08/30 15:00:49 Percentage used memory: 37.80%

2019/08/30 15:00:49 Total disk space: 26288107520 bytes
2019/08/30 15:00:49 Used disk space:   5014605824 bytes
2019/08/30 15:00:49 Free disk space:  19914551296 bytes
2019/08/30 15:00:49 Percentage disk space usage: 20.12%

2019/08/30 15:00:49 CPU index number: 0
2019/08/30 15:00:49 Number of cores: 1
2019/08/30 15:00:49 Model Name: Intel Core Processor (Broadwell, IBRS)
2019/08/30 15:00:49 Speed: 2199.99 MHz
2019/08/30 15:00:49 Current CPU utilization: [0] 100.00%
2019/08/30 15:00:49 Uptime: 2919809
2019/08/30 15:00:49 Number of processes running: 91
2019/08/30 15:00:49 ------------------------------------------------------
2019/08/30 15:00:49 Interface Name: lo
2019/08/30 15:00:49 Interface behavior or flags: up
2019/08/30 15:00:49 Interface behavior or flags: loopback
2019/08/30 15:00:49 IPv6 or IPv4 addresses: {"addr":"127.0.0.1/8"}
2019/08/30 15:00:49 IPv6 or IPv4 addresses: {"addr":"::1/128"}
2019/08/30 15:00:49 ------------------------------------------------------
2019/08/30 15:00:49 Interface Name: eth0
2019/08/30 15:00:49 Hardware(MAC) Address: 00:16:3e:3e:25:e9
2019/08/30 15:00:49 Interface behavior or flags: up
2019/08/30 15:00:49 Interface behavior or flags: broadcast
2019/08/30 15:00:49 Interface behavior or flags: multicast
2019/08/30 15:00:49 IPv6 or IPv4 addresses: {"addr":"185.247.118.201/24"}
2019/08/30 15:00:49 IPv6 or IPv4 addresses: {"addr":"fe80::216:3eff:fe3e:25e9/64"}
*/
