package internal

import (
	"testing"

	"code.gitea.io/gitea/modules/json"
)

// 1) Old documents: singular only -> populate ProjectIDs with single value
func TestIndexerDataUnmarshal_OldSingular(t *testing.T) {
	var d IndexerData
	src := `{"id":1,"project_id":42}`
	if err := json.Unmarshal([]byte(src), &d); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if len(d.ProjectIDs) != 1 || d.ProjectIDs[0] != 42 {
		t.Fatalf("expected ProjectIDs=[42], got %#v", d.ProjectIDs)
	}
}

// 2) New documents: plural only -> use project_ids
func TestIndexerDataUnmarshal_NewPlural(t *testing.T) {
	var d IndexerData
	src := `{"id":1,"project_ids":[7,8]}`
	if err := json.Unmarshal([]byte(src), &d); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if len(d.ProjectIDs) != 2 || d.ProjectIDs[0] != 7 || d.ProjectIDs[1] != 8 {
		t.Fatalf("expected ProjectIDs=[7,8], got %#v", d.ProjectIDs)
	}
}

// 3) Both present: plural non-empty should win
func TestIndexerDataUnmarshal_BothPluralWins(t *testing.T) {
	var d IndexerData
	src := `{"id":1,"project_id":42,"project_ids":[1,2]}`
	if err := json.Unmarshal([]byte(src), &d); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if len(d.ProjectIDs) != 2 || d.ProjectIDs[0] != 1 || d.ProjectIDs[1] != 2 {
		t.Fatalf("expected ProjectIDs=[1,2], got %#v", d.ProjectIDs)
	}
}

// 4) Both present: plural present but empty -> fallback to singular (edge case)
func TestIndexerDataUnmarshal_EmptyPluralFallbacksToSingular(t *testing.T) {
	var d IndexerData
	src := `{"id":1,"project_id":42,"project_ids":[]}`
	if err := json.Unmarshal([]byte(src), &d); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if len(d.ProjectIDs) != 1 || d.ProjectIDs[0] != 42 {
		t.Fatalf("expected ProjectIDs=[42], got %#v", d.ProjectIDs)
	}
}
