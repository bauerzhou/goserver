
#include "runtime.h"
#include "cgocall.h"

void ·_Cerrno(void*, int32);

void
·_Cfunc_GoString(int8 *p, String s)
{
	s = runtime·gostring((byte*)p);
	FLUSH(&s);
}

void
·_Cfunc_GoStringN(int8 *p, int32 l, String s)
{
	s = runtime·gostringn((byte*)p, l);
	FLUSH(&s);
}

void
·_Cfunc_GoBytes(int8 *p, int32 l, Slice s)
{
	s = runtime·gobytes((byte*)p, l);
	FLUSH(&s);
}

void
·_Cfunc_CString(String s, int8 *p)
{
	p = runtime·cmalloc(s.len+1);
	runtime·memmove((byte*)p, s.str, s.len);
	p[s.len] = 0;
	FLUSH(&p);
}

void
·_Cfunc__CMalloc(uintptr n, int8 *p)
{
	p = runtime·cmalloc(n);
	FLUSH(&p);
}
#pragma cgo_import_static fortytwo
extern byte *fortytwo;
void *·_Cfpvar_fp_fortytwo = fortytwo;


#pragma cgo_import_static _cgo_dbaa5be8c271_Cfunc_bridge_int_func
void _cgo_dbaa5be8c271_Cfunc_bridge_int_func(void*);

void
·_Cfunc_bridge_int_func(struct{void *y[2];}p)
{
	runtime·cgocall(_cgo_dbaa5be8c271_Cfunc_bridge_int_func, &p);
}

