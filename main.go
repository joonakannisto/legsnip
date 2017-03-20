package main

import (
  "github.com/golang/crypto/acme/autocert"
  "crypto/tls"
  "log"
  "net/http"
  "net/http/httputil"
  "net/url"
)

func main() {
        proxy := httputil.NewSingleHostReverseProxy(&url.URL{
                Scheme: "http",
                Host:   "localhost:80",
        })
  m := autocert.Manager{
        Prompt: autocert.AcceptTOS,
        // Default is allow all :)
        //HostPolicy: autocert.HostWhitelist("wizard.kyber.space"),
  }
  s := &http.Server{
        Addr: ":443",
        Handler: proxy,
        TLSConfig: &tls.Config{GetCertificate: m.GetCertificate},
  }
    err := s.ListenAndServeTLS("", "")
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
