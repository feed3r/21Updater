package test_models

const PUSH_EXPECTED_TEXT = `"Hey, brand new code has been PUSHED in our 21Updater repository by [feed3r] 
Following the commits info:

Author: "feed3r"
Message: "First Test"
Link: ""

Author
Message: "Second Test"
Link: "Second Test"`

const PUSH_HEADERS = `Request URL: http://houseoffeeder.ddns.net:9090/githubUpdate
Request method: POST
Accept: */*
Content-Type: application/x-www-form-urlencoded
User-Agent: GitHub-Hookshot/45184e4
X-GitHub-Delivery: 62be509e-f6b3-11ee-9a89-59c8095544b9
X-GitHub-Event: push
X-GitHub-Hook-ID: 468207434
X-GitHub-Hook-Installation-Target-ID: 713630393
X-GitHub-Hook-Installation-Target-Type: repository
`
const PUSH_BODY = `{
	"ref": "refs/heads/#21-manage_commits",
	"before": "2f22249a0670a5611e08a81435609307d32215b4",
	"after": "e43bc1f298cf4c44ffd72abbf1b5ecb03307e07b",
	"repository": {
	  "id": 713630393,
	  "node_id": "R_kgDOKokiuQ",
	  "name": "21Updater",
	  "full_name": "feed3r/21Updater",
	  "private": false,
	  "owner": {
		"name": "feed3r",
		"email": "feeder81@gmail.com",
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
	  "url": "https://github.com/feed3r/21Updater",
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
	  "created_at": 1698966968,
	  "updated_at": "2024-03-03T10:09:05Z",
	  "pushed_at": 1712696080,
	  "git_url": "git://github.com/feed3r/21Updater.git",
	  "ssh_url": "git@github.com:feed3r/21Updater.git",
	  "clone_url": "https://github.com/feed3r/21Updater.git",
	  "svn_url": "https://github.com/feed3r/21Updater",
	  "homepage": null,
	  "size": 7791,
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
	  "open_issues_count": 1,
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
	  "open_issues": 1,
	  "watchers": 0,
	  "default_branch": "main",
	  "stargazers": 0,
	  "master_branch": "main"
	},
	"pusher": {
	  "name": "feed3r",
	  "email": "feeder81@gmail.com"
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
	},
	"created": false,
	"deleted": false,
	"forced": false,
	"base_ref": null,
	"compare": "https://github.com/feed3r/21Updater/compare/2f22249a0670...e43bc1f298cf",
	"commits": [
	  {
		"id": "498ab88c2e035c975c9f30a63fbe9c1691e1959c",
		"tree_id": "734b68cc56eebef63abffe65bd64f3df54cfc680",
		"distinct": true,
		"message": "Manage the expected string for push event\n\nIn push event a particular message format is expected",
		"timestamp": "2024-04-09T19:59:46+02:00",
		"url": "https://github.com/feed3r/21Updater/commit/498ab88c2e035c975c9f30a63fbe9c1691e1959c",
		"author": {
		  "name": "Alessio Pascucci",
		  "email": "feeder81@gmail.com",
		  "username": "feed3r"
		},
		"committer": {
		  "name": "Alessio Pascucci",
		  "email": "feeder81@gmail.com",
		  "username": "feed3r"
		},
		"added": [
		  "test/parse_commit_test.go"
		],
		"removed": [
  
		],
		"modified": [
		  "test/test_models/commit.go"
		]
	  },
	  {
		"id": "07d21b87ce9680c08f3bc5ad4d2a366272dc7934",
		"tree_id": "5e30c0f8359b343a099c094c5038aa8c2428699d",
		"distinct": true,
		"message": "Pass logger to Parser functions\n\nPass a reference to logger in the Parser function, so that an event can be logged",
		"timestamp": "2024-04-09T22:50:13+02:00",
		"url": "https://github.com/feed3r/21Updater/commit/07d21b87ce9680c08f3bc5ad4d2a366272dc7934",
		"author": {
		  "name": "Alessio Pascucci",
		  "email": "feeder81@gmail.com",
		  "username": "feed3r"
		},
		"committer": {
		  "name": "Alessio Pascucci",
		  "email": "feeder81@gmail.com",
		  "username": "feed3r"
		},
		"added": [
  
		],
		"removed": [
  
		],
		"modified": [
		  "src/engine/github_handler.go",
		  "src/engine/issue.go",
		  "src/engine/pull_request.go",
		  "test/parse_issue_test.go",
		  "test/parse_pr_test.go"
		]
	  },
	  {
		"id": "e43bc1f298cf4c44ffd72abbf1b5ecb03307e07b",
		"tree_id": "e4053e9fc12085be69b4c29e90aaff3a1fd51075",
		"distinct": true,
		"message": "Parse push events DRAFT\n\nParse a PUSH event, building the message structure and the TELEGRAM message.\n\nActually, this implementation should be completed, but I need to push this to obtain a valid push message for testing",
		"timestamp": "2024-04-09T22:52:40+02:00",
		"url": "https://github.com/feed3r/21Updater/commit/e43bc1f298cf4c44ffd72abbf1b5ecb03307e07b",
		"author": {
		  "name": "Alessio Pascucci",
		  "email": "feeder81@gmail.com",
		  "username": "feed3r"
		},
		"committer": {
		  "name": "Alessio Pascucci",
		  "email": "feeder81@gmail.com",
		  "username": "feed3r"
		},
		"added": [
		  "test/parse_push_test.go",
		  "test/test_models/push.go"
		],
		"removed": [
		  "test/parse_commit_test.go",
		  "test/test_models/commit.go"
		],
		"modified": [
		  "src/engine/push.go",
		  "src/model/github.go",
		  "test/utils_test.go"
		]
	  }
	],
	"head_commit": {
	  "id": "e43bc1f298cf4c44ffd72abbf1b5ecb03307e07b",
	  "tree_id": "e4053e9fc12085be69b4c29e90aaff3a1fd51075",
	  "distinct": true,
	  "message": "Parse push events DRAFT\n\nParse a PUSH event, building the message structure and the TELEGRAM message.\n\nActually, this implementation should be completed, but I need to push this to obtain a valid push message for testing",
	  "timestamp": "2024-04-09T22:52:40+02:00",
	  "url": "https://github.com/feed3r/21Updater/commit/e43bc1f298cf4c44ffd72abbf1b5ecb03307e07b",
	  "author": {
		"name": "Alessio Pascucci",
		"email": "feeder81@gmail.com",
		"username": "feed3r"
	  },
	  "committer": {
		"name": "Alessio Pascucci",
		"email": "feeder81@gmail.com",
		"username": "feed3r"
	  },
	  "added": [
		"test/parse_push_test.go",
		"test/test_models/push.go"
	  ],
	  "removed": [
		"test/parse_commit_test.go",
		"test/test_models/commit.go"
	  ],
	  "modified": [
		"src/engine/push.go",
		"src/model/github.go",
		"test/utils_test.go"
	  ]
	}
  }`
