// Code generated by "stringer -type tokenType scanner.go"; DO NOT EDIT.

package json

import "strconv"

const _tokenType_name = "tokenInvalidtokenCommatokenColontokenEqualstokenKeywordtokenNumbertokenStringtokenBrackOtokenBrackCtokenBraceOtokenBraceCtokenEOF"

var _tokenType_map = map[tokenType]string{
	0:    _tokenType_name[0:12],
	44:   _tokenType_name[12:22],
	58:   _tokenType_name[22:32],
	61:   _tokenType_name[32:43],
	75:   _tokenType_name[43:55],
	78:   _tokenType_name[55:66],
	83:   _tokenType_name[66:77],
	91:   _tokenType_name[77:88],
	93:   _tokenType_name[88:99],
	123:  _tokenType_name[99:110],
	125:  _tokenType_name[110:121],
	9220: _tokenType_name[121:129],
}

func (i tokenType) String() string {
	if str, ok := _tokenType_map[i]; ok {
		return str
	}
	return "tokenType(" + strconv.FormatInt(int64(i), 10) + ")"
}
