package keeper

import (
/*	"fmt"*/

	abci "github.com/tendermint/tendermint/abci/types"

/*	"github.com/cosmos/cosmos-sdk/client"*/
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"/*
	"github.com/changtong1996/djzh/x/djzh/internal/types"*/
)

const(
	QueryVoteNum = "voteNum"
	QueryGetArticle = "getarticle"//used in /cli/query.go QueryWithData
)

// NewQuerier creates a new querier for djzh clients.不太懂这个path的作用是什么
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case QueryVoteNum:
			return queryVoteNum(ctx, k, path[1:])
		case QueryGetArticle:
			return queryGetArticle(ctx, k, path[1:])
			// TODO: Put the modules query routes
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown djzh query endpoint")
		}
	}
}

func queryVoteNum(ctx sdk.Context, k Keeper, path []string) ([]byte, error) {
	voteNum, err := k.GetAVote(ctx, path[0])
	
	if err != nil {
		return nil, err
	}

	res, err := codec.MarshalJSONIndent(k.cdc, voteNum)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryGetArticle(ctx sdk.Context, k Keeper, path []string) ([]byte, error) {
	article, err := k.GetArticle(ctx, path[0])
	
	if err != nil {
		return nil, err
	}

	res, err := codec.MarshalJSONIndent(k.cdc, article)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// TODO: Add the modules query functions
// They will be similar to the above one: queryParams()
