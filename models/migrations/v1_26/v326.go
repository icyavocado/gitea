// Copyright 2026 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package v1_26

import (
	"xorm.io/xorm"
)

// ForceIssueIndexerRebuildForProjectIDs marks the issue indexer for rebuild
// to migrate from project_id (singular) to project_ids (plural) in indexed documents
func ForceIssueIndexerRebuildForProjectIDs(x *xorm.Engine) error {
	// Delete the issue_index table to force a rebuild of the issue search index.
	// This is necessary because we changed the indexer schema from:
	//   ProjectID int64 (singular)
	// to:
	//   ProjectIDs []int64 (plural)
	//
	// By deleting the table, InitIssueIndexer will detect the missing index
	// and trigger a full reindex using PopulateIssueIndexer().
	// The new index will be populated from the project_issue table which
	// already supports multiple projects per issue.

	type IssueIndex struct {
		RepoID   int64 `xorm:"pk"`
		IsInit   bool  `xorm:"NOT NULL DEFAULT false"`
		IndexVer int64 `xorm:"NOT NULL DEFAULT 0"`
	}

	if err := x.Sync(new(IssueIndex)); err != nil {
		return err
	}

	// Drop all records to force reindex
	_, err := x.Exec("DELETE FROM issue_index")
	return err
}
