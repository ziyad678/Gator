package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}


func Read ()(Config, error){
	log.Println("Entering Read Function")
	filePath, err := getConfigFilePath()
	if err !=nil{
		log.Fatalf("Read function couldn't get file path. %v\n",err)
		return Config{},err
	}
	file,err := os.Open(filePath)
	if err !=nil{
		log.Fatalf("ReadFile couldnt open the file. %v\n",err)
		return Config{},err
	}
	config := Config{}
	dec := json.NewDecoder(file)
	err = dec.Decode(&config)
	if err !=nil{
		log.Fatalf("Couldnt deconde json. %v\n",err)
	}
	return config, nil
}
func getConfigFilePath()(string, error){
	log.Println("Entering getConfigFilePath Function")
	home, err := os.UserHomeDir()
	if err != nil{
		log.Println("Cant read user Home Directory")
		return "", err
	}
	filePath := filepath.Join(home,configFileName)
	log.Printf("Config file path retireved successfully returning %v\n",filePath)
	return filePath,nil

}

func (cfg *Config) SetUser(username string)error{
	cfg.CurrentUserName = username
	return write(*cfg)
}

func write (cfg Config) error{
	filePath, err := getConfigFilePath()
	if err !=nil{
		log.Printf("Write function couldn't get file path. %v\n",err)
		return err
	}
	file,err := os.Open(filePath)
	if err !=nil{
		log.Printf("Open func couldnt open the file. %v\n",err)
		return err
	}
	enc := json.NewEncoder(file)
	err = enc.Encode(&cfg)
	if err !=nil{
		log.Fatalf("Couldnt deconde json. %v\n",err)
	}
	return nil

}