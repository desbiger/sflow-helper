package sflow_helper

import (
	"ClickHouse"
	"net"
	"unsafe"
)

func BytesToDstIp(record ClickHouse.FinalResult) net.IP {
	dstBytes := (*[4]byte)(unsafe.Pointer(&record.IpV4DSTIP))
	dstIP := net.IP{dstBytes[3], dstBytes[2], dstBytes[1], dstBytes[0]}

	return dstIP
}
func BytesToSrcIp(record ClickHouse.FinalResult) net.IP {
	srcBytes := (*[4]byte)(unsafe.Pointer(&record.IpV4SRCIP))
	srcIP := net.IP{srcBytes[3], srcBytes[2], srcBytes[1], srcBytes[0]}

	return srcIP
}
