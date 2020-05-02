package keeper

import (
	"fmt"

/*	"github.com/tendermint/tendermint/crypto"*/
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/changtong1996/djzh/x/djzh/internal/types"
)

// Keeper of the djzh store
type Keeper struct {
	CoinKeeper bank.Keeper
	storeKey   sdk.StoreKey
	cdc        *codec.Codec
}

// NewKeeper creates a djzh keeper
func NewKeeper(coinKeeper bank.Keeper, key sdk.StoreKey, cdc *codec.Codec) Keeper {
	keeper := Keeper{
		CoinKeeper: coinKeeper,
		storeKey:   key,
		cdc:        cdc,
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}



func (k Keeper) GetArticle(ctx sdk.Context, article_id string) (types.Article, error) {
	store := ctx.KVStore(k.storeKey)
	var article types.Article
	byteKey := []byte(article_id)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &article)
	if err != nil {
		return article, err
	}
	return article, nil
}

func (k Keeper) SetArticle(ctx sdk.Context, article types.Article ) {
	article_id := article.Article_id
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(article)
	key := []byte(article_id)
	store.Set(key, bz)
}

func (k Keeper) DeleteArticle(ctx sdk.Context, article_id string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(article_id))
}




func (k Keeper) GetComment(ctx sdk.Context, comment_id string) (types.Comment, error) {
	store := ctx.KVStore(k.storeKey)
	var comment types.Comment
	byteKey := []byte(comment_id)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &comment)
	if err != nil {
		return comment, err
	}
	return comment, nil
}

func (k Keeper) SetComment(ctx sdk.Context, comment types.Comment ) {
	comment_id := comment.Comment_id
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(comment)
	key := []byte(comment_id)
	store.Set(key, bz)
}

func (k Keeper) DeleteComment(ctx sdk.Context, comment_id string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(comment_id))
}





// Get returns the pubkey from the adddress-pubkey relation
func (k Keeper) GetReturnVisit(ctx sdk.Context, return_visit_id string) (types.ReturnVisit, error) {
	store := ctx.KVStore(k.storeKey)
	var rv types.ReturnVisit
	byteKey := []byte(return_visit_id)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &rv)
	if err != nil {
		return rv, err
	}
	return rv, nil
}

func (k Keeper) SetReturnVisit(ctx sdk.Context, rv types.ReturnVisit ) {
	return_visit_id := rv.Return_visit_id
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(rv)
	key := []byte(return_visit_id)
	store.Set(key, bz)
}

func (k Keeper) DeleteReturnVisit(ctx sdk.Context, return_visit_id string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(return_visit_id))
}







// Get returns the pubkey from the adddress-pubkey relation
func (k Keeper) GetAVote(ctx sdk.Context, article_id string) (types.ArticleVote, error) {
	store := ctx.KVStore(k.storeKey)
	var aVote types.ArticleVote
	byteKey := []byte(article_id)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &aVote)
	if err != nil {
		return aVote, err
	}
	return aVote, nil
}

func (k Keeper) SetAVote(ctx sdk.Context, aVote types.ArticleVote ) {
	article_id := aVote.Article_id
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(aVote)
	key := []byte(article_id)
	store.Set(key, bz)
}

func (k Keeper) DeleteAVote(ctx sdk.Context, article_id string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(article_id))
}






// Get returns the pubkey from the adddress-pubkey relation
func (k Keeper) GetCVote(ctx sdk.Context, comment_id string) (types.CommentVote, error) {
	store := ctx.KVStore(k.storeKey)
	var cVote types.CommentVote
	byteKey := []byte(comment_id)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &cVote)
	if err != nil {
		return cVote, err
	}
	return cVote, nil
}

func (k Keeper) SetCVote(ctx sdk.Context, cVote types.CommentVote ) {
	comment_id := cVote.Comment_id
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(cVote)
	key := []byte(comment_id)
	store.Set(key, bz)
}

func (k Keeper) DeleteCVote(ctx sdk.Context, comment_id string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(comment_id))
}