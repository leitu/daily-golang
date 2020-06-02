package main

import (
        "os"
        "net/http"
        "log"
)

func main() {
        log.SetFlags(log.Lshortfile)
        http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {

                // ServeContent uses the name for mime detection
                size := int64(100 * 1024 * 1024)
                name, err := os.Create("output")
                if err != nil {
                        log.Fatal("Failed to create output")
                }
                _, err = name.Seek(size-1, 0)
                if err != nil {
                        log.Fatal("Failed to seek")
                }
                _, err = name.Write([]byte{0})
                if err != nil {
                        log.Fatal("Write failed")
                }
                err = name.Close()
                if err != nil {
                        log.Fatal("Failed to close file")
                }

                // tell the browser the returned content should be downloaded
                w.Header().Add("Content-Disposition", "Attachment")

                http.ServeFile(w, req, "output")
        })

        log.Fatal(http.ListenAndServe(":8100", nil))
}
