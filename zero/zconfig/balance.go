package zconfig

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/sero-cash/go-czero-import/c_type"
)

func Balance_dir() string {
	return filepath.Join(dir, "balance")
}

func Init_Balance_dir() {
	os.Mkdir(Balance_dir(), os.ModePerm)
}

var eroot_json = `[
	"0xeb85bf17f7f2fe9a1410f99f7e174b409f48479addd41266779308780406ae01",
	"0xeb85bf17f7f2fe9a1410f99f7e174b409f48479addd41266779308780406ae02",
	"0xeb85bf17f7f2fe9a1410f99f7e174b409f48479addd41266779308780406ae03",
	"0xeb85bf17f7f2fe9a1410f99f7e174b409f48479addd41266779308780406ae04",
	"0xeb85bf17f7f2fe9a1410f99f7e174b409f48479addd41266779308780406ae05",
	"0xeb85bf17f7f2fe9a1410f99f7e174b409f48479addd41266779308780406ae06",
	"0xeb85bf17f7f2fe9a1410f99f7e174b409f48479addd41266779308780406ae07",
	"0xeb85bf17f7f2fe9a1410f99f7e174b409f48479addd41266779308780406ae08",
	"0xeb85bf17f7f2fe9a1410f99f7e174b409f48479addd41266779308780406ae09",
	"0x75a2488445fb7ec5d37f4740b15f976ae35c358e01806701dca891221bdecfa1",
	"0x75a2488445fb7ec5d37f4740b15f976ae35c358e01806701dca891221bdecfa2",
	"0x75a2488445fb7ec5d37f4740b15f976ae35c358e01806701dca891221bdecfa3",
	"0x75a2488445fb7ec5d37f4740b15f976ae35c358e01806701dca891221bdecfa4",
	"0x75a2488445fb7ec5d37f4740b15f976ae35c358e01806701dca891221bdecfa5",
	"0x75a2488445fb7ec5d37f4740b15f976ae35c358e01806701dca891221bdecfa6",
	"0x75a2488445fb7ec5d37f4740b15f976ae35c358e01806701dca891221bdecfa7",
	"0x75a2488445fb7ec5d37f4740b15f976ae35c358e01806701dca891221bdecfa8",
	"0x75a2488445fb7ec5d37f4740b15f976ae35c358e01806701dca891221bdecfa9",
]`

func error_roots() (ret map[c_type.Uint256]bool) {
	ret = make(map[c_type.Uint256]bool)
	var erts []c_type.Uint256
	json.Unmarshal([]byte(eroot_json), &erts)

	for _, ert := range erts {
		var key c_type.Uint256
		copy(key[:], ert[:])
		ret[key] = true
	}
	return
}

var ERS = error_roots()

func IsValidRoot(root c_type.Uint256) bool {
	if _, ok := ERS[root]; ok {
		return false
	} else {
		return true
	}
}
