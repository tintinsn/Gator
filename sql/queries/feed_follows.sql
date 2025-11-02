-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
  INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
  VALUES ($1, $2, $3, $4, $5)
  RETURNING *
)
SELECT
  inserted_feed_follow.id,
  users.name  AS user_name,
  feeds.name  AS feed_name,
  inserted_feed_follow.created_at
FROM inserted_feed_follow
JOIN users ON users.id = inserted_feed_follow.user_id
JOIN feeds ON feeds.id = inserted_feed_follow.feed_id;

-- name: GetFeedByURL :one
SELECT * FROM feeds
WHERE url = $1;

-- name: GetFeedFollowsForUser :many
SELECT feed_follows.id AS id, users.name AS user_name, feeds.name AS feed_name
FROM feed_follows
JOIN users ON users.id = feed_follows.user_id
JOIN feeds ON feeds.id = feed_follows.feed_id
WHERE feed_follows.user_id = $1;


-- name: DeleteFeedFollowByUser :execrows
DELETE FROM feed_follows
WHERE user_id = $1 AND feed_id = $2;
