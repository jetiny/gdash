// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package passwd

import "encoding/base64"

const alphabet = "./ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

var bcEncoding = base64.NewEncoding(alphabet)

func base64Decode(src []byte) ([]byte, error) {
	numOfEquals := 4 - (len(src) % 4)
	for i := 0; i < numOfEquals; i++ {
		src = append(src, '=')
	}

	dst := make([]byte, bcEncoding.DecodedLen(len(src)))
	n, err := bcEncoding.Decode(dst, src)
	if err != nil {
		return nil, err
	}
	return dst[:n], nil
}
