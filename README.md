![logo2](https://github.com/user-attachments/assets/f45ea5b4-e269-4f5b-a787-d4a3870f17bf)
<p align="center">
    <img src="https://badgen.net//github/stars/custompbx/custompbx?color=ffd700" alt="">
    <img src="https://badgen.net/badge/license/MIT/cyan" alt="">
    <img src="https://badgen.net/github/last-commit/custompbx/custompbx?icon=github" alt="">
    <img alt="GitHub Created At" src="https://img.shields.io/github/created-at/custompbx/custompbx">
</p>

<h1 align="center">CustomPBX</h1>

**CustomPBX** is an API server and web interface for [FreeSWITCH](https://github.com/signalwire/freeswitch). It can manage a new installation or import an existing FreeSWITCH configuration. The Angular frontend is embedded in the Go application and distributed as a single [release binary](https://github.com/custompbx/custompbx/releases).

> [!IMPORTANT]
> CustomPBX is under active development. Validate configuration changes and keep backups before using it in production.

The **backend** uses Go 1.25 and is located in `src/custompbx`.

The **frontend** uses Angular 21 and is located in `src/cweb-app`.

---
## Requirements

For the recommended container deployment:

- Docker with Compose support

For a source build:

- Linux or WSL on amd64
- Go 1.25.12 or newer in the Go 1.25 release line
- Node.js 22
- GNU Make
- FreeSWITCH and PostgreSQL available at runtime

---
## Build from source

Run source builds from Linux or WSL. Frontend builds on native Windows and host-mounted `node_modules` are not supported.

Install **make** first (Debian/Ubuntu example):

```bash
sudo apt install -y make
```

Install dependencies and build:

```bash
make install
```

For subsequent builds:

```bash
make build
```

The compiled binary is written to `bin/cpbx`.

To run the frontend development server, set the backend WebSocket URL and start it:

```bash
export WS_BACKEND_OVERRIDE=wss://HOST:PORT/ws
make front-serve
```

Copy `config.example.json` to the ignored `config.json` for a non-container runtime. Never commit the runtime file.

WebSocket origins default to `same_origin`. Use `allow_list` with exact `allowed_origins` when the UI is hosted separately. The explicit `allow_all` policy is for development only. WebSocket timing and queue settings are documented in `config.example.json`.

Alternatively, you can utilize the precompiled binary available on the **[Releases Page](https://github.com/custompbx/custompbx/releases)**.

---
## Docker build and test targets

Docker is the recommended build path for the complete embedded frontend/backend artifact.

Useful targets:

```bash
make docker-fmt
make docker-vet
make docker-test
make docker-race
make docker-frontend-test
make docker-frontend-build
make docker-integration-test
make docker-release
```

The Docker build uses `npm ci` and embeds the container-generated frontend assets. Runtime configurations and secrets are excluded from the build context.

Run Angular unit tests with containerized Chromium:

```bash
make docker-frontend-test
```

For Linux/WSL without Docker, run `npm ci` followed by `npm run test:ci` from `src/cweb-app`.

## Docker Compose quick start

Compose starts PostgreSQL, FreeSWITCH, and CustomPBX with demo-only credentials. First create the ignored runtime configuration:

```bash
cp docker/config.example.json docker/config.json
```

PowerShell equivalent:

```powershell
Copy-Item docker/config.example.json docker/config.json
```

Then start the stack:

```bash
docker compose up -d
```

The example credentials are intended only for a local evaluation environment. Replace them before any shared or Internet-accessible deployment.

The quick-start stack creates the demo CDR table and connects `mod_cdr_pg_csv` to PostgreSQL. This bootstrap replaces only FreeSWITCH's unchanged `host=localhost dbname=cdr` sample value; operator-configured CDR database settings are preserved.

After the containers start, open `https://127.0.0.1:8080/cweb` (or the equivalent Docker host address) and accept the development self-signed certificate.

To rebuild and replace only the CustomPBX application container without recreating PostgreSQL or FreeSWITCH:

```bash
make docker-local-recreate
```

## Pull the published image

Images are available on the **[Packages Page](https://github.com/custompbx/custompbx/pkgs/container/custompbx)**.

```bash
docker pull ghcr.io/custompbx/custompbx:latest
```

---

## Documentation

For detailed instructions on **Installation** and **Configuration**, please refer to the project's **[Wiki Page](https://github.com/custompbx/custompbx/wiki)**.

If you have any questions or feedback, don't hesitate to get in touch through the **[discussions](https://github.com/custompbx/custompbx/discussions)** or by opening an **[issue](https://github.com/custompbx/custompbx/issues)**!

![system diagram](https://github.com/custompbx/doc/raw/master/img/Diagram1.png)


<details>
  <summary>Old GUI style</summary>

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

</details>
---

### GUI Screenshots


<img width="2558" height="1401" alt="Opera Instantané_2026-07-18_185430_localhost" src="https://github.com/user-attachments/assets/a35a28ba-de0f-4afc-abf5-f25faa03d81a" />

---


<img width="2558" height="1401" alt="Opera Instantané_2026-07-18_185623_localhost" src="https://github.com/user-attachments/assets/4c53404a-601f-4f5d-b2c3-2897ac1e1204" />

---


<img width="2558" height="1401" alt="Opera Instantané_2026-07-18_185658_localhost" src="https://github.com/user-attachments/assets/7f848675-e037-4c60-8d99-628995cb6c27" />

---


<img width="2558" height="1401" alt="Opera Instantané_2026-07-18_185725_localhost" src="https://github.com/user-attachments/assets/e0a4b5e5-c758-4ca3-a5ba-b67602fc405d" />

---


<img width="2558" height="1401" alt="Opera Instantané_2026-07-18_185928_localhost" src="https://github.com/user-attachments/assets/5a2ce5d1-7d06-43d0-9af7-88686f3a82b6" />

---


<img width="2558" height="1401" alt="Opera Instantané_2026-07-18_190053_localhost" src="https://github.com/user-attachments/assets/dd8ce1bb-89db-4181-9d6a-87c250c1d979" />

---


<img width="2558" height="1401" alt="Opera Instantané_2026-07-18_192213_localhost" src="https://github.com/user-attachments/assets/0857bbf4-925f-4b6c-8bff-63b267b8c73c" />

---


<img width="2558" height="1401" alt="Opera Instantané_2026-07-18_191125_localhost" src="https://github.com/user-attachments/assets/d3a1f8b3-181a-4795-89bc-c49d0b4f5ed7" />

---


<img width="2558" height="1401" alt="Opera Instantané_2026-07-18_191048_localhost" src="https://github.com/user-attachments/assets/8a9786bb-d5fc-4c1a-8ad6-65daf4cd5e6f" />

---


<img width="2558" height="1401" alt="Opera Instantané_2026-07-18_190105_localhost" src="https://github.com/user-attachments/assets/dde07e54-40ee-45aa-9685-7debd1dcb26a" />

---

<img width="2558" height="1401" alt="Opera Instantané_2026-07-18_192446_localhost" src="https://github.com/user-attachments/assets/0fc2baeb-b755-4f3e-b1d1-fc8314a14035" />

---

<img width="2558" height="1401" alt="Opera Instantané_2026-07-18_192530_localhost" src="https://github.com/user-attachments/assets/34c77b9d-0553-4a54-bce6-07314def1e8b" />

---

<img width="2558" height="1401" alt="Opera Instantané_2026-07-18_192324_localhost" src="https://github.com/user-attachments/assets/e590aaf0-aece-4a5e-899d-7efa7995e55b" />

---


<img width="2558" height="1401" alt="Opera Instantané_2026-07-18_192427_localhost" src="https://github.com/user-attachments/assets/2b8e2611-9359-4525-84cf-35aadecde036" />





