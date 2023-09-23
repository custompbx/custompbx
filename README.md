# CustomPBX

**CustomPBX** (currently in development) is an API server and Web GUI for [FreeSwitch](https://github.com/signalwire/freeswitch), offering a pure FreeSWITCH experience. It can be installed on existing systems, allowing for the import of existing configurations. The system is encapsulated within a single [binary file](https://github.com/custompbx/custompbx/releases).

Please note that this project is still in development, has not undergone extensive testing, and may potentially have security vulnerabilities.

The **Backend** is developed using Golang v.1.19 and is located in the ``src/custompbx`` directory.

The **Frontend**, built with Angular v.16, can be found in the ``src/cweb-app`` directory.

---
System Requirements:
* Linux OS (amd64)
* FreeSWITCH
* Postgres Database
---
### Build Process
To initiate the build process, execute the following command:
```
make install
```
For rebuilding after resolving dependencies, use:
```
make build
``` 
To locally run the frontend, perform the following steps: build and test the project, set the backend websocket URL using the command ``export WS_BACKGROUND_OVERRIDE=wss://HOST:PORT/ws``, and finally execute:
```
make front-serve
```  
Additional options can be found in the Makefile.

A Docker version of the project is also available, currently intended for testing purposes only.
In the ``docker-compose.yml`` file:
- Replace the token with ``- SIGNALWIRE_TOKEN=<YOUR_TOKEN_HERE>``
- Modify the host if necessary: ``WS_BACKGROUND_OVERRIDE=wss://127.0.0.1:8080/ws``
- Start the containers and open ``https://127.0.0.1:8080/cweb`` (or your Docker host), making sure to allow self-signed certificates.

You can start Docker with PostgresDB + Freeswitch + Custompbx by using the command:
```
docker compose up -d
```

The compiled binary file is located in the ``bin/`` directory and can be used as outlined in the [Documentation](https://github.com/custompbx/custompbx/wiki).

Alternatively, you can utilize the precompiled binary available on the **[Releases Page](https://github.com/custompbx/custompbx/releases)**.

If you have any questions or feedback, don't hesitate to get in touch through the **[discussions](https://github.com/custompbx/custompbx/discussions)** or by opening an **[issue](https://github.com/custompbx/custompbx/issues)**!

---
### Documentation
For detailed instructions on **Installation** and **Configuration**, please refer to the project's **[Wiki Page](https://github.com/custompbx/custompbx/wiki)**.

![system diagram](https://github.com/custompbx/doc/raw/master/img/Diagram1.png)

---
### GUI Demo
![demo](https://github.com/custompbx/doc/blob/master/img/demo_anim.gif?raw=true)