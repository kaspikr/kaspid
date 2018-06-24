// Copyright (c) 2013-2017 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package blockdag

import (
	"fmt"

	"github.com/daglabs/btcd/database"
	"github.com/daglabs/btcutil"
)

// maybeAcceptBlock potentially accepts a block into the block chain and, if
// accepted, returns whether or not it is on the main chain.  It performs
// several validation checks which depend on its position within the block chain
// before adding it.  The block is expected to have already gone through
// ProcessBlock before calling this function with it.
//
// The flags are also passed to checkBlockContext and connectBestChain.  See
// their documentation for how the flags modify their behavior.
//
// This function MUST be called with the chain state lock held (for writes).
func (b *BlockChain) maybeAcceptBlock(block *btcutil.Block, flags BehaviorFlags) (bool, error) {
	// The height of this block is one more than the referenced previous
	// block.
	nodes, err := lookupPreviousNodes(block, b)
	if err != nil {
		return false, err
	}

	firstNode := nodes[0] // TODO: (Stas) This is wrong. Modified only to satisfy compilation.
	blockHeight := firstNode.height + 1
	block.SetHeight(blockHeight)

	// The block must pass all of the validation rules which depend on the
	// position of the block within the block chain.
	err = b.checkBlockContext(block, firstNode, flags)
	if err != nil {
		return false, err
	}

	// Insert the block into the database if it's not already there.  Even
	// though it is possible the block will ultimately fail to connect, it
	// has already passed all proof-of-work and validity tests which means
	// it would be prohibitively expensive for an attacker to fill up the
	// disk with a bunch of blocks that fail to connect.  This is necessary
	// since it allows block download to be decoupled from the much more
	// expensive connection logic.  It also has some other nice properties
	// such as making blocks that never become part of the main chain or
	// blocks that fail to connect available for further analysis.
	err = b.db.Update(func(dbTx database.Tx) error {
		return dbStoreBlock(dbTx, block)
	})
	if err != nil {
		return false, err
	}

	// Create a new block node for the block and add it to the node index. Even
	// if the block ultimately gets connected to the main chain, it starts out
	// on a side chain.
	blockHeader := &block.MsgBlock().Header
	newNode := newBlockNode(blockHeader, nodes)
	newNode.status = statusDataStored

	b.index.AddNode(newNode)
	err = b.index.flushToDB()
	if err != nil {
		return false, err
	}

	// Connect the passed block to the chain while respecting proper chain
	// selection according to the chain with the most proof of work.  This
	// also handles validation of the transaction scripts.
	isMainChain, err := b.connectBestChain(newNode, block, flags)
	if err != nil {
		return false, err
	}

	// Notify the caller that the new block was accepted into the block
	// chain.  The caller would typically want to react by relaying the
	// inventory to other peers.
	b.chainLock.Unlock()
	b.sendNotification(NTBlockAccepted, block)
	b.chainLock.Lock()

	return isMainChain, nil
}

func lookupPreviousNodes(block *btcutil.Block, blockChain *BlockChain) ([]*blockNode, error) {
	header := block.MsgBlock().Header
	prevHashes := header.PrevBlocks

	nodes := make([]*blockNode, len(prevHashes))
	for _, prevHash := range prevHashes {
		node := blockChain.index.LookupNode(&prevHash)
		if node == nil {
			str := fmt.Sprintf("previous block %s is unknown", prevHashes)
			return nil, ruleError(ErrPreviousBlockUnknown, str)
		} else if blockChain.index.NodeStatus(node).KnownInvalid() {
			str := fmt.Sprintf("previous block %s is known to be invalid", prevHashes)
			return nil, ruleError(ErrInvalidAncestorBlock, str)
		}

		nodes = append(nodes, node)
	}

	return nodes, nil
}
