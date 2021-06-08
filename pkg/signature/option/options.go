// Copyright 2021 The Sigstore Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package option

import (
	"context"
	"crypto"
	"io"
)

// NoOpOptionImpl implements the RPCOption, SignOption, VerifyOption interfaces as no-ops.
type NoOpOptionImpl struct{}

func (NoOpOptionImpl) ApplyContext(ctx *context.Context)                   {}
func (NoOpOptionImpl) ApplyDigest(digest *[]byte)                          {}
func (NoOpOptionImpl) ApplyCryptoSignerOpts(cryptoOpts *crypto.SignerOpts) {}
func (NoOpOptionImpl) ApplyRand(rand *io.Reader)                           {}
