package config

import (
	"context"
	"fmt"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	userSQL := GetSecret("45782346572", "sql-user")
	pass := GetSecret("45782346572", "sql-pass")
	instanceConName := GetSecret("45782346572", "sql-instance-name")
	sqlDBname := GetSecret("45782346572", "sql-db-name")
	dsn := fmt.Sprintf("%s:%s@unix(%s)/%s?charset=utf8mb4&parseTime=True", userSQL, pass, instanceConName, sqlDBname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Errorf("Cannot connect database")
	}
	return db
}

func GetSecret(projectID, secretName string) string {
	ctx := context.Background()

	c, err := secretmanager.NewClient(ctx)

	if err != nil {
		fmt.Errorf("Fail to create a client from Secret Manager: %v", err)
		return ""
	}
	defer c.Close()

	fullSecretName := fmt.Sprintf("projects/%s/secrets/%s/versions/latest", projectID, secretName)

	request := &secretmanagerpb.AccessSecretVersionRequest{
		Name: fullSecretName,
	}
	result, err := c.AccessSecretVersion(ctx, request)

	if err != nil {
		fmt.Errorf("Cannot get the secret: %v", err)
		return err.Error()
	}
	return string(result.Payload.Data)
}
