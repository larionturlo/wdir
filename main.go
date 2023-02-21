package main

import (
	"flag"
	"fmt"
	"strings"
	"os"
)

const (
	LinkDirName = "dir" 
)

func checkOrCreateStorage(storage string) error {
	_, err := os.Stat(storage)
	if os.IsNotExist(err) {
		errMk := os.MkdirAll(storage, 0755)
		if errMk != nil && !os.IsExist(err) {
			return errMk
		}
	}
	if err != nil {
		return err
	}
	return nil
}

func setDir(key, path, storage string) error {
	if len(key) == 0 {
		return fmt.Errorf("Key is empty!")
	}
    if len(path) == 0 {
		return fmt.Errorf("Key is empty!")
	}
	wdStorage := fmt.Sprintf("%s/%s", strings.TrimRight(storage, "/"), strings.Trim(key, "/"))
	
	if err := checkOrCreateStorage(wdStorage); err != nil {
		return err;
	}

	linkDir := fmt.Sprintf("%s/%s", wdStorage, LinkDirName)

	_, err := os.Stat(linkDir)
	if err == nil {
		os.Remove(linkDir)
	}
    return os.Symlink(path, linkDir)		
}

func getDir(key, storage string) (string, error) {
	if len(key) == 0 {
		return "", fmt.Errorf("Key is empty!")
	}

	linkDir := fmt.Sprintf(
		"%s/%s/%s",
		strings.TrimRight(storage, "/"),
		strings.Trim(key, "/"),
		LinkDirName,
	)

	return os.Readlink(linkDir)
}

func main(){
	set := flag.Bool("s", false, "The flag for save working directory by key")
	key := flag.String("k", "", "A working directory key/name")
	path := flag.String("p", "", "A path to working directory")
	storagePath := flag.String("storage", "~/.config/wdir", "A directory to store workdir configs")

	flag.Parse()

	if *set {
		err := setDir(*key, *path, *storagePath)
		if err!=nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		dir, err := getDir(*key, *storagePath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Print(dir)
	}

	os.Exit(0)
}
