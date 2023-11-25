package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/utils"
)

func TestConfigExists(t *testing.T) {
	mockGitService := git.NewMockGitService(t)
	mockFileService := utils.NewMockFileService(t)

	mockGitService.On("DirectoryPath").Return("/tmp")
	mockFileService.On("FileExists", "/tmp/.stacks_config").Return(true)

	metadataService := NewMetadataService(mockGitService, mockFileService)
	assert.True(t, metadataService.ConfigExists(), "ConfigExists should return true")
}
