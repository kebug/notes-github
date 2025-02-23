package config

import (
	"log"
	"net/http"
	"net/url"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var GithubOauthConfig *oauth2.Config
var GithubOauthClient *http.Client

func InitOAuth() {
	GithubOauthClient = &http.Client{
		Transport: createTransport(),
	}

	GithubOauthConfig = &oauth2.Config{
		ClientID:     AppConfig.Github.ClientID,
		ClientSecret: AppConfig.Github.ClientSecret,
		Scopes: []string{
			"repo", // 仓库读写权限
			"user", // 用户信息权限
		},
		Endpoint:    github.Endpoint,
		RedirectURL: "http://localhost:5173/login/callback",
	}

	log.Printf("Initializing OAuth with scopes: %v", GithubOauthConfig.Scopes)
}

func createTransport() http.RoundTripper {
	if !AppConfig.Proxy.Enable {
		return http.DefaultTransport
	}

	proxyURL, err := url.Parse(AppConfig.Proxy.URL)
	if err != nil {
		log.Printf("Failed to parse proxy URL: %v", err)
		return http.DefaultTransport
	}
	log.Printf("Using proxy: %v", proxyURL)
	return &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
}

// func getGithubClient(token string) *github.Client {
// 	ts := oauth2.StaticTokenSource(
// 		&oauth2.Token{
// 			AccessToken: token,
// 			TokenType:   "Bearer", // 显式指定 token 类型
// 		},
// 	)

// 	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, config.GithubOauthClient)
// 	tc := oauth2.NewClient(ctx, ts)

// 	// 设置自定义 User-Agent
// 	client := github.NewClient(tc)
// 	client.UserAgent = "Note-Sync-App"

// 	return client
// }
