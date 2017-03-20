package main

import (
  "github.com/golang/crypto/autocert"
  "crypto/tls"
  "log"
  "net/http"
)



func HelloServer(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
    w.Write([]byte("This is an example server.\n"))
    // fmt.Fprintf(w, "This is an example server.\n")
    // io.WriteString(w, "This is an example server.\n")
}

func main() {
  m := autocert.Manager{
  	Prompt: autocert.AcceptTOS,
  	HostPolicy: autocert.HostWhitelist("porn.kyber.space"),
  }
  s := &http.Server{
  	Addr: ":4430",
    Handler:        HelloServer,
  	TLSConfig: &tls.Config{GetCertificate: m.GetCertificate},
  }
    err := s.ListenAndServeTLS(":4430", "")
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
