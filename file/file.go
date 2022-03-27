package file

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

func ReadAllText(path string) (string, error) {

	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	bytes, err := ioutil.ReadAll(file)

	if err = file.Close(); err != nil {
		return "", err
	}

	return string(bytes), nil
}

func ReadAllLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func FileExists(path string) (bool, error) {

	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}


func DirectoryExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func WriteAllText(filePath string, text string) (err error) {
	f, _ := os.Create(filePath)
	defer f.Close()
	_, err = f.WriteString(text)
	return
}

func GetCurrentDirectory() string {
	if dir, err := os.Getwd(); err == nil {
		return dir
	}
	return ""
}


func pathJoin(src string, des string) string {
	newSrcPath := strings.TrimRight(src, "\\")
	newSrcPath = strings.TrimRight(src, "/")
	newDestPath := strings.TrimLeft(des, "\\")
	newDestPath = strings.TrimLeft(des, "/")
	return newSrcPath + string(os.PathSeparator) + newDestPath
}

func PathJoin(src string, path ...string) string {
	result := src
	for _, p := range path {
		result = pathJoin(result, p)
	}
	return result
}
