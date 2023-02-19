package userplat

import (
	"bytes"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// unused protection.
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = (*regexp.Regexp)(nil)
	_ = time.Nanosecond
)

func (p *FavoriteActionRequest) IsValid() error {
	return nil
}

func (p *FavoriteListRequest) IsValid() error {
	return nil
}

func (p *CommentActionRequest) IsValid() error {
	return nil
}

func (p *CommentListRequest) IsValid() error {
	return nil
}
