package api

import (
	"archive/tar"
	"archive/zip"
	"fmt"
	"io"
	"os"
)

func TarToZip(tarPath, zipPath string) error {
	tarFile, err := os.Open(tarPath)
	if err != nil {
		return fmt.Errorf("failed to open tar file: %w", err)
	}
	defer tarFile.Close()

	tr := tar.NewReader(tarFile)

	zipFile, err := os.Create(zipPath)
	if err != nil {
		return fmt.Errorf("failed to create zip file: %w", err)
	}
	defer zipFile.Close()

	zw := zip.NewWriter(zipFile)
	defer zw.Close()

	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read tar header: %w", err)
		}

		zipHeader, err := zip.FileInfoHeader(header.FileInfo())
		if err != nil {
			return fmt.Errorf("failed to create zip header: %w", err)
		}
		zipHeader.Name = header.Name

		if header.Typeflag == tar.TypeDir {
			zipHeader.Name += "/"
			_, err := zw.CreateHeader(zipHeader)
			if err != nil {
				return fmt.Errorf("failed to create zip directory entry: %w", err)
			}
		} else if header.Typeflag == tar.TypeReg {
			fw, err := zw.CreateHeader(zipHeader)
			if err != nil {
				return fmt.Errorf("failed to create zip file entry: %w", err)
			}
			if _, err := io.Copy(fw, tr); err != nil {
				return fmt.Errorf("failed to copy file content to zip: %w", err)
			}
		}
	}

	return nil
}
