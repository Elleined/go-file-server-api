package file

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sizeTable(sizeType string) int {
	return map[string]int{
		"KB": 10, // (kilobyte)
		"MB": 20, // (Megabyte)
		"GB": 30, // (Gigabyte)
		"TB": 40, // (Terabyte)
		"PB": 50, // (Petabyte)
		"EB": 60, // (Exabyte)
	}[sizeType]
}

// MaxFileSize refer to readme for deep explanation
func MaxFileSize() int {
	maxFileSize, err := strconv.Atoi(os.Getenv("MAX_FILE_SIZE"))
	if err != nil || maxFileSize <= 0 {
		panic("MAX_FILE_SIZE is required")
	}

	sizeType := sizeTable(strings.ToUpper(os.Getenv("MAX_FILE_SIZE_TYPE")))
	return maxFileSize << sizeType
}

func MaxFileSizeStr() string {
	return fmt.Sprintf("%s %s", os.Getenv("MAX_FILE_SIZE"), os.Getenv("MAX_FILE_SIZE_TYPE"))
}
