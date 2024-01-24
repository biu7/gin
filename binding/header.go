// Copyright 2022 Gin Core Team. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"net/http"
	"net/textproto"
	"reflect"
)

type headerBinding struct{}

func (headerBinding) Name() string {
	return "header"
}

func (headerBinding) Bind(req *http.Request, obj any, tag string) error {
	if tag == "" {
		tag = "header"
	}
	if err := mapHeader(obj, req.Header, tag); err != nil {
		return err
	}

	return validate(obj)
}

func mapHeader(ptr any, h map[string][]string, tag string) error {
	return mappingByPtr(ptr, headerSource(h), tag)
}

type headerSource map[string][]string

var _ setter = headerSource(nil)

func (hs headerSource) TrySet(value reflect.Value, field reflect.StructField, tagValue string, opt setOptions) (bool, error) {
	return setByForm(value, field, hs, textproto.CanonicalMIMEHeaderKey(tagValue), opt)
}
