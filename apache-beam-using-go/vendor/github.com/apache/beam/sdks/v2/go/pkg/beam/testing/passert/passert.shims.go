// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by starcgen. DO NOT EDIT.
// File: passert.shims.go

package passert

import (
	"context"
	"fmt"
	"io"
	"reflect"

	// Library imports
	"github.com/apache/beam/sdks/v2/go/pkg/beam/core/runtime"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/core/runtime/exec"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/core/runtime/graphx/schema"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/core/sdf"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/core/typex"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/core/util/reflectx"
)

func init() {
	runtime.RegisterFunction(failIfBadEntries)
	runtime.RegisterType(reflect.TypeOf((*diffFn)(nil)).Elem())
	schema.RegisterType(reflect.TypeOf((*diffFn)(nil)).Elem())
	runtime.RegisterType(reflect.TypeOf((*elmCountCombineFn)(nil)).Elem())
	schema.RegisterType(reflect.TypeOf((*elmCountCombineFn)(nil)).Elem())
	runtime.RegisterType(reflect.TypeOf((*errFn)(nil)).Elem())
	schema.RegisterType(reflect.TypeOf((*errFn)(nil)).Elem())
	runtime.RegisterType(reflect.TypeOf((*failFn)(nil)).Elem())
	schema.RegisterType(reflect.TypeOf((*failFn)(nil)).Elem())
	runtime.RegisterType(reflect.TypeOf((*failGBKFn)(nil)).Elem())
	schema.RegisterType(reflect.TypeOf((*failGBKFn)(nil)).Elem())
	runtime.RegisterType(reflect.TypeOf((*failKVFn)(nil)).Elem())
	schema.RegisterType(reflect.TypeOf((*failKVFn)(nil)).Elem())
	runtime.RegisterType(reflect.TypeOf((*hashFn)(nil)).Elem())
	schema.RegisterType(reflect.TypeOf((*hashFn)(nil)).Elem())
	runtime.RegisterType(reflect.TypeOf((*nonEmptyFn)(nil)).Elem())
	schema.RegisterType(reflect.TypeOf((*nonEmptyFn)(nil)).Elem())
	runtime.RegisterType(reflect.TypeOf((*sumFn)(nil)).Elem())
	schema.RegisterType(reflect.TypeOf((*sumFn)(nil)).Elem())
	reflectx.RegisterStructWrapper(reflect.TypeOf((*diffFn)(nil)).Elem(), wrapMakerDiffFn)
	reflectx.RegisterStructWrapper(reflect.TypeOf((*elmCountCombineFn)(nil)).Elem(), wrapMakerElmCountCombineFn)
	reflectx.RegisterStructWrapper(reflect.TypeOf((*errFn)(nil)).Elem(), wrapMakerErrFn)
	reflectx.RegisterStructWrapper(reflect.TypeOf((*failFn)(nil)).Elem(), wrapMakerFailFn)
	reflectx.RegisterStructWrapper(reflect.TypeOf((*failGBKFn)(nil)).Elem(), wrapMakerFailGBKFn)
	reflectx.RegisterStructWrapper(reflect.TypeOf((*failKVFn)(nil)).Elem(), wrapMakerFailKVFn)
	reflectx.RegisterStructWrapper(reflect.TypeOf((*hashFn)(nil)).Elem(), wrapMakerHashFn)
	reflectx.RegisterStructWrapper(reflect.TypeOf((*nonEmptyFn)(nil)).Elem(), wrapMakerNonEmptyFn)
	reflectx.RegisterStructWrapper(reflect.TypeOf((*sumFn)(nil)).Elem(), wrapMakerSumFn)
	reflectx.RegisterFunc(reflect.TypeOf((*func(int, int) int)(nil)).Elem(), funcMakerIntIntГInt)
	reflectx.RegisterFunc(reflect.TypeOf((*func(int, func(*int) bool) error)(nil)).Elem(), funcMakerIntIterIntГError)
	reflectx.RegisterFunc(reflect.TypeOf((*func(int, func(*string) bool) error)(nil)).Elem(), funcMakerIntIterStringГError)
	reflectx.RegisterFunc(reflect.TypeOf((*func(int, typex.T) int)(nil)).Elem(), funcMakerIntTypex۰TГInt)
	reflectx.RegisterFunc(reflect.TypeOf((*func(int) error)(nil)).Elem(), funcMakerIntГError)
	reflectx.RegisterFunc(reflect.TypeOf((*func(int) int)(nil)).Elem(), funcMakerIntГInt)
	reflectx.RegisterFunc(reflect.TypeOf((*func([]byte, func(*typex.T) bool, func(*typex.T) bool, func(t typex.T), func(t typex.T), func(t typex.T)) error)(nil)).Elem(), funcMakerSliceOfByteIterTypex۰TIterTypex۰TEmitTypex۰TEmitTypex۰TEmitTypex۰TГError)
	reflectx.RegisterFunc(reflect.TypeOf((*func([]byte, func(*typex.T) bool, func(*typex.T) bool, func(*typex.T) bool) error)(nil)).Elem(), funcMakerSliceOfByteIterTypex۰TIterTypex۰TIterTypex۰TГError)
	reflectx.RegisterFunc(reflect.TypeOf((*func([]byte, func(*typex.Z) bool) error)(nil)).Elem(), funcMakerSliceOfByteIterTypex۰ZГError)
	reflectx.RegisterFunc(reflect.TypeOf((*func(typex.X, func(*typex.Y) bool) error)(nil)).Elem(), funcMakerTypex۰XIterTypex۰YГError)
	reflectx.RegisterFunc(reflect.TypeOf((*func(typex.X, typex.Y) error)(nil)).Elem(), funcMakerTypex۰XTypex۰YГError)
	reflectx.RegisterFunc(reflect.TypeOf((*func(typex.X) error)(nil)).Elem(), funcMakerTypex۰XГError)
	reflectx.RegisterFunc(reflect.TypeOf((*func() int)(nil)).Elem(), funcMakerГInt)
	exec.RegisterEmitter(reflect.TypeOf((*func(typex.T))(nil)).Elem(), emitMakerTypex۰T)
	exec.RegisterInput(reflect.TypeOf((*func(*int) bool)(nil)).Elem(), iterMakerInt)
	exec.RegisterInput(reflect.TypeOf((*func(*string) bool)(nil)).Elem(), iterMakerString)
	exec.RegisterInput(reflect.TypeOf((*func(*typex.T) bool)(nil)).Elem(), iterMakerTypex۰T)
	exec.RegisterInput(reflect.TypeOf((*func(*typex.Y) bool)(nil)).Elem(), iterMakerTypex۰Y)
	exec.RegisterInput(reflect.TypeOf((*func(*typex.Z) bool)(nil)).Elem(), iterMakerTypex۰Z)
}

func wrapMakerDiffFn(fn any) map[string]reflectx.Func {
	dfn := fn.(*diffFn)
	return map[string]reflectx.Func{
		"ProcessElement": reflectx.MakeFunc(func(a0 []byte, a1 func(*typex.T) bool, a2 func(*typex.T) bool, a3 func(t typex.T), a4 func(t typex.T), a5 func(t typex.T)) error {
			return dfn.ProcessElement(a0, a1, a2, a3, a4, a5)
		}),
	}
}

func wrapMakerElmCountCombineFn(fn any) map[string]reflectx.Func {
	dfn := fn.(*elmCountCombineFn)
	return map[string]reflectx.Func{
		"AddInput":          reflectx.MakeFunc(func(a0 int, a1 typex.T) int { return dfn.AddInput(a0, a1) }),
		"CreateAccumulator": reflectx.MakeFunc(func() int { return dfn.CreateAccumulator() }),
		"ExtractOutput":     reflectx.MakeFunc(func(a0 int) int { return dfn.ExtractOutput(a0) }),
		"MergeAccumulators": reflectx.MakeFunc(func(a0 int, a1 int) int { return dfn.MergeAccumulators(a0, a1) }),
	}
}

func wrapMakerErrFn(fn any) map[string]reflectx.Func {
	dfn := fn.(*errFn)
	return map[string]reflectx.Func{
		"ProcessElement": reflectx.MakeFunc(func(a0 int) error { return dfn.ProcessElement(a0) }),
	}
}

func wrapMakerFailFn(fn any) map[string]reflectx.Func {
	dfn := fn.(*failFn)
	return map[string]reflectx.Func{
		"ProcessElement": reflectx.MakeFunc(func(a0 typex.X) error { return dfn.ProcessElement(a0) }),
	}
}

func wrapMakerFailGBKFn(fn any) map[string]reflectx.Func {
	dfn := fn.(*failGBKFn)
	return map[string]reflectx.Func{
		"ProcessElement": reflectx.MakeFunc(func(a0 typex.X, a1 func(*typex.Y) bool) error { return dfn.ProcessElement(a0, a1) }),
	}
}

func wrapMakerFailKVFn(fn any) map[string]reflectx.Func {
	dfn := fn.(*failKVFn)
	return map[string]reflectx.Func{
		"ProcessElement": reflectx.MakeFunc(func(a0 typex.X, a1 typex.Y) error { return dfn.ProcessElement(a0, a1) }),
	}
}

func wrapMakerHashFn(fn any) map[string]reflectx.Func {
	dfn := fn.(*hashFn)
	return map[string]reflectx.Func{
		"ProcessElement": reflectx.MakeFunc(func(a0 int, a1 func(*string) bool) error { return dfn.ProcessElement(a0, a1) }),
	}
}

func wrapMakerNonEmptyFn(fn any) map[string]reflectx.Func {
	dfn := fn.(*nonEmptyFn)
	return map[string]reflectx.Func{
		"ProcessElement": reflectx.MakeFunc(func(a0 []byte, a1 func(*typex.Z) bool) error { return dfn.ProcessElement(a0, a1) }),
	}
}

func wrapMakerSumFn(fn any) map[string]reflectx.Func {
	dfn := fn.(*sumFn)
	return map[string]reflectx.Func{
		"ProcessElement": reflectx.MakeFunc(func(a0 int, a1 func(*int) bool) error { return dfn.ProcessElement(a0, a1) }),
	}
}

type callerIntIntГInt struct {
	fn func(int, int) int
}

func funcMakerIntIntГInt(fn any) reflectx.Func {
	f := fn.(func(int, int) int)
	return &callerIntIntГInt{fn: f}
}

func (c *callerIntIntГInt) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerIntIntГInt) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerIntIntГInt) Call(args []any) []any {
	out0 := c.fn(args[0].(int), args[1].(int))
	return []any{out0}
}

func (c *callerIntIntГInt) Call2x1(arg0, arg1 any) any {
	return c.fn(arg0.(int), arg1.(int))
}

type callerIntIterIntГError struct {
	fn func(int, func(*int) bool) error
}

func funcMakerIntIterIntГError(fn any) reflectx.Func {
	f := fn.(func(int, func(*int) bool) error)
	return &callerIntIterIntГError{fn: f}
}

func (c *callerIntIterIntГError) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerIntIterIntГError) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerIntIterIntГError) Call(args []any) []any {
	out0 := c.fn(args[0].(int), args[1].(func(*int) bool))
	return []any{out0}
}

func (c *callerIntIterIntГError) Call2x1(arg0, arg1 any) any {
	return c.fn(arg0.(int), arg1.(func(*int) bool))
}

type callerIntIterStringГError struct {
	fn func(int, func(*string) bool) error
}

func funcMakerIntIterStringГError(fn any) reflectx.Func {
	f := fn.(func(int, func(*string) bool) error)
	return &callerIntIterStringГError{fn: f}
}

func (c *callerIntIterStringГError) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerIntIterStringГError) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerIntIterStringГError) Call(args []any) []any {
	out0 := c.fn(args[0].(int), args[1].(func(*string) bool))
	return []any{out0}
}

func (c *callerIntIterStringГError) Call2x1(arg0, arg1 any) any {
	return c.fn(arg0.(int), arg1.(func(*string) bool))
}

type callerIntTypex۰TГInt struct {
	fn func(int, typex.T) int
}

func funcMakerIntTypex۰TГInt(fn any) reflectx.Func {
	f := fn.(func(int, typex.T) int)
	return &callerIntTypex۰TГInt{fn: f}
}

func (c *callerIntTypex۰TГInt) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerIntTypex۰TГInt) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerIntTypex۰TГInt) Call(args []any) []any {
	out0 := c.fn(args[0].(int), args[1].(typex.T))
	return []any{out0}
}

func (c *callerIntTypex۰TГInt) Call2x1(arg0, arg1 any) any {
	return c.fn(arg0.(int), arg1.(typex.T))
}

type callerIntГError struct {
	fn func(int) error
}

func funcMakerIntГError(fn any) reflectx.Func {
	f := fn.(func(int) error)
	return &callerIntГError{fn: f}
}

func (c *callerIntГError) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerIntГError) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerIntГError) Call(args []any) []any {
	out0 := c.fn(args[0].(int))
	return []any{out0}
}

func (c *callerIntГError) Call1x1(arg0 any) any {
	return c.fn(arg0.(int))
}

type callerIntГInt struct {
	fn func(int) int
}

func funcMakerIntГInt(fn any) reflectx.Func {
	f := fn.(func(int) int)
	return &callerIntГInt{fn: f}
}

func (c *callerIntГInt) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerIntГInt) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerIntГInt) Call(args []any) []any {
	out0 := c.fn(args[0].(int))
	return []any{out0}
}

func (c *callerIntГInt) Call1x1(arg0 any) any {
	return c.fn(arg0.(int))
}

type callerSliceOfByteIterTypex۰TIterTypex۰TEmitTypex۰TEmitTypex۰TEmitTypex۰TГError struct {
	fn func([]byte, func(*typex.T) bool, func(*typex.T) bool, func(t typex.T), func(t typex.T), func(t typex.T)) error
}

func funcMakerSliceOfByteIterTypex۰TIterTypex۰TEmitTypex۰TEmitTypex۰TEmitTypex۰TГError(fn any) reflectx.Func {
	f := fn.(func([]byte, func(*typex.T) bool, func(*typex.T) bool, func(t typex.T), func(t typex.T), func(t typex.T)) error)
	return &callerSliceOfByteIterTypex۰TIterTypex۰TEmitTypex۰TEmitTypex۰TEmitTypex۰TГError{fn: f}
}

func (c *callerSliceOfByteIterTypex۰TIterTypex۰TEmitTypex۰TEmitTypex۰TEmitTypex۰TГError) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerSliceOfByteIterTypex۰TIterTypex۰TEmitTypex۰TEmitTypex۰TEmitTypex۰TГError) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerSliceOfByteIterTypex۰TIterTypex۰TEmitTypex۰TEmitTypex۰TEmitTypex۰TГError) Call(args []any) []any {
	out0 := c.fn(args[0].([]byte), args[1].(func(*typex.T) bool), args[2].(func(*typex.T) bool), args[3].(func(t typex.T)), args[4].(func(t typex.T)), args[5].(func(t typex.T)))
	return []any{out0}
}

func (c *callerSliceOfByteIterTypex۰TIterTypex۰TEmitTypex۰TEmitTypex۰TEmitTypex۰TГError) Call6x1(arg0, arg1, arg2, arg3, arg4, arg5 any) any {
	return c.fn(arg0.([]byte), arg1.(func(*typex.T) bool), arg2.(func(*typex.T) bool), arg3.(func(t typex.T)), arg4.(func(t typex.T)), arg5.(func(t typex.T)))
}

type callerSliceOfByteIterTypex۰TIterTypex۰TIterTypex۰TГError struct {
	fn func([]byte, func(*typex.T) bool, func(*typex.T) bool, func(*typex.T) bool) error
}

func funcMakerSliceOfByteIterTypex۰TIterTypex۰TIterTypex۰TГError(fn any) reflectx.Func {
	f := fn.(func([]byte, func(*typex.T) bool, func(*typex.T) bool, func(*typex.T) bool) error)
	return &callerSliceOfByteIterTypex۰TIterTypex۰TIterTypex۰TГError{fn: f}
}

func (c *callerSliceOfByteIterTypex۰TIterTypex۰TIterTypex۰TГError) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerSliceOfByteIterTypex۰TIterTypex۰TIterTypex۰TГError) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerSliceOfByteIterTypex۰TIterTypex۰TIterTypex۰TГError) Call(args []any) []any {
	out0 := c.fn(args[0].([]byte), args[1].(func(*typex.T) bool), args[2].(func(*typex.T) bool), args[3].(func(*typex.T) bool))
	return []any{out0}
}

func (c *callerSliceOfByteIterTypex۰TIterTypex۰TIterTypex۰TГError) Call4x1(arg0, arg1, arg2, arg3 any) any {
	return c.fn(arg0.([]byte), arg1.(func(*typex.T) bool), arg2.(func(*typex.T) bool), arg3.(func(*typex.T) bool))
}

type callerSliceOfByteIterTypex۰ZГError struct {
	fn func([]byte, func(*typex.Z) bool) error
}

func funcMakerSliceOfByteIterTypex۰ZГError(fn any) reflectx.Func {
	f := fn.(func([]byte, func(*typex.Z) bool) error)
	return &callerSliceOfByteIterTypex۰ZГError{fn: f}
}

func (c *callerSliceOfByteIterTypex۰ZГError) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerSliceOfByteIterTypex۰ZГError) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerSliceOfByteIterTypex۰ZГError) Call(args []any) []any {
	out0 := c.fn(args[0].([]byte), args[1].(func(*typex.Z) bool))
	return []any{out0}
}

func (c *callerSliceOfByteIterTypex۰ZГError) Call2x1(arg0, arg1 any) any {
	return c.fn(arg0.([]byte), arg1.(func(*typex.Z) bool))
}

type callerTypex۰XIterTypex۰YГError struct {
	fn func(typex.X, func(*typex.Y) bool) error
}

func funcMakerTypex۰XIterTypex۰YГError(fn any) reflectx.Func {
	f := fn.(func(typex.X, func(*typex.Y) bool) error)
	return &callerTypex۰XIterTypex۰YГError{fn: f}
}

func (c *callerTypex۰XIterTypex۰YГError) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerTypex۰XIterTypex۰YГError) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerTypex۰XIterTypex۰YГError) Call(args []any) []any {
	out0 := c.fn(args[0].(typex.X), args[1].(func(*typex.Y) bool))
	return []any{out0}
}

func (c *callerTypex۰XIterTypex۰YГError) Call2x1(arg0, arg1 any) any {
	return c.fn(arg0.(typex.X), arg1.(func(*typex.Y) bool))
}

type callerTypex۰XTypex۰YГError struct {
	fn func(typex.X, typex.Y) error
}

func funcMakerTypex۰XTypex۰YГError(fn any) reflectx.Func {
	f := fn.(func(typex.X, typex.Y) error)
	return &callerTypex۰XTypex۰YГError{fn: f}
}

func (c *callerTypex۰XTypex۰YГError) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerTypex۰XTypex۰YГError) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerTypex۰XTypex۰YГError) Call(args []any) []any {
	out0 := c.fn(args[0].(typex.X), args[1].(typex.Y))
	return []any{out0}
}

func (c *callerTypex۰XTypex۰YГError) Call2x1(arg0, arg1 any) any {
	return c.fn(arg0.(typex.X), arg1.(typex.Y))
}

type callerTypex۰XГError struct {
	fn func(typex.X) error
}

func funcMakerTypex۰XГError(fn any) reflectx.Func {
	f := fn.(func(typex.X) error)
	return &callerTypex۰XГError{fn: f}
}

func (c *callerTypex۰XГError) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerTypex۰XГError) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerTypex۰XГError) Call(args []any) []any {
	out0 := c.fn(args[0].(typex.X))
	return []any{out0}
}

func (c *callerTypex۰XГError) Call1x1(arg0 any) any {
	return c.fn(arg0.(typex.X))
}

type callerГInt struct {
	fn func() int
}

func funcMakerГInt(fn any) reflectx.Func {
	f := fn.(func() int)
	return &callerГInt{fn: f}
}

func (c *callerГInt) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerГInt) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerГInt) Call(args []any) []any {
	out0 := c.fn()
	return []any{out0}
}

func (c *callerГInt) Call0x1() any {
	return c.fn()
}

type emitNative struct {
	n   exec.ElementProcessor
	fn  any
	est *sdf.WatermarkEstimator

	ctx   context.Context
	ws    []typex.Window
	et    typex.EventTime
	value exec.FullValue
}

func (e *emitNative) Init(ctx context.Context, ws []typex.Window, et typex.EventTime) error {
	e.ctx = ctx
	e.ws = ws
	e.et = et
	return nil
}

func (e *emitNative) Value() any {
	return e.fn
}

func (e *emitNative) AttachEstimator(est *sdf.WatermarkEstimator) {
	e.est = est
}

func emitMakerTypex۰T(n exec.ElementProcessor) exec.ReusableEmitter {
	ret := &emitNative{n: n}
	ret.fn = ret.invokeTypex۰T
	return ret
}

func (e *emitNative) invokeTypex۰T(val typex.T) {
	e.value = exec.FullValue{Windows: e.ws, Timestamp: e.et, Elm: val}
	if e.est != nil {
		(*e.est).(sdf.TimestampObservingEstimator).ObserveTimestamp(e.et.ToTime())
	}
	if err := e.n.ProcessElement(e.ctx, &e.value); err != nil {
		panic(err)
	}
}

type iterNative struct {
	s  exec.ReStream
	fn any

	// cur is the "current" stream, if any.
	cur exec.Stream
}

func (v *iterNative) Init() error {
	cur, err := v.s.Open()
	if err != nil {
		return err
	}
	v.cur = cur
	return nil
}

func (v *iterNative) Value() any {
	return v.fn
}

func (v *iterNative) Reset() error {
	if err := v.cur.Close(); err != nil {
		return err
	}
	v.cur = nil
	return nil
}

func iterMakerInt(s exec.ReStream) exec.ReusableInput {
	ret := &iterNative{s: s}
	ret.fn = ret.readInt
	return ret
}

func (v *iterNative) readInt(value *int) bool {
	elm, err := v.cur.Read()
	if err != nil {
		if err == io.EOF {
			return false
		}
		panic(fmt.Sprintf("broken stream: %v", err))
	}
	*value = elm.Elm.(int)
	return true
}

func iterMakerString(s exec.ReStream) exec.ReusableInput {
	ret := &iterNative{s: s}
	ret.fn = ret.readString
	return ret
}

func (v *iterNative) readString(value *string) bool {
	elm, err := v.cur.Read()
	if err != nil {
		if err == io.EOF {
			return false
		}
		panic(fmt.Sprintf("broken stream: %v", err))
	}
	*value = elm.Elm.(string)
	return true
}

func iterMakerTypex۰T(s exec.ReStream) exec.ReusableInput {
	ret := &iterNative{s: s}
	ret.fn = ret.readTypex۰T
	return ret
}

func (v *iterNative) readTypex۰T(value *typex.T) bool {
	elm, err := v.cur.Read()
	if err != nil {
		if err == io.EOF {
			return false
		}
		panic(fmt.Sprintf("broken stream: %v", err))
	}
	*value = elm.Elm.(typex.T)
	return true
}

func iterMakerTypex۰Y(s exec.ReStream) exec.ReusableInput {
	ret := &iterNative{s: s}
	ret.fn = ret.readTypex۰Y
	return ret
}

func (v *iterNative) readTypex۰Y(value *typex.Y) bool {
	elm, err := v.cur.Read()
	if err != nil {
		if err == io.EOF {
			return false
		}
		panic(fmt.Sprintf("broken stream: %v", err))
	}
	*value = elm.Elm.(typex.Y)
	return true
}

func iterMakerTypex۰Z(s exec.ReStream) exec.ReusableInput {
	ret := &iterNative{s: s}
	ret.fn = ret.readTypex۰Z
	return ret
}

func (v *iterNative) readTypex۰Z(value *typex.Z) bool {
	elm, err := v.cur.Read()
	if err != nil {
		if err == io.EOF {
			return false
		}
		panic(fmt.Sprintf("broken stream: %v", err))
	}
	*value = elm.Elm.(typex.Z)
	return true
}

// DO NOT MODIFY: GENERATED CODE