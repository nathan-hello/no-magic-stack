package routes

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/nathan-hello/htmx-template/src/auth"
	"github.com/nathan-hello/htmx-template/src/components"
	"github.com/nathan-hello/htmx-template/src/db"
)

func Todo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	claims, ok := r.Context().Value(auth.ClaimsContextKey).(*auth.CustomClaims)
	if !ok {
                http.Redirect(w, r, "/signin", http.StatusSeeOther)
		return
	}
	state := components.ClientState{
		IsAuthed: ok,
	}

	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
                        http.Redirect(w, r, "/signin", http.StatusSeeOther)
			return
		}

		body := r.FormValue("body")

		if len(body) > 255 || len(body) < 3 {
			// this should be a send htmx div like in auth.go
                        http.Redirect(w, r, "/signin", http.StatusSeeOther)
			return
		}

		row, err := db.Db().InsertTodo(ctx, db.InsertTodoParams{Body: body, Username: claims.Username})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		components.TodoRow(&row).Render(r.Context(), w)
		return
	}

	if r.Method == "DELETE" {
		id := r.URL.Query().Get("id")
		parsedId, err := strconv.ParseInt(id, 10, 0)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		err = db.Db().DeleteTodo(ctx, parsedId)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == "GET" {
		todos, err := db.Db().SelectTodosByUsername(ctx, claims.Username)
		if err != nil {
			if err != sql.ErrNoRows {
                                http.Redirect(w, r, "/signup", http.StatusSeeOther)
		        	return
			}
		}
		components.Todo(state, todos).Render(r.Context(), w)
	}

}
