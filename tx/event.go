// Copyright (c) 2018 The VeChainThor developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package tx

import (
	"github.com/stablyio/thor/thor"
)

// Event represents a contract event log. These events are generated by the LOG opcode and
// stored/indexed by the node.
type Event struct {
	// address of the contract that generated the event
	Address thor.Address
	// list of topics provided by the contract.
	Topics []thor.Bytes32
	// supplied by the contract, usually ABI-encoded
	Data []byte
}

// Events slice of event logs.
type Events []*Event
