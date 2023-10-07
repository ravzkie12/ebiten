// Copyright 2021 The Ebiten Authors
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

//go:build darwin || freebsd || linux || netbsd || openbsd

package glfw

import (
	"github.com/hajimehoshi/ebiten/v2/internal/cglfw"
)

func ToCharModsCallback(cb func(window *Window, char rune, mods ModifierKey)) CharModsCallback {
	if cb == nil {
		return nil
	}
	return func(window *cglfw.Window, char rune, mods cglfw.ModifierKey) {
		cb(theWindows.get(window), char, ModifierKey(mods))
	}
}

func ToCloseCallback(cb func(window *Window)) CloseCallback {
	if cb == nil {
		return nil
	}
	return func(window *cglfw.Window) {
		cb(theWindows.get(window))
	}
}

func ToDropCallback(cb func(window *Window, names []string)) DropCallback {
	if cb == nil {
		return nil
	}
	return func(window *cglfw.Window, names []string) {
		cb(theWindows.get(window), names)
	}
}

func ToFramebufferSizeCallback(cb func(window *Window, width int, height int)) FramebufferSizeCallback {
	if cb == nil {
		return nil
	}
	return func(window *cglfw.Window, width int, height int) {
		cb(theWindows.get(window), width, height)
	}
}

func ToMonitorCallback(cb func(monitor *Monitor, event PeripheralEvent)) MonitorCallback {
	if cb == nil {
		return nil
	}
	return func(monitor *cglfw.Monitor, event cglfw.PeripheralEvent) {
		var m *Monitor
		if monitor != nil {
			m = &Monitor{monitor}
		}
		cb(m, PeripheralEvent(event))
	}
}

func ToScrollCallback(cb func(window *Window, xoff float64, yoff float64)) ScrollCallback {
	if cb == nil {
		return nil
	}
	return func(window *cglfw.Window, xoff float64, yoff float64) {
		cb(theWindows.get(window), xoff, yoff)
	}
}

func ToSizeCallback(cb func(window *Window, width int, height int)) SizeCallback {
	if cb == nil {
		return nil
	}
	return func(window *cglfw.Window, width, height int) {
		cb(theWindows.get(window), width, height)
	}
}