# Servidor de Broadcast Concurrente

Proyecto base para la Clase sobre Sockets de Sistemas Distribuidos.

## Integrantes

- Antunez, Elias Emanuel
- Donda, Melisa Ileana
- Savallich, Milagros Antonella

## Ejecución

### Local

```bash
# Terminal 1: servidor
go run ./cmd/servidor

# Terminal 2: cliente
go run ./cmd/cliente
```

### Docker Compose

```bash
docker-compose up --build
```

## Requisitos completados

- [x] Servidor TCP concurrente
- [x] Protocolo JSON
- [x] Registro de clientes con sync.RWMutex
- [x] Broadcast a todos los clientes
- [x] Cliente interactivo (stdin + recepción paralela)
- [x] Docker + docker-compose
- [x] Bonus: descubrimiento UDP

## Captura de ejecución

(Adjuntar log o captura de pantalla con múltiples clientes conectados)

QUEDA HACER ESTO!!! NO OLVIDAR
para modificar los cambios estan en la rama desarrollo
