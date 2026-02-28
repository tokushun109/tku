package request

import (
	"strings"
	"testing"
)

func TestParseProductCSV(t *testing.T) {
	t.Run("正規ヘッダを渡したときCSV行のパースに成功する", func(t *testing.T) {
		rows, err := ParseProductCSV(strings.NewReader("ID,Name,Price,CategoryName,TargetName\n1,product,1200,cat,target\n"))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(rows) != 1 {
			t.Fatalf("expected 1 row, got %d", len(rows))
		}
		if rows[0].ID != 1 {
			t.Fatalf("unexpected id: %d", rows[0].ID)
		}
	})

	t.Run("lowerとcamelの混在ヘッダを渡したときCSV行のパースに成功する", func(t *testing.T) {
		rows, err := ParseProductCSV(strings.NewReader("id,name,price,categoryName,targetname\n2,product2,1500,cat2,target2\n"))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(rows) != 1 {
			t.Fatalf("expected 1 row, got %d", len(rows))
		}
		if rows[0].TargetName != "target2" {
			t.Fatalf("unexpected target name: %s", rows[0].TargetName)
		}
	})

	t.Run("必須ヘッダが欠けているとき失敗する", func(t *testing.T) {
		_, err := ParseProductCSV(strings.NewReader("id,name,price,categoryName\n2,product2,1500,cat2\n"))
		if err == nil {
			t.Fatalf("expected error")
		}
	})

	t.Run("categoryNameとtargetNameが空欄でもパースに成功する", func(t *testing.T) {
		rows, err := ParseProductCSV(strings.NewReader("id,name,price,categoryName,targetName\n1,product,1200,,\n"))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(rows) != 1 {
			t.Fatalf("expected 1 row, got %d", len(rows))
		}
		if rows[0].CategoryName != "" {
			t.Fatalf("unexpected category name: %s", rows[0].CategoryName)
		}
		if rows[0].TargetName != "" {
			t.Fatalf("unexpected target name: %s", rows[0].TargetName)
		}
	})
}
