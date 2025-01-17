// 版权 @2022 凹语言 作者。保留所有权利。

package compiler_wat

import (
	_ "embed"
)

//go:embed base.wat
var __base_wat_data string

// wasm 内存 1 页大小
const _WASM_PAGE_SIZE = 65536

// WASM 约定栈和内存管理
// 相关全局变量地址必须和 base_wasm.go 保持一致
const (
	// ;; heap 和 stack 状态(__heap_base 只读)
	// ;; | 0 <-- stack --> | <-- static-data --> | <-- heap --> |
	// (global $$stack_prt (mut i32) (i32.const 1024)) ;; index=0
	// (global $$heap_base i32 (i32.const 2048))       ;; index=1
	__stack_ptr_index = 0
	__heap_base_index = 1
)

// 内置函数名字
const (
	// 栈函数
	_waStackPtr   = "$$StackPtr"
	_waStackAlloc = "$$StackAlloc"

	// 堆管理函数
	_waHeapPtr   = "$$HeapPtr"
	_waHeapAlloc = "$$HeapAlloc"
	_waHeapFree  = "$$HeapFree"

	// 输出函数
	_waPrintString = "$$waPuts"
	_waPrintRune   = "$$waPrintRune"
	_waPrintInt32  = "$$waPrintI32"

	// 开始函数
	_waStart = "_start"
)

// todo: PrintI32 => waPrintI32, 以 wa 开头

const modBaseWat_wa = `
(import "wa_js_env" "waPrintI32" (func $$waPrintI32 (param $x i32)))
(import "wa_js_env" "waPrintRune" (func $$waPrintRune (param $ch i32)))
(import "wa_js_env" "waPuts" (func $$waPuts (param $str i32) (param $len i32)))

(memory $memory 1)

(export "memory" (memory $memory))
(export "_start" (func $_start))

(func $_start
	;; {{$_start/body/begin}}
	;; (call $main.init)
	(call $main)
	;; {{$_start/body/end}}
)
`

const modBaseWat_wasi = `
(import "wasi_snapshot_preview1" "fd_write"
	(func $$FdWrite (param i32 i32 i32 i32) (result i32))
)

(memory $memory 1)

(export "memory" (memory $memory))
(export "_start" (func $_start))
;; (export "main.main" (func $main.main))

;; | 0 <-- stack --> | <-- static-data --> | <-- heap --> |
(global $$stack_prt (mut i32) (i32.const 1024)) ;; index=0
(global $$heap_base i32 (i32.const 2048))       ;; index=1

(func $$StackPtr (result i32)
	(global.get $$stack_prt)
)

(func $$StackAlloc (param $size i32) (result i32)
	;; $$stack_prt -= $size
	(global.set $$stack_prt (i32.sub (global.get $$stack_prt) (local.get  $size)))
	;; return $$stack_prt
	(return (global.get $$stack_prt))
)

(func $$HeapPtr (result i32)
	(global.get $$heap_base)
)

(func $$HeapAlloc (param $size i32) (result i32)
	;; {{$$HeapAlloc/body/begin}}
	unreachable
	;; {{$$HeapAlloc/body/end}}
)

(func $$HeapFree (param $ptr i32)
	;; {{$$HeapFree/body/begin}}
	unreachable
	;; {{$$HeapFree/body/end}}
)

(func $$Puts (param $str i32) (param $len i32)
	;; {{$$Puts/body/begin}}

	(local $sp i32)
	(local $p_iov i32)
	(local $p_nwritten i32)
	(local $stdout i32)

	(local.set $sp (global.get $$stack_prt))

	(local.set $p_iov (call $$StackAlloc (i32.const 8)))

	(local.set $p_nwritten (call $$StackAlloc (i32.const 4)))

	(i32.store offset=0 align=1 (local.get $p_iov) (local.get $str))
	(i32.store offset=4 align=1 (local.get $p_iov) (local.get $len))

	(local.set $stdout (i32.const 1))

	(call $$FdWrite
		(local.get $stdout)
		(local.get $p_iov) (i32.const 1)
		(local.get $p_nwritten)
	)

	(global.set $$stack_prt (local.get $sp))
	drop

	;; {{$$Puts/body/end}}
)

(func $$waPrintRune (param $ch i32)
	;; {{$$waPrintRune/body/begin}}

	(local $sp i32)
	(local $p_ch i32)

	(local.set $sp (global.get $$stack_prt))

	(local.set $p_ch (call $$StackAlloc (i32.const 4)))
	(i32.store offset=0 align=1 (local.get $p_ch) (local.get $ch))

	(call $$Puts (local.get $p_ch) (i32.const 1))

	(global.set $$stack_prt (local.get $sp))

	;; {{$$waPrintRune/body/begin}}
)

(func $$waPrintI32 (param $x i32)
	;; if $x == 0 { print '0'; return }
	(i32.eq (local.get $x) (i32.const 0))
	if
		(call $$waPrintRune (i32.const 48)) ;; '0'
		(return)
	end

	;; if $x < 0 { $x = 0-$x; print '-'; }
	(i32.lt_s (local.get $x) (i32.const 0))
	if 
		(local.set $x (i32.sub (i32.const 0) (local.get $x)))
		(call $$waPrintRune (i32.const 45)) ;; '-'
	end

	local.get $x
	call $$$print_i32
)

(func $$$print_i32 (param $x i32)
	;; {{$$$print_i32/body/begin}}

	(local $div i32)
	(local $rem i32)

	;; if $x == 0 { print '0'; return }
	(i32.eq (local.get $x) (i32.const 0))
	if
		(return)
	end

	;; print_i32($x / 10)
	;; puchar($x%10 + '0')
	(local.set $div (i32.div_s (local.get $x) (i32.const 10)))
	(local.set $rem (i32.rem_s (local.get $x) (i32.const 10)))
	(call $$$print_i32 (local.get $div))
	(call $$waPrintRune (i32.add (local.get $rem) (i32.const 48))) ;; '0'

	;; {{$$$print_i32/body/end}}
)

(func $_start
	;; {{$_start/body/begin}}
	;; (call $main.init)
	(call $main)
	;; {{$_start/body/end}}
)
`
