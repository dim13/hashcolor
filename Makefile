# $Id$

PROG=hashcolor
SRCS=main.c crc24.c color.c
BINDIR=/usr/local/bin
NOMAN=

#afterinstall:
patch:
	patch `which uxterm` uxterm.diff

.include <bsd.prog.mk>