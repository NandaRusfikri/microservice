package schemas

import "errors"

type CallAPIDto struct {
	Method       string
	Url          string
	ContentType  string
	Headers      map[string]interface{}
	BodyRequest  string
	BodyResponse string
	HttpCode     int
}

func (d *CallAPIDto) Validate() error {
	if d.Method == "" {
		return errors.New("method required")
	}
	if d.Url == "" {
		return errors.New("url required")
	}

	return nil
}
