package model

import (
	"fmt"
	"github.com/davy66666/rpc_service/internal/types"

	g "github.com/doug-martin/goqu/v9"
)

func UserGameBetAmountFindOne(ex g.Ex) (types.UserGameBetAmount, error) {

	var data types.UserGameBetAmount
	query, _, _ := meta.Dialect.From("user_game_bet_amounts").Select(UserGameBetAmountFields...).Where(ex).Limit(1).ToSQL()
	fmt.Println(query)
	err := meta.SqlxDb.Get(&data, query)

	return data, err
}

func GetUserGameBetAmount(ex g.Ex) ([]*types.UserGameBetAmount, error) {

	var data []*types.UserGameBetAmount
	query, _, _ := meta.Dialect.From("user_game_bet_amounts").Select(UserGameBetAmountFields...).Where(ex).ToSQL()
	fmt.Println(query)
	err := meta.SqlxDb.Select(&data, query)

	return data, err
}

func UserGameBetAmountUpdate(ex g.Ex, record g.Record) error {

	query, _, _ := meta.Dialect.Update("user_game_bet_amounts").Set(record).Where(ex).Limit(1).ToSQL()
	fmt.Println(query)
	_, err := meta.SqlxDb.Exec(query)

	return err
}
