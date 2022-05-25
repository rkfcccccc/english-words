package freepik

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/rkfcccccc/english_words/services/dictionary/pkg/pictureapi"

	"github.com/PuerkitoBio/goquery"
)

const searchUrl = "https://www.freepik.com/search"

type repository struct{}

func NewPictureRepository() pictureapi.Repository {
	return &repository{}
}

func (repo *repository) GetName() string {
	return "freepik"
}

func (repo *repository) Search(ctx context.Context, query string) ([]pictureapi.Picture, error) {
	request, err := http.NewRequestWithContext(ctx, "GET", searchUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext: %v", err)
	}

	params := request.URL.Query()
	params.Add("format", "search")
	params.Add("type", "photo")
	params.Add("orientation", "landscape")
	params.Add("selection", "1")
	params.Add("query", query)

	request.URL.RawQuery = params.Encode()

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("DefaultClient.Do: %v", err)
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("bad status code %s", response.Status)
	}

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, fmt.Errorf("goquery.NewDocumentFromReader: %v", err)
	}

	result := []pictureapi.Picture{}
	document.Find("figure.showcase__item").Each(func(i int, s *goquery.Selection) {
		imageUrl := s.AttrOr("data-image", "")

		parsed, err := url.Parse(imageUrl)
		if err != nil {
			return
		}

		parsed.RawQuery = "?size=500"
		result = append(result, pictureapi.Picture{Url: parsed.String(), Source: "www.freepik.com"})
	})

	return result, nil
}
