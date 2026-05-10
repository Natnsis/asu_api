package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
)

type UserHandler struct {
	DB sql.DB
}
