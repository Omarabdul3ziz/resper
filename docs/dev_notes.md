What to do next?

- is it okay for the reader/writer to deal with strings? or should it be a io_stream instead? 
    so instead of returning a string from the connection it should return `bufio.NewReader`
- better to start reading the resources now. and finish the encoders.
- draw a diagram to explain how each component work
- read the code carefully and enhance.
- after implementing the resp. you need to implement `prepareCommand` 
- and then you can register commands to your server to do things
- it is a good idea to build the client yourself but try to use the existing technologies first.
---
Commands
- Echo
- Set/Get
- Curl

-----
Plan the day?
- [x] 2 sessions on nimdays
- [ ] 2 sessions on howtobuild
- break
- [ ] 2 sessions implementing
- [ ] 2 sessions enhancing
- ask for review

---
What to expect?
- encoding/decoding 
    - decodeString("*3\r\n:1\r\n:2\r\n:3\r\n\r\n") > [1, 2, 3]
    - encodeString("str", "Hi") > "+Hello, World\r\n"
  
---
The terms used on the data conversion context.
- Serialize/Deserialize: the border term used generally to mean convert to text.
- Encode/Decode: converting data from one representation to the other.
- Marshaling/Unmarshaling: convert data structure to some format.


---
- [ ] first finish server. reader. decoder in appropriate way. & make sure communication with redis-cli is working.
- [ ] implement the writer, encoder to respond
- [ ] implement the client.