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

### Docker Compose (interactivo)

**1. Levantar solo el servidor** (en background):
```bash
make docker-up 
```

**2. Conectar clientes interactivos** (en terminales separadas):
```bash
# Terminal 2: Cliente 1
make docker-cliente1

# Terminal 3: Cliente 2
make docker-cliente2

# Terminal 4: Cliente 3
make docker-cliente3
```

Cada cliente abre un prompt. Escribí un mensaje y presioná Enter para enviarlo a todos.

**3. Ver logs del servidor**:
```bash
make docker-logs
```

**4. Detener todo**:
```bash
make docker-down
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

servidor-1  | 2026/05/23 00:35:56 Servidor de broadcast escuchando en :4000
servidor-1  | 2026/05/23 00:36:26 Cliente conectado: Cliente1 desde 172.18.0.3:51014
servidor-1  | 2026/05/23 00:36:46 Cliente conectado: Cliente2 desde 172.18.0.4:56928
servidor-1  | 2026/05/23 00:36:56 Cliente conectado: Cliente3 desde 172.18.0.5:47576

