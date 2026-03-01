package request

import (
	"strings"
	"testing"
)

func TestParseProductCSV(t *testing.T) {
	t.Run("正規ヘッダを渡したときCSV行のパースに成功する", func(t *testing.T) {
		rows, err := ParseProductCSV(strings.NewReader("UUID,Name,Price,CategoryName,TargetName\n11111111-1111-4111-8111-111111111111,product,1200,cat,target\n"))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(rows) != 1 {
			t.Fatalf("expected 1 row, got %d", len(rows))
		}
		if rows[0].UUID != "11111111-1111-4111-8111-111111111111" {
			t.Fatalf("unexpected uuid: %s", rows[0].UUID)
		}
	})

	t.Run("lowerとcamelの混在ヘッダを渡したときCSV行のパースに成功する", func(t *testing.T) {
		rows, err := ParseProductCSV(strings.NewReader("uuid,name,price,categoryName,targetname\n22222222-2222-4222-8222-222222222222,product2,1500,cat2,target2\n"))
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
		_, err := ParseProductCSV(strings.NewReader("uuid,name,price,categoryName\n22222222-2222-4222-8222-222222222222,product2,1500,cat2\n"))
		if err == nil {
			t.Fatalf("expected error")
		}
	})

	t.Run("categoryNameとtargetNameが空欄でもパースに成功する", func(t *testing.T) {
		rows, err := ParseProductCSV(strings.NewReader("uuid,name,price,categoryName,targetName\n11111111-1111-4111-8111-111111111111,product,1200,,\n"))
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
