package storage

var sqlTables = map[string]string{
	"user": `CREATE TABLE IF NOT EXISTS user (
		id VARCHAR(255) PRIMARY KEY,
		username VARCHAR(255) NOT NULL UNIQUE,
		nickname VARCHAR(255) DEFAULT '',
		avatar VARCHAR(255) DEFAULT '',
		email VARCHAR(255) UNIQUE NOT NULL,
		role  VARCHAR(20) NOT NULL,

		last_seen_at DATETIME DEFAULT CURRENT_DATETIME,
		is_diabled BOOL NOT NULL DEFAULT 'false',

		created DATETIME NOT NULL,
		updated DATETIME NOT NULL
	);
	CREATE INDEX IF NOT EXISTS user_username
		ON user (username);
	CREATE INDEX IF NOT EXISTS user_role
		ON user (role);
	CREATE INDEX IF NOT EXISTS user_email
		ON user (email);`,

	"user_profile": `CREATE TABLE IF NOT EXISTS user_profile (
			id VARCHAR(255) PRIMARY KEY,

			tagline VARCHAR(255),
			cover VARCHAR(255),
			location VARCHAR(255),
			avail_for TEXT,
			about  TEXT,
			
			website VARCHAR(255),
			twitter VARCHAR(255),
			github VARCHAR(255),
			zhihu VARCHAR(255),
			weibo VARCHAR(255),
			facebook VARCHAR(255),
			stackoverflow VARCHAR(255),
		
			updated DATETIME
		);`,

	"user_skills": `CREATE TABLE IF NOT EXISTS user_skills (
			user_id          VARCHAR(255),
			skill_id         INTEGER   
		);
	CREATE INDEX IF NOT EXISTS user_skills_userid
		ON user_skills (user_id);
	CREATE INDEX IF NOT EXISTS user_skills_skillid
		ON user_skills (skill_id);
	`,

	"sessions": `CREATE TABLE IF NOT EXISTS sessions (
			sid              VARCHAR(255) primary key,   
			user_id          VARCHAR(255)
		);
	`,

	"posts": `CREATE TABLE IF NOT EXISTS posts (
		id 			VARCHAR(255) PRIMARY KEY,
		creator 	VARCHAR(255) NOT NULL,
		slug 		VARCHAR(64) NOT NULL,
		title 		VARCHAR(255) NOT NULL,
		md   		TEXT,
		url  		VARCHAR(255),
		cover 		VARCHAR(255),
		brief 		TEXT,
		likes 		INTEGER DEFAULT 0,
		views 		INTEGER DEFAULT 0,
		comments 	INTEGER DEFAULT 0,
		status 		tinyint NOT NULL,
		created 	DATETIME NOT NULL,
		updated 	DATETIME
	);
	CREATE INDEX IF NOT EXISTS posts_creator
		ON posts (creator);
	CREATE INDEX IF NOT EXISTS posts_created
		ON posts (created);
	CREATE UNIQUE INDEX IF NOT EXISTS posts_creator_slug
		ON posts (creator, slug);
	`,

	"likes": `CREATE TABLE IF NOT EXISTS likes (
		user_id          VARCHAR(255),
		story_id       	 VARCHAR(255),
		created          DATETIME NOT NULL
	);
	CREATE INDEX IF NOT EXISTS likes_userid
		ON likes (user_id);
	CREATE INDEX IF NOT EXISTS likes_storyid
		ON likes (story_id);
	`,

	"tags": `CREATE TABLE IF NOT EXISTS tags (
		id 		VARCHAR(255) PRIMARY KEY,
		creator VARCHAR(255) NOT NULL,
		title 	VARCHAR(255) NOT NULL,
		name  	VARCHAR(255) NOT NULL,
		icon  	VARCHAR(255),
		cover 	VARCHAR(255),
		md	 	TEXT,
		follower_count INTEGER DEFAULT 0,
		created DATETIME NOT NULL,
		updated DATETIME
	);
	CREATE UNIQUE INDEX IF NOT EXISTS tags_name
		ON tags (name);
	CREATE INDEX IF NOT EXISTS tags_created
		ON tags (created);
	`,

	"tag_post": `CREATE TABLE IF NOT EXISTS tag_post (
		tag_id           VARCHAR(255), 
		post_id          VARCHAR(255)
	);
	CREATE INDEX IF NOT EXISTS tag_post_tagid
		ON tag_post (tag_id);
	CREATE INDEX IF NOT EXISTS tag_post_postid
		ON tag_post (post_id);
	`,

	"comments": `CREATE TABLE IF NOT EXISTS comments (
		id           VARCHAR(255) PRIMARY KEY, 
		target_id    VARCHAR(255),
		creator      VARCHAR(255),
		MD           TEXT,
		likes        INTEGER DEFAULT 0,
		created DATETIME NOT NULL,
		updated DATETIME
	);
	CREATE INDEX IF NOT EXISTS comments_targetid
		ON comments (target_id);
	CREATE INDEX IF NOT EXISTS comments_creator
		ON comments (creator);
	`,

	"comments_count": `CREATE TABLE IF NOT EXISTS comments_count (
		story_id 	VARCHAR(255) PRIMARY KEY, 
        count       INTEGER DEFAULT 0   
	);
	`,

	"bookmarks": `CREATE TABLE IF NOT EXISTS bookmarks (
		user_id          VARCHAR(255), 
		story_id         VARCHAR(255),
		created          DATETIME
	);
	CREATE INDEX IF NOT EXISTS bookmarks_userid
		ON bookmarks (user_id);
	`,
}
