package services

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"note-sync/config"
	"note-sync/models"

	"github.com/google/go-github/v45/github"
	"golang.org/x/oauth2"
)

func getGithubClient(token string) *github.Client {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, config.GithubOauthClient)
	tc := oauth2.NewClient(ctx, ts)
	return github.NewClient(tc)
}

func SyncArticleToGithub(article *models.Article, token string) error {
	client := getGithubClient(token)

	ctx := context.Background()

	// 打印更详细的调试信息
	log.Printf("Syncing article: %s", article.Title)

	// 检查仓库访问权限
	repo, resp, err := client.Repositories.Get(ctx, config.AppConfig.Github.RepoOwner, config.AppConfig.Github.RepoName)
	if err != nil {
		log.Printf("Error checking repo access: %v, Response: %+v", err, resp)
		return fmt.Errorf("repository access error: %v", err)
	}
	log.Printf("Repository full name: %s", *repo.FullName)
	log.Printf("Repository permissions: %+v", repo.Permissions)

	// 生成文件路径
	filePath := fmt.Sprintf("articles/%d-%s.md", article.ID, article.Title)
	article.GitPath = filePath

	// 创建或更新文件
	opts := &github.RepositoryContentFileOptions{
		Message: github.String(fmt.Sprintf("Update article: %s", article.Title)),
		Content: []byte(article.Content),
		Branch:  github.String("master"), // 显式指定分支
	}

	// 检查文件是否已存在
	_, _, resp, err = client.Repositories.GetContents(
		ctx,
		config.AppConfig.Github.RepoOwner,
		config.AppConfig.Github.RepoName,
		filePath,
		nil,
	)

	// 改进错误处理
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			// 文件不存在是正常的，继续创建
			log.Printf("File does not exist, will create new: %s", filePath)
		} else {
			// 其他错误需要处理
			log.Printf("Error checking file existence: %v, Response: %+v", err, resp)
			return fmt.Errorf("failed to check file existence: %v", err)
		}
	}

	var createOrUpdateErr error
	if resp == nil || resp.StatusCode == 404 {
		// 文件不存在，创建新文件
		_, _, createOrUpdateErr = client.Repositories.CreateFile(
			ctx,
			config.AppConfig.Github.RepoOwner,
			config.AppConfig.Github.RepoName,
			filePath,
			opts,
		)
	} else {
		// 文件存在，获取SHA
		fileContent, _, _, err := client.Repositories.GetContents(
			ctx,
			config.AppConfig.Github.RepoOwner,
			config.AppConfig.Github.RepoName,
			filePath,
			nil,
		)
		if err != nil {
			log.Printf("Error getting file content: %v", err)
			return fmt.Errorf("failed to get file content: %v", err)
		}
		opts.SHA = fileContent.SHA

		// 更新文件
		_, _, createOrUpdateErr = client.Repositories.UpdateFile(
			ctx,
			config.AppConfig.Github.RepoOwner,
			config.AppConfig.Github.RepoName,
			filePath,
			opts,
		)
	}

	if createOrUpdateErr != nil {
		log.Printf("Error syncing to GitHub: %v", createOrUpdateErr)
		return fmt.Errorf("failed to sync to GitHub: %v", createOrUpdateErr)
	}

	// 更新文章的 GitPath
	if err := models.DB.Model(article).Update("git_path", filePath).Error; err != nil {
		log.Printf("Error updating article git_path: %v", err)
		return fmt.Errorf("failed to update article git_path: %v", err)
	}

	return nil
}

func DeleteArticleFromGithub(article *models.Article, token string) error {
	client := getGithubClient(token)
	ctx := context.Background()

	_, _, err := client.Repositories.DeleteFile(
		ctx,
		config.GithubRepoOwner,
		config.GithubRepoName,
		article.GitPath,
		&github.RepositoryContentFileOptions{
			Message: github.String(fmt.Sprintf("Delete article: %s", article.Title)),
		},
	)

	return err
}

func SyncQuestionToGithub(question *models.Question, token string) error {
	client := getGithubClient(token)
	ctx := context.Background()

	// 创建Issue
	issueRequest := &github.IssueRequest{
		Title: github.String(question.Title),
		Body:  github.String(formatQuestionBody(question)),
	}
	log.Printf("owner: %s, repo: %s", config.GithubRepoOwner, config.GithubRepoName)
	issue, _, err := client.Issues.Create(
		ctx,
		config.AppConfig.Github.RepoOwner,
		config.AppConfig.Github.RepoName,
		issueRequest,
	)
	log.Printf("issue: %+v", issue)

	if err != nil {
		return err
	}

	// 保存Issue URL
	question.IssueURL = issue.GetHTMLURL()
	return models.DB.Save(question).Error
}

func UpdateQuestionInGithub(question *models.Question, token string) error {
	client := getGithubClient(token)
	ctx := context.Background()

	// 从IssueURL中提取Issue编号
	issueNumber := extractIssueNumber(question.IssueURL)
	if issueNumber == 0 {
		return fmt.Errorf("invalid issue URL")
	}

	// 更新Issue
	issueRequest := &github.IssueRequest{
		Title: github.String(question.Title),
		Body:  github.String(formatQuestionBody(question)),
	}

	_, _, err := client.Issues.Edit(
		ctx,
		config.GithubRepoOwner,
		config.GithubRepoName,
		issueNumber,
		issueRequest,
	)

	return err
}

func DeleteQuestionFromGithub(question *models.Question, token string) error {
	client := getGithubClient(token)
	ctx := context.Background()

	issueNumber := extractIssueNumber(question.IssueURL)
	if issueNumber == 0 {
		return fmt.Errorf("invalid issue URL")
	}

	// GitHub API不支持直接删除Issue，我们将其关闭
	issueRequest := &github.IssueRequest{
		State: github.String("closed"),
	}

	_, _, err := client.Issues.Edit(
		ctx,
		config.GithubRepoOwner,
		config.GithubRepoName,
		issueNumber,
		issueRequest,
	)

	return err
}

func formatQuestionBody(question *models.Question) string {
	return fmt.Sprintf("## 问题描述\n\n%s\n\n## 解答\n\n%s", question.Content, question.Answer)
}

func extractIssueNumber(issueURL string) int {
	// 简单的URL解析，从URL中提取Issue编号
	parts := strings.Split(issueURL, "/")
	if len(parts) == 0 {
		return 0
	}

	numberStr := parts[len(parts)-1]
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		return 0
	}

	return number
}
