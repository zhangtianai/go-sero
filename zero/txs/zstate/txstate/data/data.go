package data

import (
	"github.com/sero-cash/go-czero-import/keys"
	"github.com/sero-cash/go-sero/zero/localdb"
)

type Data struct {
	Num   uint64
	Cur   Current
	Block StateBlock
	//G2ins  map[keys.Uint256]bool
	G2outs map[keys.Uint256]*localdb.OutState

	//Dirty_last_out bool
	Dirty_G2ins  map[keys.Uint256]bool
	Dirty_G2outs map[keys.Uint256]bool
}

func NewData(num uint64) (ret Data) {
	ret.Num = num
	return
}

func (state *Data) clear_dirty() {
	//state.Dirty_last_out = false
	state.Dirty_G2ins = make(map[keys.Uint256]bool)
	state.Dirty_G2outs = make(map[keys.Uint256]bool)
}

func (state *Data) Clear() {
	state.Cur = NewCur()
	state.G2outs = make(map[keys.Uint256]*localdb.OutState)
	state.Block = StateBlock{}
	state.clear_dirty()
}

func (self *Data) appendDel(del *keys.Uint256) {
	if del == nil {
		panic("set_last_out but del is nil")
	}
	self.Block.Dels = append(self.Block.Dels, *del)
	//self.Dirty_last_out = true
}

func (self *Data) appendRoot(root *keys.Uint256) {
	if root == nil {
		panic("set_last_out but root is nil")
	}
	self.Cur.Index = self.Cur.Index + int64(1)
	self.Block.Roots = append(self.Block.Roots, *root)
	//self.Dirty_last_out = true
}

func (self *Data) addInByNilOrRoot(in *keys.Uint256) {
	//self.G2ins[*in] = true
	self.Dirty_G2ins[*in] = true
}

func (self *Data) addOutByRoot(k *keys.Uint256, out *localdb.OutState) {
	self.G2outs[*k] = out
	self.Dirty_G2outs[*k] = true
}

func (self *Data) AddOut(root *keys.Uint256, out *localdb.OutState) {
	self.addOutByRoot(root, out)
	self.appendRoot(root)
	if self.Cur.Index != int64(out.Index) {
		panic("add out but cur.index != current_index")
	}
	if self.Cur.Index < 0 {
		panic("add out but cur.index < 0")
	}
	return
}

func (self *Data) AddNil(in *keys.Uint256) {
	self.addInByNilOrRoot(in)
	self.appendDel(in)
}

func (self *Data) AddDel(in *keys.Uint256) {
	self.appendDel(in)
}
