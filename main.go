package main

import (
	"embed"
	"io/fs"
	"log"
	"net"
	"net/http"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/Fovir-GitHub/grain-128aeadv2-go/internal/handler"
	"github.com/Fovir-GitHub/grain-128aeadv2-go/internal/service"
)

//go:embed frontend/*
var frontendFS embed.FS

// TODO:
// - Refactor code
// - Error handling
func main() {
	sub, err := fs.Sub(frontendFS, "frontend")
	if err != nil {
		log.Fatal(err)
	}

	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal(err)
	}

	addr := "http://localhost:" + strconv.Itoa(ln.Addr().(*net.TCPAddr).Port)
	log.Println(addr)
	go openBrowser(addr)

	mux := http.NewServeMux()
	srv := service.New()
	handler := handler.New(srv)
	handler.Register(mux)

	mux.Handle("/", http.FileServer(http.FS(sub)))
	log.Fatal(http.Serve(ln, mux))
}

func openBrowser(url string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}
	if err := cmd.Start(); err != nil {
		log.Println("open browser failed:", err)
	}
}
