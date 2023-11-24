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

func ConfigExists() bool {
	return utils.FileExists(".git/.stacks_config")
}

func UpdateConfig(config Config) {
	b, err := json.Marshal(config)
	utils.CheckError(err)

	configPath := fmt.Sprintf("%s/.stacks_config", git.DirectoryPath())
	utils.WriteByteArrayToFile(b, configPath)
}

func GetConfig() Config {
	if !ConfigExists() {
		currentBranch := git.GetCurrentBranch()
		config := Config{Trunk: currentBranch}
		UpdateConfig(config)
		return config
	}

	configPath := fmt.Sprintf("%s/.stacks_config", git.DirectoryPath())
	ba := utils.ReadFileToByteArray(configPath)

	var config Config
	utils.CheckError(json.Unmarshal(ba, &config))

	return config
}

func CacheExists() bool {
	cachePath := fmt.Sprintf("%s/.stacks_cache", git.DirectoryPath())
	return utils.FileExists(cachePath)
}

func GetCache() Cache {
	if !CacheExists() {
		// TODO: add trunk here at minimum
		// could also build graph and populate it fully
		cache := Cache{Branches: []Branch{}}
		UpdateCache(cache)
		return cache
	}

	ba := utils.ReadFileToByteArray(".git/.stacks_cache")
	var cache Cache
	utils.CheckError(json.Unmarshal(ba, &cache))

	return cache
}

func UpdateCache(cache Cache) {
	b, err := json.Marshal(cache)
	utils.CheckError(err)

	utils.WriteByteArrayToFile(b, ".git/.stacks_cache")
}

func StoreContinueInfo(branch string, queue *queue.Queue) {
	var branches []string
	for !queue.IsEmpty() {
		stack := queue.Pop().(*StackNode)
		branches = append(branches, stack.Name)
	}

	continueInfo := ContinueInfo{ContinueBranch: branch, Branches: branches}
	b, err := json.Marshal(continueInfo)
	utils.CheckError(err)

	continePath := fmt.Sprintf("%s/.stacks_continue", git.DirectoryPath())
	utils.WriteByteArrayToFile(b, continePath)
}

func ContinueInfoExists() bool {
	continePath := fmt.Sprintf("%s/.stacks_continue", git.DirectoryPath())
	return utils.FileExists(continePath)
}

func GetContinueInfo() ContinueInfo {
	continePath := fmt.Sprintf("%s/.stacks_continue", git.DirectoryPath())
	ba := utils.ReadFileToByteArray(continePath)
	var continueInfo ContinueInfo
	utils.CheckError(json.Unmarshal(ba, &continueInfo))

	return continueInfo
}

func RemoveContinueInfo() {
	continePath := fmt.Sprintf("%s/.stacks_continue", git.DirectoryPath())
	utils.RemoveFile(continePath)
}
