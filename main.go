package main

import (
	"tinygo.org/x/bluetooth"
	"log"
)

func main() {
	adapter := bluetooth.DefaultAdapter
	// Enable adapter
	err := adapter.Enable()
	if err != nil {
		panic("failed to enable BLE adapter")
	}

    // Start scanning and define callback for scan results
	err = adapter.Scan(onScan)
	if err != nil {
		panic("failed to register scan callback")
	}
}

func onScan(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
	log.Println("found device:", device.Address.String(), device.RSSI, device.LocalName(), device.AdvertisementPayload)
}
