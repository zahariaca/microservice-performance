package models

import (
	"fmt"
	"log"
	"time"
)

type Units struct {
	scale uint64
	base  string
	units []string
}

var (
	BinaryUnits = &Units{
		scale: 1024,
		base:  "",
		units: []string{"KB", "MB", "GB", "TB", "PB"},
	}
	TimeUnitsUs = &Units{
		scale: 1000,
		base:  "us",
		units: []string{"ms", "s"},
	}
	TimeUnitsS = &Units{
		scale: 60,
		base:  "s",
		units: []string{"m", "h"},
	}
)

func formatUnits(n float64, m *Units, prec int) string {
	amt := n
	unit := m.base

	scale := float64(m.scale) * 0.85

	for i := 0; i < len(m.units) && amt >= scale; i++ {
		amt /= float64(m.scale)
		unit = m.units[i]
	}
	return fmt.Sprintf("%.*f%s", prec, amt, unit)
}

func UnitToFloat(n float64) float64 {
	m := TimeUnitsUs
	if n >= 1000000.0 {
		n /= 1000000.0
		m = TimeUnitsS
		log.Println("INSIDE IF")
	}

	amt := n

	scale := float64(m.scale) * 0.85

	for i := 0; i < len(m.units) && amt >= scale; i++ {
		amt /= float64(m.scale)
	}
	return amt
}

func formatBinary(n float64) string {
	return formatUnits(n, BinaryUnits, 2)
}

func formatTimeUs(n float64) string {
	units := TimeUnitsUs
	if n >= 1000000.0 {
		n /= 1000000.0
		units = TimeUnitsS
	}
	return formatUnits(n, units, 2)
}

func ConvertToMs(n float64) float64 {
	duration := time.Duration(n) * time.Millisecond

	return float64(duration.Milliseconds()) / float64(time.Microsecond)
}
