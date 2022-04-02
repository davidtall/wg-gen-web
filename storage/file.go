package storage

import (
	"encoding/json"
	"github.com/vx3r/wg-gen-web/model"
	"github.com/vx3r/wg-gen-web/util"
	"os"
	"path/filepath"
	"strings"
)

// Serialize write interface to disk
func Serialize(id string, c interface{}) error {
	b, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	return util.WriteFile(filepath.Join(os.Getenv("WG_CONF_DIR"), id), b)
}

// Deserialize read interface from disk
func Deserialize(id string) (interface{}, error) {
	path := filepath.Join(os.Getenv("WG_CONF_DIR"), id)
	data, err := util.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if strings.HasPrefix(id, "server") {
		var s *model.Server
		err = json.Unmarshal(data, &s)
		if err != nil {
			return nil, err
		}
		if s.Name == "" {
			s.Name = strings.TrimRight(id, ".json")
		}
		s.IsCurrentServer = false
		if id == GetServerFileName() {
			s.IsCurrentServer = true
		}
		return s, nil
	}

	// if not the server, must be client
	var c *model.Client
	err = json.Unmarshal(data, &c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
func GetServerFileName() string {
	filename := os.Getenv("SERVER_FILE")
	if filename == "" {
		hostname, err := os.Hostname()
		if err != nil {
			return "server.json"
		}
		filename = "server-" + hostname + ".json"
	}
	return filename
}
