package main

import (
	"fmt"
	"strconv"
)

var mem = make(map[int]int64)

func applyMAsk(in int) int64 {
	strbin := strconv.FormatInt(int64(in), 2)
	if len(strbin) != len(mask) {
		zero := ""
		for i := 0; i < (len(mask) - len(strbin)); i++ {
			zero += "0"
		}
		strbin = zero + strbin
	}
	newBin := make([]rune, len(strbin))
	runeMask := []rune(mask)
	for idx, runeBin := range strbin {
		if runeMask[idx] != 'X' {
			newBin[idx] = runeMask[idx]
		} else {
			newBin[idx] = runeBin

		}

	}
	newInt, err := strconv.ParseInt(string(newBin), 2, 64)
	if err != nil {
		panic("lol")
	}
	return newInt

}

func parta() {

	mask = "00000X110010111111X000100XX01010000X"
	mem[20690] = applyMAsk(435)
	mem[54036] = applyMAsk(231)
	mem[27099] = applyMAsk(118644255)
	mem[55683] = applyMAsk(22299263)
	mem[26119] = applyMAsk(2279399)
	mask = "00X000X0001X111111101X1111XX11X001XX"
	mem[42072] = applyMAsk(1658073)
	mem[63234] = applyMAsk(2277)
	mask = "1001X010011011111110101101X0XX11X010"
	mem[31090] = applyMAsk(52291)
	mem[31244] = applyMAsk(377352406)
	mem[10621] = applyMAsk(18801757)
	mem[31666] = applyMAsk(5100853)
	mask = "10X0110X11XX101XX1000011001001010100"
	mem[18680] = applyMAsk(80608039)
	mem[13197] = applyMAsk(7957847)
	mem[17080] = applyMAsk(117501010)
	mask = "1000110011111X11X1XXXX1X000X010011X1"
	mem[25308] = applyMAsk(257586)
	mem[14518] = applyMAsk(62108102)
	mem[21633] = applyMAsk(1544993)
	mem[36955] = applyMAsk(1363)
	mem[45764] = applyMAsk(49755959)
	mem[40967] = applyMAsk(425)
	mem[47858] = applyMAsk(611686)
	mask = "0010111010X0111111011X0110X0101010X1"
	mem[7451] = applyMAsk(1208)
	mem[31918] = applyMAsk(769)
	mem[29313] = applyMAsk(1888678)
	mem[52254] = applyMAsk(32237487)
	mask = "00X001001111X11111X010000X0110XX0X11"
	mem[61531] = applyMAsk(15796066)
	mem[305] = applyMAsk(130785)
	mem[25845] = applyMAsk(197912)
	mem[29251] = applyMAsk(374061)
	mem[37177] = applyMAsk(17950)
	mask = "100100X00110111111100110001X1X100X00"
	mem[40491] = applyMAsk(66538375)
	mem[42244] = applyMAsk(240009051)
	mem[18805] = applyMAsk(33518831)
	mem[17072] = applyMAsk(518835559)
	mask = "XX100100X1101X11010001X11001100XX1XX"
	mem[16935] = applyMAsk(1124623)
	mem[45248] = applyMAsk(155461)
	mem[37224] = applyMAsk(5755511)
	mask = "00X011101110101X10X1XXX1100X0001000X"
	mem[6440] = applyMAsk(116801)
	mem[193] = applyMAsk(7318437)
	mem[58568] = applyMAsk(8082803)
	mem[43695] = applyMAsk(909697)
	mem[29001] = applyMAsk(27290)
	mem[29210] = applyMAsk(91241)
	mask = "XXX01X001111111111101010000001XX1011"
	mem[21289] = applyMAsk(354401446)
	mem[33814] = applyMAsk(1605382)
	mem[16967] = applyMAsk(242083755)
	mem[60470] = applyMAsk(22550)
	mem[16485] = applyMAsk(3945104)
	mem[37687] = applyMAsk(86474)
	mem[51031] = applyMAsk(5255)
	mask = "00100100111101111100100X0X001XX10011"
	mem[34832] = applyMAsk(191857526)
	mem[30126] = applyMAsk(180246093)
	mem[310] = applyMAsk(1895)
	mem[49300] = applyMAsk(117732)
	mask = "00000100011011X11X10010XX1XX0X1101XX"
	mem[54544] = applyMAsk(1368)
	mem[30126] = applyMAsk(596855)
	mem[18483] = applyMAsk(124319430)
	mem[63246] = applyMAsk(95337119)
	mem[3917] = applyMAsk(1620395)
	mask = "00XX010001101111X110011X010100X10XX1"
	mem[31090] = applyMAsk(203896198)
	mem[36989] = applyMAsk(203)
	mem[8762] = applyMAsk(372392)
	mem[59728] = applyMAsk(486751)
	mask = "00101X00X11X11111X100010XXXX0011011X"
	mem[59728] = applyMAsk(30591660)
	mem[43720] = applyMAsk(315507593)
	mem[39732] = applyMAsk(42157)
	mem[3440] = applyMAsk(242110717)
	mem[36955] = applyMAsk(871544)
	mem[51251] = applyMAsk(2489781)
	mask = "00X00X0X01111111110000001100100X0011"
	mem[51149] = applyMAsk(12451455)
	mem[17566] = applyMAsk(351620601)
	mem[33842] = applyMAsk(1119118)
	mem[23677] = applyMAsk(100601411)
	mem[12826] = applyMAsk(2474316)
	mask = "00X00100XX1111111110111X1X010X10010X"
	mem[43163] = applyMAsk(27012)
	mem[53314] = applyMAsk(2717910)
	mem[20842] = applyMAsk(239857)
	mem[43816] = applyMAsk(3173699)
	mem[11343] = applyMAsk(37315312)
	mem[37493] = applyMAsk(262038)
	mem[25824] = applyMAsk(13598271)
	mask = "00X0X00001101X1X10100000110001110101"
	mem[39732] = applyMAsk(1402)
	mem[50014] = applyMAsk(32437274)
	mem[10770] = applyMAsk(192187204)
	mask = "000101110110X111111011010X11000000XX"
	mem[51283] = applyMAsk(2490405)
	mem[33814] = applyMAsk(471881)
	mem[15119] = applyMAsk(3807095)
	mask = "000X00XX0X10111111X0011XX11X10X00010"
	mem[10405] = applyMAsk(46099021)
	mem[42308] = applyMAsk(1001)
	mem[57329] = applyMAsk(2310)
	mask = "11X10110011XX100XX1010101100010X0000"
	mem[40240] = applyMAsk(28185370)
	mem[43296] = applyMAsk(2212)
	mem[15632] = applyMAsk(3512122)
	mem[61953] = applyMAsk(2534700)
	mem[58797] = applyMAsk(258533)
	mask = "00000X00011X1X11X1X001000010010X1110"
	mem[13671] = applyMAsk(66116)
	mem[5234] = applyMAsk(46868488)
	mem[48068] = applyMAsk(259070)
	mem[35833] = applyMAsk(1904)
	mask = "XX1101000XX01111X1100XX10X1000011000"
	mem[58276] = applyMAsk(827)
	mem[29197] = applyMAsk(6552)
	mem[21249] = applyMAsk(173)
	mem[5723] = applyMAsk(4730123)
	mem[59627] = applyMAsk(3299104)
	mem[17008] = applyMAsk(74955518)
	mask = "0X000100111111XX11001000000110X00100"
	mem[53231] = applyMAsk(909153)
	mem[28837] = applyMAsk(1739162)
	mem[21336] = applyMAsk(3932)
	mem[32899] = applyMAsk(872661)
	mem[29051] = applyMAsk(228916)
	mask = "0X1X11000111111X101111100000X10X00XX"
	mem[23121] = applyMAsk(4940)
	mem[64259] = applyMAsk(339599819)
	mem[268] = applyMAsk(2533)
	mem[5725] = applyMAsk(1430)
	mem[56946] = applyMAsk(618)
	mask = "001X1X0001111111101001111001X101XX10"
	mem[46780] = applyMAsk(339675)
	mem[57420] = applyMAsk(10161)
	mem[32105] = applyMAsk(5534)
	mask = "X1000X10011001111110110100X01010X011"
	mem[47922] = applyMAsk(892051565)
	mem[50583] = applyMAsk(2962439)
	mem[43673] = applyMAsk(107)
	mask = "000001001X11011111101X00XX1111100111"
	mem[17938] = applyMAsk(29693823)
	mem[27809] = applyMAsk(17197)
	mem[62755] = applyMAsk(6590924)
	mem[26483] = applyMAsk(15837)
	mem[5245] = applyMAsk(486)
	mem[8213] = applyMAsk(1239)
	mask = "0010X0000XX11XX1100010X0X11000101XX1"
	mem[3842] = applyMAsk(3541)
	mem[55663] = applyMAsk(76779528)
	mem[29851] = applyMAsk(2801)
	mask = "XX011X0001111111X1XX0000000001X00010"
	mem[20066] = applyMAsk(97384)
	mem[35212] = applyMAsk(10209)
	mem[15847] = applyMAsk(499740)
	mem[9349] = applyMAsk(9638367)
	mask = "001011X00XXX11X111100X1XX00010100X01"
	mem[52845] = applyMAsk(1056563)
	mem[30126] = applyMAsk(13918626)
	mem[17709] = applyMAsk(25538089)
	mem[1413] = applyMAsk(459461)
	mem[59577] = applyMAsk(52944410)
	mask = "X1011001X0XX111011100XX010101X01X010"
	mem[56449] = applyMAsk(144)
	mem[8753] = applyMAsk(984864)
	mem[23728] = applyMAsk(173703761)
	mem[34970] = applyMAsk(28269)
	mem[32500] = applyMAsk(49931)
	mask = "X0000100111X11111110X10X0101001100X1"
	mem[64582] = applyMAsk(6646737)
	mem[37177] = applyMAsk(10)
	mem[57474] = applyMAsk(313623)
	mem[17322] = applyMAsk(147838906)
	mem[28766] = applyMAsk(15110001)
	mem[49] = applyMAsk(80836580)
	mask = "001010000X11111X10X0X111X0XX00X100X1"
	mem[53163] = applyMAsk(15243)
	mem[61002] = applyMAsk(406400)
	mem[28930] = applyMAsk(465647779)
	mask = "0010000X11101111100010X00XX1XX01XX01"
	mem[1315] = applyMAsk(625209)
	mem[44187] = applyMAsk(14395)
	mask = "X0010XX0011011111110110X0101XX0X00XX"
	mem[31859] = applyMAsk(95408)
	mem[16534] = applyMAsk(121119590)
	mem[26550] = applyMAsk(8188494)
	mem[37302] = applyMAsk(407378)
	mask = "001XX1XXX1111X1110100001X00000111001"
	mem[37574] = applyMAsk(31364)
	mem[26443] = applyMAsk(2676291)
	mem[22192] = applyMAsk(26966115)
	mask = "001X010X001X011110001101XXX0X101XX11"
	mem[41368] = applyMAsk(50472035)
	mem[25252] = applyMAsk(3850)
	mem[21011] = applyMAsk(7912441)
	mem[55890] = applyMAsk(2474497)
	mask = "0010010000X11XX111X0101X100000101000"
	mem[57489] = applyMAsk(10006848)
	mem[7880] = applyMAsk(30889)
	mem[54742] = applyMAsk(14408)
	mask = "0010100X00X1X11X101000X1X00X001100XX"
	mem[28474] = applyMAsk(137340532)
	mem[57910] = applyMAsk(3261)
	mem[35212] = applyMAsk(974067528)
	mem[24595] = applyMAsk(15641)
	mask = "0X0X01X00110X11111X0X001X1011010X101"
	mem[1515] = applyMAsk(4597)
	mem[20626] = applyMAsk(483632)
	mem[50912] = applyMAsk(101611112)
	mem[62450] = applyMAsk(463312)
	mask = "00101X000X0111011110X10X10011X100001"
	mem[5378] = applyMAsk(132014)
	mem[13345] = applyMAsk(2058543)
	mem[42684] = applyMAsk(2824)
	mem[34576] = applyMAsk(6385683)
	mem[27201] = applyMAsk(2519)
	mem[9632] = applyMAsk(202081)
	mask = "X0X1010001X011X111100001001X01XX1XX1"
	mem[1538] = applyMAsk(2389067)
	mem[4972] = applyMAsk(19131)
	mem[23129] = applyMAsk(256828081)
	mem[17188] = applyMAsk(185346747)
	mem[44295] = applyMAsk(143437003)
	mem[44830] = applyMAsk(5686)
	mem[46528] = applyMAsk(4177799)
	mask = "X0010010001X1111110011X011XX11X1X011"
	mem[7033] = applyMAsk(2748)
	mem[2431] = applyMAsk(17997007)
	mem[13924] = applyMAsk(90861)
	mem[63656] = applyMAsk(497878)
	mem[61841] = applyMAsk(891)
	mem[10405] = applyMAsk(6177)
	mem[55811] = applyMAsk(43078384)
	mask = "X110XX0011111XX1X11001X000000001X000"
	mem[62283] = applyMAsk(8553774)
	mem[14788] = applyMAsk(308418)
	mem[5878] = applyMAsk(2324)
	mask = "0000010001101111101XX1X00001001101XX"
	mem[23816] = applyMAsk(69720)
	mem[29524] = applyMAsk(197631)
	mask = "10X100100011X111110011010X0X10X10XX1"
	mem[5288] = applyMAsk(1072)
	mem[34681] = applyMAsk(1902)
	mem[47529] = applyMAsk(1012160)
	mem[42117] = applyMAsk(232642695)
	mem[7153] = applyMAsk(420427964)
	mem[23129] = applyMAsk(10261)
	mem[24545] = applyMAsk(1661292)
	mask = "X1011X010XX0X1X1X010001001001X001111"
	mem[3984] = applyMAsk(17460969)
	mem[43208] = applyMAsk(1626)
	mem[12288] = applyMAsk(3244)
	mem[1261] = applyMAsk(685777140)
	mem[35662] = applyMAsk(3875)
	mem[13197] = applyMAsk(807702837)
	mem[8450] = applyMAsk(39850899)
	mask = "001X0X00011X11111000000110010X0XXX1X"
	mem[35167] = applyMAsk(3384)
	mem[1969] = applyMAsk(3362919)
	mem[4732] = applyMAsk(99083530)
	mem[58162] = applyMAsk(1382314)
	mask = "10000X10011011X11010000X11X1X0010100"
	mem[50583] = applyMAsk(4112)
	mem[4097] = applyMAsk(907)
	mem[45785] = applyMAsk(1275731)
	mem[31108] = applyMAsk(7733)
	mem[50267] = applyMAsk(2625942)
	mask = "X0010X1001X011111110110011000111X0X0"
	mem[35870] = applyMAsk(1100551)
	mem[8514] = applyMAsk(8042956)
	mem[10848] = applyMAsk(96032)
	mem[44678] = applyMAsk(213384)
	mem[25743] = applyMAsk(3586812)
	mem[34074] = applyMAsk(991022)
	mask = "X010X1001111X1111110010X000011000011"
	mem[61953] = applyMAsk(3703)
	mem[41415] = applyMAsk(250960289)
	mem[24262] = applyMAsk(14129393)
	mask = "000X010XX11011X111100100110X001000X0"
	mem[51393] = applyMAsk(320156165)
	mem[27955] = applyMAsk(21751009)
	mem[61468] = applyMAsk(8941693)
	mem[24188] = applyMAsk(176466079)
	mem[10717] = applyMAsk(2950)
	mask = "0010XX00X11110101000X0110000X101X100"
	mem[37149] = applyMAsk(18981413)
	mem[12384] = applyMAsk(479738)
	mem[17072] = applyMAsk(5196)
	mem[59325] = applyMAsk(170080)
	mem[3269] = applyMAsk(86268393)
	mem[48598] = applyMAsk(18530)
	mem[11287] = applyMAsk(4082)
	mask = "0X101100X11X11111X10X0X01000X1110010"
	mem[47267] = applyMAsk(12410)
	mem[8609] = applyMAsk(6923289)
	mem[28364] = applyMAsk(23091829)
	mem[63780] = applyMAsk(858)
	mem[21558] = applyMAsk(48929393)
	mem[46110] = applyMAsk(74033138)
	mask = "001011X0011111111010X10100X10X000011"
	mem[30364] = applyMAsk(14013071)
	mem[23121] = applyMAsk(5777)
	mem[54108] = applyMAsk(11707710)
	mask = "0010XX00X11X11111XX010X00X000X010011"
	mem[29453] = applyMAsk(3480476)
	mem[7516] = applyMAsk(869816189)
	mem[57136] = applyMAsk(130673464)
	mem[8609] = applyMAsk(2000)
	mem[45543] = applyMAsk(1014823)
	mem[3249] = applyMAsk(75)
	mem[14460] = applyMAsk(18422415)
	mask = "000X011000111X1X111010X0X10X11X10011"
	mem[14556] = applyMAsk(7853751)
	mem[29755] = applyMAsk(535169084)
	mem[24262] = applyMAsk(4027)
	mem[34051] = applyMAsk(13187123)
	mask = "000X011X0X1X1111X110000000XX0000X11X"
	mem[13879] = applyMAsk(2383)
	mem[57329] = applyMAsk(749)
	mem[54544] = applyMAsk(3055190)
	mask = "00110XX00111111X10010001XX1000101000"
	mem[4852] = applyMAsk(429814346)
	mem[55439] = applyMAsk(7610)
	mem[31685] = applyMAsk(811508716)
	mem[38296] = applyMAsk(185763)
	mem[16482] = applyMAsk(3668)
	mem[47529] = applyMAsk(3803)
	mask = "X01011000111111X1011010101101X011X01"
	mem[58499] = applyMAsk(851439)
	mem[38516] = applyMAsk(3082)
	mem[32500] = applyMAsk(364520)
	mask = "001X110001111X10X000001011X001000101"
	mem[44653] = applyMAsk(157371860)
	mem[2226] = applyMAsk(58088617)
	mem[10098] = applyMAsk(67459)
	mem[45739] = applyMAsk(3994)
	mem[4180] = applyMAsk(206930963)
	mask = "001XX0000XX1111010X00011000001011011"
	mem[53876] = applyMAsk(843104)
	mem[56118] = applyMAsk(1019)
	mem[39503] = applyMAsk(6758)
	mem[24134] = applyMAsk(9483199)
	mem[25914] = applyMAsk(26956)
	mem[10098] = applyMAsk(63837172)
	mem[40642] = applyMAsk(2366588)
	mask = "0XX10X1001X1111X1X101010010XX01000X0"
	mem[16432] = applyMAsk(17158914)
	mem[29927] = applyMAsk(9292527)
	mem[57922] = applyMAsk(24395252)
	mem[48327] = applyMAsk(253)
	mem[15450] = applyMAsk(496726)
	mem[57027] = applyMAsk(518857449)
	mask = "0101XXX0011X11111110X0XX01X010000111"
	mem[39393] = applyMAsk(570)
	mem[38893] = applyMAsk(21253926)
	mask = "0010010X01XX1X1101100000010100110101"
	mem[64325] = applyMAsk(416581774)
	mem[26376] = applyMAsk(1666947)
	mem[6276] = applyMAsk(90042)
	mask = "0X00001101101111X11X010010XX1X1XX01X"
	mem[20354] = applyMAsk(2180)
	mem[50761] = applyMAsk(7237731)
	mem[54710] = applyMAsk(5718)
	mem[43883] = applyMAsk(2618938)
	mem[59235] = applyMAsk(22130448)
	mem[59325] = applyMAsk(14410783)
	mask = "00000001X0101X11X110101001X11100X0X1"
	mem[24262] = applyMAsk(10756242)
	mem[59282] = applyMAsk(296121)
	mem[15931] = applyMAsk(49)
	mem[44067] = applyMAsk(339152264)
	mem[22192] = applyMAsk(2750756)
	mem[897] = applyMAsk(639)
	mask = "001011X011111X1110X01X01X0XX0101X101"
	mem[21410] = applyMAsk(5056)
	mem[43472] = applyMAsk(198924166)
	mem[50343] = applyMAsk(5363196)
	mem[7486] = applyMAsk(773744)
	mem[49418] = applyMAsk(77311216)
	mask = "001X0000011011111X00000X1011X1X01XX0"
	mem[19633] = applyMAsk(5522082)
	mem[4682] = applyMAsk(51724569)
	mem[36252] = applyMAsk(260)
	mask = "001011X0111X1X1110XX000X100001X100X0"
	mem[34373] = applyMAsk(803)
	mem[61841] = applyMAsk(25585959)
	mem[29051] = applyMAsk(2011)
	mem[53885] = applyMAsk(4255251)
	mem[55135] = applyMAsk(49781551)
	mem[11748] = applyMAsk(5712)
	mask = "000001000X1011111110X11X0XX11101100X"
	mem[28073] = applyMAsk(257781932)
	mem[32292] = applyMAsk(7788)
	mem[47529] = applyMAsk(21491591)
	mem[26354] = applyMAsk(3991)
	mem[46496] = applyMAsk(225777)
	mem[19054] = applyMAsk(6818)
	mem[46391] = applyMAsk(1804050)
	mask = "0X01X1110011111101X011X1X10010011100"
	mem[8848] = applyMAsk(3301953)
	mem[21325] = applyMAsk(828483041)
	mem[35954] = applyMAsk(393891988)
	mask = "001010XX11X0X11X1010001000001111X1X1"
	mem[14556] = applyMAsk(48978)
	mem[17078] = applyMAsk(3023995)
	mem[41895] = applyMAsk(1263)
	mem[26354] = applyMAsk(982)
	mem[47494] = applyMAsk(9997)
	mem[42458] = applyMAsk(139205796)
	mask = "01011X0X01XX11X11X10001X010X10001100"
	mem[30326] = applyMAsk(230268)
	mem[13671] = applyMAsk(406)
	mem[13219] = applyMAsk(816366)
	mask = "00X0111000101101XX10X010X100X0101X01"
	mem[9172] = applyMAsk(5603)
	mem[10540] = applyMAsk(399)
	mem[6994] = applyMAsk(2969)
	mem[41827] = applyMAsk(157730)
	mem[57564] = applyMAsk(713362)
	mem[16823] = applyMAsk(335722)
	mem[38893] = applyMAsk(724)
	mask = "X1X10110011111X01X101X111101X0X1X000"
	mem[5336] = applyMAsk(623)
	mem[31859] = applyMAsk(88241437)
	mem[7282] = applyMAsk(19257561)
	mem[37684] = applyMAsk(5467271)
	mem[50354] = applyMAsk(2120264)
	mask = "0000X100111X1111X1X001X0000X00110001"
	mem[61468] = applyMAsk(9124391)
	mem[35212] = applyMAsk(23096803)
	mem[9949] = applyMAsk(9454)
	mem[1331] = applyMAsk(995)
	mem[39238] = applyMAsk(74423)
	mask = "00100100111X1111111011X1XX010X1X0010"
	mem[39294] = applyMAsk(114968517)
	mem[23155] = applyMAsk(1708)
	mem[29927] = applyMAsk(7693420)
	mem[26742] = applyMAsk(2017)
	mask = "001001001X101111X1X0XX01100X0X0101XX"
	mem[29056] = applyMAsk(30646)
	mem[59210] = applyMAsk(113022)
	mem[43000] = applyMAsk(144138476)
	mem[35167] = applyMAsk(721)
	mem[30809] = applyMAsk(507151422)
	mask = "0001X0100101111010X0100000X0X0100010"
	mem[34554] = applyMAsk(49221)
	mem[7437] = applyMAsk(62877)
	mem[59828] = applyMAsk(184498)
	mem[45586] = applyMAsk(20089049)
	mem[48248] = applyMAsk(98197865)
	mem[44772] = applyMAsk(113026522)
	mask = "00X0X000111X11111010X10X1011010X0X01"
	mem[40499] = applyMAsk(11427785)
	mem[60906] = applyMAsk(496319403)
	mem[55126] = applyMAsk(270707060)
	mask = "001XXXX10011X1111000110X0X000XX10001"
	mem[17938] = applyMAsk(535040)
	mem[59138] = applyMAsk(250862772)
	mem[8507] = applyMAsk(41576622)
	mem[14146] = applyMAsk(1026)
	mem[32774] = applyMAsk(2975)
	mem[39952] = applyMAsk(440004)
	mem[13671] = applyMAsk(2552)
	mask = "00000100X110111X111001X0X10X0XX100X1"
	mem[43163] = applyMAsk(1156)
	mem[35680] = applyMAsk(6039549)
	mem[15378] = applyMAsk(657)
	mem[52635] = applyMAsk(15396)
	mem[25926] = applyMAsk(210)
	mem[28827] = applyMAsk(206264701)
	mask = "001X100X001X1111100011X1XX1000100X01"
	mem[5753] = applyMAsk(1286)
	mem[6271] = applyMAsk(6375)
	mem[47122] = applyMAsk(5264524)
	mask = "0X01100XXX10111X1X100X10000X1000001X"
	mem[51736] = applyMAsk(2480)
	mem[16722] = applyMAsk(663122146)
	mem[18483] = applyMAsk(119830)
	mem[13423] = applyMAsk(136699070)
	mem[65442] = applyMAsk(23484946)
	mem[50742] = applyMAsk(716)
	mem[14788] = applyMAsk(129808)
	mask = "100100000X101111111001000101X00X01X0"
	mem[35279] = applyMAsk(11923915)
	mem[12886] = applyMAsk(216539704)
	mem[29197] = applyMAsk(258471)
	mem[34051] = applyMAsk(24342647)
	mem[34556] = applyMAsk(219870381)
	mask = "0010X100011011X11110X1101X010X1001X1"
	mem[17497] = applyMAsk(684)
	mem[8762] = applyMAsk(255129)
	mem[44486] = applyMAsk(49869056)
	mem[35680] = applyMAsk(210517)
	mem[18805] = applyMAsk(17289)
	mask = "001011101X1011111X01111010001000XX01"
	mem[31918] = applyMAsk(89826257)
	mem[31090] = applyMAsk(57962)
	mem[53807] = applyMAsk(2817)
	mem[12784] = applyMAsk(2137)
	mem[31369] = applyMAsk(64658)
	mask = "X000X1X0011011111010X0X111XX00X1X100"
	mem[63989] = applyMAsk(3267)
	mem[23677] = applyMAsk(2600)
	mem[42781] = applyMAsk(125518)
	mem[37480] = applyMAsk(3172)
	mem[23573] = applyMAsk(8287963)
	mask = "0010110001111XXX10XX0X1X0X0000010100"
	mem[1654] = applyMAsk(317418946)
	mem[46425] = applyMAsk(459126)
	mem[31666] = applyMAsk(9190945)
	mem[12079] = applyMAsk(168185843)
	mem[39256] = applyMAsk(178728)
	mem[52287] = applyMAsk(59458806)
	mask = "000001000110111X1010X00000X100110X00"
	mem[55820] = applyMAsk(2071)
	mem[33429] = applyMAsk(821261571)
	mem[16244] = applyMAsk(3215)
	mask = "100110100110X1111110X01101X1XX10X011"
	mem[47796] = applyMAsk(98)
	mem[31040] = applyMAsk(1031334)
	mem[9622] = applyMAsk(1580517)
	mask = "00100X001110XX11110000011101X0011001"
	mem[42781] = applyMAsk(312404)
	mem[39942] = applyMAsk(123252858)
	mem[20867] = applyMAsk(1356)
	mem[14667] = applyMAsk(576)
	mem[35502] = applyMAsk(298476332)
	mem[53427] = applyMAsk(365745)
	mask = "00010X100110111111100010XX00XX000100"
	mem[53876] = applyMAsk(603517)
	mem[10405] = applyMAsk(16459102)
	mem[45543] = applyMAsk(4443)
	mem[41543] = applyMAsk(1411)
	mem[62450] = applyMAsk(6470215)
	mask = "100X110X11111X110X0010000X0X10001X10"
	mem[5336] = applyMAsk(650575)
	mem[50124] = applyMAsk(3080229)
	mem[51618] = applyMAsk(156)
	mem[42185] = applyMAsk(1366)
	mask = "00X0X1101110101110XX1011X001001110X0"
	mem[30736] = applyMAsk(333574460)
	mem[13675] = applyMAsk(8643742)
	mem[12826] = applyMAsk(453315)
	mask = "000X00100X1X1X111XX0101X0X100111000X"
	mem[54995] = applyMAsk(183737953)
	mem[63234] = applyMAsk(679)
	mem[61488] = applyMAsk(337)
	mask = "0XX11111X0111X1X01001111X00X11001XX1"
	mem[10621] = applyMAsk(122118726)
	mem[6109] = applyMAsk(9210)
	mem[15688] = applyMAsk(184799)
	mem[25564] = applyMAsk(367237)
	mask = "00110000011X1111100X000X00XXX1XX1000"
	mem[17938] = applyMAsk(33020705)
	mem[2666] = applyMAsk(88651117)
	mem[21482] = applyMAsk(161753)

	bigo := int64(0)
	for _, v := range mem {
		bigo += v
	}
	fmt.Println("", bigo)
}