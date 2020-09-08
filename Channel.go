package main

import "github.com/Traliaa/go-zigbee/nwkmgr"

type Channel struct {
	ID       string
	device   *Device
	endpoint *nwkmgr.NwkSimpleDescriptorT
}
