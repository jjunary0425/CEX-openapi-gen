// Package main provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package main

// Ticker defines model for ticker.
type Ticker struct {
	// Ask Lowest sell order
	Ask *float32 `json:"ask,omitempty"`

	// Bid Highest buy order
	Bid *float32 `json:"bid,omitempty"`

	// High Last 24 hours price high
	High *string `json:"high,omitempty"`

	// Last Last price
	Last *string `json:"last,omitempty"`

	// Low Last 24 hours price low
	Low *string `json:"low,omitempty"`

	// Timestamp Unix timestamp
	Timestamp *string `json:"timestamp,omitempty"`

	// Volume Last 24 hours volume
	Volume *string `json:"volume,omitempty"`

	// Volume30d Last 30 days volume
	Volume30d *string `json:"volume30d,omitempty"`
}
