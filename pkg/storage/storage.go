package storage

import (
	"os"
	"path/filepath"
	"strings"

	"doctools/pkg/config"
)

type Identifier interface {
	GetID() string
}

type Resource interface {
	Content() []byte
}

type IdentifiedResource interface {
	Identifier
	Resource
}

type Saveable interface {
	Save(what IdentifiedResource) error
}

type Repository struct {
	cfg  config.Configuration
	repo string
	path string
}

func (x Repository) ListIDs() ([]string, error) {
	var ids []string
	entries, err := os.ReadDir(x.path)
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := strings.TrimSuffix(entry.Name(), ".md")
		ids = append(ids, name)
	}
	return ids, err
}

func (x Repository) PathTo(what Identifier) string {
	return filepath.Join(x.path, what.GetID()+".md")
}

func (x Repository) GetByID(id string) ([]byte, error) {
	return os.ReadFile(filepath.Join(x.path, id+".md"))
}

func (x Repository) Save(what IdentifiedResource) error {
	dest := filepath.Join(x.path, what.GetID()+".md")
	return os.WriteFile(dest, what.Content(), 0622)
}

func (x Repository) Initialize(cfg config.Configuration) error {
	if err := InitializeProject(cfg); err != nil {
		return PathError("initializing storage: %w", err)
	}

	repoDir, err := getRepoPath(cfg, x.repo)
	if err != nil {
		if nfo, err := os.Stat(repoDir); err != nil {
			if err := os.Mkdir(repoDir, 0722); err != nil {
				return PathError("creating x.repo subdir (%v): %w", repoDir, err)
			}
		} else if nfo.IsDir() {
			return nil
		}
	}

	if _, err := os.Stat(repoDir); err != nil {
		return PathError("x.repo subdir not created (%s): %v", repoDir, err)
	}

	return nil
}

func GetRepo(cfg config.Configuration, repo string) (Repository, error) {
	x := Repository{repo: repo, cfg: cfg}
	path, err := getRepoPath(cfg, repo)
	if err != nil {
		return x, err
	}
	x = Repository{path: path}
	return x, nil
}

func InitializeProject(cfg config.Configuration) error {
	docPath, err := getRoot(cfg)
	if err != nil {
		if _, err := os.Stat(docPath); err != nil {
			if err := os.Mkdir(docPath, 0722); err != nil {
				return PathError("doc dir creation (%s): %w", docPath, err)
			}
		} else {
			return PathError("doc dir resolution (%s): %w", docPath, err)
		}
	}

	if nfo, err := os.Stat(docPath); err != nil {
		return PathError("doc path %s missing: %w", docPath, err)
	} else if !nfo.IsDir() {
		return PathError("%s is not a directory", docPath)
	}

	return nil
}

func getRoot(cfg config.Configuration) (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return currentDir, PathError("unable to determine working directory: %w", err)
	}
	docPath := filepath.Join(currentDir, cfg.DocPath)

	if nfo, err := os.Stat(docPath); err != nil {
		return docPath, PathError("doc path %s missing: %w", docPath, err)
	} else if !nfo.IsDir() {
		return docPath, PathError("%s is not a directory", docPath)
	}

	return docPath, nil
}

func getRepoPath(cfg config.Configuration, repo string) (string, error) {
	docsDir, err := getRoot(cfg)
	if err != nil {
		return docsDir, PathError("getting repo root for %s: %w", repo, err)
	}
	repoDir := filepath.Join(docsDir, repo)

	if nfo, err := os.Stat(repoDir); err != nil {
		return repoDir, PathError("doc path %s missing: %w", repoDir, err)
	} else if !nfo.IsDir() {
		return repoDir, PathError("%s is not a directory", repoDir)
	}

	return repoDir, nil
}
