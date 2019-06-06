package tests

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"testing"
)

type Config struct {
	Test Test `yaml:"test" json:"test"`
}

type Test struct {
	User []string `yaml:"user" json:"user"`
	Db   Db       `yaml:"db" json:"db"`
	Web  Web      `yaml:"web" json:"web"`
}

type Web struct {
	Port string `yaml:"port" json:"port"`
	Host string `yaml:"host" json:"host"`
}

type Db struct {
	Host     string `yaml:"host" json:"host"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
}

//read yaml config
//注：path为yaml或yml文件的路径
func YamlToStruct(path string) (*Config, error) {
	conf := &Config{}
	if f, err := os.Open(path); err != nil {
		return nil, err
	} else {
		yaml.NewDecoder(f).Decode(conf)
	}
	return conf, nil
}

func StructToJson(conf *Config) (string, error) {
	yamlJson, err := json.Marshal(conf)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return "", err
	}
	return string(yamlJson), nil
}

func YAMLToMAP(path string) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if f, err := ioutil.ReadFile(path); err != nil {
		return nil, err
	} else {
		err = yaml.Unmarshal([]byte(string(f)), &m)
		return m, nil
	}
}

func JSONToYAML(j []byte) ([]byte, error) {
	// Convert the JSON to an object.
	var jsonObj interface{}
	// We are using yaml.Unmarshal here (instead of json.Unmarshal) because the
	// Go JSON library doesn't try to pick the right number type (int, float,
	// etc.) when unmarshalling to interface{}, it just picks float64
	// universally. go-yaml does go through the effort of picking the right
	// number type, so we can preserve number type throughout this process.
	err := yaml.Unmarshal(j, &jsonObj)
	if err != nil {
		return nil, err
	}

	// Marshal this object into YAML.
	return yaml.Marshal(jsonObj)
}

func TestYaml(t *testing.T) {
	conf, err := YamlToStruct("test.yaml")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	// 1 yaml to json
	yamlJson, err := StructToJson(conf)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(yamlJson)

	// 2 yaml to map
	yamlMap, err := YAMLToMAP("test.yaml")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	// 修改
	yamlMap["title"] = "world"
	fmt.Println(yamlMap)

	d, err := yaml.Marshal(&yamlMap)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(string(d))

	// 3 json to yaml
	jsonYaml, err := JSONToYAML([]byte(yamlJson))
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(string(jsonYaml))
}
