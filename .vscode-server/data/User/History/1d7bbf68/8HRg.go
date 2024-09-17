package model

import "github.com/google/uuid"

type Post struct {
   Id      uuid.UUID `json:"id"`
   User    string `json:"user"`
   Message string `json:"message"`
   Url     string `json:"url"`
   Type    string `json:"type"`
}

type User struct {
   Username string `json:"username"`
   Password string `json:"password"`
   Age      int64  `json:"age"`
   Gender   string `json:"gender"`
}