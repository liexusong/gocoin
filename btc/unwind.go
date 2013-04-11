package btc

import (
	//"fmt"
)

type txUnwindData struct {
	us *UnspentDb
}


func NewUnwindBuffer(us *UnspentDb) (ub *txUnwindData) {
	ub = new(txUnwindData)
	ub.us = us
	return 
}

func (u *txUnwindData)NewHeight(height uint32) {
	if height > UnwindBufferMaxHistory {
		u.us.db.UnwindDel(height-UnwindBufferMaxHistory)
	}
}

func (u *txUnwindData)addToDeleted(height uint32, txin *TxPrevOut, txout *TxOut) {
	u.us.db.UnwindAdd(height, 0, txin, txout)
}


func (u *txUnwindData)addToAdded(height uint32, txin *TxPrevOut, newout *TxOut) {
	u.us.db.UnwindAdd(height, 1, txin, newout)
}


func (u *txUnwindData)UnwindBlock(height uint32, db *UnspentDb) {
	u.us.db.UnwindNow(height)
}


