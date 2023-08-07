How to use?
- Run the server `go run server/server.go`

Choose a client to call it with?
- Natively with telnet
    - `telnet localhost 8888`
    - Type commands in resp `+ping\r\n`
- With help of redis-cli
    - `redis-cli 8888`
    - Type commands in normal string `ping`
- The client implementation
    - `go run client/client.go`
    - Type commands in normal string `ping`