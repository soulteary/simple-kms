package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/denisbrodbeck/machineid"
	"github.com/google/uuid"
	"github.com/soulteary/simple-kms/internal/cmd"
	"github.com/soulteary/simple-kms/internal/define"
	"github.com/soulteary/simple-kms/internal/filler"
	"github.com/soulteary/simple-kms/internal/home"
	"github.com/soulteary/simple-kms/internal/transformer"

	cloudid "github.com/soulteary/go-cloud-id"
)

func main() {
	cliMode := flag.Bool("generate", false, "generate ak/sk pairs")
	flag.Parse()

	if *cliMode {
		cmd.Generate()
		os.Exit(0)
	}

	// 1. generate a random identifier
	id := uuid.New().String()

	// 2. try to fetch the machine id
	machineID, err := machineid.ID()
	if err == nil {
		id = machineID
	}

	// 3. try to fetch the cloud instance id
	sn, err := cloudid.GetAliyunSerialNumber()
	if err == nil {
		fmt.Println("use stable serial number 🚀")
		id = sn
	}

	// 4. generate a encrypted id with accesskey
	ak := cmd.GetAccessKey()
	// 4.1 generate seed and padding
	seed, padding := filler.GetFillers()
	if ak != "" {
		id = transformer.Encode(id, ak, seed, padding)
	}
	// `id`` ready for exposing

	// homepage for fun
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if strings.HasPrefix(req.RequestURI, "/favicon.ico") {
			w.Header().Set("Content-Type", "image/x-icon")
			w.Write([]byte(""))
		} else {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Write(home.Template)
		}
	})

	// api for sharing basic data
	http.HandleFunc(define.API_DATA, func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte(fmt.Sprintf("%s\n%s", id, time.Now().Format("2006-01-02 15:04:05"))))
	})

	// api for sharing seed
	http.HandleFunc(define.API_SEED, func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte(fmt.Sprintf("%s\n", string(seed))))
	})

	// api for sharing padding
	http.HandleFunc(define.API_PADDING, func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte(fmt.Sprintf("%s\n", padding)))
	})

	http.ListenAndServe(define.DEFAULT_PORT, nil)
}
