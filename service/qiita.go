package service

import (
	"encoding/json"
	"fmt"
	"github.com/tonouchi510/Jeeek/domain"
	"github.com/tonouchi510/Jeeek/domain/repository"
	"github.com/tonouchi510/Jeeek/service/model"
	"io/ioutil"
	"net/http"
)

type qiitaService struct{
	baseUrl    string
}

func NewQiitaService() repository.QiitaRepository {
	return &qiitaService{baseUrl: "https://qiita.com/api/v2"}
}

func (s qiitaService) GetArticleByUserId(userID string) (res *domain.Activity, err error) {
	resp, err := http.Get(s.baseUrl + "/users/" + userID + "/items?page=1&per_page=1")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("StatusCode=%d", resp.StatusCode)
	}
	var posts []model.QiitaPost
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &posts)
	if err != nil {
		return nil, err
	}
	post := posts[0]

	var tags []string
	for _, tag := range posts[0].Tags {
		tags = append(tags, tag.Name)
	}
	content := domain.Content{
		Subject: post.Title,
		Url: post.Url,
		Comment: "",
	}

	res = &domain.Activity{
		ID:        post.ID,
		Category:  2,
		Content:   content,
		Rank:      0,
		Tags:      tags,
	}
	return res, nil
}

