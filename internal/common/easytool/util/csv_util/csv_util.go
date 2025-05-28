package csv_util

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"github.com/saintfish/chardet"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"os"
)

// ReadCSVByReader read csv by io.reader
func ReadCSVByReader(reader io.Reader) ([]map[string]any, error) {
	content, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	charset := detectCharset(content)
	var csvReader *csv.Reader
	switch charset.Language {
	case "zh":
		transformReader := transform.NewReader(bytes.NewReader(content), simplifiedchinese.GBK.NewDecoder())
		bufReader := bufio.NewReader(transformReader)
		csvReader = csv.NewReader(bufReader)
		break
	default:
		csvReader = csv.NewReader(bytes.NewReader(content))
	}
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	sliceMap := toSliceMap(records)
	return sliceMap, nil
}

// ReadCSVByFile read file
func ReadCSVByFile(file *os.File) ([]map[string]any, error) {
	defer file.Close()
	return ReadCSVByReader(file)
}

// ReadCSVByPath read csv local file by path
func ReadCSVByPath(path string) ([]map[string]any, error) {
	open, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return ReadCSVByFile(open)
}

func detectCharset(content []byte) *chardet.Result {
	// 创建探测器并检测字符编码
	detector := chardet.NewTextDetector()
	result, _ := detector.DetectBest(content)
	return result
}

// toSliceMap convert two-dimension array to slice_util map
func toSliceMap(records [][]string) []map[string]any {

	//获得key
	keys := records[0]
	result := make([]map[string]any, len(records)-1)
	for i := 0; i < len(records)-1; i++ {
		record := records[i+1]
		item := make(map[string]any)
		for k := 0; k < len(keys); k++ {
			key := keys[k]
			item[key] = record[k]
		}
		result[i] = item
	}
	return result
}
