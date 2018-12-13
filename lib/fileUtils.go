package lib

import (
	"github.com/beego/bee/logger"
	"github.com/disintegration/imaging"
	"github.com/google/uuid"
	_ "image/jpeg"
	"mime/multipart"
	"os"
)

func SaveFile(user string, fileHeader *multipart.FileHeader) {
	fileName := uuid.New().String()
	dir := "./uploads/" + user + "/"
	file, err := fileHeader.Open()
	defer file.Close()
	if err != nil {
		beeLogger.Log.Error("Error read file " + fileName + " " + err.Error())
		return
	}

	imageData, err := imaging.Decode(file)
	if err != nil {
		beeLogger.Log.Error("Error decode image file " + fileName + " " + err.Error())
		return
	}

	resizedImage := imaging.Resize(imageData, 800, 0, imaging.Lanczos)

	beeLogger.Log.Hint("new file upload:" + fileName)
	err = os.Mkdir(dir, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		beeLogger.Log.Error("Error create dir " + dir + " " + err.Error())
		return
	}
	err = imaging.Save(resizedImage, dir+fileName+".png")
	if err != nil {
		beeLogger.Log.Error("Error save file " + fileName + " " + err.Error())
		return
	}
	AddImageToUser(user, fileName)
}

/*func ResizeImage(imagePath string, width int) ([]byte, error) {
	imageData, err := imaging.Open(imagePath)
	if err != nil {
		return nil, errors.New("Error decode image file " + imagePath + " : " + err.Error())
	}
	resizedImage := imaging.Resize(imageData, width, 0, imaging.Lanczos)
	var file os.File
	imaging.Encode(&file, resizedImage, imaging.PNG)
	return file., nil
}*/
