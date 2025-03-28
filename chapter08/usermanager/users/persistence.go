package users

import (
	"encoding/gob"
	"encoding/json"
	"os"
)

type Persistence interface {
	Load() (map[int]User, error)
	Store(users map[int]User) error
}

type JSONFile struct {
	name string
}

func NewJsonFile(name string) JSONFile {
	return JSONFile{name + ".json"}
}

func (j JSONFile) Load() (map[int]User, error) {
	bytes, err := os.ReadFile(j.name)
	if err != nil {
		if os.IsNotExist(err) {
			return map[int]User{}, nil
		}
		return nil, err
	}
	var users map[int]User
	err = json.Unmarshal(bytes, &users)
	return users, err
}

func (j JSONFile) Store(users map[int]User) error {
	bytes, err := json.Marshal(users)
	if err != nil {
		return err
	}
	err = os.WriteFile(j.name, bytes, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

type GOBFile struct {
	name string
}

func NewGOBFile(name string) GOBFile {
	return GOBFile{name + ".gob"}
}

func (g GOBFile) Load() (map[int]User, error) {
	file, err := os.Open(g.name)
	if err != nil {
		if os.IsNotExist(err) {
			return map[int]User{}, nil
		}
		return nil, err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	var users map[int]User
	err = decoder.Decode(&users)
	return users, err
}

func (g GOBFile) Store(users map[int]User) error {
	file, err := os.Create(g.name)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	return encoder.Encode(users)
}
