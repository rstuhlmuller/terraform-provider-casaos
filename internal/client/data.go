package casaos

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type AppGrid struct {
}
type AppGridResponse struct {
}

func (c *Client) GetAppManagementWebAppGrid(app AppGrid) (*AppGridResponse, error) {
	rb, err := json.Marshal(app)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/v2/app_management/web/appgrid", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}
	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, errors.New("unable to login")
	}

	ar := AppGridResponse{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	return &ar, nil
}
