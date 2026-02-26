package request

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

var (
	errInvalidCSVHeader = errors.New("invalid csv header")
	errInvalidCSVRecord = errors.New("invalid csv record")
)

type ProductCSVRow struct {
	ID           uint
	Name         string
	Price        int
	CategoryName string
	TargetName   string
}

func ParseProductCSV(r io.Reader) ([]ProductCSVRow, error) {
	reader := csv.NewReader(r)

	header, err := reader.Read()
	if err != nil {
		if errors.Is(err, io.EOF) {
			return nil, errInvalidCSVHeader
		}
		return nil, fmt.Errorf("%w: %v", errInvalidCSVHeader, err)
	}

	columnMap, err := buildProductCSVColumnMap(header)
	if err != nil {
		return nil, err
	}

	rows := make([]ProductCSVRow, 0)
	for lineNo := 2; ; lineNo++ {
		record, err := reader.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("%w: line=%d: %v", errInvalidCSVRecord, lineNo, err)
		}
		if isEmptyCSVRecord(record) {
			continue
		}

		row, err := parseProductCSVRecord(record, columnMap, lineNo)
		if err != nil {
			return nil, err
		}
		rows = append(rows, row)
	}

	return rows, nil
}

func buildProductCSVColumnMap(header []string) (map[string]int, error) {
	columnMap := make(map[string]int, len(header))
	for idx, column := range header {
		normalized := normalizeCSVHeader(column)
		if normalized == "" {
			continue
		}
		if _, exists := columnMap[normalized]; exists {
			return nil, fmt.Errorf("%w: duplicated column=%s", errInvalidCSVHeader, column)
		}
		columnMap[normalized] = idx
	}

	required := []string{"id", "name", "price", "categoryname", "targetname"}
	for _, key := range required {
		if _, ok := columnMap[key]; !ok {
			return nil, fmt.Errorf("%w: missing column=%s", errInvalidCSVHeader, key)
		}
	}

	return columnMap, nil
}

func parseProductCSVRecord(record []string, columnMap map[string]int, lineNo int) (ProductCSVRow, error) {
	idText := getCSVColumn(record, columnMap["id"])
	name := getCSVColumn(record, columnMap["name"])
	priceText := getCSVColumn(record, columnMap["price"])
	categoryName := getCSVColumn(record, columnMap["categoryname"])
	targetName := getCSVColumn(record, columnMap["targetname"])

	if name == "" || categoryName == "" || targetName == "" {
		return ProductCSVRow{}, fmt.Errorf("%w: line=%d", errInvalidCSVRecord, lineNo)
	}

	id, err := strconv.ParseUint(idText, 10, 64)
	if err != nil || id == 0 {
		return ProductCSVRow{}, fmt.Errorf("%w: line=%d", errInvalidCSVRecord, lineNo)
	}

	price, err := strconv.Atoi(priceText)
	if err != nil {
		return ProductCSVRow{}, fmt.Errorf("%w: line=%d", errInvalidCSVRecord, lineNo)
	}

	return ProductCSVRow{
		ID:           uint(id),
		Name:         name,
		Price:        price,
		CategoryName: categoryName,
		TargetName:   targetName,
	}, nil
}

func normalizeCSVHeader(v string) string {
	normalized := strings.TrimSpace(strings.TrimPrefix(v, "\ufeff"))
	normalized = strings.ToLower(normalized)
	normalized = strings.NewReplacer("_", "", " ", "").Replace(normalized)
	return normalized
}

func getCSVColumn(record []string, idx int) string {
	if idx < 0 || idx >= len(record) {
		return ""
	}
	return strings.TrimSpace(record[idx])
}

func isEmptyCSVRecord(record []string) bool {
	for _, value := range record {
		if strings.TrimSpace(value) != "" {
			return false
		}
	}
	return true
}
