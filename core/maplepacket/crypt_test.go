package maplepacket

import "testing"

func TestGetPacketLength(t *testing.T) {
	packetHeader := [4]byte{5, 6, 12, 56}
	v := (int(packetHeader[0]) + (int(packetHeader[1]) << 8)) ^ (int(packetHeader[2]) + (int(packetHeader[3]) << 8))
	if v != GetPacketLength(packetHeader[:]) {
		t.Errorf("Not equal")
	}
}

func testGetpacketHeader() {

}
