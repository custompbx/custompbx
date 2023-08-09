# CustomPBX

**CustomPBX** (under development) is API server and Web GUI for [FreeSwitch](https://github.com/signalwire/freeswitch).
Providing pure FreeSWITCH experience and can be installed on existing systems with import existing configuration.
Built in the single [binary file](https://github.com/custompbx/custompbx/releases).

The project is under development, not tested well and can have security issues.  

**Backend** created with Golang v.1.19 located in ``src/custompbx``.

**Frontend** with Angular v.15 located in ``src/cweb-app``.

---
Service requirements:  
* OS Linux (amd64)  
* FreeSWITCH  
* Postgres Database
---
### Build
To build run:  
```
make install
```
To rebuild after resolved dependencies:  
```
make build
``` 
To run frontend locally build+test set backend websocket url ``export WS_BACKGROUND_OVERRIDE=wss://HOST:PORT/ws`` and run:  
```
make front-serve
```  
Check Makefile for more options.

Also Docker version available currently for test only. You can start DB + Freeswitch + Custompbx by replacing token ``- SIGNALWIRE_TOKEN=<YOUR_TOKEN_HERE>``  in ``docker-compose.yml`` and command:
```
docker compose up -d
```

Built binary file can be found in ``bin/`` directory and can be used according [Documentation](https://github.com/custompbx/custompbx/wiki).

Or just use precompiled binary from **[Releases Page](https://github.com/custompbx/custompbx/releases)**

Feel free to contact via **[discussions](https://github.com/custompbx/custompbx/discussions)** or open an **[issue](https://github.com/custompbx/custompbx/issues)**!

---
### Documentation
For **Installation** and **Configuration** manuals follow project's **[Wiki Page](https://github.com/custompbx/custompbx/wiki)**

![scheme](https://github.com/custompbx/doc/raw/master/img/Diagram1.png)

---
### GUI Demo
![demo](https://github.com/custompbx/doc/blob/master/img/demo_anim.gif?raw=true)
