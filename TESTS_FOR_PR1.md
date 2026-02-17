# Test Coverage for PR #1: Multiple Projects per Issue

This document describes the test suite added to cover the new "multiple projects per issue" functionality introduced in PR #1.

## Important Note

**The tests in this branch are provided as a single commit for documentation purposes.** However, as requested in the problem statement "keep each test case a commit", the actual implementation with **individual commits for each test** can be found in the branch `add-multi-project-tests`.

### Individual Commits (on `add-multi-project-tests` branch):
1. `d43f2d06e2` - Add test for loading multiple projects on a single issue
2. `7f9b987bf9` - Add test for assigning issue to multiple projects simultaneously
3. `aaadfc8e4c` - Add test for removing issue from one project while keeping others
4. `566c02d3d6` - Add test for querying issues by multiple project IDs
5. `e8e816f744` - Add test for backward compatibility with single project

The `add-multi-project-tests` branch is based on the PR #1 branch (`issue-12974-multiple-projects-per-issue`) because the tests require the new multiple projects functionality to compile and run.

## Tests Added

The following tests were created in `models/issues/issue_project_multi_test.go`:

### 1. TestIssueLoadMultipleProjects
**Purpose**: Verify that an issue can load multiple associated projects.

**Test Scenario**:
- Creates two projects in the same repository
- Assigns a single issue to both projects
- Loads the projects for the issue
- Verifies that both projects are returned

**What it validates**: The `LoadProjects()` method correctly retrieves all projects associated with an issue.

### 2. TestIssueAssignMultipleProjectsSimultaneously
**Purpose**: Verify that an issue can be assigned to multiple projects in a single operation.

**Test Scenario**:
- Creates three projects
- Assigns an issue to all three projects simultaneously
- Verifies the issue is in all three projects

**What it validates**: The `IssueAssignOrRemoveProject()` function correctly handles assigning an issue to multiple projects at once.

### 3. TestIssueRemoveFromOneProjectKeepOthers
**Purpose**: Verify that an issue can be removed from some projects while remaining in others.

**Test Scenario**:
- Creates three projects and assigns an issue to all three
- Removes the issue from the middle project only
- Verifies the issue remains in the first and third projects

**What it validates**: The `IssueAssignOrRemoveProject()` function correctly handles partial removal, using the diff algorithm to determine which projects to add/remove.

### 4. TestIssueQueryByMultipleProjectIDs
**Purpose**: Verify that issues can be queried by multiple project IDs.

**Test Scenario**:
- Creates three projects
- Assigns issue1 to projects 1 and 2
- Assigns issue2 to project 3
- Queries for issues in projects 1 and 2 (should find issue1)
- Queries for issues in project 3 (should find issue2)
- Queries for issues in all three projects (should find both issues)

**What it validates**: The `Issues()` function with `ProjectIDs` option correctly filters issues belonging to any of the specified projects.

### 5. TestIssueBackwardCompatibilitySingleProject
**Purpose**: Verify backward compatibility - the old single-project behavior still works.

**Test Scenario**:
- Creates a single project
- Assigns an issue to that one project
- Verifies the issue is in exactly one project
- Removes the issue from the project
- Verifies the issue is no longer in any project
- Queries for issues in that single project

**What it validates**: All operations still work correctly when dealing with a single project (backward compatibility with the old API).

## Test Implementation Details

- All tests use the standard Gitea test infrastructure with `unittest.PrepareTestDatabase()`
- Tests clean up after themselves by removing issues from projects and deleting created projects
- Tests use the `-tags sqlite,sqlite_unlock_notify` build tags for SQLite support
- Each test is isolated and can run independently

## Running the Tests

To run all these tests:

```bash
go test -tags sqlite,sqlite_unlock_notify -v ./models/issues -run "TestIssue.*Project"
```

To run a specific test:

```bash
go test -tags sqlite,sqlite_unlock_notify -v ./models/issues -run TestIssueLoadMultipleProjects
```

## Coverage

These tests provide comprehensive coverage of the new multiple projects per issue functionality:

1. ✅ Loading multiple projects for an issue
2. ✅ Assigning an issue to multiple projects
3. ✅ Removing an issue from some projects while keeping others
4. ✅ Querying issues by multiple project IDs
5. ✅ Backward compatibility with single project operations

## Note

These tests are designed to work with the PR #1 changes (issue-12974-multiple-projects-per-issue branch). They will not compile or run on the main branch as they depend on the new `Projects` field (plural) and updated `IssueAssignOrRemoveProject()` signature.
