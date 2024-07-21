-- name: CreateUser :one
INSERT INTO users(id, email, username)
VALUES (@user_id, @email, @username)
RETURNING created_at;

-- name: UserByEmail :one
SELECT * FROM users WHERE email = @email;

-- name: UserExistsByEmail :one
SELECT EXISTS (
    SELECT 1 FROM users WHERE email = @email
);

-- name: UserExistsByUsername :one
SELECT EXISTS (
    SELECT 1 FROM users WHERE username ILIKE @username
);

-- name: CreatePost :one
INSERT INTO posts (id, user_id, content)
VALUES (@post_id, @user_id, @content)
RETURNING created_at;

-- name: Posts :many
SELECT posts.*, users.username
FROM posts
INNER JOIN users ON posts.user_id = user_id
ORDER BY posts.id DESC;

-- name: Post :one
SELECT posts.*, users.username
FROM posts
INNER JOIN users ON posts.user_id = user_id
WHERE posts.id = @post_id;

-- name: CreateComment :one
INSERT INTO comments (id, user_id, post_id, content)
VALUES (@comment_id, @user_id, @post_id, @content)
RETURNING created_at;

-- name: Comments :many
SELECT comments.*, users.username
FROM comments
INNER JOIN users ON comments.user_id = user_id
WHERE comments.post_id = @post_id
ORDER BY comments.id DESC;