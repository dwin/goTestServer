package model

import (
	"time"

	u "github.com/dwin/goTestServer/app/utils"
	_ "github.com/lib/pq"
)

type Token struct {
	ID        int64
	TokenID   string
	ClientIP  string
	CreatedAt time.Time
}

func SaveNewToken(t, ip string) error {
	sStmt := "INSERT INTO token (token_id, client_ip, created) values ($1, $2, $3)"
	db := getDBconn()
	values := Token{
		TokenID:   t,
		ClientIP:  ip,
		CreatedAt: time.Now(),
	}
	pStmt, err := db.Prepare(sStmt)
	if err != nil {
		u.Log.Error().Err(err).Msg("Unable to prepare Token statement to DB")
		return err
	}
	_, err = pStmt.Exec(values.TokenID, values.ClientIP, values.CreatedAt)
	if err != nil {
		u.Log.Error().Err(err).Msg("Unable to execute statement to insert new Token to DB")
		return err
	}
	pStmt.Close()
	db.Close()
	return err
}

/*
func GetToken(t string) TokenData {
	db := getDB()
	var tokenData TokenData
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tokens"))
		bytes := b.Get([]byte(t))
		jsoniter.Unmarshal(bytes, &tokenData)
		u.Log.Info().Msgf("data: %v", tokenData)
		return nil
	})

		var json []byte
		err := jsoniter.Unmarshal(data, &json)
		if err != nil {
			u.Log.Error().Err(err).Msg("Unable to unmarshall token")
			return nil
		}

	return tokenData
}
*/
