package useruser

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

func (p *RelationActionRequest) IsValid() error {
	return nil
}

func (p *RelationFollowListRequest) IsValid() error {
	return nil
}

func (p *RelationFollowerListRequest) IsValid() error {
	return nil
}

func (p *RelationFriendListRequest) IsValid() error {
	return nil
}