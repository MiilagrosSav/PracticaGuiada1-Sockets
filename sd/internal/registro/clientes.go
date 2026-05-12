package registro

import (
	"net"
	//importamos el paqueete sync para usar mutexes y proteger el acceso al mapa de clientes
	"sync"
)

// RegistroClientes mantiene el listado de conexiones activas de forma segura
type RegistroClientes struct {
	//TODO 1: AGREGAR UN CAMPO sync.RWMutex para proteger el mapa
	mu       sync.RWMutex
	clientes map[string]net.Conn
}

// NuevoRegistro crea un registro vacío
func NuevoRegistro() *RegistroClientes {
	//TODO 2: INICIALIZAR EL MAPA DE CLIENTES
	return &RegistroClientes{
		clientes: make(map[string]net.Conn),
	}
}

// Agregar añade un cliente al registro
func (r *RegistroClientes) Agregar(nombre string, conexion net.Conn) {
	//TODO 3: BLOQUEAR PARA ESCRITURA, AGREGAR AL MAPA, DESBLOQUEAR
	r.mu.Lock()
	defer r.mu.Unlock()
	r.clientes[nombre] = conexion
}

// Eliminar remueve un cliente del registro
func (r *RegistroClientes) Eliminar(nombre string) {
	//TODO 4: BLOQUEAR PARA ESCRITURA, ELIMINAR DEL MAPA, DESBLOQUEAR
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.clientes, nombre)
}

// ObtenerConexiones devuelve una copia de todas las conexiones activas
func (r *RegistroClientes) ObtenerConexiones() map[string]net.Conn {
	//TODO 5: BLOQUEAR PARA LECTURA, COPIAR LAS CONEXIONES A UN MAPA, DESBLOQUEAR
	r.mu.RLock()
	defer r.mu.RUnlock()

	conexiones := make(map[string]net.Conn, len(r.clientes))
	for nombre, conexion := range r.clientes {
		conexiones[nombre] = conexion
	}
	return conexiones
}

// Cantidad devuelve el número de clientes conectados
func (r *RegistroClientes) Cantidad() int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.clientes)
}

// Nombres devuelve un slice con los nombres de los clientes
func (r *RegistroClientes) Nombres() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	nombres := make([]string, 0, len(r.clientes))
	for nombre := range r.clientes {
		nombres = append(nombres, nombre)
	}
	return nombres
}
