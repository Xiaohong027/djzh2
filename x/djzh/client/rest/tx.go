package rest

import (
	/*"bytes"*/
	"net/http"
/*
	"github.com/gorilla/mux"*/

	"github.com/cosmos/cosmos-sdk/client/context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/changtong1996/djzh/x/djzh/internal/types"
)



type CreateArticleReq struct {
	BaseReq      rest.BaseReq   `json:"base_req"`
	Creator      string         `json:"creator"`           // address of the article creator
	A_text       string         `json:"a_text"`             
	A_title      string         `json:"a_title"`
	Tag          string         `json:"tag"`
	Article_id   string         `json:"article_id"`
	Tid			 string         `json:"tid"`
	Uid          string         `json:"uid"`
	A_timestamp  string         `json:"a_timestamp"`
	Reward       sdk.Coins      `json:"reward"`            // reward of the article
}

func CreateArticleHandler(cliCtx context.CLIContext) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateArticleReq

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// create the message
		msg := types.NewMsgCreateArticle(addr, req.A_text, req.A_title, req.Tag, req.Article_id, req.Tid, req.Uid, req.A_timestamp, req.Reward)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}



type CreateCommentReq struct {
	BaseReq      rest.BaseReq    `json:"base_req"`
	Creator      string          `json:"creator"`           // address of the article creator
	Comment_id   string          `json:"comment_id"`                     // id of the comment
	Article_id   string          `json:"article_id"`                     // id of the article
	Tid          string          `json:"tid"`                            // id of the transaction
	Uid          string          `json:"uid"`							 // id of the user
	C_timestamp  string          `json:"c_timestamp"`                    // timestamp of the comment 
	C_text       string          `json:"c_text"`                         // context of the comment
	Reward       sdk.Coins       `json:"reward"`
}

func CreateCommentHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateCommentReq

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// create the message
		msg := types.NewMsgCreateComment(addr, req.Comment_id, req.Article_id, req.Tid, req.Uid, req.C_timestamp, req.C_text, req.Reward)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}



type CreateReturnVisitReq struct {
	BaseReq          rest.BaseReq     `json:"base_req"`
	Creator          string           `json:"creator"`           // address of the article creator
	Return_visit_id  string           `json:"return_visit_id"`
	Article_id       string           `json:"article_id"`
	Tid              string           `json:"tid"`
	Uid              string           `json:"uid"`
	Rv_timestamp     string           `json:"rv_timestamp"`
	Rv_text			 string           `json:"rv_text"`
	Flag             string           `json:"flag"` 
	Reward           sdk.Coins        `json:"reward"`            // reward of the article

}


func CreateReturnVisitHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateReturnVisitReq

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// create the message
		msg := types.NewMsgCreateReturnVisit(addr, req.Return_visit_id, req.Article_id, req.Tid, req.Uid, req.Rv_timestamp, req.Rv_text, req.Flag, req.Reward)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}


type CreateAVoteReq struct {
	BaseReq          rest.BaseReq     `json:"base_req"`
	Creator          string           `json:"creator"`           // address of the article creator
	Article_id       string           `json:"article_id"`        // id of the article
	VoteUP           int              `json:"voteUP"`
	VoteDOWN         int              `json:"voteDOWN"`
	Num              int              `json:"num"`
}


func CreateAVoteHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateAVoteReq

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// create the message
		msg := types.NewMsgCreateAVote(addr, req.Article_id, req.VoteUP, req.VoteDOWN, req.Num)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}



type CreateCVoteReq struct {
	BaseReq          rest.BaseReq     `json:"base_req"`
	Creator          string           `json:"creator"`           // address of the comment creator
	Comment_id       string           `json:"comment_id"`        // id of the comment
	VoteUP           int              `json:"voteUP"`
	VoteDOWN         int              `json:"voteDOWN"`
	Num              int              `json:"num"`
}


func CreateCVoteHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateCVoteReq

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// create the message
		msg := types.NewMsgCreateCVote(addr, req.Comment_id, req.VoteUP, req.VoteDOWN, req.Num)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}





/*func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router) {
	// r.HandleFunc(
	// TODO: Define the Rest route ,
	// Call the function which should be executed for this route),
	// ).Methods("POST")
}*/

/*
// Action TX body
type <Action>Req struct {
	BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`
	// TODO: Define more types if needed
}

func <Action>RequestHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req <Action>Req
		vars := mux.Vars(r)

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		// TODO: Define the module tx logic for this action

		utils.WriteGenerateStdTxResponse(w, cliCtx, BaseReq, []sdk.Msg{msg})
	}
}
*/