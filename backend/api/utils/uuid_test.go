package utils

import (
	"testing"
	"github.com/google/uuid"
)

func TestGenerateUUID(t *testing.T) {
	// UUIDが正常に生成されることをテスト
	uuidString, err := GenerateUUID()
	
	if err != nil {
		t.Errorf("GenerateUUID() returned error: %v", err)
	}
	
	if uuidString == "" {
		t.Error("GenerateUUID() returned empty string")
	}
	
	// 生成されたUUIDが有効な形式かチェック
	_, parseErr := uuid.Parse(uuidString)
	if parseErr != nil {
		t.Errorf("Generated UUID is not valid format: %s, error: %v", uuidString, parseErr)
	}
}

func TestGenerateUUIDUniqueness(t *testing.T) {
	// 複数回実行して異なるUUIDが生成されることをテスト
	uuid1, err1 := GenerateUUID()
	uuid2, err2 := GenerateUUID()
	
	if err1 != nil || err2 != nil {
		t.Errorf("GenerateUUID() returned errors: %v, %v", err1, err2)
	}
	
	if uuid1 == uuid2 {
		t.Errorf("GenerateUUID() generated same UUID twice: %s", uuid1)
	}
}