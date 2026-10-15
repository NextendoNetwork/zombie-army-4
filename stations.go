package main

import (
	nex "github.com/NextendoNetwork/nextendo-nex"
	"net"
)

// The IP-wide NNCS cache cannot distinguish two clients behind one router.
// Build both candidates from this player's own public ReplaceURL, preserving
// its UDP port and CID. Never substitute another connection's observed port.
func preserveUDPStations(c *nex.Connection) {
	var local, public *nex.StationURL
	for _, station := range c.Stations() {
		ip := net.ParseIP(station.Get("address"))
		if ip == nil {
			continue
		}
		if ip.IsPrivate() || ip.IsLoopback() {
			local = station
		} else {
			public = station
		}
	}
	if local == nil || public == nil || public.GetInt("CID") == 0 ||
		public.GetInt("port") <= 0 || public.GetInt("port") > 65535 {
		return
	}
	lan, pub := public.Copy(), public.Copy()
	lan.Set("address", local.Get("address"))
	lan.Remove("type")
	lan.Remove("Pa")
	pub.SetInt("type", 11)
	pub.Set("Pa", pub.Get("address"))
	c.SetStations([]*nex.StationURL{lan, pub})
}
