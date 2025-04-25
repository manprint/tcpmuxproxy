# Proxy

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
GOOS=linux GOARCH=amd64 go build -o ./builds/proxy-linux-amd64
GOOS=linux GOARCH=arm64 go build -o ./builds/proxy-linux-arm64
GOOS=windows GOARCH=amd64 go build -o ./builds/proxy-win-x64.exe
GOOS=darwin GOARCH=amd64 go build -o ./builds/proxy-macos-darwin-amd64
GOOS=darwin GOARCH=arm64 go build -o ./builds/proxy-macos-darwin-arm64
```

## Extra - ProxyTunnel

Soluzione con proxytunnel

windows: `https://github.com/proxytunnel/proxytunnel/releases/download/v1.12.3/proxytunnel-v1.12.3-x86_64-windows-msys.zip`

mac: brew `install proxytunnel`

alpine: `apk add proxytunnel`

debian: `apt install proxytunnel`

```
proxytunnel -p sub.mydomain.tld:443 -d level.sub.mydomain.tld:5900 -a 6000 -P admin:abc
```

## Extra - Socat

**Nota bene**: non supporta autenticazione

```
socat TCP-LISTEN:6000,fork,reuseaddr PROXY:sub.mydomain.tld:level.sub.mydomain.tld:5900,proxyport=443
```