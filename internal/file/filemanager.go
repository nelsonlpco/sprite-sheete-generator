package file

import (
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

type Manager struct {
	basePath string
	images   []string
	size     uint
}

func (f Manager) OpenImageByIndex(index uint) *os.File {
	path := fmt.Sprintf("%s/%s", f.basePath, f.images[index])
	fileImage, err := os.Open(path)
	if err != nil {
		log.Panicf("path: %v, error: %v", path, err)
	}

	return fileImage
}

func (f Manager) Size() uint {
	return f.size
}

func (f Manager) Save(outputPath string, output *image.RGBA) {
	outputFile, err := os.Create(outputPath)
	if err != nil {
		log.Fatalf("error on save file: %v", outputPath)
	}
	defer outputFile.Close()

	enc := png.Encoder{
		CompressionLevel: png.BestCompression,
	}
	err = enc.Encode(outputFile, output)
	if err != nil {
		log.Fatalf("error on encode image: %v", outputPath)
	}
}

func New(basePath string) *Manager {
	files := listFiles(basePath)
	return &Manager{
		basePath: basePath,
		images:   files,
		size:     uint(len(files)),
	}
}

func listFiles(path string) []string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	var names []string

	for _, f := range files {
		names = append(names, f.Name())
	}

	sort.Slice(names, func(i int, j int) bool {
		return names[i] < names[j]
	})

	return names
}
