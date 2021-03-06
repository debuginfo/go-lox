//  MIT License
//
//  Copyright (c) 2019 Marco Pacini
//
//  Permission is hereby granted, free of charge, to any person obtaining a copy
//  of this software and associated documentation files (the "Software"), to deal
//  in the Software without restriction, including without limitation the rights
//  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//  copies of the Software, and to permit persons to whom the Software is
//  furnished to do so, subject to the following conditions:
//
//  The above copyright notice and this permission notice shall be included in all
//  copies or substantial portions of the Software.
//
//  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//  SOFTWARE.

package ast

import (
	"fmt"
	"math"
)

type Expr interface {
	Accept(ExprVisitor) error
}

type ExprVisitor interface {
	visitAssign(Assign) error
	visitBinary(Binary) error
	visitCall(Call) error
	visitGet(Get) error
	visitGrouping(Grouping) error
	visitLiteral(Literal) error
	visitLogical(Logical) error
	visitSet(Set) error
	visitUnary(Unary) error
	visitVariable(Variable) error
}

type Assign struct {
	Variable
	Token
	Expr
}

func (a Assign) Accept(visitor ExprVisitor) error {
	return visitor.visitAssign(a)
}

type Binary struct {
	Left     Expr
	Operator Token
	Right    Expr
}

func (b Binary) Accept(visitor ExprVisitor) error {
	return visitor.visitBinary(b)
}

type Call struct {
	Callee    Expr
	Arguments []Expr
}

func (c Call) Accept(visitor ExprVisitor) error {
	return visitor.visitCall(c)
}

type Get struct {
	Name   Token
	Object Expr
}

func (g Get) Accept(visitor ExprVisitor) error {
	return visitor.visitGet(g)
}

type Grouping struct {
	Expr
}

func (g Grouping) Accept(visitor ExprVisitor) error {
	return visitor.visitGrouping(g)
}

type Literal struct {
	Value interface{}
}

func (l Literal) Bool() bool {
	if b, ok := l.Value.(bool); ok {
		return b
	} else {
		if l.Value == nil {
			return false
		} else {
			return true
		}
	}
}

func (l Literal) String() string {
	if s, ok := l.Value.(string); ok {
		return s
	}

	if f, ok := l.Value.(float64); ok {
		if f == math.Trunc(f) {
			return fmt.Sprintf("%d", int64(f))
		}

		return fmt.Sprintf("%f", f)
	}

	if b, ok := l.Value.(bool); ok {
		return fmt.Sprintf("%t", b)
	}

	if l.Value == nil {
		return "nil"
	}

	return fmt.Sprintf("%v", l.Value)
}

func (l Literal) Accept(visitor ExprVisitor) error {
	return visitor.visitLiteral(l)
}

type Logical struct {
	Left     Expr
	Operator Token
	Right    Expr
}

func (l Logical) Accept(visitor ExprVisitor) error {
	return visitor.visitLogical(l)
}

type Set struct {
	Object Expr
	Name   Token
	Value  Expr
}

func (s Set) Accept(visitor ExprVisitor) error {
	return visitor.visitSet(s)
}

type Unary struct {
	Operator Token
	Right    Expr
}

func (u Unary) Accept(visitor ExprVisitor) error {
	return visitor.visitUnary(u)
}

type Variable struct {
	Token
}

func (v Variable) Accept(visitor ExprVisitor) error {
	return visitor.visitVariable(v)
}
