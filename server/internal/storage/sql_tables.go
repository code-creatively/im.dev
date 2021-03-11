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

	"sessions": `CREATE TABLE IF NOT EXISTS sessions (
			sid              VARCHAR(255) primary key,   
			user_id          VARCHAR(255)
		);
	`,

	"story": `CREATE TABLE IF NOT EXISTS story (
		id 			VARCHAR(255) PRIMARY KEY,
		type        VARCHAR(1) NOT NULL,
		creator 	VARCHAR(255) NOT NULL,
		slug 		VARCHAR(64) DEFAULT '',
		title 		VARCHAR(255) DEFAULT '',
		md   		TEXT DEFAULT '',
		url  		VARCHAR(255),
		cover 		VARCHAR(255),
		brief 		TEXT,
		status 		tinyint NOT NULL,
		created 	DATETIME NOT NULL,
		updated 	DATETIME
	);
	CREATE INDEX IF NOT EXISTS story_type
		ON story (type);
	CREATE INDEX IF NOT EXISTS story_creator
		ON story (creator);
	CREATE INDEX IF NOT EXISTS story_created
		ON story (created);
	`,

	"likes": `CREATE TABLE IF NOT EXISTS likes (
		story_id       	 VARCHAR(255),
		story_type       VARCHAR(1),
		user_id          VARCHAR(255),
		created          DATETIME NOT NULL
	);
	CREATE INDEX IF NOT EXISTS likes_userid
		ON likes (user_id);
	CREATE INDEX IF NOT EXISTS likes_storyid
		ON likes (story_id);
	`,

	"likes_count": `CREATE TABLE IF NOT EXISTS likes_count (
		story_id       	 VARCHAR(255) PRIMARY KEY,
		count            INTEGER
	);
	`,

	"follows": `CREATE TABLE IF NOT EXISTS follows (
		user_id       	 VARCHAR(255),
		target_id        VARCHAR(255),
		target_type       VARCHAR(1),
		weight			 TINYINT DEFAULT 1,
		created          DATETIME NOT NULL
	);
	CREATE INDEX IF NOT EXISTS follows_userid
		ON follows (user_id);
	CREATE INDEX IF NOT EXISTS follows_targetid
		ON follows (target_id);
	`,

	"follows_count": `CREATE TABLE IF NOT EXISTS follows_count (
		target_id       	VARCHAR(255) PRIMARY KEY,
		count            	INTEGER
	);
	`,

	"tags": `CREATE TABLE IF NOT EXISTS tags (
		id 		VARCHAR(255) PRIMARY KEY,
		creator VARCHAR(255) NOT NULL,
		title 	VARCHAR(255) NOT NULL,
		name  	VARCHAR(255) NOT NULL,
		icon  	VARCHAR(255),
		cover 	VARCHAR(255),
		md	 	TEXT,
		created DATETIME NOT NULL,
		updated DATETIME
	);
	CREATE UNIQUE INDEX IF NOT EXISTS tags_name
		ON tags (name);
	CREATE INDEX IF NOT EXISTS tags_created
		ON tags (created);
	`,

	"tags_using": `CREATE TABLE IF NOT EXISTS tags_using (
		tag_id           VARCHAR(255), 
		target_type      VARCHAR(1),
		target_id        VARCHAR(255),
		target_creator  VARCHAR(255)
	);
	CREATE INDEX IF NOT EXISTS tags_using_tagid
		ON tags_using (tag_id);
	CREATE INDEX IF NOT EXISTS tags_using_targetid
		ON tags_using (target_id);
	CREATE INDEX IF NOT EXISTS tags_using_idtype
		ON tags_using (tag_id,target_type);
	`,

	"comments": `CREATE TABLE IF NOT EXISTS comments (
		id           VARCHAR(255) PRIMARY KEY, 
		story_id    VARCHAR(255),
		creator      VARCHAR(255),
		MD           TEXT,
		created DATETIME NOT NULL,
		updated DATETIME
	);
	CREATE INDEX IF NOT EXISTS comments_storyid
		ON comments (story_id);
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

	"series_post": `CREATE TABLE IF NOT EXISTS series_post (
		series_id       VARCHAR(255), 
		post_id         VARCHAR(255),
		priority        TINYINT
	);
	CREATE INDEX IF NOT EXISTS series_post_seriesid
		ON series_post (series_id);
	CREATE INDEX IF NOT EXISTS series_post_postid
		ON series_post (post_id);
	`,
}
