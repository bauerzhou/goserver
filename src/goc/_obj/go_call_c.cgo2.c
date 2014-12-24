#line 3 "/Users/bauerzhou/work/goserver/src/goc/go_call_c.go"
 typedef int (*intFunc) ();

 int
 bridge_int_func(intFunc f)
 {
		return f();
 }

 int fortytwo()
 {
	    return 42;
 }



// Usual nonsense: if x and y are not equal, the type will be invalid
// (have a negative array count) and an inscrutable error will come
// out of the compiler and hopefully mention "name".
#define __cgo_compile_assert_eq(x, y, name) typedef char name[(x-y)*(x-y)*-2+1];

// Check at compile time that the sizes we use match our expectations.
#define __cgo_size_assert(t, n) __cgo_compile_assert_eq(sizeof(t), n, _cgo_sizeof_##t##_is_not_##n)

__cgo_size_assert(char, 1)
__cgo_size_assert(short, 2)
__cgo_size_assert(int, 4)
typedef long long __cgo_long_long;
__cgo_size_assert(__cgo_long_long, 8)
__cgo_size_assert(float, 4)
__cgo_size_assert(double, 8)

#include <errno.h>
#include <string.h>

void
_cgo_dbaa5be8c271_Cfunc_bridge_int_func(void *v)
{
	struct {
		intFunc p0;
		int r;
	} __attribute__((__packed__)) *a = v;
	a->r = bridge_int_func((void*)a->p0);
}

