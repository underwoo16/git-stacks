package stacks

import (
	"encoding/json"
	"fmt"

	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/queue"
	"github.com/underwoo16/git-stacks/utils"
)

type Config struct {
	Trunk string
}

type Cache struct {
	Branches []Branch
}

type Branch struct {
	Name                 string
	BranchRevision       string
	ParentBranchName     string
	ParentBranchRevision string
	Children             []string
}

type ContinueInfo struct {
	ContinueBranch string
	OriginalBranch string
	Branches       []string
}

type MetadataService struct {
	gitService *git.GitService
}

func NewMetadataService(gitService *git.GitService) *MetadataService {
	return &MetadataService{gitService: gitService}
}

func (m *MetadataService) ConfigExists() bool {
	gitPath := fmt.Sprintf("%s/.stacks_config", m.gitService.DirectoryPath())
	return utils.FileExists(gitPath)
}

func (m *MetadataService) UpdateConfig(config Config) {
	gitService := git.NewGitService()
	b, err := json.Marshal(config)
	utils.CheckError(err)

	configPath := fmt.Sprintf("%s/.stacks_config", gitService.DirectoryPath())
	utils.WriteByteArrayToFile(b, configPath)
}

func (m *MetadataService) GetConfig() Config {
	gitService := git.NewGitService()
	if !m.ConfigExists() {
		currentBranch := gitService.GetCurrentBranch()
		config := Config{Trunk: currentBranch}
		m.UpdateConfig(config)
		return config
	}

	configPath := fmt.Sprintf("%s/.stacks_config", gitService.DirectoryPath())
	ba := utils.ReadFileToByteArray(configPath)

	var config Config
	utils.CheckError(json.Unmarshal(ba, &config))

	return config
}

func (m *MetadataService) CacheExists() bool {
	gitService := git.NewGitService()
	cachePath := fmt.Sprintf("%s/.stacks_cache", gitService.DirectoryPath())
	return utils.FileExists(cachePath)
}

func (m *MetadataService) GetCache() Cache {
	if !m.CacheExists() {
		// TODO: add trunk here at minimum
		// could also build graph and populate it fully
		cache := Cache{Branches: []Branch{}}
		m.UpdateCache(cache)
		return cache
	}

	cachePath := fmt.Sprintf("%s/.stacks_cache", m.gitService.DirectoryPath())
	ba := utils.ReadFileToByteArray(cachePath)
	var cache Cache
	utils.CheckError(json.Unmarshal(ba, &cache))

	return cache
}

func (m *MetadataService) UpdateCache(cache Cache) {
	b, err := json.Marshal(cache)
	utils.CheckError(err)

	cachePath := fmt.Sprintf("%s/.stacks_cache", m.gitService.DirectoryPath())
	utils.WriteByteArrayToFile(b, cachePath)
}

func (m *MetadataService) StoreContinueInfo(branch string, queue *queue.Queue) {
	var branches []string
	for !queue.IsEmpty() {
		stack := queue.Pop().(*StackNode)
		branches = append(branches, stack.Name)
	}

	continueInfo := ContinueInfo{ContinueBranch: branch, Branches: branches}
	b, err := json.Marshal(continueInfo)
	utils.CheckError(err)

	gitService := git.NewGitService()
	continePath := fmt.Sprintf("%s/.stacks_continue", gitService.DirectoryPath())
	utils.WriteByteArrayToFile(b, continePath)
}

func (m *MetadataService) ContinueInfoExists() bool {
	gitService := git.NewGitService()
	continePath := fmt.Sprintf("%s/.stacks_continue", gitService.DirectoryPath())
	return utils.FileExists(continePath)
}

func (m *MetadataService) GetContinueInfo() ContinueInfo {
	gitService := git.NewGitService()
	continePath := fmt.Sprintf("%s/.stacks_continue", gitService.DirectoryPath())
	ba := utils.ReadFileToByteArray(continePath)
	var continueInfo ContinueInfo
	utils.CheckError(json.Unmarshal(ba, &continueInfo))

	return continueInfo
}

func (m *MetadataService) RemoveContinueInfo() {
	gitService := git.NewGitService()
	continePath := fmt.Sprintf("%s/.stacks_continue", gitService.DirectoryPath())
	utils.RemoveFile(continePath)
}
