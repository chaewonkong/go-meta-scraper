package surfio

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/headzoo/surf/agent"
	"github.com/headzoo/surf/browser"
	"gopkg.in/headzoo/surf.v1"
)

// Scraper 스크래퍼
type Scraper struct {
	TargetURL string `json:"target_url"`
	Bow       *browser.Browser
}

// New constructor
func New(url string) *Scraper {
	bow := surf.NewBrowser()
	bow.AddRequestHeader("Accept", "text/html")
	bow.AddRequestHeader("Accept-Charset", "utf8")
	bow.SetUserAgent(agent.Chrome())

	return &Scraper{
		TargetURL: url,
		Bow:       bow,
	}
}

// GetMetaContentByProp property명으로 meta tag의 content 가져오기
func (s *Scraper) GetMetaContentByProp(prop string) (contents []string, err error) {
	err = s.Bow.Open(s.TargetURL)
	if err != nil {
		return
	}

	s.Bow.Dom().Find("meta").Each(func(i int, s *goquery.Selection) {
		if result, ok := s.Attr("property"); ok && result == prop {
			imgURL, ok := s.Attr("content")
			if ok {
				contents = append(contents, imgURL)
			}
		}

		// name, content 형식도 대응
		if result, ok := s.Attr("name"); ok && result == prop {
			imgURL, ok := s.Attr("content")
			if ok {
				contents = append(contents, imgURL)
			}
		}
	})

	return
}
