package main

import (
  "github.com/golang/crypto/acme/autocert"
  "crypto/tls"
  "log"
  "net/http"
  "ftm"
  "html"
)





func main() {


  http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
  	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
  })

  m := autocert.Manager{
  	Prompt: autocert.AcceptTOS,
  	HostPolicy: autocert.HostWhitelist("porn.kyber.space"),
  }
  s := &http.Server{
  	Addr: ":4430",
  	TLSConfig: &tls.Config{GetCertificate: m.GetCertificate},
  }
    err := s.ListenAndServeTLS(":4430", "")
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
