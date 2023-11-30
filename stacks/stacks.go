package stacks

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/underwoo16/git-stacks/git"
	"github.com/underwoo16/git-stacks/metadata"
	"github.com/underwoo16/git-stacks/utils"
)

type StackNode struct {
	Name         string
	Parent       *StackNode
	Children     []*StackNode
	RefSha       string
	ParentBranch string
	ParentRefSha string
}

type StackService interface {
	GetGraph() *StackNode
	CacheGraphToDisk(trunk *StackNode)
	NeedsSync(stack *StackNode) bool
	GetCurrentStackNode() *StackNode
	CreateStack(name string, parent string, parentRefSha string)
}

type stackService struct {
	gitService      git.GitService
	metadataService metadata.MetadataService
	fileService     utils.FileService
}

func NewStackService(gitService git.GitService, metadataService metadata.MetadataService, fileService utils.FileService) *stackService {
	return &stackService{gitService: gitService, metadataService: metadataService, fileService: fileService}
}

func (s *stackService) readStack(ref string) *StackNode {
	out := s.gitService.Show(ref)
	items := strings.Fields(out)
	name := s.GetNameFromRef(ref)
	return &StackNode{
		Name:         name,
		RefSha:       s.gitService.RevParse(name),
		ParentBranch: s.GetNameFromRef(items[0]),
		ParentRefSha: items[1],
		Children:     []*StackNode{},
	}
}

func (s *stackService) GetNameFromRef(ref string) string {
	ref = strings.Replace(ref, "refs/heads/", "", -1)
	return strings.Replace(ref, "refs/stacks/", "", -1)
}

func (s *stackService) convertHeadToStack(ref string) string {
	return strings.Replace(ref, "refs/heads", "refs/stacks", -1)
}

func (s *stackService) getStacks() []*StackNode {
	var existingStacks []*StackNode

	stacksPath := fmt.Sprintf("%s/refs/stacks", s.gitService.DirectoryPath())
	if !s.fileService.FileExists(stacksPath) {
		return existingStacks
	}

	err := filepath.Walk(stacksPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if info.IsDir() {
			return nil
		}

		index := strings.Index(path, "refs/stacks/")
		ref := path[index:]
		stack := s.readStack(ref)
		existingStacks = append(existingStacks, stack)

		return nil
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return existingStacks
}

func (s *stackService) UpdateStack(stack *StackNode) {
	tempFileName := strings.ReplaceAll(stack.Name, "/", "-")
	tempFilePath := fmt.Sprintf("%s/temp-%s", s.gitService.DirectoryPath(), tempFileName)

	hashObject := fmt.Sprintf("%s\n%s", stack.ParentBranch, stack.ParentRefSha)
	s.fileService.WriteToFile(tempFilePath, hashObject)

	objectSha := s.gitService.CreateHashObject(tempFilePath)
	s.fileService.RemoveFile(tempFilePath)

	newRef := fmt.Sprintf("refs/stacks/%s", stack.Name)
	s.gitService.UpdateRef(newRef, objectSha)

	s.CacheGraphToDisk(stack)
}

func (s *stackService) CreateStack(name string, parentBranch string, parentRefSha string) {
	currentStack := s.GetCurrentStackNode()

	tempFileName := strings.ReplaceAll(name, "/", "-")
	tempFilePath := fmt.Sprintf("%s/temp-%s", s.gitService.DirectoryPath(), tempFileName)

	hashObject := fmt.Sprintf("%s\n%s", parentBranch, parentRefSha)
	s.fileService.WriteToFile(tempFilePath, hashObject)

	objectSha := s.gitService.CreateHashObject(tempFilePath)
	s.fileService.RemoveFile(tempFilePath)

	newRef := fmt.Sprintf("refs/stacks/%s", name)
	s.gitService.UpdateRef(newRef, objectSha)

	s.gitService.CreateAndCheckoutBranch(name)

	currentStack.Children = append(currentStack.Children, &StackNode{
		Name:         name,
		ParentBranch: parentBranch,
		ParentRefSha: parentRefSha,
		Children:     []*StackNode{},
	})

	s.CacheGraphToDisk(currentStack)
}

func (s *stackService) StackExists(ref string) bool {
	name := s.GetNameFromRef(ref)
	return s.fileService.FileExists(fmt.Sprintf("%s/refs/stacks/%s", s.gitService.DirectoryPath(), name))
}

func (s *stackService) GetCurrentStackNode() *StackNode {
	currentBranch := s.gitService.GetCurrentBranch()

	trunk := s.GetGraph()
	return s.findStack(trunk, currentBranch)
}

func (s *stackService) findStack(node *StackNode, branch string) *StackNode {
	if node.Name == branch {
		return node
	}

	for _, child := range node.Children {
		if found := s.findStack(child, branch); found != nil {
			return found
		}
	}

	return nil
}

func (s *stackService) NeedsSync(stack *StackNode) bool {
	if stack.Parent == nil {
		return false
	}

	actualParentSha := s.gitService.RevParse(stack.ParentBranch)
	return stack.ParentRefSha != actualParentSha
}
