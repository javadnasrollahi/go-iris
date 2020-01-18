package hfuncs

import (
	"errors"
	"strings"
)

func FileType(filename string) (string, error) {
	FType := strings.Split(filename, ".")
	if len(FType) < 2 {
		return "", errors.New("NotSupportedType")

	}
	return FType[len(FType)-1], nil

}

func IsImage(filename string) error {
	ft, err := FileType(filename)
	//log.Println(ft)
	if err != nil {
		return err
	}
	switch ft {
	case "png", "PNG", "JPG", "JPEG", "jpg", "jpeg":
		return nil

	}
	return errors.New("NotImage")
}
func IsValidType(filename string) error {
	ft, err := FileType(filename)
	if err != nil {
		return err
	}
	switch ft {
	case "png":
	case "PNG":
	case "JPG":
	case "JPEG":
	case "jpg":
	case "jpeg":
	case "mp3":
	case "wma":
	case "zip":
	case "tar":
	case "gzip":
	case "mp4":
	case "avi":
		return nil

	}
	return errors.New("NotImage")
}
