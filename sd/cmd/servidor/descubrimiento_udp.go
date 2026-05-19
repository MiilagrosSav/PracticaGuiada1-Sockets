package main

import (
	"encoding/json"
	"log"
	"net"
	"time"
)

const (
	udpMulticastAddr     = "239.0.0.1:9999"
	udpDiscoveryInterval = 2 * time.Second
)

type anuncioServidor struct {
	Tipo      string `json:"tipo"`
	Direccion string `json:"direccion"`
}

func iniciarDescubrimientoUDP(puerto string) {
	destino, err := net.ResolveUDPAddr("udp4", udpMulticastAddr)
	if err != nil {
		log.Printf("No se pudo resolver UDP multicast: %v", err)
		return
	}

	conexion, err := net.DialUDP("udp4", nil, destino)
	if err != nil {
		log.Printf("No se pudo abrir UDP multicast: %v", err)
		return
	}
	defer conexion.Close()

	host := obtenerIPLocal()
	anuncio := anuncioServidor{
		Tipo:      "descubrimiento",
		Direccion: host + ":" + puerto,
	}

	ticker := time.NewTicker(udpDiscoveryInterval)
	defer ticker.Stop()

	for range ticker.C {
		payload, err := json.Marshal(anuncio)
		if err != nil {
			log.Printf("No se pudo codificar anuncio UDP: %v", err)
			continue
		}
		if _, err := conexion.Write(payload); err != nil {
			log.Printf("No se pudo enviar anuncio UDP: %v", err)
		}
	}
}

func obtenerIPLocal() string {
	conexion, err := net.Dial("udp4", "8.8.8.8:80")
	if err != nil {
		return "127.0.0.1"
	}
	defer conexion.Close()

	host, _, err := net.SplitHostPort(conexion.LocalAddr().String())
	if err != nil {
		return "127.0.0.1"
	}
	return host
}
