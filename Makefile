# $Id$

PROG=hashcolor
BINDIR=/usr/local/bin
NOMAN=

#afterinstall:
patch:
	patch `which uxterm` uxterm.patch

.include <bsd.prog.mk>
