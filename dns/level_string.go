// Code generated by "stringer -type=level util.go"; DO NOT EDIT.

package dns

import "strconv"

const _level_name = "DISABLEBLOCKSPEEDUPlevelEnd"

var _level_index = [...]uint8{0, 7, 12, 19, 27}

func (i level) String() string {
	if i < 0 || i >= level(len(_level_index)-1) {
		return "level(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _level_name[_level_index[i]:_level_index[i+1]]
}
