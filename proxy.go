package main

import (
    "bufio"
    "crypto/tls"
    "encoding/base64"
    "encoding/json"
    "flag"
    "fmt"
    "io"
    "log"
    "net"
    "net/http"
    "net/url"
    "os"
    "strings"
    "time"
)

var (
    useJSONLog bool
)

func logMsg(msg string, fields ...any) {
    if useJSONLog {
        data := map[string]any{"msg": msg, "ts": time.Now().Format(time.RFC3339)}
        for i := 0; i < len(fields)-1; i += 2 {
            key := fmt.Sprintf("%v", fields[i])
            val := fields[i+1]
            data[key] = val
        }
        json.NewEncoder(os.Stdout).Encode(data)
    } else {
        log.Printf(msg, fields...)
    }
}

func isBrokenPipe(err error) bool {
    return err != nil && strings.Contains(err.Error(), "broken pipe")
}

func handleConnection(localConn net.Conn, proxyAddr, remoteHost, proxyUser, proxyPass string, timeout time.Duration, useTLS bool) {
    defer localConn.Close()

    logMsg("Connessione accettata", "client", localConn.RemoteAddr(), "target", remoteHost)

    var proxyConn net.Conn
    var err error

    d := net.Dialer{Timeout: timeout}
    if useTLS {
        proxyConn, err = tls.DialWithDialer(&d, "tcp", proxyAddr, &tls.Config{
            InsecureSkipVerify: true,
        })
    } else {
        proxyConn, err = d.Dial("tcp", proxyAddr)
    }
    if err != nil {
        logMsg("Errore connessione al proxy", "proxy", proxyAddr, "error", err.Error())
        return
    }
    defer proxyConn.Close()

    headers := make(http.Header)
    if proxyUser != "" && proxyPass != "" {
        auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", proxyUser, proxyPass)))
        headers.Set("Proxy-Authorization", "Basic "+auth)
    }

    req := &http.Request{
        Method: "CONNECT",
        URL:    &url.URL{},
        Host:   remoteHost,
        Header: headers,
    }

    err = req.Write(proxyConn)
    if err != nil {
        logMsg("Errore invio richiesta CONNECT", "error", err.Error())
        return
    }

    resp, err := http.ReadResponse(bufio.NewReader(proxyConn), req)
    if err != nil {
        logMsg("Errore risposta dal proxy", "error", err.Error())
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        logMsg("Proxy ha rifiutato la connessione", "status", resp.Status)
        return
    }

    logMsg("Tunnel stabilito", "client", localConn.RemoteAddr(), "remote", remoteHost)

    done := make(chan struct{})

    go func() {
        _, err := io.Copy(proxyConn, localConn)
        if err != nil && !isBrokenPipe(err) {
            logMsg("Errore relay verso proxy", "error", err.Error())
        }
        close(done)
    }()

    _, err = io.Copy(localConn, proxyConn)
    if err != nil && !isBrokenPipe(err) {
        logMsg("Errore relay verso client", "error", err.Error())
    }

    <-done
    logMsg("Connessione chiusa", "client", localConn.RemoteAddr())
}

func main() {
    localPort := flag.Int("listen", 0, "Porta locale da esporre")
    proxyHost := flag.String("proxy", "", "Proxy (host:porta)")
    remote := flag.String("remote", "", "Destinazione remota (host:porta)")
    proxyUser := flag.String("user", "", "Username proxy (opzionale)")
    proxyPass := flag.String("pass", "", "Password proxy (opzionale)")
    timeoutSec := flag.Int("timeout", 30, "Timeout connessioni in secondi")
    tlsFlag := flag.Bool("tls", false, "Usa TLS verso il proxy (HTTPS)")
    jsonLogFlag := flag.Bool("jsonlog", true, "Log in formato JSON")

    flag.Parse()

    useJSONLog = *jsonLogFlag
    timeout := time.Duration(*timeoutSec) * time.Second

    if *localPort == 0 || *proxyHost == "" || *remote == "" {
        fmt.Println("Parametri mancanti.")
        flag.Usage()
        os.Exit(1)
    }

    listenAddr := fmt.Sprintf(":%d", *localPort)
    listener, err := net.Listen("tcp", listenAddr)
    if err != nil {
        log.Fatalf("Errore ascolto: %v", err)
    }

    logMsg("In ascolto", "porta", *localPort, "proxy", *proxyHost, "destinazione", *remote)

    for {
        conn, err := listener.Accept()
        if err != nil {
            logMsg("Errore accettando connessione", "error", err.Error())
            continue
        }
        go handleConnection(conn, *proxyHost, *remote, *proxyUser, *proxyPass, timeout, *tlsFlag)
    }
}
