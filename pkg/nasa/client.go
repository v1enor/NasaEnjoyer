package nasa

import (
	"NasaEnjoyer/domain"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type NASAClient struct {
	apikey string
	URl    string
	client *http.Client
}

func NewNASAClient(apikey string, URl string, client *http.Client) *NASAClient {
	return &NASAClient{
		apikey: apikey,
		URl:    URl,
		client: client,
	}
}

func (n *NASAClient) GetAPOD(ctx context.Context) (*domain.APOD, error) {
	url := n.URl + "/planetary/apod?api_key=" + n.apikey
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, domain.ErrCreateRequest
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, domain.ErrFetchData
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, domain.ErrUnexpectedStatusCode
	}

	var apod domain.APOD
	err = json.NewDecoder(response.Body).Decode(&apod)
	if err != nil {
		return nil, domain.ErrDecodeResponse
	}

	return &apod, nil
}

func (n *NASAClient) DownloadImage(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, domain.ErrRequestCreation
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, domain.ErrImageDownload
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, domain.ErrUnexpectedStatusCode
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, domain.ErrFetchData
	}

	return data, nil
}
