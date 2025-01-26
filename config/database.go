package config

import (
	"context"
	"fmt"
	"log"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {

}

func GetSecret(projectID, secretName string) (string, error) {
	ctx := context.Background()

	c, err := secretmanager.NewClient(ctx)

	if err != nil {
		log.Fatalf("Fail to create a client from Secret Manager: &v", err)
	}
	defer c.Close()

	fullSecretName := fmt.Sprintf("projects/%s/secrets/%s/versions/latest", projectID, secretName)

	request := &secretmanagerpb.AccessSecretVersionRequest{
		Name: fullSecretName,
	}
	result, err := c.AccessSecretVersion(ctx, request)

	if err != nil {
		log.Fatalf("Cannot get the secret: %v", err)
	}
	return string(result.Payload.Data), err
}
