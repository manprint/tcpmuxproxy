# Proxy

## Compose Go-Alpine

```
services:
  proxytunnel-alpine-go:
    image: fabiop85/proxytunnel:alpine-go
    container_name: proxytunnel-alpine-go
    restart: always
    ports:
      - "6000:6000"
    network_mode: "bridge"
    command: >
      /proxy -listen 6000 -proxy sub.mydomain.tld:443 -remote level.sub.mydomain.tld:5900 -user admin -pass abc
```

## Compose Go-Debian

```
services:
  proxytunnel-go:
    image: fabiop85/proxytunnel:go
    container_name: proxytunnel-go
    restart: always
    ports:
      - "6000:6000"
    network_mode: "bridge"
    command: >
      proxy -listen 6000 -proxy sub.mydomain.tld:443 -remote level.sub.mydomain.tld:5900 -user admin -pass abc
```

## Compose proxytunnel native

```
services:
  proxytunnel:
    image: fabiop85/proxytunnel:latest
    container_name: proxytunnel
    network_mode: bridge
    restart: always
    ports:
      - 6000:6000
    environment:
      - BASIC_AUTH=admin:admin
      - PROXY_HOST=3lev.mydomain.tld:443
      - PROXY_REMOTE=4lev.3lev.mydomain.tld:5900
      - LOCAL_PORT=6000
```

## Help

```
Usage of ./proxy-linux-amd64:
  -jsonlog
    	Log in formato JSON (default true)
  -listen int
    	Porta locale da esporre
  -pass string
    	Password proxy (opzionale)
  -proxy string
    	Proxy (host:porta)
  -remote string
    	Destinazione remota (host:porta)
  -timeout int
    	Timeout connessioni in secondi (default 30)
  -tls
    	Usa TLS verso il proxy (HTTPS)
  -user string
    	Username proxy (opzionale)
```

## Example

```
proxy -listen 6000 -proxy sub.mydomain.tld:443 -remote level.sub.mydomain.tld:5900
```
```
proxy -listen 6000 -proxy sub.mydomain.tld:443 -remote level.sub.mydomain.tld:5900 -user admin -pass abc
```

## Build

```
# Compilazione statica per Linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./builds/proxy-linux-amd64
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ./builds/proxy-linux-arm64

# Compilazione statica per Windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./builds/proxy-win-x64.exe
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o ./builds/proxy-win-x86.exe

# Compilazione statica per macOS
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./builds/proxy-macos-darwin-amd64
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o ./builds/proxy-macos-darwin-arm64
```

## Extra - ProxyTunnel

Soluzione con proxytunnel

windows: `https://github.com/proxytunnel/proxytunnel/releases/download/v1.12.3/proxytunnel-v1.12.3-x86_64-windows-msys.zip`

mac: `brew install proxytunnel`

alpine: `apk add proxytunnel`

debian: `apt install proxytunnel`

```
proxytunnel -p sub.mydomain.tld:443 -d level.sub.mydomain.tld:5900 -a 6000 -P test:test
./proxytunnel.exe -v -p sub.mydomain.tld:443 -d level.sub.mydomain.tld:5900 -a 0.0.0.0:7000 -P test:test
```

## Android Termux

Install android from Fdroid (not google store!)

```
$ pkg in proot-distro
$ proot-distro install debian
$ proot-distro login Debian
root@localhost:~# apt update
root@localhost:~# apt install proxytunnel
```

```
proot-distro login debian --bind /storage/emulated/0:/root/storage
```

## Extra - Socat

**Nota bene**: non supporta autenticazione

```
socat TCP-LISTEN:6000,fork,reuseaddr PROXY:sub.mydomain.tld:level.sub.mydomain.tld:5900,proxyport=443
```