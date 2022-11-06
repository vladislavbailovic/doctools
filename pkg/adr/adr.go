package adr

import (
	"doctools/pkg/config"
	"doctools/pkg/storage"
	"fmt"
	"strconv"
	"strings"
)

const (
	AdrRepository string = "adr"
)

type entity struct {
	data Data
}

func (x entity) GetID() string {
	return fmt.Sprintf("adr-%03d", x.data.Number)
}

func (x entity) Content() []byte {
	return []byte(x.data.String())
}

func numberFromEntityId(id string) (int, error) {
	numeric := strings.TrimPrefix(id, "adr-")
	return strconv.Atoi(numeric)
}

type Repository struct {
	storage.Repository
}

func (x Repository) NextID() (uint, error) {
	var max uint = 0
	ids, err := x.ListIDs()
	if err != nil {
		return max, err
	}
	for _, raw := range ids {
		id, err := numberFromEntityId(raw)
		if err != nil {
			continue
		}
		if id > int(max) {
			max = uint(id)
		}
	}
	return max + 1, nil
}

func Save(data Data, repo storage.Saveable) error {
	if err := repo.Save(entity{data: data}); err != nil {
		return err
	}
	return nil
}

func GetRepo(cfg config.Configuration) (Repository, error) {
	store, err := storage.GetRepo(cfg, AdrRepository)
	repo := Repository{store}
	if err != nil {
		return repo, err
	}
	return repo, nil
}
