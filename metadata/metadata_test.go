package metadata

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/utils"
)

func TestConfigExists(t *testing.T) {
	mockGitService := git.NewMockGitService(t)
	mockFileService := utils.NewMockFileService(t)

	mockGitService.On("DirectoryPath").Return("/tmp")
	mockCall := mockFileService.On("FileExists", "/tmp/.stacks_config").Return(true)

	metadataService := NewMetadataService(mockGitService, mockFileService)
	assert.True(t, metadataService.ConfigExists(), "ConfigExists should return true")

	mockGitService.AssertExpectations(t)
	mockFileService.AssertExpectations(t)

	mockCall.Unset()

	mockFileService.On("FileExists", "/tmp/.stacks_config").Return(false)
	assert.False(t, metadataService.ConfigExists(), "ConfigExists should return false")

	mockFileService.AssertExpectations(t)
	mockGitService.AssertExpectations(t)
}

// write a test for GetConfig
func TestGetConfig(t *testing.T) {
	mockGitService := git.NewMockGitService(t)
	mockFileService := utils.NewMockFileService(t)

	mockGitService.On("DirectoryPath").Return("/tmp")
	mockCall := mockFileService.On("FileExists", "/tmp/.stacks_config").Return(true)
	mockReadFileToByteArray := mockFileService.On("ReadFileToByteArray", "/tmp/.stacks_config").Return([]byte(`{"Trunk":"master"}`), nil)

	metadataService := NewMetadataService(mockGitService, mockFileService)
	config := metadataService.GetConfig()
	assert.Equal(t, "master", config.Trunk, "GetConfig should return the correct config")

	mockGitService.AssertExpectations(t)
	mockFileService.AssertExpectations(t)

	mockCall.Unset()
	mockReadFileToByteArray.Unset()

	mockFileService.On("FileExists", "/tmp/.stacks_config").Return(false)
	mockFileService.On("WriteByteArrayToFile", []byte(`{"Trunk":"develop"}`), "/tmp/.stacks_config").Return(nil)
	mockGitService.On("GetCurrentBranch").Return("develop")

	config = metadataService.GetConfig()
	assert.Equal(t, "develop", config.Trunk, "GetConfig should return default config based on current branch")

	mockFileService.AssertExpectations(t)
	mockGitService.AssertExpectations(t)
}

func TestUpdateConfig(t *testing.T) {
	mockGitService := git.NewMockGitService(t)
	mockFileService := utils.NewMockFileService(t)

	mockGitService.On("DirectoryPath").Return("/tmp")
	mockFileService.On("WriteByteArrayToFile", []byte(`{"Trunk":"develop"}`), "/tmp/.stacks_config").Return(nil)

	metadataService := NewMetadataService(mockGitService, mockFileService)
	metadataService.UpdateConfig(Config{Trunk: "develop"})

	mockGitService.AssertExpectations(t)
	mockFileService.AssertExpectations(t)
}

func TestCacheExists(t *testing.T) {
	mockGitService := git.NewMockGitService(t)
	mockFileService := utils.NewMockFileService(t)

	mockGitService.On("DirectoryPath").Return("/tmp")
	mockCall := mockFileService.On("FileExists", "/tmp/.stacks_cache").Return(true)

	metadataService := NewMetadataService(mockGitService, mockFileService)
	assert.True(t, metadataService.CacheExists(), "CacheExists should return true")

	mockGitService.AssertExpectations(t)
	mockFileService.AssertExpectations(t)

	mockCall.Unset()

	mockFileService.On("FileExists", "/tmp/.stacks_cache").Return(false)
	assert.False(t, metadataService.CacheExists(), "CacheExists should return false")

	mockFileService.AssertExpectations(t)
	mockGitService.AssertExpectations(t)
}

func TestGetCache(t *testing.T) {
	mockGitService := git.NewMockGitService(t)
	mockFileService := utils.NewMockFileService(t)

	mockGitService.On("DirectoryPath").Return("/tmp")
	mockCall := mockFileService.On("FileExists", "/tmp/.stacks_cache").Return(true)

	branch := Branch{Name: "master"}
	mockReadFileToByteArray := mockFileService.On("ReadFileToByteArray", "/tmp/.stacks_cache").Return([]byte(`{"Branches":[{"Name": "master"}]}`), nil)

	metadataService := NewMetadataService(mockGitService, mockFileService)
	cache := metadataService.GetCache()
	assert.Equal(t, []Branch{branch}, cache.Branches, "GetCache should return the correct cache")

	mockGitService.AssertExpectations(t)
	mockFileService.AssertExpectations(t)

	mockCall.Unset()
	mockReadFileToByteArray.Unset()

	mockFileService.On("FileExists", "/tmp/.stacks_cache").Return(false)
	branch = Branch{Name: "develop"}
	cache = Cache{Branches: []Branch{branch}}
	b, _ := json.Marshal(cache)
	mockFileService.On("WriteByteArrayToFile", b, "/tmp/.stacks_cache").Return(nil)
	mockGitService.On("GetCurrentBranch").Return("develop")

	cache = metadataService.GetCache()
	assert.Equal(t, []Branch{branch}, cache.Branches, "GetCache should return default cache based on current branch")

	mockFileService.AssertExpectations(t)
	mockGitService.AssertExpectations(t)
}

func TestUpdateCache(t *testing.T) {
	mockGitService := git.NewMockGitService(t)
	mockFileService := utils.NewMockFileService(t)

	mockGitService.On("DirectoryPath").Return("/tmp")
	branch := Branch{Name: "develop"}
	cache := Cache{Branches: []Branch{branch}}
	b, _ := json.Marshal(cache)
	mockFileService.On("WriteByteArrayToFile", b, "/tmp/.stacks_cache").Return(nil)

	metadataService := NewMetadataService(mockGitService, mockFileService)
	metadataService.UpdateCache(cache)

	mockGitService.AssertExpectations(t)
	mockFileService.AssertExpectations(t)
}

func TestContinueInfoExists(t *testing.T) {
	mockGitService := git.NewMockGitService(t)
	mockFileService := utils.NewMockFileService(t)

	mockGitService.On("DirectoryPath").Return("/tmp")
	mockCall := mockFileService.On("FileExists", "/tmp/.stacks_continue").Return(true)

	metadataService := NewMetadataService(mockGitService, mockFileService)
	assert.True(t, metadataService.ContinueInfoExists(), "ContinueInfoExists should return true")

	mockGitService.AssertExpectations(t)
	mockFileService.AssertExpectations(t)

	mockCall.Unset()

	mockFileService.On("FileExists", "/tmp/.stacks_continue").Return(false)
	assert.False(t, metadataService.ContinueInfoExists(), "ContinueInfoExists should return false")

	mockFileService.AssertExpectations(t)
	mockGitService.AssertExpectations(t)
}

func TestGetContinueInfo(t *testing.T) {
	mockGitService := git.NewMockGitService(t)
	mockFileService := utils.NewMockFileService(t)

	mockGitService.On("DirectoryPath").Return("/tmp")

	mockContinueInfo := ContinueInfo{Branches: []string{"master"}, ContinueBranch: "develop"}
	b, _ := json.Marshal(mockContinueInfo)
	mockFileService.On("ReadFileToByteArray", "/tmp/.stacks_continue").Return(b, nil)

	metadataService := NewMetadataService(mockGitService, mockFileService)
	continueInfo := metadataService.GetContinueInfo()
	assert.Equal(t, []string{"master"}, continueInfo.Branches, "GetContinueInfo should return the correct branches")
	assert.Equal(t, "develop", continueInfo.ContinueBranch, "GetContinueInfo should return the correct continue branch")

	mockGitService.AssertExpectations(t)
	mockFileService.AssertExpectations(t)
}

func TestStoreContinueInfo(t *testing.T) {
	mockGitService := git.NewMockGitService(t)
	mockFileService := utils.NewMockFileService(t)

	mockGitService.On("DirectoryPath").Return("/tmp")

	mockContinueInfo := ContinueInfo{Branches: []string{"master"}, ContinueBranch: "develop"}
	b, _ := json.Marshal(mockContinueInfo)
	mockFileService.On("WriteByteArrayToFile", b, "/tmp/.stacks_continue").Return(nil)

	metadataService := NewMetadataService(mockGitService, mockFileService)
	metadataService.StoreContinueInfo("develop", []string{"master"})

	mockGitService.AssertExpectations(t)
	mockFileService.AssertExpectations(t)
}

func TestRemoveContinueInfo(t *testing.T) {
	mockGitService := git.NewMockGitService(t)
	mockFileService := utils.NewMockFileService(t)

	mockGitService.On("DirectoryPath").Return("/tmp")

	mockFileService.On("RemoveFile", "/tmp/.stacks_continue").Return(nil)

	metadataService := NewMetadataService(mockGitService, mockFileService)
	metadataService.RemoveContinueInfo()

	mockGitService.AssertExpectations(t)
	mockFileService.AssertExpectations(t)
}
