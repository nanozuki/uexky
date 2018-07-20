package resolver

import (
	"context"
	"errors"

	"gitlab.com/abyss.club/uexky/model"
)

// SliceInfoResolver ...
type SliceInfoResolver struct {
	SliceInfo *model.SliceInfo
}

// FirstCursor ...
func (si *SliceInfoResolver) FirstCursor(ctx context.Context) (string, error) {
	return si.SliceInfo.FirstCursor, nil
}

// LastCursor ...
func (si *SliceInfoResolver) LastCursor(ctx context.Context) (string, error) {
	return si.SliceInfo.LastCursor, nil
}

// SliceQuery for api, different from model.SliceQuery
type SliceQuery struct {
	Before *string
	After  *string
	Limit  int32
}

// Parse to model.SliceQuery
func (sq *SliceQuery) Parse(reverse bool) (*model.SliceQuery, error) {
	if (sq.Before == nil && sq.After == nil) ||
		(sq.Before != nil && sq.After != nil) {
		return nil, errors.New("Invalid query")
	}
	if sq.Before != nil {
		return &model.SliceQuery{
			Limit:  int(sq.Limit),
			Desc:   !reverse,
			Cursor: *(sq.Before),
		}, nil
	}
	return &model.SliceQuery{
		Limit:  int(sq.Limit),
		Desc:   reverse,
		Cursor: *(sq.After),
	}, nil
}