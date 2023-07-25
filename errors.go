package surfio

const ImgNotFound = "Image Not Found"

type ImgNotFoundError struct {
	Message string
	Code    int
}

func (i *ImgNotFoundError) Error() string {
	return i.Message
}

func NewImgNotFoundError() error {
	return &ImgNotFoundError{
		Message: ImgNotFound,
		Code:    400,
	}
}
