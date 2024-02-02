package models

import "fmt"

type RawUsageStats struct {
	CpuUsagePercentage    string
	MemoryUsagePercentage string
	MemoryUsage           string
	MaxAllocatedMemory    string
	ProjectName           string
}

func (rus *RawUsageStats) String() string {
	return fmt.Sprintf("%s %s", rus.CpuUsagePercentage, rus.MemoryUsagePercentage)
}

type UsageStats struct {
	ServerName        string
	MinCpuUsage       float64
	MaxCpuUsage       float64
	MinMemoryUsage    float64
	MaxMemoryUsage    float64
	MinMemoryUsageMiB float64
	MaxMemoryUsageMiB float64
}

func (us *UsageStats) String() string {
	return fmt.Sprintf("ServerName: %s"+
		" MinCpuUsage: %.2f %% MaxCpuUsage: %.2f %%"+
		" MinMemoryUsage: %.2f %% MaxMemoryUsage: %.2f %%"+
		" MinMemoryUsage: %.2f MiB MaxMemoryUsage: %.2f MiB",
		us.ServerName,
		us.MinCpuUsage,
		us.MaxCpuUsage,
		us.MinMemoryUsage,
		us.MaxMemoryUsage,
		us.MinMemoryUsageMiB,
		us.MaxMemoryUsageMiB)
}
