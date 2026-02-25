// Copyright 2026 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package repo

import (
	"testing"

	issues_model "code.gitea.io/gitea/models/issues"
	project_model "code.gitea.io/gitea/models/project"
	"code.gitea.io/gitea/models/unittest"
	"code.gitea.io/gitea/services/contexttest"

	"github.com/stretchr/testify/assert"
)

func TestRetrieveProjectsDataForIssueWriter_SelectedProjectID(t *testing.T) {
	unittest.PrepareTestEnv(t)
	ctx, _ := contexttest.MockContext(t, "user2/repo1/issues/1")
	contexttest.LoadUser(t, ctx, 2)
	contexttest.LoadRepo(t, ctx, 1)

	for _, tc := range []struct {
		name               string
		issue              *issues_model.Issue
		expectedSelectedID string
	}{
		{
			name:               "nil issue",
			issue:              nil,
			expectedSelectedID: "",
		},
		{
			name:               "issue with no projects",
			issue:              &issues_model.Issue{Projects: nil},
			expectedSelectedID: "",
		},
		{
			name:               "issue with one project",
			issue:              &issues_model.Issue{Projects: []*project_model.Project{{ID: 1}}},
			expectedSelectedID: "1",
		},
		{
			name: "issue with multiple projects",
			issue: &issues_model.Issue{Projects: []*project_model.Project{
				{ID: 1},
				{ID: 2},
				{ID: 3},
			}},
			expectedSelectedID: "1,2,3",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			d := &IssuePageMetaData{
				Repository:   ctx.Repo.Repository,
				Issue:        tc.issue,
				ProjectsData: &issueSidebarProjectsData{},
			}
			d.retrieveProjectsDataForIssueWriter(ctx)
			assert.Equal(t, tc.expectedSelectedID, d.ProjectsData.SelectedProjectID)
		})
	}
}
