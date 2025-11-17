package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const component_path_base = "../components/%s/"

func getComponents(w http.ResponseWriter, r *http.Request) []string {
	var file_paths []string

	path_value := r.PathValue("params")
	params := strings.SplitSeq(path_value, "&")

	for component_name := range params {
		filepath := filepath.Join(component_path_base, component_name)
		
		if _, err := os.Stat(filepath); os.IsNotExist(err) {
			http.Error(w, "Component Not Found!", http.StatusNotFound)
			return nil
		}

		file_paths = append(file_paths, filepath)
	}

	return file_paths
}

func bundleComponents(component_file_paths []string) {
	// error here!
	temp_dir := filepath.Join(os.TempDir(), "bundle-"+string(md5.Sum([]byte(component_file_paths[0])))
	os.MkdirAll(temp_dir, 0755)
	defer os.RemoveAll(temp_dir)

	entryFile := filepath.Join(temp_dir, "entry.js")

	var imports []string
	for i, comp := range component_file_paths {
		comp = strings.TrimSpace(comp)
		absPath, _ := filepath.Abs(filepath.Join(comp, comp+".js"))
		fmt.Println(absPath)
	}


	outputPath := filepath.Join(temp_dir, "bundle.js")

	cmd := exec.Command("npx", "esbuild", en)
}

func main() {
	router := http.NewServeMux()

	router.Handle("GET /", http.FileServer(http.Dir("./static")))
	router.HandleFunc("GET /components/{params}", func(w http.ResponseWriter, r *http.Request) {
		file_paths := getComponents(w, r)

	})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Starting server on port ", server.Addr)
	log.Fatalf("Server failed with error: %s", server.ListenAndServe().Error())
}
