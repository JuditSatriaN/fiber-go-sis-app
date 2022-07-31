package custom

import (
	"embed"
	"path"
	"path/filepath"
	"strings"
)

func ReadFileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func GetAllFilenames(fs *embed.FS, dir string) (out []string, err error) {
	if len(dir) == 0 {
		dir = "."
	}

	entries, err := fs.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		fp := path.Join(dir, entry.Name())

		if entry.IsDir() {
			res, err := GetAllFilenames(fs, fp)
			if err != nil {
				return nil, err
			}

			out = append(out, res...)

			continue
		}

		out = append(out, fp)
	}

	return
}

func GetAllContentFiles(fs *embed.FS, dir string) (out []string, err error) {
	if len(dir) == 0 {
		dir = "."
	}

	entries, err := fs.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		fp := path.Join(dir, entry.Name())

		if entry.IsDir() {
			res, err := GetAllContentFiles(fs, fp)
			if err != nil {
				return nil, err
			}

			out = append(out, res...)

			continue
		}

		contents, _ := fs.ReadFile(fp)
		out = append(out, string(contents))
	}

	return
}
