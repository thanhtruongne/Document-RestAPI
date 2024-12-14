package common

import (
	"encoding/json"
	"errors"
	"fmt"
)

type ImageStruct struct {
	Id        int    `json:"id" gorm:"column:id;"`
	Url       string `json:"url" gorm:"column:url;"`
	Width     int    `json:"width" gorm:"column:width;"`
	Height    int    `json:"height" gorm:"column:height;"`
	Cloudname string `json:"cloud_name,omitempty" gorm:"-"`
	Extension string `json:"extension,omitempty	" gorm:"-"`
}

func (ImageStruct) TableName() string { return "image" }

func (image *ImageStruct) FullFill(temp string) {
	image.Url = fmt.Sprintf("%s%s", temp, image.Url)
}

func (image *ImageStruct) Scan(val interface{}) error {
	bytes, ok := val.([]byte)

	if !ok {
		return errors.New(fmt.Sprintf("Failed to unmarshal data: ", val))

	}
	var img ImageStruct

	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}
	*image = img
	return nil

}
