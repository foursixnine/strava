// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package Strava provides constants for using OAuth2 to access Strava.
package strava

import (
	"golang.org/x/oauth2"
)

// Endpoint is Strava's OAuth 2.0 endpoint.
var Endpoint = oauth2.Endpoint{
            AuthURL:  "http://www.strava.com/oauth/authorize",
            TokenURL: "http://www.strava.com/oauth/token",
}

