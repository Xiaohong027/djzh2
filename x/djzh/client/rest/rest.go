package rest

import (
	"fmt"
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

const(
	article_id = "article_id"
)

// RegisterRoutes registers djzh-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
/*	registerQueryRoutes(cliCtx, r)
	registerTxRoutes(cliCtx, r)*/
	r.HandleFunc(fmt.Sprintf("/djzh/articles/getarticle"), GetArticleHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/djzh/articles/createarticle"), CreateArticleHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/djzh/articles/createcomment"), CreateCommentHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/djzh/articles/createreturnvisit"), CreateReturnVisitHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/djzh/articles/createavote"), CreateAVoteHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/djzh/articles/createcvote"), CreateCVoteHandler(cliCtx)).Methods("POST")
}
