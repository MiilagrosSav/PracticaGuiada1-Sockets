package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"sd-broadcast/pkg/protocolo"
)

func main() {
	direccionServidor := os.Getenv("SERVIDOR")
	if direccionServidor == "" {
		direccionDescubierta, ok := descubrirServidorUDP()
		if ok {
			log.Printf("Servidor descubierto en %s", direccionDescubierta)
			direccionServidor = direccionDescubierta
		} else {
			direccionServidor = "localhost:4000"
		}
	}

	nombre := os.Getenv("NOMBRE")
	if nombre == "" {
		fmt.Print("Ingrese su nombre: ")
		lector := bufio.NewReader(os.Stdin)
		nombreBytes, _, _ := lector.ReadLine()
		nombre = string(nombreBytes)
	}

	// TODO 20: conectar al servidor usando net.Dial("tcp", direccionServidor)
	// Manejar errores y usar defer conexion.Close()
	conexion, err := net.Dial("tcp", direccionServidor)
	if err != nil {
		log.Fatalf("No se pudo conectar al servidor: %v", err)
	}
	defer conexion.Close()

	// TODO 21: enviar mensaje de identificación con protocolo.Codificar
	// mensaje de tipo "identificacion" con Emisor = nombre
	identificacion := protocolo.NuevoMensaje(nombre, "", "identificacion")
	if err := protocolo.Codificar(conexion, identificacion); err != nil {
		log.Fatalf("No se pudo enviar identificación: %v", err)
	}

	// TODO 22: iniciar una goroutine que escuche mensajes del servidor en paralelo
	// La goroutine debe usar protocolo.Decodificar en un bucle e imprimir los mensajes recibidos
	// Si hay error, imprimir y retornar (el servidor cerró la conexión)
	go recibirMensajes(conexion)

	// TODO 23: en el hilo principal, leer líneas de stdin y enviar mensajes de tipo "broadcast"
	// Usar bufio.NewReader(os.Stdin) y ReadString('\n')
	// Para cada línea, crear un Mensaje y enviarlo con protocolo.Codificar
	lector := bufio.NewReader(os.Stdin)
	for {
		linea, err := lector.ReadString('\n')
		if err != nil {
			log.Printf("Error leyendo stdin: %v", err)
			break
		}
		contenido := strings.TrimSpace(linea)
		if contenido == "" {
			continue
		}
		mensaje := protocolo.NuevoMensaje(nombre, contenido, "broadcast")
		if err := protocolo.Codificar(conexion, mensaje); err != nil {
			log.Printf("Error enviando mensaje: %v", err)
			break
		}
	}

	log.Println("Cliente finalizado")
}

// recibirMensajes lee continuamente desde la conexión e imprime en consola
func recibirMensajes(conexion net.Conn) {
	// TODO 24: implementar bucle infinito de protocolo.Decodificar
	// Imprimir Emisor, Contenido y Timestamp de cada mensaje recibido
	// Si Decode retorna error, imprimir "Desconectado del servidor" y retornar
	for {
		mensaje, err := protocolo.Decodificar(conexion)
		if err != nil {
			log.Println("Desconectado del servidor")
			return
		}
		fmt.Printf("[%s] %s: %s\n", mensaje.Timestamp.Format("2006-01-02 15:04:05"), mensaje.Emisor, mensaje.Contenido)
	}
}
