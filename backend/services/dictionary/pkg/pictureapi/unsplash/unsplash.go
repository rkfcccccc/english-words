package unsplash

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/rkfcccccc/english_words/services/dictionary/pkg/pictureapi"

	"github.com/PuerkitoBio/goquery"
)

const searchUrl = "https://unsplash.com/s/photos"

type repository struct{}

func NewPictureRepository() pictureapi.Repository {
	return &repository{}
}

func (repo *repository) GetName() string {
	return "unsplash"
}

func (repo *repository) Search(ctx context.Context, query string) ([]pictureapi.Picture, error) {
	requestUrl := fmt.Sprintf("%s/%s", searchUrl, query)
	request, err := http.NewRequestWithContext(ctx, "GET", requestUrl, nil)

	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext: %v", err)
	}

	params := request.URL.Query()
	params.Add("orientation", "landscape")

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
	document.Find("figure[itemprop='image']").Each(func(i int, s *goquery.Selection) {
		srcset := s.Find("img[srcset]").AttrOr("srcset", "")
		sources := strings.Split(srcset, ", ")

		if len(sources) == 0 {
			return
		}

		imageUrl := strings.Split(sources[0], " ")[0]

		parsed, err := url.Parse(imageUrl)
		if err != nil {
			return
		}

		parsed.RawQuery = "w=500"
		result = append(result, pictureapi.Picture{Url: parsed.String(), Source: "unsplash.com"})
	})

	return result, nil
}
