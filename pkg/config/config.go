package config

import (
	"github.com/ghodss/yaml"
	"github.com/jalilbengoufa/pixicoreAPI/pkg/server"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
)

type configContent []byte

type ConfigFile struct {
	Servers *server.Servers
	Path    string
}

//InitConfig create config if it does not exist
func InitConfig() (*ConfigFile, error) {
	filePath := "servers-config.yaml"

	confFile := new(ConfigFile)
	confFile.Path = filePath
	err := confFile.ReadYamlConfig()
	switch err.(type) {
	case nil:
		return confFile, nil
	case *os.PathError:
		log.Warnln(err)
		return confFile, nil
	// for all unmanaged kind of errors
	default:
		return nil, err
	}
}

//WriteConfig write the yaml config file
func (configFile *ConfigFile) WriteYamlConfig() {

	if _, err := os.Stat(configFile.Path); os.IsNotExist(err) {
		f, err := os.Create(configFile.Path)
		log.Fatalln(err)
		f.Close()
	}
	filename, _ := filepath.Abs(configFile.Path)
	yamlServers, err := yaml.Marshal(&configFile.Servers)

	f, err := os.Create(filename)
	if err != nil {
		log.Errorln(err)
	}
	f.Write(yamlServers)
	f.Sync()
	f.Close()
}

//ReadConfig read the yaml config file
func (configFile *ConfigFile) ReadYamlConfig() (error) {
	// If the config does not exist
	if _, err := os.Stat(configFile.Path); os.IsNotExist(err) {
		log.Infoln("Configfile in does not exist . Init servers list in-memory")
		emptyServers := make(server.Servers)
		configFile.Servers = &emptyServers
		//If the config file are not empty, do more checks to load his content
	} else {

		filename, _ := filepath.Abs(configFile.Path)

		yamlFile, err := ioutil.ReadFile(filename)
		if err != nil {
			return err
		}
		//if cofig file is empty create and use servers in memory
		if len(yamlFile) == 0 {
			emptyServers := make(server.Servers)
			configFile.Servers = &emptyServers
		} else {
			// Try to parse the config file as a server list
			err = yaml.Unmarshal(yamlFile, &configFile.Servers)
			if err != nil {
				log.Infoln("corrupted config file")
				return err
			}

		}
		// if the server list are nil in the config file, make a new list of servers
		if *configFile.Servers == nil {
			emptyServers := make(server.Servers)
			configFile.Servers = &emptyServers
		}
	}
	return nil
}

//GetServers return config of the all the servers
func (configFile *ConfigFile) GetServers() (*server.Servers, error) {
	log.Info(configFile)

	if configFile.Servers == nil {
		return configFile.Servers, new(server.NilServerListError)
	}
	return configFile.Servers, nil
}
