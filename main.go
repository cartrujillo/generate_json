package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	service := flag.String("service", "", "")
	image := flag.String("image", "", "")
	port := flag.Int("port", 0, "")
	entrypoint := flag.String("entrypoint", "", "")
	dbPort := flag.String("db_port", "", "")
	dbHost := flag.String("db_host", "", "")
	dbUser := flag.String("db_user", "", "")
	dbName := flag.String("db_name", "", "")
	dbPassword := flag.String("db_password", "", "")
	tier := flag.String("tier", "", "")
	jwtKey := flag.String("jwt_key", "", "")
	idFilesPath := flag.String("id_files_path", "", "")
	debug := flag.String("debug", "", "")
	flag.Parse()

	taskDefinition := map[string]interface{}{
		"family":      *service,
		"networkMode": "bridge",
		"containerDefinitions": []map[string]interface{}{
			{
				"name":   *service,
				"image":  *image,
				"cpu":    0,
				"memory": 128,
				"portMappings": []map[string]interface{}{
					{
						"containerPort": *port,
						"hostPort":      *port,
						"protocol":      "tcp",
					},
				},
				"essential":  true,
				"entryPoint": []string{*entrypoint},
				"environment": []map[string]interface{}{
					{
						"name":  "PORT",
						"value": strconv.FormatInt(int64(*port), 10),
					},
					{
						"name":  "DB_PORT",
						"value": *dbPort,
					},
					{
						"name":  "DB_HOST",
						"value": *dbHost,
					},
					{
						"name":  "DB_NAME",
						"value": *dbName,
					},
					{
						"name":  "DB_PASSWORD",
						"value": *dbPassword,
					},
					{
						"name":  "DB_USER",
						"value": *dbUser,
					},
					{
						"name":  "TIER",
						"value": *tier,
					},
					{
						"name":  "JWT_KEY",
						"value": *jwtKey,
					},
					{
						"name":  "ID_FILES_PATH",
						"value": *idFilesPath,
					},
					{
						"name":  "DEBUG",
						"value": *debug,
					},
				},
			},
		},
		"volumes":                 []interface{}{},
		"placementConstraints":    []interface{}{},
		"requiresCompatibilities": []string{"EC2"},
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	jsonFile, err := os.Create(dir + "/ci-cd/task_definition.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	jsonData, err := json.Marshal(taskDefinition)
	jsonFile.Write(jsonData)
	jsonFile.Close()
	fmt.Println("JSON task definition data written to", jsonFile.Name())
}
