// SPDX-License-Identifier: LGPL-3.0-or-later
// Copyright 2019 DNA Dev team
//
/*
 * Copyright (C) 2018 The ontology Authors
 * This file is part of The ontology library.
 *
 * The ontology is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The ontology is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The ontology.  If not, see <http://www.gnu.org/licenses/>.
 */

package actor

import (
	cactor "github.com/DNAProject/DNA/consensus/actor"
	"github.com/ontio/ontology-eventbus/actor"
)

var consensusSrvPid *actor.PID

func SetConsensusPid(actr *actor.PID) {
	consensusSrvPid = actr
}

//start consensus to consensus actor
func ConsensusSrvStart() error {
	if consensusSrvPid != nil {
		consensusSrvPid.Tell(&cactor.StartConsensus{})
	}
	return nil
}

//halt consensus to consensus actor
func ConsensusSrvHalt() error {
	if consensusSrvPid != nil {
		consensusSrvPid.Tell(&cactor.StopConsensus{})
	}
	return nil
}
