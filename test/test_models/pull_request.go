package test_models

const PR_EXPECTED_TEXT = `Hey, in our repository 21Updater a new PULL REQUEST has been opened by feed3r having title "Feeder bot stub".
Text says "This is the text of the newly created Pull Request"
You can see the changes here: "https://api.github.com/repos/feed3r/21Updater/pulls/3"`

const PR_HEADER = `Request URL: http://houseoffeeder.ddns.net:9090/githubUpdate
Request method: POST
Accept: */*
Content-Type: application/json
User-Agent: GitHub-Hookshot/1d1fb3d
X-GitHub-Delivery: c5c1f2e0-96a0-11ee-916d-a578274848c8
X-GitHub-Event: pull_request
X-GitHub-Hook-ID: 445243377
X-GitHub-Hook-Installation-Target-ID: 713630393
X-GitHub-Hook-Installation-Target-Type: repository`

const PR_BODY = `{
	"action": "synchronize",
	"number": 3,
	"pull_request": {
	  "url": "https://api.github.com/repos/feed3r/21Updater/pulls/3",
	  "id": 1629166543,
	  "node_id": "PR_kwDOKokiuc5hGxvP",
	  "html_url": "https://github.com/feed3r/21Updater/pull/3",
	  "diff_url": "https://github.com/feed3r/21Updater/pull/3.diff",
	  "patch_url": "https://github.com/feed3r/21Updater/pull/3.patch",
	  "issue_url": "https://api.github.com/repos/feed3r/21Updater/issues/3",
	  "number": 3,
	  "state": "open",
	  "locked": false,
	  "title": "Feeder bot stub",
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
	  "body": "This is the text of the newly created Pull Request",
	  "created_at": "2023-12-04T21:46:56Z",
	  "updated_at": "2023-12-09T14:39:35Z",
	  "closed_at": null,
	  "merged_at": null,
	  "merge_commit_sha": "a4fd4125e6eeb56def1c8733be47911204bbeb6e",
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
	  "commits_url": "https://api.github.com/repos/feed3r/21Updater/pulls/3/commits",
	  "review_comments_url": "https://api.github.com/repos/feed3r/21Updater/pulls/3/comments",
	  "review_comment_url": "https://api.github.com/repos/feed3r/21Updater/pulls/comments{/number}",
	  "comments_url": "https://api.github.com/repos/feed3r/21Updater/issues/3/comments",
	  "statuses_url": "https://api.github.com/repos/feed3r/21Updater/statuses/48e646db0ca59051f0986c1f448608098f3e7ca4",
	  "head": {
		"label": "feed3r:feeder_bot_stub",
		"ref": "feeder_bot_stub",
		"sha": "48e646db0ca59051f0986c1f448608098f3e7ca4",
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
		  "updated_at": "2023-11-16T09:48:58Z",
		  "pushed_at": "2023-12-09T14:39:33Z",
		  "git_url": "git://github.com/feed3r/21Updater.git",
		  "ssh_url": "git@github.com:feed3r/21Updater.git",
		  "clone_url": "https://github.com/feed3r/21Updater.git",
		  "svn_url": "https://github.com/feed3r/21Updater",
		  "homepage": null,
		  "size": 9805,
		  "stargazers_count": 0,
		  "watchers_count": 0,
		  "language": null,
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
		  "open_issues_count": 3,
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
		  "open_issues": 3,
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
		"sha": "a5fd97b9e2c4fba2b17973d7c1a93bfb37e26e52",
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
		  "updated_at": "2023-11-16T09:48:58Z",
		  "pushed_at": "2023-12-09T14:39:33Z",
		  "git_url": "git://github.com/feed3r/21Updater.git",
		  "ssh_url": "git@github.com:feed3r/21Updater.git",
		  "clone_url": "https://github.com/feed3r/21Updater.git",
		  "svn_url": "https://github.com/feed3r/21Updater",
		  "homepage": null,
		  "size": 9805,
		  "stargazers_count": 0,
		  "watchers_count": 0,
		  "language": null,
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
		  "open_issues_count": 3,
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
		  "open_issues": 3,
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
		  "href": "https://api.github.com/repos/feed3r/21Updater/pulls/3"
		},
		"html": {
		  "href": "https://github.com/feed3r/21Updater/pull/3"
		},
		"issue": {
		  "href": "https://api.github.com/repos/feed3r/21Updater/issues/3"
		},
		"comments": {
		  "href": "https://api.github.com/repos/feed3r/21Updater/issues/3/comments"
		},
		"review_comments": {
		  "href": "https://api.github.com/repos/feed3r/21Updater/pulls/3/comments"
		},
		"review_comment": {
		  "href": "https://api.github.com/repos/feed3r/21Updater/pulls/comments{/number}"
		},
		"commits": {
		  "href": "https://api.github.com/repos/feed3r/21Updater/pulls/3/commits"
		},
		"statuses": {
		  "href": "https://api.github.com/repos/feed3r/21Updater/statuses/48e646db0ca59051f0986c1f448608098f3e7ca4"
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
	  "commits": 18,
	  "additions": 866,
	  "deletions": 1,
	  "changed_files": 14
	},
	"before": "5f76029e3778208ca57422b6120f04af3a2b7662",
	"after": "48e646db0ca59051f0986c1f448608098f3e7ca4",
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
	  "updated_at": "2023-11-16T09:48:58Z",
	  "pushed_at": "2023-12-09T14:39:33Z",
	  "git_url": "git://github.com/feed3r/21Updater.git",
	  "ssh_url": "git@github.com:feed3r/21Updater.git",
	  "clone_url": "https://github.com/feed3r/21Updater.git",
	  "svn_url": "https://github.com/feed3r/21Updater",
	  "homepage": null,
	  "size": 9805,
	  "stargazers_count": 0,
	  "watchers_count": 0,
	  "language": null,
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
	  "open_issues_count": 3,
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
	  "open_issues": 3,
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
