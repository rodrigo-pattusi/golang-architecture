package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("rumi.txt")

	var perr *os.PathError

	switch {
	case errors.Is(err, os.ErrPermission) && errors.As(err, &perr):
		err = fmt.Errorf("you do not have permission to open the file: %w and the path is %s", err, perr.Path)
		log.Println(err)
	case errors.Is(err, os.ErrPermission) && errors.As(err, &perr):
		err = fmt.Errorf("the file does not exist: %w and the path %s", err, perr.Path)
		log.Println(err)
	case errors.As(err, &perr):
		err = fmt.Errorf("here is the original error %w and the path error %s", err, perr.Path)
		log.Println(err)
	case err != nil:
		err = fmt.Errorf("file couldn't be opened: %w", err)
		log.Println(err)
	}

	//if errors.Is(err, os.ErrPermission) && errors.As(err, &perr) {
	//	err = fmt.Errorf("you do not have permission to open the file: %w and the path is %s", err, perr.Path)
	//	log.Println(err)
	//} else if errors.Is(err, os.ErrPermission) && errors.As(err, &perr) {
	//	err = fmt.Errorf("the file does not exist: %w and the path %s", err, perr.Path)
	//	log.Println(err)
	//} else if errors.As(err, &perr) {
	//	// fmt.Errorf("here is the original error %w and the path %s", err, perr.Path) // RETURNS A STRING
	//	err = fmt.Errorf("here is the original error %w and the path error %s", err, perr.Path)
	//	log.Println(err)
	//} else if err != nil {
	//	err = fmt.Errorf("file couldn't be opened: %w", err)
	//	log.Println(err)
	//}
	defer f.Close()
}
