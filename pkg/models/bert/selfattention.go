// Copyright 2020 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bert

import (
	"encoding/gob"

	"github.com/nlpodyssey/spago/ag"
	"github.com/nlpodyssey/spago/mat"
	"github.com/nlpodyssey/spago/mat/float"
	"github.com/nlpodyssey/spago/nn"
	"github.com/nlpodyssey/spago/nn/attention/multiheadattention"
	"github.com/nlpodyssey/spago/nn/normalization/layernorm"
)

var _ nn.Model = &SelfAttentionBlock{}

// SelfAttentionBlock implements the self-attention block of a BERT model.
type SelfAttentionBlock struct {
	nn.Module
	// Attention is the multi-head attention module.
	Attention *multiheadattention.Model
	// Norm is the layer normalization module.
	Norm *layernorm.Model
}

func init() {
	gob.Register(&SelfAttentionBlock{})
}

// SelfAttentionBlockConfig is the configuration of a SelfAttentionBlock.
type SelfAttentionBlockConfig struct {
	Dim        int
	NumOfHeads int
}

// NewSelfAttentionBlock creates a new SelfAttentionBlock.
func NewSelfAttentionBlock[T float.DType](c SelfAttentionBlockConfig) *SelfAttentionBlock {
	return &SelfAttentionBlock{
		Attention: multiheadattention.New[T](c.Dim, c.NumOfHeads, false, false),
		Norm:      layernorm.New[T](c.Dim, 1e-5),
	}
}

// Forward returns the output of the model.
func (m SelfAttentionBlock) Forward(xs []mat.Tensor) []mat.Tensor {
	att, _, _ := m.Attention.Forward(nil, xs, xs)

	residual := att // reuse the same slice to avoid allocation
	for i := range residual {
		residual[i] = ag.Add(xs[i], att[i])
	}

	return m.Norm.Forward(residual...)
}
