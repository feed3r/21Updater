package test_models

const PR_SYNCH_HEADER = `Request URL: http://houseoffeeder.ddns.net:9090/githubUpdate
Request method: POST
Accept: */*
Content-Type: application/x-www-form-urlencoded
User-Agent: GitHub-Hookshot/be8a274
X-GitHub-Delivery: 6eb9cc90-f84b-11ee-9d9f-d02a5660463b
X-GitHub-Event: pull_request
X-GitHub-Hook-ID: 468207434
X-GitHub-Hook-Installation-Target-ID: 713630393
X-GitHub-Hook-Installation-Target-Type: repository`

const PR_SYNCH_BODY = `{
	"action": "synchronize",
	"number": 23,
	"pull_request": {
	  "url": "https://api.github.com/repos/feed3r/21Updater/pulls/23",
	  "id": 1818836484,
	  "node_id": "PR_kwDOKokiuc5saT4E",
	  "html_url": "https://github.com/feed3r/21Updater/pull/23",
	  "diff_url": "https://github.com/feed3r/21Updater/pull/23.diff",
	  "patch_url": "https://github.com/feed3r/21Updater/pull/23.patch",
	  "issue_url": "https://api.github.com/repos/feed3r/21Updater/issues/23",
	  "number": 23,
	  "state": "open",
	  "locked": false,
	  "title": "Fix the Push message and add branch ref",
	  "user": {
		"login": "feed3r",
		"id": 4531376,
		"node_id": "MDQ6VXNlcjQ1MzEzNzY=",
		"avatar_url": "https://avatars.githubusercontent.com/u/4531376?v=4",
		"gravatar_id": "",
		"url": "https://api.github.com/users/feed3r",
		"html_url": "https://github.com/feed3r",
		"followers_url": "https://api.github.com/users/feed3r/followers",
		"following_url": "https://api.github.com/users/feed3r/following{/other_user}",
		"gists_url": "https://api.github.com/users/feed3r/gists{/gist_id}",
		"starred_url": "https://api.github.com/users/feed3r/starred{/owner}{/repo}",
		"subscriptions_url": "https://api.github.com/users/feed3r/subscriptions",
		"organizations_url": "https://api.github.com/users/feed3r/orgs",
		"repos_url": "https://api.github.com/users/feed3r/repos",
		"events_url": "https://api.github.com/users/feed3r/events{/privacy}",
		"received_events_url": "https://api.github.com/users/feed3r/received_events",
		"type": "User",
		"site_admin": false
	  },
	  "body": "Fix Push message management\r\n\r\nAdd the branch reference to message sent to Telegram",
	  "created_at": "2024-04-11T21:33:21Z",
	  "updated_at": "2024-04-11T21:35:35Z",
	  "closed_at": null,
	  "merged_at": null,
	  "merge_commit_sha": "257914ae0fe32039ed11c08b654b29eb21b8a2d0",
	  "assignee": null,
	  "assignees": [
  
	  ],
	  "requested_reviewers": [
  
	  ],
	  "requested_teams": [
  
	  ],
	  "labels": [
  
	  ],
	  "milestone": null,
	  "draft": false,
	  "commits_url": "https://api.github.com/repos/feed3r/21Updater/pulls/23/commits",
	  "review_comments_url": "https://api.github.com/repos/feed3r/21Updater/pulls/23/comments",
	  "review_comment_url": "https://api.github.com/repos/feed3r/21Updater/pulls/comments{/number}",
	  "comments_url": "https://api.github.com/repos/feed3r/21Updater/issues/23/comments",
	  "statuses_url": "https://api.github.com/repos/feed3r/21Updater/statuses/537fb74d21b03cf135f7e40145848851b1c83473",
	  "head": {
		"label": "feed3r:test-commits",
		"ref": "test-commits",
		"sha": "537fb74d21b03cf135f7e40145848851b1c83473",
		"user": {
		  "login": "feed3r",
		  "id": 4531376,
		  "node_id": "MDQ6VXNlcjQ1MzEzNzY=",
		  "avatar_url": "https://avatars.githubusercontent.com/u/4531376?v=4",
		  "gravatar_id": "",
		  "url": "https://api.github.com/users/feed3r",
		  "html_url": "https://github.com/feed3r",
		  "followers_url": "https://api.github.com/users/feed3r/followers",
		  "following_url": "https://api.github.com/users/feed3r/following{/other_user}",
		  "gists_url": "https://api.github.com/users/feed3r/gists{/gist_id}",
		  "starred_url": "https://api.github.com/users/feed3r/starred{/owner}{/repo}",
		  "subscriptions_url": "https://api.github.com/users/feed3r/subscriptions",
		  "organizations_url": "https://api.github.com/users/feed3r/orgs",
		  "repos_url": "https://api.github.com/users/feed3r/repos",
		  "events_url": "https://api.github.com/users/feed3r/events{/privacy}",
		  "received_events_url": "https://api.github.com/users/feed3r/received_events",
		  "type": "User",
		  "site_admin": false
		},
		"repo": {
		  "id": 713630393,
		  "node_id": "R_kgDOKokiuQ",
		  "name": "21Updater",
		  "full_name": "feed3r/21Updater",
		  "private": false,
		  "owner": {
			"login": "feed3r",
			"id": 4531376,
			"node_id": "MDQ6VXNlcjQ1MzEzNzY=",
			"avatar_url": "https://avatars.githubusercontent.com/u/4531376?v=4",
			"gravatar_id": "",
			"url": "https://api.github.com/users/feed3r",
			"html_url": "https://github.com/feed3r",
			"followers_url": "https://api.github.com/users/feed3r/followers",
			"following_url": "https://api.github.com/users/feed3r/following{/other_user}",
			"gists_url": "https://api.github.com/users/feed3r/gists{/gist_id}",
			"starred_url": "https://api.github.com/users/feed3r/starred{/owner}{/repo}",
			"subscriptions_url": "https://api.github.com/users/feed3r/subscriptions",
			"organizations_url": "https://api.github.com/users/feed3r/orgs",
			"repos_url": "https://api.github.com/users/feed3r/repos",
			"events_url": "https://api.github.com/users/feed3r/events{/privacy}",
			"received_events_url": "https://api.github.com/users/feed3r/received_events",
			"type": "User",
			"site_admin": false
		  },
		  "html_url": "https://github.com/feed3r/21Updater",
		  "description": "Some experiments in creating a Telegram BOT to receive the updates from a GitHub (in particular, from THIS repository)",
		  "fork": false,
		  "url": "https://api.github.com/repos/feed3r/21Updater",
		  "forks_url": "https://api.github.com/repos/feed3r/21Updater/forks",
		  "keys_url": "https://api.github.com/repos/feed3r/21Updater/keys{/key_id}",
		  "collaborators_url": "https://api.github.com/repos/feed3r/21Updater/collaborators{/collaborator}",
		  "teams_url": "https://api.github.com/repos/feed3r/21Updater/teams",
		  "hooks_url": "https://api.github.com/repos/feed3r/21Updater/hooks",
		  "issue_events_url": "https://api.github.com/repos/feed3r/21Updater/issues/events{/number}",
		  "events_url": "https://api.github.com/repos/feed3r/21Updater/events",
		  "assignees_url": "https://api.github.com/repos/feed3r/21Updater/assignees{/user}",
		  "branches_url": "https://api.github.com/repos/feed3r/21Updater/branches{/branch}",
		  "tags_url": "https://api.github.com/repos/feed3r/21Updater/tags",
		  "blobs_url": "https://api.github.com/repos/feed3r/21Updater/git/blobs{/sha}",
		  "git_tags_url": "https://api.github.com/repos/feed3r/21Updater/git/tags{/sha}",
		  "git_refs_url": "https://api.github.com/repos/feed3r/21Updater/git/refs{/sha}",
		  "trees_url": "https://api.github.com/repos/feed3r/21Updater/git/trees{/sha}",
		  "statuses_url": "https://api.github.com/repos/feed3r/21Updater/statuses/{sha}",
		  "languages_url": "https://api.github.com/repos/feed3r/21Updater/languages",
		  "stargazers_url": "https://api.github.com/repos/feed3r/21Updater/stargazers",
		  "contributors_url": "https://api.github.com/repos/feed3r/21Updater/contributors",
		  "subscribers_url": "https://api.github.com/repos/feed3r/21Updater/subscribers",
		  "subscription_url": "https://api.github.com/repos/feed3r/21Updater/subscription",
		  "commits_url": "https://api.github.com/repos/feed3r/21Updater/commits{/sha}",
		  "git_commits_url": "https://api.github.com/repos/feed3r/21Updater/git/commits{/sha}",
		  "comments_url": "https://api.github.com/repos/feed3r/21Updater/comments{/number}",
		  "issue_comment_url": "https://api.github.com/repos/feed3r/21Updater/issues/comments{/number}",
		  "contents_url": "https://api.github.com/repos/feed3r/21Updater/contents/{+path}",
		  "compare_url": "https://api.github.com/repos/feed3r/21Updater/compare/{base}...{head}",
		  "merges_url": "https://api.github.com/repos/feed3r/21Updater/merges",
		  "archive_url": "https://api.github.com/repos/feed3r/21Updater/{archive_format}{/ref}",
		  "downloads_url": "https://api.github.com/repos/feed3r/21Updater/downloads",
		  "issues_url": "https://api.github.com/repos/feed3r/21Updater/issues{/number}",
		  "pulls_url": "https://api.github.com/repos/feed3r/21Updater/pulls{/number}",
		  "milestones_url": "https://api.github.com/repos/feed3r/21Updater/milestones{/number}",
		  "notifications_url": "https://api.github.com/repos/feed3r/21Updater/notifications{?since,all,participating}",
		  "labels_url": "https://api.github.com/repos/feed3r/21Updater/labels{/name}",
		  "releases_url": "https://api.github.com/repos/feed3r/21Updater/releases{/id}",
		  "deployments_url": "https://api.github.com/repos/feed3r/21Updater/deployments",
		  "created_at": "2023-11-02T23:16:08Z",
		  "updated_at": "2024-03-03T10:09:05Z",
		  "pushed_at": "2024-04-11T21:35:34Z",
		  "git_url": "git://github.com/feed3r/21Updater.git",
		  "ssh_url": "git@github.com:feed3r/21Updater.git",
		  "clone_url": "https://github.com/feed3r/21Updater.git",
		  "svn_url": "https://github.com/feed3r/21Updater",
		  "homepage": null,
		  "size": 7814,
		  "stargazers_count": 0,
		  "watchers_count": 0,
		  "language": "Go",
		  "has_issues": true,
		  "has_projects": true,
		  "has_downloads": true,
		  "has_wiki": false,
		  "has_pages": false,
		  "has_discussions": false,
		  "forks_count": 0,
		  "mirror_url": null,
		  "archived": false,
		  "disabled": false,
		  "open_issues_count": 2,
		  "license": {
			"key": "cc0-1.0",
			"name": "Creative Commons Zero v1.0 Universal",
			"spdx_id": "CC0-1.0",
			"url": "https://api.github.com/licenses/cc0-1.0",
			"node_id": "MDc6TGljZW5zZTY="
		  },
		  "allow_forking": true,
		  "is_template": false,
		  "web_commit_signoff_required": false,
		  "topics": [
  
		  ],
		  "visibility": "public",
		  "forks": 0,
		  "open_issues": 2,
		  "watchers": 0,
		  "default_branch": "main",
		  "allow_squash_merge": true,
		  "allow_merge_commit": true,
		  "allow_rebase_merge": true,
		  "allow_auto_merge": false,
		  "delete_branch_on_merge": false,
		  "allow_update_branch": false,
		  "use_squash_pr_title_as_default": false,
		  "squash_merge_commit_message": "COMMIT_MESSAGES",
		  "squash_merge_commit_title": "COMMIT_OR_PR_TITLE",
		  "merge_commit_message": "PR_TITLE",
		  "merge_commit_title": "MERGE_MESSAGE"
		}
	  },
	  "base": {
		"label": "feed3r:main",
		"ref": "main",
		"sha": "a29ed371af1383f2183671d9c5016a7d3084a592",
		"user": {
		  "login": "feed3r",
		  "id": 4531376,
		  "node_id": "MDQ6VXNlcjQ1MzEzNzY=",
		  "avatar_url": "https://avatars.githubusercontent.com/u/4531376?v=4",
		  "gravatar_id": "",
		  "url": "https://api.github.com/users/feed3r",
		  "html_url": "https://github.com/feed3r",
		  "followers_url": "https://api.github.com/users/feed3r/followers",
		  "following_url": "https://api.github.com/users/feed3r/following{/other_user}",
		  "gists_url": "https://api.github.com/users/feed3r/gists{/gist_id}",
		  "starred_url": "https://api.github.com/users/feed3r/starred{/owner}{/repo}",
		  "subscriptions_url": "https://api.github.com/users/feed3r/subscriptions",
		  "organizations_url": "https://api.github.com/users/feed3r/orgs",
		  "repos_url": "https://api.github.com/users/feed3r/repos",
		  "events_url": "https://api.github.com/users/feed3r/events{/privacy}",
		  "received_events_url": "https://api.github.com/users/feed3r/received_events",
		  "type": "User",
		  "site_admin": false
		},
		"repo": {
		  "id": 713630393,
		  "node_id": "R_kgDOKokiuQ",
		  "name": "21Updater",
		  "full_name": "feed3r/21Updater",
		  "private": false,
		  "owner": {
			"login": "feed3r",
			"id": 4531376,
			"node_id": "MDQ6VXNlcjQ1MzEzNzY=",
			"avatar_url": "https://avatars.githubusercontent.com/u/4531376?v=4",
			"gravatar_id": "",
			"url": "https://api.github.com/users/feed3r",
			"html_url": "https://github.com/feed3r",
			"followers_url": "https://api.github.com/users/feed3r/followers",
			"following_url": "https://api.github.com/users/feed3r/following{/other_user}",
			"gists_url": "https://api.github.com/users/feed3r/gists{/gist_id}",
			"starred_url": "https://api.github.com/users/feed3r/starred{/owner}{/repo}",
			"subscriptions_url": "https://api.github.com/users/feed3r/subscriptions",
			"organizations_url": "https://api.github.com/users/feed3r/orgs",
			"repos_url": "https://api.github.com/users/feed3r/repos",
			"events_url": "https://api.github.com/users/feed3r/events{/privacy}",
			"received_events_url": "https://api.github.com/users/feed3r/received_events",
			"type": "User",
			"site_admin": false
		  },
		  "html_url": "https://github.com/feed3r/21Updater",
		  "description": "Some experiments in creating a Telegram BOT to receive the updates from a GitHub (in particular, from THIS repository)",
		  "fork": false,
		  "url": "https://api.github.com/repos/feed3r/21Updater",
		  "forks_url": "https://api.github.com/repos/feed3r/21Updater/forks",
		  "keys_url": "https://api.github.com/repos/feed3r/21Updater/keys{/key_id}",
		  "collaborators_url": "https://api.github.com/repos/feed3r/21Updater/collaborators{/collaborator}",
		  "teams_url": "https://api.github.com/repos/feed3r/21Updater/teams",
		  "hooks_url": "https://api.github.com/repos/feed3r/21Updater/hooks",
		  "issue_events_url": "https://api.github.com/repos/feed3r/21Updater/issues/events{/number}",
		  "events_url": "https://api.github.com/repos/feed3r/21Updater/events",
		  "assignees_url": "https://api.github.com/repos/feed3r/21Updater/assignees{/user}",
		  "branches_url": "https://api.github.com/repos/feed3r/21Updater/branches{/branch}",
		  "tags_url": "https://api.github.com/repos/feed3r/21Updater/tags",
		  "blobs_url": "https://api.github.com/repos/feed3r/21Updater/git/blobs{/sha}",
		  "git_tags_url": "https://api.github.com/repos/feed3r/21Updater/git/tags{/sha}",
		  "git_refs_url": "https://api.github.com/repos/feed3r/21Updater/git/refs{/sha}",
		  "trees_url": "https://api.github.com/repos/feed3r/21Updater/git/trees{/sha}",
		  "statuses_url": "https://api.github.com/repos/feed3r/21Updater/statuses/{sha}",
		  "languages_url": "https://api.github.com/repos/feed3r/21Updater/languages",
		  "stargazers_url": "https://api.github.com/repos/feed3r/21Updater/stargazers",
		  "contributors_url": "https://api.github.com/repos/feed3r/21Updater/contributors",
		  "subscribers_url": "https://api.github.com/repos/feed3r/21Updater/subscribers",
		  "subscription_url": "https://api.github.com/repos/feed3r/21Updater/subscription",
		  "commits_url": "https://api.github.com/repos/feed3r/21Updater/commits{/sha}",
		  "git_commits_url": "https://api.github.com/repos/feed3r/21Updater/git/commits{/sha}",
		  "comments_url": "https://api.github.com/repos/feed3r/21Updater/comments{/number}",
		  "issue_comment_url": "https://api.github.com/repos/feed3r/21Updater/issues/comments{/number}",
		  "contents_url": "https://api.github.com/repos/feed3r/21Updater/contents/{+path}",
		  "compare_url": "https://api.github.com/repos/feed3r/21Updater/compare/{base}...{head}",
		  "merges_url": "https://api.github.com/repos/feed3r/21Updater/merges",
		  "archive_url": "https://api.github.com/repos/feed3r/21Updater/{archive_format}{/ref}",
		  "downloads_url": "https://api.github.com/repos/feed3r/21Updater/downloads",
		  "issues_url": "https://api.github.com/repos/feed3r/21Updater/issues{/number}",
		  "pulls_url": "https://api.github.com/repos/feed3r/21Updater/pulls{/number}",
		  "milestones_url": "https://api.github.com/repos/feed3r/21Updater/milestones{/number}",
		  "notifications_url": "https://api.github.com/repos/feed3r/21Updater/notifications{?since,all,participating}",
		  "labels_url": "https://api.github.com/repos/feed3r/21Updater/labels{/name}",
		  "releases_url": "https://api.github.com/repos/feed3r/21Updater/releases{/id}",
		  "deployments_url": "https://api.github.com/repos/feed3r/21Updater/deployments",
		  "created_at": "2023-11-02T23:16:08Z",
		  "updated_at": "2024-03-03T10:09:05Z",
		  "pushed_at": "2024-04-11T21:35:34Z",
		  "git_url": "git://github.com/feed3r/21Updater.git",
		  "ssh_url": "git@github.com:feed3r/21Updater.git",
		  "clone_url": "https://github.com/feed3r/21Updater.git",
		  "svn_url": "https://github.com/feed3r/21Updater",
		  "homepage": null,
		  "size": 7814,
		  "stargazers_count": 0,
		  "watchers_count": 0,
		  "language": "Go",
		  "has_issues": true,
		  "has_projects": true,
		  "has_downloads": true,
		  "has_wiki": false,
		  "has_pages": false,
		  "has_discussions": false,
		  "forks_count": 0,
		  "mirror_url": null,
		  "archived": false,
		  "disabled": false,
		  "open_issues_count": 2,
		  "license": {
			"key": "cc0-1.0",
			"name": "Creative Commons Zero v1.0 Universal",
			"spdx_id": "CC0-1.0",
			"url": "https://api.github.com/licenses/cc0-1.0",
			"node_id": "MDc6TGljZW5zZTY="
		  },
		  "allow_forking": true,
		  "is_template": false,
		  "web_commit_signoff_required": false,
		  "topics": [
  
		  ],
		  "visibility": "public",
		  "forks": 0,
		  "open_issues": 2,
		  "watchers": 0,
		  "default_branch": "main",
		  "allow_squash_merge": true,
		  "allow_merge_commit": true,
		  "allow_rebase_merge": true,
		  "allow_auto_merge": false,
		  "delete_branch_on_merge": false,
		  "allow_update_branch": false,
		  "use_squash_pr_title_as_default": false,
		  "squash_merge_commit_message": "COMMIT_MESSAGES",
		  "squash_merge_commit_title": "COMMIT_OR_PR_TITLE",
		  "merge_commit_message": "PR_TITLE",
		  "merge_commit_title": "MERGE_MESSAGE"
		}
	  },
	  "_links": {
		"self": {
		  "href": "https://api.github.com/repos/feed3r/21Updater/pulls/23"
		},
		"html": {
		  "href": "https://github.com/feed3r/21Updater/pull/23"
		},
		"issue": {
		  "href": "https://api.github.com/repos/feed3r/21Updater/issues/23"
		},
		"comments": {
		  "href": "https://api.github.com/repos/feed3r/21Updater/issues/23/comments"
		},
		"review_comments": {
		  "href": "https://api.github.com/repos/feed3r/21Updater/pulls/23/comments"
		},
		"review_comment": {
		  "href": "https://api.github.com/repos/feed3r/21Updater/pulls/comments{/number}"
		},
		"commits": {
		  "href": "https://api.github.com/repos/feed3r/21Updater/pulls/23/commits"
		},
		"statuses": {
		  "href": "https://api.github.com/repos/feed3r/21Updater/statuses/537fb74d21b03cf135f7e40145848851b1c83473"
		}
	  },
	  "author_association": "OWNER",
	  "auto_merge": null,
	  "active_lock_reason": null,
	  "merged": false,
	  "mergeable": null,
	  "rebaseable": null,
	  "mergeable_state": "unknown",
	  "merged_by": null,
	  "comments": 0,
	  "review_comments": 0,
	  "maintainer_can_modify": false,
	  "commits": 1,
	  "additions": 12,
	  "deletions": 1,
	  "changed_files": 2
	},
	"before": "0fe2dd03750d635ef23d7138cb8155c64de93ffc",
	"after": "537fb74d21b03cf135f7e40145848851b1c83473",
	"repository": {
	  "id": 713630393,
	  "node_id": "R_kgDOKokiuQ",
	  "name": "21Updater",
	  "full_name": "feed3r/21Updater",
	  "private": false,
	  "owner": {
		"login": "feed3r",
		"id": 4531376,
		"node_id": "MDQ6VXNlcjQ1MzEzNzY=",
		"avatar_url": "https://avatars.githubusercontent.com/u/4531376?v=4",
		"gravatar_id": "",
		"url": "https://api.github.com/users/feed3r",
		"html_url": "https://github.com/feed3r",
		"followers_url": "https://api.github.com/users/feed3r/followers",
		"following_url": "https://api.github.com/users/feed3r/following{/other_user}",
		"gists_url": "https://api.github.com/users/feed3r/gists{/gist_id}",
		"starred_url": "https://api.github.com/users/feed3r/starred{/owner}{/repo}",
		"subscriptions_url": "https://api.github.com/users/feed3r/subscriptions",
		"organizations_url": "https://api.github.com/users/feed3r/orgs",
		"repos_url": "https://api.github.com/users/feed3r/repos",
		"events_url": "https://api.github.com/users/feed3r/events{/privacy}",
		"received_events_url": "https://api.github.com/users/feed3r/received_events",
		"type": "User",
		"site_admin": false
	  },
	  "html_url": "https://github.com/feed3r/21Updater",
	  "description": "Some experiments in creating a Telegram BOT to receive the updates from a GitHub (in particular, from THIS repository)",
	  "fork": false,
	  "url": "https://api.github.com/repos/feed3r/21Updater",
	  "forks_url": "https://api.github.com/repos/feed3r/21Updater/forks",
	  "keys_url": "https://api.github.com/repos/feed3r/21Updater/keys{/key_id}",
	  "collaborators_url": "https://api.github.com/repos/feed3r/21Updater/collaborators{/collaborator}",
	  "teams_url": "https://api.github.com/repos/feed3r/21Updater/teams",
	  "hooks_url": "https://api.github.com/repos/feed3r/21Updater/hooks",
	  "issue_events_url": "https://api.github.com/repos/feed3r/21Updater/issues/events{/number}",
	  "events_url": "https://api.github.com/repos/feed3r/21Updater/events",
	  "assignees_url": "https://api.github.com/repos/feed3r/21Updater/assignees{/user}",
	  "branches_url": "https://api.github.com/repos/feed3r/21Updater/branches{/branch}",
	  "tags_url": "https://api.github.com/repos/feed3r/21Updater/tags",
	  "blobs_url": "https://api.github.com/repos/feed3r/21Updater/git/blobs{/sha}",
	  "git_tags_url": "https://api.github.com/repos/feed3r/21Updater/git/tags{/sha}",
	  "git_refs_url": "https://api.github.com/repos/feed3r/21Updater/git/refs{/sha}",
	  "trees_url": "https://api.github.com/repos/feed3r/21Updater/git/trees{/sha}",
	  "statuses_url": "https://api.github.com/repos/feed3r/21Updater/statuses/{sha}",
	  "languages_url": "https://api.github.com/repos/feed3r/21Updater/languages",
	  "stargazers_url": "https://api.github.com/repos/feed3r/21Updater/stargazers",
	  "contributors_url": "https://api.github.com/repos/feed3r/21Updater/contributors",
	  "subscribers_url": "https://api.github.com/repos/feed3r/21Updater/subscribers",
	  "subscription_url": "https://api.github.com/repos/feed3r/21Updater/subscription",
	  "commits_url": "https://api.github.com/repos/feed3r/21Updater/commits{/sha}",
	  "git_commits_url": "https://api.github.com/repos/feed3r/21Updater/git/commits{/sha}",
	  "comments_url": "https://api.github.com/repos/feed3r/21Updater/comments{/number}",
	  "issue_comment_url": "https://api.github.com/repos/feed3r/21Updater/issues/comments{/number}",
	  "contents_url": "https://api.github.com/repos/feed3r/21Updater/contents/{+path}",
	  "compare_url": "https://api.github.com/repos/feed3r/21Updater/compare/{base}...{head}",
	  "merges_url": "https://api.github.com/repos/feed3r/21Updater/merges",
	  "archive_url": "https://api.github.com/repos/feed3r/21Updater/{archive_format}{/ref}",
	  "downloads_url": "https://api.github.com/repos/feed3r/21Updater/downloads",
	  "issues_url": "https://api.github.com/repos/feed3r/21Updater/issues{/number}",
	  "pulls_url": "https://api.github.com/repos/feed3r/21Updater/pulls{/number}",
	  "milestones_url": "https://api.github.com/repos/feed3r/21Updater/milestones{/number}",
	  "notifications_url": "https://api.github.com/repos/feed3r/21Updater/notifications{?since,all,participating}",
	  "labels_url": "https://api.github.com/repos/feed3r/21Updater/labels{/name}",
	  "releases_url": "https://api.github.com/repos/feed3r/21Updater/releases{/id}",
	  "deployments_url": "https://api.github.com/repos/feed3r/21Updater/deployments",
	  "created_at": "2023-11-02T23:16:08Z",
	  "updated_at": "2024-03-03T10:09:05Z",
	  "pushed_at": "2024-04-11T21:35:34Z",
	  "git_url": "git://github.com/feed3r/21Updater.git",
	  "ssh_url": "git@github.com:feed3r/21Updater.git",
	  "clone_url": "https://github.com/feed3r/21Updater.git",
	  "svn_url": "https://github.com/feed3r/21Updater",
	  "homepage": null,
	  "size": 7814,
	  "stargazers_count": 0,
	  "watchers_count": 0,
	  "language": "Go",
	  "has_issues": true,
	  "has_projects": true,
	  "has_downloads": true,
	  "has_wiki": false,
	  "has_pages": false,
	  "has_discussions": false,
	  "forks_count": 0,
	  "mirror_url": null,
	  "archived": false,
	  "disabled": false,
	  "open_issues_count": 2,
	  "license": {
		"key": "cc0-1.0",
		"name": "Creative Commons Zero v1.0 Universal",
		"spdx_id": "CC0-1.0",
		"url": "https://api.github.com/licenses/cc0-1.0",
		"node_id": "MDc6TGljZW5zZTY="
	  },
	  "allow_forking": true,
	  "is_template": false,
	  "web_commit_signoff_required": false,
	  "topics": [
  
	  ],
	  "visibility": "public",
	  "forks": 0,
	  "open_issues": 2,
	  "watchers": 0,
	  "default_branch": "main"
	},
	"sender": {
	  "login": "feed3r",
	  "id": 4531376,
	  "node_id": "MDQ6VXNlcjQ1MzEzNzY=",
	  "avatar_url": "https://avatars.githubusercontent.com/u/4531376?v=4",
	  "gravatar_id": "",
	  "url": "https://api.github.com/users/feed3r",
	  "html_url": "https://github.com/feed3r",
	  "followers_url": "https://api.github.com/users/feed3r/followers",
	  "following_url": "https://api.github.com/users/feed3r/following{/other_user}",
	  "gists_url": "https://api.github.com/users/feed3r/gists{/gist_id}",
	  "starred_url": "https://api.github.com/users/feed3r/starred{/owner}{/repo}",
	  "subscriptions_url": "https://api.github.com/users/feed3r/subscriptions",
	  "organizations_url": "https://api.github.com/users/feed3r/orgs",
	  "repos_url": "https://api.github.com/users/feed3r/repos",
	  "events_url": "https://api.github.com/users/feed3r/events{/privacy}",
	  "received_events_url": "https://api.github.com/users/feed3r/received_events",
	  "type": "User",
	  "site_admin": false
	}
  }`
