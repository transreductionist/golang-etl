package buildenv

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

func FilePaths() {
    _ = os.Setenv("DATA_CSV", filepath.Dir("C:\\Users\\Aaron\\go-modules\\USERetl"))
    _ = os.Setenv("LOG_TXT", filepath.Dir("C:\\Users\\Aaron\\go-modules\\USERetl"))
}

func GetFilePath(category string, fileName string) string {
    fileExtension := strings.Trim(strings.ToUpper(filepath.Ext(fileName)), ".")
    fileCategory := strings.ToUpper(category)
    environment := fmt.Sprintf("%s_%s", fileCategory, fileExtension)
    fileRoot := os.Getenv(environment)
    filePath := fmt.Sprintf("%s\\%s", fileRoot, fileName)
    return filePath
}
