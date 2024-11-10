package types

// for image save request
type Images struct{
	ThumbNail string `json:"thumbnail"`
	OtherImages []string `json:"otherImages"`
}

type Urls struct{
	Thumbnail_image_url string `json:"thumbnail_image_url"`
	Other_images_urls []string `json:"other_images_urls"`
	Errors []string `json:"errors"`
}

type ImagesPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            Images `json:"input"`
  }


// for image request
type Url struct{
	Link string `json:"link"`
}

type ImageRequestPayload struct{
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            Url `json:"input"`

  }

type Image struct{
	Image string `json:"image"`
}