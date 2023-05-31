// SPDX-License-Identifier: BUSL-1.1

// Copyright (C) 2023, Curiosity Research. All rights reserved.
// Use of this software is covered by the Business Source License included
// in the LICENSE file in the license folder of this repository and at www.mariadb.com/bsl11.

// Any use of the Licensed Work in violation of this License will automatically
// terminate your rights under this License for the current and all other
// versions of the Licensed Work.

// This License does not grant you any right in any trademark or logo of
// Licensor or its affiliates (provided that you may use a trademark or logo of
// Licensor as expressly required by this License).

// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN "AS IS" BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package indexer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/engine"
	"github.com/ethereum/go-ethereum/game"
	"github.com/gorilla/websocket"
)

const (
	SNAPSHOT               = "SNAPSHOT"
	DIFF                   = "DIFF"
	READ_BUFFER_SIZE       = 1024
	WRITE_BUFFER_SIZE      = 1024
	INDEX_WS_PORT          = 9450
	INDEX_WS_TICK_INTERVAL = 1
	TEXT_MESSAGE           = 1
)

type WSMessageType = string

type WSMessage struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type WS struct {
	ecsDiff engine.ECSUpdateArray
}

func (ws *WS) Init() {
	ws.ecsDiff = engine.InitializeEcsUpdateArray()
}

func (ws *WS) PushECSDiff(ecsDiff engine.ECSUpdateArray) {
	for _, ecsUpdate := range ecsDiff {
		ws.ecsDiff.AddUpdate(ecsUpdate.Entity, ecsUpdate.Component, ecsUpdate.Value)
	}
}

// Start WS Server
func (ws *WS) Run() {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  READ_BUFFER_SIZE,
		WriteBufferSize: WRITE_BUFFER_SIZE,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Upgrade normal http protocl to websocket
		websocket, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println("Connected!")

		ws.publishMessage(websocket)
	})

	go func() {
		http.ListenAndServe(fmt.Sprintf("%s%d", ":", INDEX_WS_PORT), nil)
	}()
}

// Serve real time ECSDiffs to client
func (ws *WS) publishMessage(conn *websocket.Conn) {
	// Tick messge to client evert 1s
	ticker := time.NewTicker(INDEX_WS_TICK_INTERVAL * time.Second)
	quit := make(chan struct{})

	// get the worlds state
	latestWorldEcsArr := game.ExportWorldState(&game.MainSnapshotWorld)

	serializedLastestWorldEcsArr, err := json.Marshal(latestWorldEcsArr)

	if err != nil {
		fmt.Println("Serialize message error in publishMessage")
	}

	fullSnapshot := &WSMessage{SNAPSHOT, string(serializedLastestWorldEcsArr)}
	fullSnapshotStr, parseFullSnapshotError := json.Marshal(fullSnapshot)

	// ecs update array
	if parseFullSnapshotError != nil {
		log.Println(parseFullSnapshotError)
	}

	conn.WriteMessage(TEXT_MESSAGE, []byte(string(fullSnapshotStr)))

	go func() {
		for {
			select {
			case <-ticker.C:
				// Diff Message to send
				ecsUpdateArray := ws.fetchECSDiff()
				ecsUpdateArrayStr, parseECSUpdateArrayError := json.Marshal(ecsUpdateArray)

				// Reset ecs diff for being consumed
				ws.resetECSDiff()

				if parseECSUpdateArrayError != nil {
					log.Println(parseECSUpdateArrayError)
					return
				}

				diff := &WSMessage{DIFF, string(ecsUpdateArrayStr)}
				diffStr, parseDiffError := json.Marshal(diff)

				if parseDiffError != nil {
					log.Println(parseDiffError)
					return
				}

				if err := conn.WriteMessage(1, []byte(string(diffStr))); err != nil {
					log.Println(err)
					return
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func (ws *WS) fetchECSDiff() engine.ECSUpdateArray {
	return ws.ecsDiff
}

func (ws *WS) resetECSDiff() {
	ws.ecsDiff = engine.InitializeEcsUpdateArray()
}

// Single WS instance for other module call
var ws = &WS{}

// Get Single WS instance
func GetWSInstance() *WS {
	return ws
}
