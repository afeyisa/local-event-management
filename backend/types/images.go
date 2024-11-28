package types

// for images to based
type Images struct{
	ThumbNail string `json:"thumbnail"`
	OtherImages []string `json:"otherImages"`
}

// for image links to returned
type Urls struct{
	Thumbnail_image_url string `json:"thumbnail_image_url"`
	Other_images_urls []string `json:"other_images_urls"`
	Errors []string `json:"errors"`
}

// for image saving payload
type ImagesPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            Images `json:"input"`
  }


// for image request url 
type Url struct{
	Link string `json:"link"`
}

// for request payload
type ImageRequestPayload struct{
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            Url `json:"input"`

  }

// for ecoding the image to base64 and serve 
type Image struct{
	Image string `json:"image"`
}

//thumanal image link to deleted
type ThumbNailUrl struct {
	Thumbnail_image_url string `json:"thumbnail_image_url"`
}

// othr image link to be delete
type ImageUrl struct{
	URL string `json:"image_url"`
}
