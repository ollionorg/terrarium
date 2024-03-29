// Copyright (c) Ollion
// SPDX-License-Identifier: Apache-2.0

package git

import (
	"context"
	"net/http"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type GitClient struct {
	client *github.Client
}

func GithubClient(ctx context.Context, token string) Git {
	var tc *http.Client

	if token != "" {
		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
		tc = oauth2.NewClient(ctx, ts)
	} else {
		// No token provided, create a client without authentication for public access.
		tc = http.DefaultClient
	}

	client := github.NewClient(tc)
	return &GitClient{client: client}
}

func (gh *GitClient) FetchCommitSHA(ctx context.Context, owner, repo, ref string) (string, error) {
	commit, _, err := gh.client.Repositories.GetCommitSHA1(ctx, owner, repo, ref, "")
	if err != nil {
		return "", err
	}
	return commit, nil
}

func (gh *GitClient) GetContents(ctx context.Context, owner, repo, ref, path string) (*github.RepositoryContent, []*github.RepositoryContent, error) {
	content, list, _, err := gh.client.Repositories.GetContents(ctx, owner, repo, path, &github.RepositoryContentGetOptions{Ref: ref})
	if err != nil {
		return nil, nil, err
	}
	return content, list, nil
}
