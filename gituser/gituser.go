package gituser

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
)

type TargetScope int

const TargetScopeAuto TargetScope = 0
const TargetScopeLocal TargetScope = 1
const TargetScopeGlobal TargetScope = 2
const TargetScopeSystem TargetScope = 3

func mapTargetScopeToConfigScope(target TargetScope) (config.Scope, error) {
	switch target {
	case TargetScopeLocal:
		return config.LocalScope, nil
	case TargetScopeGlobal:
		return config.GlobalScope, nil
	case TargetScopeSystem:
		return config.SystemScope, nil
	}
	return 0, errors.New("not supported target")
}

type GitUser struct {
	//scope config.Scope
	targetScope TargetScope

	//config config.Config
	Name                 string
	Email                string
	InsertUsernameTarget []string
}

func New(scope TargetScope) (GitUser, error) {
	gu := GitUser{
		targetScope: scope,
	}

	if err := gu.loadUser(); err != nil {
		return GitUser{}, fmt.Errorf("failed to New. %s", err)
	}
	return gu, nil
}

func (g *GitUser) loadUser() error {
	c, err := g.getConfig()
	if err != nil {
		return fmt.Errorf("user information load from git repository or git config: %s", err)
	}
	g.setConfigToStruct(*c)
	return nil
}

func (g *GitUser) getConfig() (*config.Config, error) {
	if g.targetScope == TargetScopeAuto {
		c, err := getConfigFromLocalRepo()
		if err == nil && c.User.Name != "" {
			return c, nil
		}
		c, err = config.LoadConfig(config.GlobalScope)
		if err == nil && c.User.Name != "" {
			return c, nil
		}
		c, err = config.LoadConfig(config.SystemScope)
		if err == nil && c.User.Name != "" {
			return c, nil
		}
		return nil, fmt.Errorf("failed to config load from auto scope target: %s", err)

	} else if g.targetScope == TargetScopeLocal {
		c, err := getConfigFromLocalRepo()
		if err != nil {
			return nil, fmt.Errorf("config get error: %s", err)
		}
		return c, nil
	} else {
		configScope, err := mapTargetScopeToConfigScope(g.targetScope)
		if err != nil {
			return nil, fmt.Errorf("failed to config load from global scope or system scope: %s", err)
		}
		c, err := config.LoadConfig(configScope)
		if err != nil {
			return nil, fmt.Errorf("failed to config load from global scope or system scope: %s", err)
		}
		return c, nil
	}
}

func getLocalRepo() (*git.Repository, error) {
	isNotRootDir := true
	path, _ := os.Getwd()
	if path == `/` {
		isNotRootDir = false
	}

	repo, err := git.PlainOpen(path)
	for err != nil && isNotRootDir {
		path = filepath.Dir(path)
		repo, err = git.PlainOpen(path)
		if path == `/` {
			isNotRootDir = false
		}
	}
	if err != nil {
		return nil, fmt.Errorf("git repository is not found: %s", err)
	}

	return repo, nil
}

func getConfigFromLocalRepo() (*config.Config, error) {
	repo, err := getLocalRepo()
	if err != nil {
		return nil, fmt.Errorf("failed to config load from local scope: %s", err)
	}
	c, err := repo.Storer.Config()
	if err != nil {
		return nil, fmt.Errorf("failed to config load from local scope: %s", err)
	}
	return c, nil
}

func (g *GitUser) SetConfig() error {
	c, err := g.genModifiedConfig()

	if err != nil {
		return err
	}
	if g.targetScope != TargetScopeLocal {
		configScope, err := mapTargetScopeToConfigScope(g.targetScope)
		if err != nil {
			return fmt.Errorf("config path not exists: %s", err)
		}
		paths, err := config.Paths(configScope)
		if err != nil {
			return fmt.Errorf("config path not exists: %s", err)
		}
		b, err := c.Marshal()
		if err != nil {
			return fmt.Errorf("config serialize failed: %s", err)
		}
		err = os.WriteFile(paths[0], b, 0664)
		if err != nil {
			return fmt.Errorf("failed to write config file: %s", err)
		}
		return nil
	} else {
		r, err := getLocalRepo()
		if err != nil {
			return fmt.Errorf("failed to open git repository: %s", err)
		}

		err = r.Storer.SetConfig(c)
		if err != nil {
			return fmt.Errorf("failed to set config to git repository: %s", err)
		}
		return nil
	}
}

func (g *GitUser) setConfigToStruct(c config.Config) {
	g.Name = c.User.Name
	g.Email = c.User.Email
	g.InsertUsernameTarget = []string{}
	for _, urls := range c.URLs {
		g.InsertUsernameTarget = append(g.InsertUsernameTarget, urls.InsteadOf)
	}
}

func (g *GitUser) genModifiedConfig() (*config.Config, error) {
	c, err := g.getConfig()
	if err != nil {
		return &config.Config{}, err
	}
	c.User.Name = g.Name
	c.User.Email = g.Email
	c.URLs = map[string]*config.URL{}
	for _, url := range g.InsertUsernameTarget {
		// expect url is `https://example.com`
		parsed := strings.Split(url, `://`)
		host := parsed[1]
		prefix := parsed[0] + `://`
		URLWithUsername := prefix + g.Name + `@` + host

		c.URLs[URLWithUsername] = &config.URL{
			Name:      URLWithUsername,
			InsteadOf: url,
		}
	}
	//url := `https://example.com`
	//parsed := strings.Split(url, `://`)
	//host := parsed[1]
	//prefix := parsed[0] + `://`
	////c.URLs["https://github.com"].ApplyInsteadOf(prefix + g.Name + `@` + host)
	//c.URLs["https://example.com"] = &config.URL{
	//	Name:      url,
	//	InsteadOf: prefix + g.Name + `@` + host,
	//}
	return c, nil

}
