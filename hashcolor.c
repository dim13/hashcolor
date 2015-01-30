/* $Id$ */
/*
 * Copyright (c) 2015 Dimitri Sokolyuk <demon@dim13.org>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

#include <stdio.h>
#include <string.h>

#define CRC24_INIT	0x0B704CEL
#define CRC24_POLY	0x1864CFBL

long
crc24(char *s)
{
	long crc;
	int i;

	for (crc = CRC24_INIT; *s; s++) {
		crc ^= *s << 0x10;
		for (i = 0; i < 8; i++) {
			crc <<= 1;
			if (crc & 0x1000000)
				crc ^= CRC24_POLY;
		}
	}

	return crc;
}

long
shade(long c)
{
	unsigned char r = c >> 0x10;
	unsigned char g = c >> 0x08;
	unsigned char b = c;

	r >>= 2;
	g >>= 2;
	b >>= 2;

	return (r << 0x10) | (g << 0x8) | b;
}

long
tint(long c)
{
	unsigned char r = c >> 0x10;
	unsigned char g = c >> 0x08;
	unsigned char b = c;

	r += (0xFF - r) >> 1;
	g += (0xFF - g) >> 1;
	b += (0xFF - b) >> 1;

	return (r << 0x10) | (g << 0x8) | b;
}

#if defined (__linux__)
#define strlcpy(d,s,l) (strncpy(d,s,(l)-strlen(d)-1),(d)[(l)-1]='\0')
#define strlcat(d,s,l) strncat(d,s,(l)-strlen(d)-1)
#endif

int
main(int argc, char **argv)
{
	char arg[256] = {};
	long color;

	while (*++argv)
		strlcat(arg, *argv, sizeof(arg));

	color = crc24(arg);

	printf("-fg #%.6x -bg #%.6x\n", tint(color), shade(color));

	return 0;
}
