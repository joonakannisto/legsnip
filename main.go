package main

import (
  "github.com/golang/crypto/blob/master/acme/autocert"
  "net/http"
)

m := autocert.Manager{
	Prompt: autocert.AcceptTOS,
	HostPolicy: autocert.HostWhitelist("porn.kyber.space"),
}
s := &http.Server{
	Addr: ":https",
	TLSConfig: &tls.Config{GetCertificate: m.GetCertificate},
}

func HelloServer(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
    w.Write([]byte("This is an example server.\n"))
    // fmt.Fprintf(w, "This is an example server.\n")
    // io.WriteString(w, "This is an example server.\n")
}

func main() {
    http.HandleFunc("/hello", HelloServer)
    err := s.ListenAndServeTLS(":4430", "")
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
