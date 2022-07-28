## gRPC assignment

base url - `localhost:9091`

---
To update/add proto files :

- `make clean` to remove existing stubs from `/pb`
- `make gen` to regenerate stubs

---
To run server :
  
- `make server_start`
  
---
To run client and invoke methods :

- `make client_start arg=___`
Where arg is one of :
 - `sum`
 - `primes`
 - `max`
 - `avg`