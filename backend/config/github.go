package config

import "os"

var (
	GithubRepoOwner = os.Getenv("GITHUB_REPO_OWNER")
	GithubRepoName  = os.Getenv("GITHUB_REPO_NAME")
)
