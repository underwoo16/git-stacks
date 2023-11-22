package stacks

import (
	"encoding/json"

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

func ConfigExists() bool {
	return utils.FileExists(".git/.stacks_config")
}

func UpdateConfig(config Config) {
	b, err := json.Marshal(config)
	utils.CheckError(err)

	// TODO: use dynamic .git path
	utils.WriteByteArrayToFile(b, ".git/.stacks_config")
}

func GetConfig() Config {
	if !ConfigExists() {
		currentBranch := git.GetCurrentBranch()
		config := Config{Trunk: currentBranch}
		UpdateConfig(config)
		return config
	}

	// TODO: use dynamic .git path
	ba := utils.ReadFileToByteArray(".git/.stacks_config")

	var config Config
	utils.CheckError(json.Unmarshal(ba, &config))

	return config
}

func CacheExists() bool {
	return utils.FileExists(".git/.stacks_cache")
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
