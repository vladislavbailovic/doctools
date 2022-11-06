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
	return numberToEntityId(x.data.Number)
}

func (x entity) Content() []byte {
	return []byte(x.data.String())
}

func numberToEntityId(num uint) string {
	return fmt.Sprintf("adr-%03d", num)
}

func numberFromEntityId(id string) (int, error) {
	numeric := strings.TrimPrefix(id, "adr-")
	return strconv.Atoi(numeric)
}

type Repository struct {
	storage.Repository
}

func (x Repository) GetByNumber(num uint) (Data, error) {
	var data Data
	id := numberToEntityId(num)
	raw, err := x.GetByID(id)
	if err != nil {
		return data, err
	}
	data, err = parseData(string(raw))
	if err != nil {
		return data, err
	}
	data.Number = num
	return data, nil
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

func (x Repository) PathToADR(data Data) string {
	return x.PathTo(entity{data: data})
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
