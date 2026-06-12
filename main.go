package main

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"github.com/Fovir-GitHub/grain-128aeadv2-go/internal/handler"
	"github.com/Fovir-GitHub/grain-128aeadv2-go/internal/service"
)

//go:embed frontend/*
var frontendFS embed.FS

func main() {
	sub, err := fs.Sub(frontendFS, "frontend")
	if err != nil {
		slog.Error("create sub fs failed", "err", err)
		os.Exit(1)
	}

	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		slog.Error("listen failed", "err", err)
		os.Exit(1)
	}

	port := ln.Addr().(*net.TCPAddr).Port
	addr := fmt.Sprintf("http://localhost:%d", port)
	slog.Info("starting server", "addr", addr)

	go func() {
		if err := openBrowser(addr); err != nil {
			slog.Warn("open browser failed", "err", err)
		}
	}()

	mux := http.NewServeMux()
	srv := service.New()
	h := handler.New(srv)
	h.Register(mux)
	mux.Handle("/", http.FileServer(http.FS(sub)))

	server := &http.Server{
		Handler: mux,
	}
	if err := server.Serve(ln); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("server stopped", "err", err)
		os.Exit(1)
	}
}

func openBrowser(url string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}

	err := cmd.Start()
	return err
}
