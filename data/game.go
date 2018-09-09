package data

import (
	"time"
)

type Game struct {
	Id        int
	Uuid1     string
	Uuid2	  string
	UserId1    int
	UserId2		int

	CreatedAt time.Time
	JoinAt time.Time
	CloseAt time.Time

	LeftCards string
	User1Cards string
	User2Cards string
	Status int
	Result int
}
// format the CreatedAt date to display nicely on the screen
func (game *Game) CreatedAtDate() string {
	return game.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

func (post *Post) CreatedAtDate() string {
	return post.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

// get the number of posts in a thread
func (game *Game) IsPlayer(user_id int) (bool) {
	rows, err := Db.Query("SELECT count(*) FROM games where user1_id = ? or user2_id = ?", user_id , user_id)
	if err != nil {
		return false
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return false // ???
		}
	}

	if count > 0 {
		return true
	}else {
		return false
	}
}

// Create a new thread
func (user *User) CreateGame() (game Game, err error) {
	statement := "insert into games (uuid1, user_id1, created_at,status) values (?, ?, ?, ?) returning uuid1, user_id1, created_at, status"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return Game{}, err
	}
	defer stmt.Close()

	rst, err := stmt.Exec(user.Uuid, user.Id, time.Now(), 1)
	if err != nil {
		p("User.CreateGame stmt.Exec failed.")
		return Game{}, err
	}
	lastid, err := rst.LastInsertId()
	if err != nil {
		p("User.CreateGame get lastid failed.")
		return Game{}, err
	}

	statement = "select id, uuid1, user_id1, created_at, status from games where id=?"
	stmt, err = Db.Prepare(statement)
	if err != nil {
		p("User.CreateGame query prepare failed.")
		return Game{}, err
	}
	// use QueryRow to return a row and scan the returned id into the Session struct
	err = stmt.QueryRow(lastid).Scan(&game.Id, &game.Uuid1, &game.UserId1, &game.CreatedAt, &game.Status)

	if err != nil {
		p("User.CreateGame query row failed for id=", lastid)
		return Game{}, err
	}
	return game, nil
}



// Get all new start games.
func Games() (games []Game, err error) {
	rows, err := Db.Query("SELECT id, uuid1, user_id1, created_at, status FROM games where status=1 ORDER BY created_at DESC")
	if err != nil {
		return
	}
	for rows.Next() {
		conv := Game{}
		if err = rows.Scan(&conv.Id, &conv.Uuid1, &conv.UserId1, &conv.CreatedAt, &conv.Status); err != nil {
			return
		}
		games = append(games, conv)
	}


	rows, err := Db.Query("SELECT id, uuid1, user_id1, created_at, status FROM games where status=1 ORDER BY created_at DESC")
	if err != nil {
		return
	}
	for rows.Next() {
		conv := Game{}
		if err = rows.Scan(&conv.Id, &conv.Uuid1, &conv.UserId1, &conv.CreatedAt, &conv.Status); err != nil {
			return
		}
		games = append(games, conv)
	}



	rows.Close()
	return
}

// Get a game by the UUID
func GameByUUID(uuid string) (conv Game, err error) {
	conv = Game{}
	err = Db.QueryRow("SELECT id, uuid1, uuid2, user_id1, user_id2, created_at, join_at, close_at, left_cards, user1_cards, user2_cards, result, status FROM games WHERE uuid1 = ? or uuid2 = ?", uuid, uuid).
		Scan(&conv.Id, &conv.Uuid1, &conv.Uuid2, &conv.UserId1, &conv.UserId2, &conv.CreatedAt, &conv.JoinAt, &conv.CloseAt, &conv.LeftCards, &conv.User1Cards, &conv.User2Cards, &conv.Result, &conv.Status)
	return
}

// Get the user who started this game
func (game *Game) User() (user User) {
	user = User{}
	Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = ?", game.Uuid1).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}


