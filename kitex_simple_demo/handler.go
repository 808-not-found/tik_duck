package main

import (
	"context"

	minimal_demo "github.com/808-not-found/tik_duck/kitex_simple_demo/kitex_gen/minimal_demo"
)

// AddServiceImpl implements the last service interface defined in the IDL.
type AddServiceImpl struct{}

// Add implements the AddServiceImpl interface.
func (s *AddServiceImpl) Add(
	ctx context.Context,
	req *minimal_demo.AddRequest,
) (resp *minimal_demo.AddResponse, err error) {
	// TODO: Your code here...
	resp = &minimal_demo.AddResponse{Res: req.A + req.B}
	return
}
