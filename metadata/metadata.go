package metadata

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/underwoo16/git-stacks/git"
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

type MetadataService interface {
	ConfigExists() bool
	UpdateConfig(config Config)
	GetConfig() Config
	CacheExists() bool
	GetCache() Cache
	UpdateCache(cache Cache)
	ContinueInfoExists() bool
	GetContinueInfo() ContinueInfo
	RemoveContinueInfo()
	StoreContinueInfo(branch string, branches []string)
}

type metadataService struct {
	gitService git.GitService
}

func NewMetadataService(gitService git.GitService) *metadataService {
	return &metadataService{gitService: gitService}
}

func (m *metadataService) ConfigExists() bool {
	gitPath := fmt.Sprintf("%s/.stacks_config", m.gitService.DirectoryPath())
	return utils.FileExists(gitPath)
}

func (m *metadataService) UpdateConfig(config Config) {
	b, err := json.Marshal(config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	configPath := fmt.Sprintf("%s/.stacks_config", m.gitService.DirectoryPath())
	utils.WriteByteArrayToFile(b, configPath)
}

func (m *metadataService) GetConfig() Config {
	if !m.ConfigExists() {
		currentBranch := m.gitService.GetCurrentBranch()
		config := Config{Trunk: currentBranch}
		m.UpdateConfig(config)
		return config
	}

	configPath := fmt.Sprintf("%s/.stacks_config", m.gitService.DirectoryPath())
	ba := utils.ReadFileToByteArray(configPath)

	var config Config
	err := json.Unmarshal(ba, &config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return config
}

func (m *metadataService) CacheExists() bool {
	cachePath := fmt.Sprintf("%s/.stacks_cache", m.gitService.DirectoryPath())
	return utils.FileExists(cachePath)
}

func (m *metadataService) GetCache() Cache {
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
	err := json.Unmarshal(ba, &cache)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return cache
}

func (m *metadataService) UpdateCache(cache Cache) {
	b, err := json.Marshal(cache)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cachePath := fmt.Sprintf("%s/.stacks_cache", m.gitService.DirectoryPath())
	utils.WriteByteArrayToFile(b, cachePath)
}

func (m *metadataService) StoreContinueInfo(branch string, branches []string) {
	continueInfo := ContinueInfo{ContinueBranch: branch, Branches: branches}
	b, err := json.Marshal(continueInfo)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	continePath := fmt.Sprintf("%s/.stacks_continue", m.gitService.DirectoryPath())
	utils.WriteByteArrayToFile(b, continePath)
}

func (m *metadataService) ContinueInfoExists() bool {
	continePath := fmt.Sprintf("%s/.stacks_continue", m.gitService.DirectoryPath())
	return utils.FileExists(continePath)
}

func (m *metadataService) GetContinueInfo() ContinueInfo {
	continePath := fmt.Sprintf("%s/.stacks_continue", m.gitService.DirectoryPath())
	ba := utils.ReadFileToByteArray(continePath)
	var continueInfo ContinueInfo
	err := json.Unmarshal(ba, &continueInfo)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return continueInfo
}

func (m *metadataService) RemoveContinueInfo() {
	continePath := fmt.Sprintf("%s/.stacks_continue", m.gitService.DirectoryPath())
	utils.RemoveFile(continePath)
}
