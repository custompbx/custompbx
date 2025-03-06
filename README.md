![logo2](https://github.com/user-attachments/assets/f45ea5b4-e269-4f5b-a787-d4a3870f17bf)
<p align="center">
    <img src="https://badgen.net//github/stars/custompbx/custompbx?color=ffd700" alt="">
    <img src="https://badgen.net/badge/license/MIT/cyan" alt="">
    <img src="https://badgen.net/github/last-commit/custompbx/custompbx?icon=github" alt="">
    <img alt="GitHub Created At" src="https://img.shields.io/github/created-at/custompbx/custompbx">
</p>

<h1 align="center">CustomPBX</h1>

**CustomPBX** is an API server and Web GUI for [FreeSwitch](https://github.com/signalwire/freeswitch), offering a pure FreeSWITCH experience. It can be installed on existing systems, allowing for the import of existing configurations. The system is encapsulated within a single [binary file](https://github.com/custompbx/custompbx/releases).

Please note that this project is still in development, has not undergone extensive testing.

The **Backend** is developed using Golang v.1.24 and is located in the ``src/custompbx`` directory.

The **Frontend**, built with Angular v.17, can be found in the ``src/cweb-app`` directory.

---
System Requirements:
* Linux OS (amd64)
* FreeSWITCH
* Postgres Database
---
### Build Process
Install **make** first (apt example):
```
sodo apt install -y make
```

To install golang, node, all dependencies and initiate the build process, execute the following command:
```
make install
```
For rebuilding after resolving dependencies, use:
```
make build
``` 
To locally run the frontend, perform the following steps: build and test the project, set the backend websocket URL using the command ``export WS_BACKEND_OVERRIDE=wss://HOST:PORT/ws``, and finally execute:
```
make front-serve
```  
Additional options can be found in the Makefile.
- install-golang
- install-node
- install-dep (install dependencies for back and front)
- dep-front
- dep-back

The compiled binary file is located in the ``bin/`` directory and can be used as outlined in the [Documentation](https://github.com/custompbx/custompbx/wiki).

Alternatively, you can utilize the precompiled binary available on the **[Releases Page](https://github.com/custompbx/custompbx/releases)**.

---
#### Build with Docker (TEST ONLY)
A Docker version of the project is also available, currently intended for testing purposes only.
In the ``.env`` file:
- Replace the token with yours ``SIGNALWIRE_TOKEN=<YOUR_TOKEN_HERE>``
- Start the containers and open ``https://127.0.0.1:8080/cweb`` (or your Docker host), making sure to allow self-signed certificates.

You can start Docker with PostgresDB + Freeswitch + Custompbx by using the command:
```
docker compose up -d
```

---
### Documentation
For detailed instructions on **Installation** and **Configuration**, please refer to the project's **[Wiki Page](https://github.com/custompbx/custompbx/wiki)**.

If you have any questions or feedback, don't hesitate to get in touch through the **[discussions](https://github.com/custompbx/custompbx/discussions)** or by opening an **[issue](https://github.com/custompbx/custompbx/issues)**!

![system diagram](https://github.com/custompbx/doc/raw/master/img/Diagram1.png)

---
### GUI Demo
GIF
![demo](https://github.com/custompbx/doc/blob/master/img/demo_anim.gif?raw=true)

---

![1](https://github.com/user-attachments/assets/3a6c238b-015b-4abf-86f7-cd6c74b94608)

---

![2](https://github.com/user-attachments/assets/eb692658-838f-4bfc-957a-38b21b81e6ff)

---

![3](https://github.com/user-attachments/assets/ecb77fc3-f2ae-4377-a880-2ea60b0f5e9b)

---

![4](https://github.com/user-attachments/assets/4d3d1621-f2f8-44b6-a5cf-d5de9e956f9a)

---

![5](https://github.com/user-attachments/assets/5b9e7d32-efb0-437d-a613-f6c5c44b4e0c)

---

![6](https://github.com/user-attachments/assets/d3188dce-9237-4085-83f8-db62ffdc5164)

---

![7](https://github.com/user-attachments/assets/bd19000e-e661-4370-b490-a8bc28b03b71)

---

![8](https://github.com/user-attachments/assets/8101e3ac-85f7-428c-8bb2-2dcc00863454)

---

![9](https://github.com/user-attachments/assets/0bf2b11e-8ca9-4fcb-9b2c-99b37d83bae3)

---

![10](https://github.com/user-attachments/assets/4be13bef-df08-4b0c-9ead-d8ca35c7117c)

---

![11](https://github.com/user-attachments/assets/6591410f-c699-470e-b71e-3c45de33ba8a)

---

![12](https://github.com/user-attachments/assets/8b4f7117-0487-467d-a5e7-c08bb5ceee31)

---

![13](https://github.com/user-attachments/assets/1f8cd30d-1bad-482e-8844-7d894c31ae1a)

---

![14](https://github.com/user-attachments/assets/45847e99-8f7f-45e3-a721-5df134c2cbfa)
