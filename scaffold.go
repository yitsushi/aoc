package aoc

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// Scaffold generates a Solution directory structure for a day.
func Scaffold(templateDir string, outputDir string, variables interface{}) error {
	ensureDirectory(outputDir)

	return filepath.Walk(
		templateDir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				ensureDirectory(path)

				return nil
			}

			dir := filepath.Join(
				outputDir,
				strings.TrimPrefix(
					strings.TrimSuffix(path, filepath.Base(path)),
					templateDir,
				),
			)
			ensureDirectory(dir)

			targetFile := strings.TrimSuffix(filepath.Join(dir, filepath.Base(path)), ".tmpl")

			f, err := os.Create(targetFile)
			if err != nil {
				return err
			}

			defer f.Close()

			tmpl, err := template.ParseFiles(path)
			if err != nil {
				return err
			}

			return tmpl.Execute(f, variables)
		},
	)
}
