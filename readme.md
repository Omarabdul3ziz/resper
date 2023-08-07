# RespEr
Is a server/client project that communicate over tcp socket with resp as the communication protocol. 


What expected to be learnt from this project?
- Natively dealing with sockets over tcp.
- Handling reading/writing over bytes stream.
- Handling strings buy Encoding/Decoding messages.


How this project is structured?
- `client/` is the client code.
- `server/` is the server code.
- `pkg/ioutil/` is the code for reader/writer.
- `pkg/codec/` is the code for the encoder/decoder.
- `pkg/handlers/` is the handlers for the requested operations.


The expected flow?
- running the server will open a tcp socket on a port and keep listening for any connection to this port.
- running the client will open a shell that will prompt you for any message to send for the server.
- any message is prompted in the client will be encoded first to be in a resp schema.
- the generated string will be written over the connection with ioutil pkg.
- the reversed process will be done on the server side, handled with the handlers.
- the same process will be repeated to send the response back to the client.


References and materials used?
- Build your own X from scratch: Redis. [here](https://build-your-own.org/redis/)
- Nim Days book. Days 2, 12 and 13. [here](https://xmonader.github.io/nimdays/day12_resp.html)


Development Notes:
- [Dev Notes](./docs/dev_notes.md)