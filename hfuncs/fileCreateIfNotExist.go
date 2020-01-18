package hfuncs

import "os"

func CreateFolderIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0750)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	return nil
}
