package stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/metadata"
	"github.com/underwoo16/git-stacks/utils"
)

func TestGetNameFromRef(t *testing.T) {
	mockGitService := git.NewMockGitService(t)
	mockFileService := utils.NewMockFileService(t)
	mockMetadataService := metadata.NewMockMetadataService(t)

	stackService := NewStackService(mockGitService, mockMetadataService, mockFileService)
	assert.Equal(t, "feature/branch", stackService.GetNameFromRef("refs/heads/feature/branch"))

	assert.Equal(t, "feature_one", stackService.GetNameFromRef("refs/stacks/feature_one"))
}

func TestStackExists(t *testing.T) {
	mockGitService := git.NewMockGitService(t)
	mockFileService := utils.NewMockFileService(t)
	mockMetadataService := metadata.NewMockMetadataService(t)

	mockGitService.On("DirectoryPath").Return("/tmp")
	mockFileService.On("FileExists", "/tmp/refs/stacks/feature/branch").Return(true)

	stackService := NewStackService(mockGitService, mockMetadataService, mockFileService)
	assert.True(t, stackService.StackExists("feature/branch"))

	mockGitService.AssertExpectations(t)
	mockFileService.AssertExpectations(t)

	mockGitService.On("DirectoryPath").Return("/tmp")
	mockFileService.On("FileExists", "/tmp/refs/stacks/feature/branch/2").Return(false)

	assert.False(t, stackService.StackExists("feature/branch/2"))

	mockGitService.AssertExpectations(t)
	mockFileService.AssertExpectations(t)
}

func TestGetCurrentStackNode(t *testing.T) {
	mockGitService := git.NewMockGitService(t)
	mockFileService := utils.NewMockFileService(t)
	mockMetadataService := metadata.NewMockMetadataService(t)

	mockMetadataService.On("CacheExists").Return(true)
	mockMetadataService.On("GetConfig").Return(metadata.Config{
		Trunk: "develop",
	})
	mockMetadataService.On("GetCache").Return(metadata.Cache{
		Branches: []metadata.Branch{
			{
				Name:                 "develop",
				BranchRevision:       "1234567890",
				ParentBranchName:     "",
				ParentBranchRevision: "",
				Children:             []string{"feature/branch"},
			},
			{
				Name:                 "feature/branch",
				BranchRevision:       "1112131415",
				ParentBranchName:     "develop",
				ParentBranchRevision: "0987654321",
				Children:             []string{"feature/branch/2"},
			},
			{
				Name:                 "feature/branch/2",
				BranchRevision:       "1617181920",
				ParentBranchName:     "feature/branch",
				ParentBranchRevision: "0987654321",
				Children:             []string{},
			},
		},
	})

	mockGitService.On("GetCurrentBranch").Return("feature/branch")

	stackService := NewStackService(mockGitService, mockMetadataService, mockFileService)
	stackNode := stackService.GetCurrentStackNode()
	assert.Equal(t, "feature/branch", stackNode.Name)
	assert.Equal(t, "1112131415", stackNode.RefSha)
	assert.Equal(t, "feature/branch/2", stackNode.Children[0].Name)
	assert.Equal(t, "develop", stackNode.Parent.Name)
}

func TestNeedsSync(t *testing.T) {
	mockGitService := git.NewMockGitService(t)
	mockFileService := utils.NewMockFileService(t)
	mockMetadataService := metadata.NewMockMetadataService(t)

	mockGitService.On("RevParse", "develop").Return("1234567890")

	stackService := NewStackService(mockGitService, mockMetadataService, mockFileService)

	stackNode := StackNode{
		Name:         "feature/branch",
		RefSha:       "1234567890",
		ParentBranch: "develop",
		ParentRefSha: "0987654321",
		Parent: &StackNode{
			Name:   "develop",
			RefSha: "0987654321",
		},
		Children: []*StackNode{},
	}

	assert.True(t, stackService.NeedsSync(&stackNode))

	mockGitService.AssertExpectations(t)
}
