package models

import (
	"context"
	"note-sync/config"

	"github.com/google/go-github/v45/github"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string `json:"username"`
	GithubID    int64  `json:"github_id"`
	AvatarURL   string `json:"avatar_url"`
	GithubToken string `json:"-"` // OAuth token
}

func CreateOrUpdateUser(accessToken string) (*User, error) {
	// 使用 OAuth token 获取用户信息
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, config.GithubOauthClient)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	githubUser, _, err := client.Users.Get(ctx, "")
	if err != nil {
		return nil, err
	}

	var user User
	result := DB.Where("github_id = ?", githubUser.GetID()).First(&user)
	if result.Error != nil {
		// 创建新用户
		user = User{
			Username:    githubUser.GetLogin(),
			GithubID:    githubUser.GetID(),
			AvatarURL:   githubUser.GetAvatarURL(),
			GithubToken: accessToken,
		}
		DB.Create(&user)
	} else {
		// 更新现有用户的token
		user.GithubToken = accessToken
		DB.Save(&user)
	}

	return &user, nil
}

func CreateOrUpdateUserWithPAT(oauthToken, patToken string) (*User, error) {
	// 使用 OAuth token 获取用户信息
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: oauthToken})
	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, config.GithubOauthClient)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	githubUser, _, err := client.Users.Get(ctx, "")
	if err != nil {
		return nil, err
	}

	var user User
	result := DB.Where("github_id = ?", githubUser.GetID()).First(&user)
	if result.Error != nil {
		// 创建新用户
		user = User{
			Username:    githubUser.GetLogin(),
			GithubID:    githubUser.GetID(),
			AvatarURL:   githubUser.GetAvatarURL(),
			GithubToken: oauthToken,
		}
		DB.Create(&user)
	} else {
		// 更新现有用户的token
		user.GithubToken = oauthToken
		DB.Save(&user)
	}

	return &user, nil
}
