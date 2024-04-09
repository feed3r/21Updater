package test_models

const COMMIT_EXPECTED_TEXT = `"Hey, brand new code has been PUSHED in our 21Updater repository by [feed3r] 
Following the commits info:

Message: "First Test"
Link: ""

Message: "Second Test"
Link: "Second Test"`

const COMMIT_HEADERS = `Content-Length: 7347
Accept: */*
Content-Type: application/json
X-Github-Delivery: caa0788c-8ca4-11ee-96a0-f969ee385b69
X-Github-Hook-Id: 445243377
Connection: close
User-Agent: GitHub-Hookshot/3f0c934
X-Github-Event: push
X-Github-Hook-Installation-Target-Id: 713630393
X-Github-Hook-Installation-Target-Type: repository
`
const COMMIT_BODY = `{
	"after": "e55baded02bd4a2d5b2c4604ea8fc5b1f42338ec",
	"base_ref": null,
	"before": "fac80a356e400903ac11133fb1cb4c1d6c689f15",
	"commits": [
	  {
		"added": [
		  ".vscode/launch.json"
		],
		"author": {
		  "email": "feeder81@gmail.com",
		  "name": "Alessio Pascucci",
		  "username": "feed3r"
		},
		"committer": {
		  "email": "feeder81@gmail.com",
		  "name": "Alessio Pascucci",
		  "username": "feed3r"
		},
		"distinct": true,
		"id": "e55baded02bd4a2d5b2c4604ea8fc5b1f42338ec",
		"message": "second test",
		"modified": [
		  "src/__debug_bin",
		  "src/engine/github_handler.go"
		],
		"removed": [],
		"timestamp": "2023-11-26T22:31:39+01:00",
		"tree_id": "e3c97d4d7d69434b79b3f4984aef7724672f8780",
		"url": "https://github.com/feed3r/21Updater/commit/e55baded02bd4a2d5b2c4604ea8fc5b1f42338ec"
	  }
	],
	"compare": "https://github.com/feed3r/21Updater/compare/fac80a356e40...e55baded02bd",
	"created": false,
	"deleted": false,
	"forced": false,
	"head_commit": {
	  "added": [
		".vscode/launch.json"
	  ],
	  "author": {
		"email": "feeder81@gmail.com",
		"name": "Alessio Pascucci",
		"username": "feed3r"
	  },
	  "committer": {
		"email": "feeder81@gmail.com",
		"name": "Alessio Pascucci",
		"username": "feed3r"
	  },
	  "distinct": true,
	  "id": "e55baded02bd4a2d5b2c4604ea8fc5b1f42338ec",
	  "message": "second test",
	  "modified": [
		"src/__debug_bin",
		"src/engine/github_handler.go"
	  ],
	  "removed": [],
	  "timestamp": "2023-11-26T22:31:39+01:00",
	  "tree_id": "e3c97d4d7d69434b79b3f4984aef7724672f8780",
	  "url": "https://github.com/feed3r/21Updater/commit/e55baded02bd4a2d5b2c4604ea8fc5b1f42338ec"
	},
	"pusher": {
	  "email": "feeder81@gmail.com",
	  "name": "feed3r"
	},
	"ref": "refs/heads/feeder_bot_stub",
	"repository": {
	  "allow_forking": true,
	  "archive_url": "https://api.github.com/repos/feed3r/21Updater/{archive_format}{/ref}",
	  "archived": false,
	  "assignees_url": "https://api.github.com/repos/feed3r/21Updater/assignees{/user}",
	  "blobs_url": "https://api.github.com/repos/feed3r/21Updater/git/blobs{/sha}",
	  "branches_url": "https://api.github.com/repos/feed3r/21Updater/branches{/branch}",
	  "clone_url": "https://github.com/feed3r/21Updater.git",
	  "collaborators_url": "https://api.github.com/repos/feed3r/21Updater/collaborators{/collaborator}",
	  "comments_url": "https://api.github.com/repos/feed3r/21Updater/comments{/number}",
	  "commits_url": "https://api.github.com/repos/feed3r/21Updater/commits{/sha}",
	  "compare_url": "https://api.github.com/repos/feed3r/21Updater/compare/{base}...{head}",
	  "contents_url": "https://api.github.com/repos/feed3r/21Updater/contents/{+path}",
	  "contributors_url": "https://api.github.com/repos/feed3r/21Updater/contributors",
	  "created_at": 1698966968,
	  "default_branch": "main",
	  "deployments_url": "https://api.github.com/repos/feed3r/21Updater/deployments",
	  "description": "Some experiments in creating a Telegram BOT to receive the updates from a GitHub (in particular, from THIS repository)",
	  "disabled": false,
	  "downloads_url": "https://api.github.com/repos/feed3r/21Updater/downloads",
	  "events_url": "https://api.github.com/repos/feed3r/21Updater/events",
	  "fork": false,
	  "forks": 0,
	  "forks_count": 0,
	  "forks_url": "https://api.github.com/repos/feed3r/21Updater/forks",
	  "full_name": "feed3r/21Updater",
	  "git_commits_url": "https://api.github.com/repos/feed3r/21Updater/git/commits{/sha}",
	  "git_refs_url": "https://api.github.com/repos/feed3r/21Updater/git/refs{/sha}",
	  "git_tags_url": "https://api.github.com/repos/feed3r/21Updater/git/tags{/sha}",
	  "git_url": "git://github.com/feed3r/21Updater.git",
	  "has_discussions": false,
	  "has_downloads": true,
	  "has_issues": true,
	  "has_pages": false,
	  "has_projects": true,
	  "has_wiki": false,
	  "homepage": null,
	  "hooks_url": "https://api.github.com/repos/feed3r/21Updater/hooks",
	  "html_url": "https://github.com/feed3r/21Updater",
	  "id": 713630393,
	  "is_template": false,
	  "issue_comment_url": "https://api.github.com/repos/feed3r/21Updater/issues/comments{/number}",
	  "issue_events_url": "https://api.github.com/repos/feed3r/21Updater/issues/events{/number}",
	  "issues_url": "https://api.github.com/repos/feed3r/21Updater/issues{/number}",
	  "keys_url": "https://api.github.com/repos/feed3r/21Updater/keys{/key_id}",
	  "labels_url": "https://api.github.com/repos/feed3r/21Updater/labels{/name}",
	  "language": null,
	  "languages_url": "https://api.github.com/repos/feed3r/21Updater/languages",
	  "license": {
		"key": "cc0-1.0",
		"name": "Creative Commons Zero v1.0 Universal",
		"node_id": "MDc6TGljZW5zZTY=",
		"spdx_id": "CC0-1.0",
		"url": "https://api.github.com/licenses/cc0-1.0"
	  },
	  "master_branch": "main",
	  "merges_url": "https://api.github.com/repos/feed3r/21Updater/merges",
	  "milestones_url": "https://api.github.com/repos/feed3r/21Updater/milestones{/number}",
	  "mirror_url": null,
	  "name": "21Updater",
	  "node_id": "R_kgDOKokiuQ",
	  "notifications_url": "https://api.github.com/repos/feed3r/21Updater/notifications{?since,all,participating}",
	  "open_issues": 1,
	  "open_issues_count": 1,
	  "owner": {
		"avatar_url": "https://avatars.githubusercontent.com/u/4531376?v=4",
		"email": "feeder81@gmail.com",
		"events_url": "https://api.github.com/users/feed3r/events{/privacy}",
		"followers_url": "https://api.github.com/users/feed3r/followers",
		"following_url": "https://api.github.com/users/feed3r/following{/other_user}",
		"gists_url": "https://api.github.com/users/feed3r/gists{/gist_id}",
		"gravatar_id": "",
		"html_url": "https://github.com/feed3r",
		"id": 4531376,
		"login": "feed3r",
		"name": "feed3r",
		"node_id": "MDQ6VXNlcjQ1MzEzNzY=",
		"organizations_url": "https://api.github.com/users/feed3r/orgs",
		"received_events_url": "https://api.github.com/users/feed3r/received_events",
		"repos_url": "https://api.github.com/users/feed3r/repos",
		"site_admin": false,
		"starred_url": "https://api.github.com/users/feed3r/starred{/owner}{/repo}",
		"subscriptions_url": "https://api.github.com/users/feed3r/subscriptions",
		"type": "User",
		"url": "https://api.github.com/users/feed3r"
	  },
	  "private": false,
	  "pulls_url": "https://api.github.com/repos/feed3r/21Updater/pulls{/number}",
	  "pushed_at": 1701034304,
	  "releases_url": "https://api.github.com/repos/feed3r/21Updater/releases{/id}",
	  "size": 8,
	  "ssh_url": "git@github.com:feed3r/21Updater.git",
	  "stargazers": 0,
	  "stargazers_count": 0,
	  "stargazers_url": "https://api.github.com/repos/feed3r/21Updater/stargazers",
	  "statuses_url": "https://api.github.com/repos/feed3r/21Updater/statuses/{sha}",
	  "subscribers_url": "https://api.github.com/repos/feed3r/21Updater/subscribers",
	  "subscription_url": "https://api.github.com/repos/feed3r/21Updater/subscription",
	  "svn_url": "https://github.com/feed3r/21Updater",
	  "tags_url": "https://api.github.com/repos/feed3r/21Updater/tags",
	  "teams_url": "https://api.github.com/repos/feed3r/21Updater/teams",
	  "topics": [],
	  "trees_url": "https://api.github.com/repos/feed3r/21Updater/git/trees{/sha}",
	  "updated_at": "2023-11-16T09:48:58Z",
	  "url": "https://github.com/feed3r/21Updater",
	  "visibility": "public",
	  "watchers": 0,
	  "watchers_count": 0,
	  "web_commit_signoff_required": false
	},
	"sender": {
	  "avatar_url": "https://avatars.githubusercontent.com/u/4531376?v=4",
	  "events_url": "https://api.github.com/users/feed3r/events{/privacy}",
	  "followers_url": "https://api.github.com/users/feed3r/followers",
	  "following_url": "https://api.github.com/users/feed3r/following{/other_user}",
	  "gists_url": "https://api.github.com/users/feed3r/gists{/gist_id}",
	  "gravatar_id": "",
	  "html_url": "https://github.com/feed3r",
	  "id": 4531376,
	  "login": "feed3r",
	  "node_id": "MDQ6VXNlcjQ1MzEzNzY=",
	  "organizations_url": "https://api.github.com/users/feed3r/orgs",
	  "received_events_url": "https://api.github.com/users/feed3r/received_events",
	  "repos_url": "https://api.github.com/users/feed3r/repos",
	  "site_admin": false,
	  "starred_url": "https://api.github.com/users/feed3r/starred{/owner}{/repo}",
	  "subscriptions_url": "https://api.github.com/users/feed3r/subscriptions",
	  "type": "User",
	  "url": "https://api.github.com/users/feed3r"
	}
  }`
