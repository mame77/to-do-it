package user

import "time"

// User はユーザー情報を保持します
type User struct {
    ID        int       `json:"id"`
    Username  string    `json:"username"`
    Email     string    `json:"email"`
    PasswordHash string `json:"-"` // パスワードハッシュはJSON出力しない
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}