/*
 * This file is part of GyroidOS
 * Copyright(c) 2013 - 2023 Fraunhofer AISEC
 * Fraunhofer-Gesellschaft zur Förderung der angewandten Forschung e.V.
 *
 * This program is free software; you can redistribute it and/or modify it
 * under the terms and conditions of the GNU General Public License,
 * version 2 (GPL 2), as published by the Free Software Foundation.
 *
 * This program is distributed in the hope it will be useful, but WITHOUT
 * ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
 * FITNESS FOR A PARTICULAR PURPOSE. See the GPL 2 license for more details.
 *
 * You should have received a copy of the GNU General Public License along with
 * this program; if not, see <http://www.gnu.org/licenses/>
 *
 * The full GNU General Public License is included in this distribution in
 * the file called "COPYING".
 *
 * Contact Information:
 * Fraunhofer AISEC <gyroidos@aisec.fraunhofer.de>
 */

package main

// Install github packages with "go get [url]"
import (
	"context"
	"encoding/binary"
	"fmt"
	"net"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	// local modules
)

func control(msg proto.Message) ([]byte, error) {

	payload, err := proto.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal: %v", err)
	}

	// Establish connection with timeout
	var d net.Dialer
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	conn, err := d.DialContext(ctx, "unix", "/run/socket/cml-control")
	if err != nil {
		return nil, fmt.Errorf("error dialing: %w", err)
	}
	defer conn.Close()
	err = conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		return nil, fmt.Errorf("failed to set read deadline: %w", err)
	}

	// Send request
	err = send(conn, payload)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	// Read reply
	data, err := receive(conn)
	if err != nil {
		return nil, fmt.Errorf("failed to receive: %w", err)
	}

	response, err := toJson(data)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to json: %w", err)
	}

	return response, nil
}

func toJson(data []byte) ([]byte, error) {
	response := new(DaemonToController)
	err := proto.Unmarshal(data, response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	json, err := protojson.Marshal(response)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal: %w", err)
	}

	return json, nil
}

func receive(conn net.Conn) ([]byte, error) {

	// Read header
	buf := make([]byte, 4)
	n, err := conn.Read(buf)
	if err != nil {
		return nil, fmt.Errorf("failed to read header: %w", err)
	}
	if n != 4 {
		return nil, fmt.Errorf("read %v bytes (expected 4)", n)
	}

	len := binary.BigEndian.Uint32(buf)

	// Read payload
	payload := make([]byte, len)
	n, err = conn.Read(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to read payload: %w", err)
	}
	if uint32(n) != len {
		return nil, fmt.Errorf("failed to read payload (received %v, expected %v bytes)",
			n, len)
	}

	return payload, nil
}

func send(conn net.Conn, payload []byte) error {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf[0:4], uint32(len(payload)))
	n, err := conn.Write(buf)
	if err != nil {
		return fmt.Errorf("failed to send header: %w", err)
	}
	if n != len(buf) {
		return fmt.Errorf("could only send %v of %v bytes", n, len(buf))
	}
	n, err = conn.Write(payload)
	if err != nil {
		return fmt.Errorf("failed to send response: %w", err)
	}
	if n != len(payload) {
		return fmt.Errorf("could only send %v of %v bytes", n, len(payload))
	}

	return nil
}

func controlRun(msg proto.Message) (*RunResponseJson, error) {

	payload, err := proto.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal: %v", err)
	}

	// Establish connection with timeout
	var d net.Dialer
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	conn, err := d.DialContext(ctx, "unix", "/run/socket/cml-control")
	if err != nil {
		return nil, fmt.Errorf("error dialing: %w", err)
	}
	err = conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		return nil, fmt.Errorf("failed to set read deadline: %w", err)
	}
	defer conn.Close()

	// Send request
	err = send(conn, payload)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	// Read reply
	responses := &RunResponseJson{
		Code: "EXEC_OUTPUT",
	}
	for {
		log.Warn("Reading data")
		data, err := receive(conn)
		if err != nil {
			return nil, fmt.Errorf("failed to receive: %w", err)
		}

		var m DaemonToController
		err = proto.Unmarshal(data, &m)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal protobuf: %w", err)
		}

		if *m.Code == DaemonToController_EXEC_END {
			break
		}

		if *m.Code != DaemonToController_EXEC_OUTPUT {
			return nil, fmt.Errorf("unexpected response: %v", m.Code.String())
		}

		responses.Data = append(responses.Data, *&m.ExecOutput)
	}

	return responses, nil
}
