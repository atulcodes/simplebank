// package token

// import (
// 	"testing"
// 	"time"

// 	"github.com/simplebankapp/db/util"
// 	"github.com/stretchr/testify/require"
// )

// func TestPasetoMaker(t *testing.T) {
// 	maker, err := NewPasetoMaker(util.RandomString(32))
// 	require.NoError(t, err)

// 	username := util.RandomOwner()
// 	duration := time.Minute

// 	issuedAt := time.Now()
// 	espiredAt := time.Now().Add(duration)

// 	token, err := maker.CreateToken(username, duration)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, token)

// 	payload, err := maker.VerifyToken(token)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, payload)
// 	require.NotZero(t, payload.ID)
// 	require.Equal(t, payload.Username, username)
// 	require.WithinDuration(t, payload.IssuedAt, issuedAt, time.Second) 
// 	require.WithinDuration(t, payload.ExpiredAt, espiredAt, time.Second)
// }

// func TestExpiredPasetoToken(t *testing.T) {
// 	maker, err := NewPasetoMaker(util.RandomString(32))
// 	require.NoError(t, err)

// 	token, err := maker.CreateToken(util.RandomOwner(), -time.Minute)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, token)

// 	payload, err := maker.VerifyToken(token)
// 	require.Error(t, err)
// 	require.EqualError(t, err, ErrExpiredToken.Error())
// 	require.Nil(t, payload)
// }