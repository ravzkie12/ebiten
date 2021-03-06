// Copyright 2018 The Ebiten Authors
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

package audio

import (
	"io"
)

type (
	dummyDriver struct{}
	dummyPlayer struct{}
)

func (d *dummyDriver) NewPlayer() io.WriteCloser {
	return &dummyPlayer{}
}

func (d *dummyDriver) Close() error {
	return nil
}

func (p *dummyPlayer) Write(b []byte) (int, error) {
	return len(b), nil
}

func (p *dummyPlayer) Close() error {
	return nil
}

func init() {
	writerDriverForTesting = &dummyDriver{}
}

type dummyHook struct {
	updates []func() error
}

func (h *dummyHook) OnSuspendAudio(f func()) {
}

func (h *dummyHook) OnResumeAudio(f func()) {
}

func (h *dummyHook) AppendHookOnBeforeUpdate(f func() error) {
	h.updates = append(h.updates, f)
}

func init() {
	hookForTesting = &dummyHook{}
}

func UpdateForTesting() error {
	for _, f := range hookForTesting.(*dummyHook).updates {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}

func PlayersNumForTesting() int {
	c := CurrentContext()
	c.m.Lock()
	n := len(c.players)
	c.m.Unlock()
	return n
}

func ResetContext() {
	theContext = nil
}
