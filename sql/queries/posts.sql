-- name: CreatePost :one
INSERT INTO posts (
	id,
	title,
	url,
	description,
	published_at,
	feed_id	
) VALUES ( $1, $2, $3, $4, $5, $6 )
ON CONFLICT (url) DO NOTHING
RETURNING *;

-- name: GetPostForUser :many
SELECT posts.* 
FROM posts
INNER JOIN feeds ON posts.feed_id = feeds.id
INNER JOIN feed_follows ON feeds.id = feed_follows.feed_id
WHERE feed_follows.user_id = $1
ORDER BY posts.published_at DESC
LIMIT $2;
