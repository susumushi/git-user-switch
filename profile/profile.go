package profile

import (
	"encoding/json"
	"errors"
	"fmt"
	"git-user-switch/utils"
	"os"
)

type Profiles []Profile
type Profile struct {
	Name                 string
	Email                string
	NickName             string
	InsertUsernameTarget []string
}

func (ps *Profiles) Load(config string) error {
	absPath, err := utils.PathParser(config)
	if err != nil {
		return fmt.Errorf("failed to load user profile: %s", err)
	}
	bs, err := os.ReadFile(absPath)
	if err != nil {
		return fmt.Errorf("failed to load user profile: %s", err)
	}
	if err = json.Unmarshal(bs, ps); err != nil {
		return fmt.Errorf("failed to load user profile: %s", err)
	}
	return nil
}

func (ps *Profiles) Save(config string) error {
	bs, err := json.Marshal(ps)
	if err != nil {
		return fmt.Errorf("failed to write user profile: %s", err)
	}
	absPath, err := utils.PathParser(config)
	if err != nil {
		return fmt.Errorf("failed to write user profile: %s", err)
	}
	if err = os.WriteFile(absPath, bs, 0664); err != nil {
		return fmt.Errorf("failed to write user profile: %s", err)
	}
	return nil
}

func (ps *Profiles) Flush(config string) error {
	*ps = Profiles{}
	bs, err := json.Marshal(ps)
	if err != nil {
		return fmt.Errorf("failed to flush profile: %s", err)
	}
	absPath, err := utils.PathParser(config)
	if err != nil {
		return fmt.Errorf("failed to flush user profile: %s", err)
	}
	if err = os.WriteFile(absPath, bs, 0664); err != nil {
		return fmt.Errorf("failed to flush user profile: %s", err)
	}
	return nil
}

func (ps *Profiles) Set(p Profile) error {
	for _, cp := range *ps {
		if cp.Name == p.Name {
			return errors.New("profile exists. do not update")
		}
	}
	modify := append(*ps, p)
	*ps = modify
	return nil
}

func (ps *Profiles) Delete(p Profile) error {
	var modify Profiles
	for index, cp := range *ps {
		if cp.Name == p.Name {
			modify = append((*ps)[:index], (*ps)[index+1:]...)
			break
		}
	}
	*ps = modify
	return nil
}
