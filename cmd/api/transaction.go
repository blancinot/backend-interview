package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gustvision/backend-interview/pkg/account"
	"github.com/gustvision/backend-interview/pkg/account/dto"
	"github.com/rs/zerolog/log"
	uuid "github.com/satori/go.uuid"
)

func (h *handler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid path", http.StatusNotFound)
		return
	}

	ctx := r.Context()
	logger := log.With().Str("method", "create_transaction").Logger()

	var req dto.CreateTransactionReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || len(req.AccountID) == 0 {
		logger.Error().Err(err).Msg("invalid payload")
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	// # Get associated account
	acc, err := h.account.Fetch(ctx, account.Filter{ID: req.AccountID})
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch account")
		http.Error(w, "failed to fetch account", http.StatusInternalServerError)
		return
	}

	// If req.Amout < 0 -> withdrawal
	// If req.Amout > 0 -> deposit
	if acc.Total+req.Amount < 0 {
		logger.Error().Err(err).Msg("not enough money on account")
		http.Error(w, "not enough money on account", http.StatusBadRequest)
		return
	}

	// # Insert new valid transaction
	transaction := account.Transaction{
		ID:        uuid.NewV4().String(),
		Amount:    req.Amount,
		AccountID: req.AccountID,
		CreatedAt: time.Now().Unix(),
	}
	err = h.account.InsertTransaction(ctx, transaction)
	if err != nil {
		logger.Error().Err(err).Msg("failed to insert new transaction")
		http.Error(w, "failed to insert new transaction", http.StatusInternalServerError)
		return
	}

	// # Update account total
	resp := dto.CreateTransactionResp{
		Transaction: transaction,
	}
	resp.AccountTotal, err = h.account.UpdateAccountTotal(
		ctx,
		account.Filter{ID: req.AccountID},
		req.Amount,
	)
	if err != nil {
		logger.Error().Err(err).Msg("failed to update account total")
		http.Error(w, "failed to update account total", http.StatusInternalServerError)
		return
	}

	// #Marshal results.
	raw, err := json.Marshal(resp)
	if err != nil {
		logger.Error().Err(err).Msg("failed to marshal transaction response")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// #Write response
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(raw); err != nil {
		logger.Error().Err(err).Msg("failed to write response")
		return
	}

	logger.Info().Msg("success")
}

func (h *handler) GetMaxTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "invalid path", http.StatusNotFound)
		return
	}
	ctx := r.Context()
	logger := log.With().Str("method", "get_max_transaction").Logger()

	var req dto.GetMaxTransactionReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || len(req.UserID) == 0 || req.From > req.To {
		logger.Error().Err(err).Msg("invalid payload")
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	transaction, err := h.account.FetchMaxTransaction(
		ctx,
		account.FilterTransaction{UserID: req.UserID},
		req.From,
		req.To)
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch max transaction")
		http.Error(w, "failed to fetch max transaction", http.StatusInternalServerError)
		return
	}

	// #Marshal results.
	raw, err := json.Marshal(transaction)
	if err != nil {
		logger.Error().Err(err).Msg("failed to marshal transaction")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// #Write response
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(raw); err != nil {
		logger.Error().Err(err).Msg("failed to write response")
		return
	}

	logger.Info().Msg("success")
}
