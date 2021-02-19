package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers cosmos-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	// this line is used by starport scaffolding # 1
	r.HandleFunc("/cosmos/vote", createVoteHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/cosmos/vote", listVoteHandler(cliCtx, "cosmos")).Methods("GET")
	r.HandleFunc("/cosmos/vote/{key}", getVoteHandler(cliCtx, "cosmos")).Methods("GET")
	r.HandleFunc("/cosmos/vote", setVoteHandler(cliCtx)).Methods("PUT")
	r.HandleFunc("/cosmos/vote", deleteVoteHandler(cliCtx)).Methods("DELETE")

	r.HandleFunc("/cosmos/poll", createPollHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/cosmos/poll", listPollHandler(cliCtx, "cosmos")).Methods("GET")
	r.HandleFunc("/cosmos/poll/{key}", getPollHandler(cliCtx, "cosmos")).Methods("GET")
	r.HandleFunc("/cosmos/poll", setPollHandler(cliCtx)).Methods("PUT")
	r.HandleFunc("/cosmos/poll", deletePollHandler(cliCtx)).Methods("DELETE")

}
