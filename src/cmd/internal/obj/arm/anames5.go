// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package arm

var cnames5 = []string{
	"NONE",
	"REG",
	"REGREG",
	"REGREG2",
	"SHIFT",
	"FREG",
	"PSR",
	"FCR",
	"RCON",
	"NCON",
	"SCON",
	"LCON",
	"LCONADDR",
	"ZFCON",
	"SFCON",
	"LFCON",
	"RACON",
	"LACON",
	"SBRA",
	"LBRA",
	"HAUTO",
	"FAUTO",
	"HFAUTO",
	"SAUTO",
	"LAUTO",
	"HOREG",
	"FOREG",
	"HFOREG",
	"SOREG",
	"ROREG",
	"SROREG",
	"LOREG",
	"PC",
	"SP",
	"HREG",
	"ADDR",
	"C_TLS_LE",
	"C_TLS_IE",
	"TEXTSIZE",
	"GOK",
	"NCLASS",
	"SCOND = (1<<4)-1",
	"SBIT = 1<<4",
	"PBIT = 1<<5",
	"WBIT = 1<<6",
	"FBIT = 1<<7",
	"UBIT = 1<<7",
	"SCOND_XOR = 14",
	"SCOND_EQ = 0 ^ C_SCOND_XOR",
	"SCOND_NE = 1 ^ C_SCOND_XOR",
	"SCOND_HS = 2 ^ C_SCOND_XOR",
	"SCOND_LO = 3 ^ C_SCOND_XOR",
	"SCOND_MI = 4 ^ C_SCOND_XOR",
	"SCOND_PL = 5 ^ C_SCOND_XOR",
	"SCOND_VS = 6 ^ C_SCOND_XOR",
	"SCOND_VC = 7 ^ C_SCOND_XOR",
	"SCOND_HI = 8 ^ C_SCOND_XOR",
	"SCOND_LS = 9 ^ C_SCOND_XOR",
	"SCOND_GE = 10 ^ C_SCOND_XOR",
	"SCOND_LT = 11 ^ C_SCOND_XOR",
	"SCOND_GT = 12 ^ C_SCOND_XOR",
	"SCOND_LE = 13 ^ C_SCOND_XOR",
	"SCOND_NONE = 14 ^ C_SCOND_XOR",
	"SCOND_NV = 15 ^ C_SCOND_XOR",
}
