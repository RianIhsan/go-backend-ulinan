package cloudinary

import (
	"context"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"time"
	"ulinan/config"
)

func Uploader(input interface{}) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cloudName := config.BootConfig().Cloudinary.CCName
	apiKey := config.BootConfig().Cloudinary.CCAPIKey
	apiSecret := config.BootConfig().Cloudinary.CCAPISecret
	folderName := config.BootConfig().Cloudinary.CCFolder

	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		return "", err
	}

	uploadParams := uploader.UploadParams{
		Folder: folderName,
	}

	uploadResult, err := cld.Upload.Upload(ctx, input, uploadParams)
	if err != nil {
		return "", err
	}

	return uploadResult.SecureURL, nil

}
