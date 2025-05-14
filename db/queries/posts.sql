-- name: CreatePost :exec
INSERT INTO posts (title, slug, content, author_id)
VALUES ($1, $2, $3, $4);

-- name: ListPosts :many
SELECT * FROM posts ORDER BY created_at DESC;

-- name: GetPostBySlug :one
SELECT * FROM posts WHERE slug = $1;

-- name: UpdatePost :exec
UPDATE posts SET title = $1, content = $2 WHERE id = $3 AND author_id = $4;

-- name: DeletePost :exec
DELETE FROM posts WHERE id = $1 AND author_id = $2;
