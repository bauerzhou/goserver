// Created by cgo - DO NOT EDIT

package main

import "unsafe"

import "syscall"

import _ "runtime/cgo"

type _ unsafe.Pointer

func _Cerrno(dst *error, x int32) { *dst = syscall.Errno(x) }
type _Ctype_int int32

type _Ctype_intFunc *[0]byte

type _Ctype_void [0]byte
var _Cfpvar_fp_fortytwo unsafe.Pointer

func _Cfunc_bridge_int_func(*[0]byte) _Ctype_int
