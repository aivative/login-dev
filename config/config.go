package config

import (
	"encoding/json"
	"log"

	"github.com/tidwall/gjson"
)

var (
	MongoConf  map[string]TMongodbMetadata
	SVCConf    map[string]TServiceMetadata
	APIKeyConf map[string]TAPIKeyMetadata
)

type TServiceMetadata struct {
	Name string `json:"name,omitempty"`
	Host string `json:"host,omitempty"`
	Port string `json:"port,omitempty"`
	URI  string `json:"uri,omitempty"`
}

type TMongodbMetadata struct {
	Type     string `json:"type,omitempty"`
	Name     string `json:"name,omitempty"`
	DBName   string `json:"db_name,omitempty"`
	CollName string `json:"coll_name,omitempty"`
	URI      string `json:"uri,omitempty"`
}

type TAPIKeyMetadata struct {
	Name string `json:"name,omitempty"`
	Key  string `json:"key,omitempty"`
}

type TRawConfig struct {
	Data []TConfigs `json:"data"`
}

type TConfigs struct {
	Kind     string      `json:"kind"`
	Metadata interface{} `json:"metadata"`
}

func ParseMongoConfig(Config string) map[string]TMongodbMetadata {
	mongoMetaStr := gjson.Get(Config, `data.#(kind=="db").metadata.#(type=="mongodb")#`).String()
	var mongoConf []TMongodbMetadata
	if err := json.Unmarshal([]byte(mongoMetaStr), &mongoConf); err != nil {
		log.Fatalln("MONGO CONFIG CAN'T BE PARSED")
		return nil
	}

	MongoConf := make(map[string]TMongodbMetadata)
	for _, conf := range mongoConf {
		MongoConf[conf.Name] = conf
	}

	return MongoConf
}

func ParseServiceConfig(Config string) map[string]TServiceMetadata {
	svcMetaStr := gjson.Get(Config, `data.#(kind=="service").metadata`).String()
	var svcConf []TServiceMetadata
	if err := json.Unmarshal([]byte(svcMetaStr), &svcConf); err != nil {
		log.Fatalln("SERVICE CONFIG CAN'T BE PARSED")
		return nil
	}

	ServiceConf := make(map[string]TServiceMetadata)
	for _, conf := range svcConf {
		ServiceConf[conf.Name] = conf
	}

	return ServiceConf
}

func ParseAPIKeyConfig(Config string) map[string]TAPIKeyMetadata {
	svcMetaStr := gjson.Get(Config, `data.#(kind=="apikey").metadata`).String()
	var svcConf []TAPIKeyMetadata
	if err := json.Unmarshal([]byte(svcMetaStr), &svcConf); err != nil {
		log.Fatalln("API KEY CONFIG CAN'T BE PARSED")
		return nil
	}

	APIKeyConf := make(map[string]TAPIKeyMetadata)
	for _, conf := range svcConf {
		APIKeyConf[conf.Name] = conf
	}

	return APIKeyConf
}
