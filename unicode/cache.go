package unicode

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const Url = "https://unicode.org/Public/emoji/latest/emoji-test.txt"

func Download(url string) (content []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func CacheDir() string {
	HOME := os.Getenv("USERPROFILE")
	path := filepath.Join(HOME, ".cache/goemoji")
	if f, err := os.Stat(path); os.IsNotExist(err) || !f.IsDir() {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			panic("fail to mkdir")
		}
	}
	return path
}

func CachePath() string {
	return filepath.Join(CacheDir(), "emoji-test.txt")
}

func GetCache() []byte {
	path := CachePath()
	bytes, err := os.ReadFile(path)
	if err == nil {
		// return cache
		return bytes
	}

	// download
	fmt.Printf("download: %s <== %s ...", path, Url)
	bytes, err = Download(Url)
	if err != nil {
		panic("fail to download")
	}

	// write cache
	err = os.WriteFile(path, bytes, 0644)
	if err != nil {
		panic("fail to write cache")
	}

	return bytes
}
