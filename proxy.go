package main

import (
    "bufio"
    "flag"
    "fmt"
    "io"
    "log"
    "net"
    "net/http"
    "net/url"
    "os"
)

func handleConnection(localConn net.Conn, proxyAddr, remoteHost string) {
    defer localConn.Close()

    log.Printf("Connessione accettata da %s, inoltro tramite proxy %s verso %s", localConn.RemoteAddr(), proxyAddr, remoteHost)

    proxyConn, err := net.Dial("tcp", proxyAddr)
    if err != nil {
        log.Printf("Errore nel connettersi al proxy %s: %v", proxyAddr, err)
        return
    }
    defer proxyConn.Close()

    log.Printf("Connessione al proxy %s stabilita", proxyAddr)

    req := &http.Request{
        Method: "CONNECT",
        URL:    &url.URL{},
        Host:   remoteHost,
        Header: make(http.Header),
    }

    err = req.Write(proxyConn)
    if err != nil {
        log.Printf("Errore nell'invio della richiesta CONNECT: %v", err)
        return
    }

    log.Printf("Richiesta CONNECT inviata al proxy %s per %s", proxyAddr, remoteHost)

    resp, err := http.ReadResponse(bufio.NewReader(proxyConn), req)
    if err != nil {
        log.Printf("Errore nella lettura della risposta del proxy: %v", err)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        log.Printf("Proxy ha risposto con codice %d", resp.StatusCode)
        return
    }

    log.Printf("Proxy ha risposto con successo, avvio il relay tra connessioni")

    go func() {
        _, err := io.Copy(proxyConn, localConn)
        if err != nil {
            log.Printf("Errore durante il relay verso il proxy: %v", err)
        }
    }()
    _, err = io.Copy(localConn, proxyConn)
    if err != nil {
        log.Printf("Errore durante il relay verso il client: %v", err)
    }
}

func main() {
    localPort := flag.Int("listen", 0, "Porta su cui rimanere in ascolto (obbligatorio)")
    proxyHost := flag.String("proxy", "", "Indirizzo del proxy (host:porta) (obbligatorio)")
    remote := flag.String("remote", "", "Host remoto da raggiungere tramite proxy (host:porta) (obbligatorio)")

    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "Utilizzo: %s -listen <porta> -proxy <proxy_host:porta> -remote <remote_host:porta>\n", os.Args[0])
        flag.PrintDefaults()
    }

    flag.Parse()

    if *localPort == 0 || *proxyHost == "" || *remote == "" {
        log.Println("Errore: tutti i parametri -listen, -proxy e -remote sono obbligatori")
        flag.Usage()
        os.Exit(1)
    }

    listenAddr := fmt.Sprintf(":%d", *localPort)
    listener, err := net.Listen("tcp", listenAddr)
    if err != nil {
        log.Fatalf("Errore nell'ascolto su %s: %v", listenAddr, err)
    }
    log.Printf("In ascolto su %s, inoltro a %s via proxy %s", listenAddr, *remote, *proxyHost)

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Printf("Errore accettando connessione: %v", err)
            continue
        }
        go handleConnection(conn, *proxyHost, *remote)
    }
}
