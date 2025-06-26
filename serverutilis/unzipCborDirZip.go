package serverutilis

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func UnzipCborDirZip(zipFilePath string) error {
	r, err := zip.OpenReader(zipFilePath)
	if err != nil {
		return fmt.Errorf("cant open zip reader: %w", err)
	}
	defer r.Close()
	parentDir := filepath.Dir(zipFilePath)
	folderName := strings.TrimSuffix(filepath.Base(zipFilePath), filepath.Ext(zipFilePath))
	unzippedFolderPath := filepath.Join(parentDir, folderName)
	os.MkdirAll(unzippedFolderPath, 0777)
	for _, file := range r.File {
		log.Printf("unzipping %s\n", file.Name)
		readcloser, err := file.Open()
		if err != nil {
			return fmt.Errorf("cant open file %s : %w", file.Name, err)
		}
		defer readcloser.Close()

		if file.FileInfo().IsDir() {
			return fmt.Errorf("directory named %s was found inside cbor zip file, cbor zip file invalid format", file.Name)
		}

		unzippedFile, err := os.Create(filepath.Join(unzippedFolderPath, file.Name))
		if err != nil {
			return fmt.Errorf("cant create file %s for writing unzipped file : %w", file.Name, err)
		}
		_, err = io.Copy(unzippedFile, readcloser)
		if err != nil {
			return fmt.Errorf("cant extract file %s to folder %s during unzipping: %w", file.Name, folderName, err)
		}
		log.Printf("extracted file %s\n", file.Name)
	}
	return nil
}
