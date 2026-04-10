package database

import(
	"time"
	"go-chat-app/pkg/models"
)

func CreateUser(username, email, passwordHash string) (*models.User, error){
	result, err:=DB.Exec(
		"INSERT INTO users(username, email, password) VALUES(?, ?, ?)",
		username, 
		email,
		passwordHash,
	)
	if err!=nil{
		return nil, err 
	}
	id, err:=result.LastInsertId()
	if err!=nil{
		return nil, err
	}

	user := &models.User{
		ID : int(id),
		UserName : username,
		Email : email,
		Password : passwordHash,
		CreatedAt : time.Now(),
	}
	
	return user, nil
}

func GetUserByUserName(username string) (*models.User, error){
	user := &models.User{}
	err := DB.QueryRow(
		"SELECT id, username, email, password, created_at FROM users WHERE username=?", username,
		).Scan(&user.ID, &user.UserName, &user.Email, &user.Password, &user.CreatedAt)

	if err!=nil{
		return nil, err 
	}

	return user, nil
}

func GetUserById(id int) (*models.User, error){
	user := &models.User{}
	err := DB.QueryRow(
		"SELECT id, username, email, password, created_at FROM users WHERE id=?", id,
		).Scan(&user.ID, &user.UserName, &user.Email, &user.Password, &user.CreatedAt)
	
	if err!=nil{
		return nil, err  
	}
	return user, nil
}

func GetUserByEmail(email string) (*models.User, error){
	user := &models.User{}
	err := DB.QueryRow(
		"SELECT id, username, email, password, created_at FROM users WHERE email=?", email,
		).Scan(&user.ID, &user.UserName, &user.Email, &user.Password, &user.CreatedAt)
	
	if err!=nil{
		return nil, err  
	}
	return user, nil
}

func SaveMessages(userId, roomId int, username, content string) (*models.Message, error){
	result, err:=DB.Exec(
		"INSERT INTO messages (user_id, username, content, room_id) VALUES(?,?,?,?)",
		userId,
		username,
		content,
		roomId,
	)
	if err!=nil{
		return nil, err 
	}
	id, err := result.LastInsertId()
	if err!=nil{
		return nil, err 
	}
	message := &models.Message{
		ID : int(id),
		UserID : userId,
		UserName : username,
		Content : content, 
		RoomID : int(roomId),
		CreatedAt : time.Now(),
	}
	return message, nil
}

func GetMessageByRoom(roomId, limit int) ([]*models.Message, error){
	rows, err:=DB.Query(
		"SELECT id, user_id, username, content, room_id, created_at FROM messages WHERE roomId = ? ORDER BY created_at DESC LIMIT ?",
		roomId, 
		limit,
	)
	if err!=nil{
		return nil,err 
	}
	defer rows.Close()
	
	var messages []*models.Message
	for rows.Next(){
		msg := &models.Message{}
		err := rows.Scan(&msg.ID, &msg.UserID, &msg.UserName, &msg.Content, &msg.RoomID, &msg.CreatedAt)
		if err!=nil{
			return nil, err
		}
		messages = append(messages, msg)
	}

	for i,j := 0, len(messages)-1; i < j; i,j = i+1, j-1{
		messages[j], messages[i] = messages[i], messages[j]
	}

	return messages, nil
}