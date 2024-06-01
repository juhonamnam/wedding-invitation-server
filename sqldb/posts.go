package sqldb

import (
	"fmt"
	"time"

	"github.com/juhonamnam/wedding-invitation-server/types"
)

func initializePostsTable() error {
	_, err := sqlDb.Exec(`
		CREATE TABLE IF NOT EXISTS posts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(20),
			content VARCHAR(200),
			password VARCHAR(20),
			timestamp INTEGER,
			valid BOOLEAN DEFAULT TRUE
		)
	`)
	if err != nil {
		return err
	}

	_, err = sqlDb.Exec(`
		CREATE INDEX IF NOT EXISTS posts_timestamp
		ON posts (timestamp)
	`)

	if err != nil {
		return err
	}

	_, err = sqlDb.Exec(`
		CREATE INDEX IF NOT EXISTS posts_valid
		ON posts (valid)
	`)

	return err
}

func GetPosts(offset, limit int) (*types.PostsGetResponse, error) {
	rows, err := sqlDb.Query(`
		SELECT id, name, content, timestamp
		FROM posts
		WHERE valid = TRUE
		ORDER BY timestamp DESC
		LIMIT ? OFFSET ?
	`, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	total, err := sqlDb.Query(`
		SELECT COUNT(*)
		FROM posts
		WHERE valid = TRUE
	`)
	if err != nil {
		return nil, err
	}
	defer total.Close()

	postsGetResponse := &types.PostsGetResponse{
		Posts: []types.PostGet{},
	}

	for total.Next() {
		err = total.Scan(&postsGetResponse.Total)

		if err != nil {
			return nil, err
		}
	}

	for rows.Next() {
		post := types.PostGet{}
		err := rows.Scan(&post.Id, &post.Name, &post.Content, &post.Timestamp)
		if err != nil {
			return nil, err
		}
		postsGetResponse.Posts = append(postsGetResponse.Posts, post)
	}

	return postsGetResponse, nil
}

func CreatePost(name, content, password string) error {
	_, err := sqlDb.Exec(`
		INSERT INTO posts (name, content, password, timestamp)
		VALUES (?, ?, ?, ?)
	`, name, content, password, time.Now().Unix())
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func DeletePost(id int, password string) error {
	result, err := sqlDb.Exec(`
		UPDATE posts
		SET valid = FALSE
		WHERE id = ? AND password = ?
	`, id, password)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}
