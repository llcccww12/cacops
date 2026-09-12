package util

import (
	"fmt"
)

// ByteConverter 字节转换器结构体
type ByteConverter struct{}

// ConvertBytes 将字节转换为MB或GB，自动选择合适的单位
func (bc *ByteConverter) ConvertBytes(bytes uint64) (float64, string) {
	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
		TB = GB * 1024
	)

	switch {
	case bytes >= TB:
		return float64(bytes) / float64(TB), "TB"
	case bytes >= GB:
		return float64(bytes) / float64(GB), "GB"
	case bytes >= MB:
		return float64(bytes) / float64(MB), "MB"
	case bytes >= KB:
		return float64(bytes) / float64(KB), "KB"
	default:
		return float64(bytes), "B"
	}
}

// ConvertToMB 将字节转换为MB
func (bc *ByteConverter) ConvertToMB(bytes uint64) float64 {
	const MB = 1024 * 1024
	return float64(bytes) / float64(MB)
}

// ConvertToGB 将字节转换为GB
func (bc *ByteConverter) ConvertToGB(bytes uint64) float64 {
	const GB = 1024 * 1024 * 1024
	return float64(bytes) / float64(GB)
}

// FormatBytes 格式化字节显示，保留指定位数小数
func (bc *ByteConverter) FormatBytes(bytes uint64, precision int) string {
	value, unit := bc.ConvertBytes(bytes)
	return fmt.Sprintf("%.*f %s", precision, value, unit)
}

// 简单的函数版本（不需要结构体）

// BytesToMB 将字节转换为MB
func BytesToMB(bytes uint64) float64 {
	return float64(bytes) / (1024 * 1024)
}

// BytesToGB 将字节转换为GB
func BytesToGB(bytes uint64) float64 {
	return float64(bytes) / (1024 * 1024 * 1024)
}

// FormatBytesAuto 自动选择合适的单位格式化字节
func FormatBytesAuto(bytes uint64) string {
	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
		TB = GB * 1024
	)

	var value float64
	var unit string

	switch {
	case bytes >= TB:
		value = float64(bytes) / float64(TB)
		unit = "TB"
	case bytes >= GB:
		value = float64(bytes) / float64(GB)
		unit = "GB"
	case bytes >= MB:
		value = float64(bytes) / float64(MB)
		unit = "MB"
	case bytes >= KB:
		value = float64(bytes) / float64(KB)
		unit = "KB"
	default:
		value = float64(bytes)
		unit = "B"
	}

	// 根据数值大小决定小数位数
	precision := 2
	if value < 10 {
		precision = 3
	} else if value >= 100 {
		precision = 1
	}

	return fmt.Sprintf("%.*f %s", precision, value, unit)
}
