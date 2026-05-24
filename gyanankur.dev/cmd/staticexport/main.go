package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gyanankur.dev/internal/handlers"
)

func main() {
	outDir := "public"

	if err := os.RemoveAll(outDir); err != nil {
		log.Fatalf("clean public dir: %v", err)
	}
	if err := os.MkdirAll(filepath.Join(outDir, "static"), 0o755); err != nil {
		log.Fatalf("create public dir: %v", err)
	}

	h, err := handlers.New()
	if err != nil {
		log.Fatalf("init handlers: %v", err)
	}

	var buf strings.Builder
	if err := h.RenderIndex(&buf); err != nil {
		log.Fatalf("render index: %v", err)
	}

	html := strings.ReplaceAll(buf.String(), "/static/", "static/")

	if err := os.WriteFile(filepath.Join(outDir, "index.html"), []byte(html), 0o644); err != nil {
		log.Fatalf("write index.html: %v", err)
	}

	if err := copyDir("web/static", filepath.Join(outDir, "static")); err != nil {
		log.Fatalf("copy static assets: %v", err)
	}

	log.Printf("static site exported to %s/", outDir)
}

func copyDir(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dst, rel)

		if info.IsDir() {
			return os.MkdirAll(target, 0o755)
		}

		return copyFile(path, target)
	})
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		return err
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}
