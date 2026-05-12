package main

import (
	"encoding/json"
	"net"
	"time"
)

const (
	udpMulticastAddr    = "239.0.0.1:9999"
	udpDiscoveryTimeout = 5 * time.Second
)

type anuncioServidor struct {
	Tipo      string `json:"tipo"`
	Direccion string `json:"direccion"`
}

func descubrirServidorUDP() (string, bool) {
	destino, err := net.ResolveUDPAddr("udp4", udpMulticastAddr)
	if err != nil {
		return "", false
	}

	conexion, err := net.ListenMulticastUDP("udp4", nil, destino)
	if err != nil {
		return "", false
	}
	defer conexion.Close()

	_ = conexion.SetReadBuffer(1024)
	_ = conexion.SetReadDeadline(time.Now().Add(udpDiscoveryTimeout))

	buffer := make([]byte, 1024)
	for {
		n, _, err := conexion.ReadFromUDP(buffer)
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				return "", false
			}
			continue
		}

		var anuncio anuncioServidor
		if err := json.Unmarshal(buffer[:n], &anuncio); err != nil {
			continue
		}
		if anuncio.Tipo != "descubrimiento" || anuncio.Direccion == "" {
			continue
		}
		return anuncio.Direccion, true
	}
}
