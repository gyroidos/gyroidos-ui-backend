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

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ContainerConfigJson struct {
	Name           string                   `json:"name"`
	GuestOs        string                   `json:"guestOs"`
	GuestOsVersion string                   `json:"guestOsVersion"`
	Color          int                      `json:"color"`
	Type           string                   `json:"type"`
	TokenType      string                   `json:"tokenType"`
	UsbConfigs     []ContainerUsbConfigJson `json:"usbConfigs"`
}

type ContainerUsbConfigJson struct {
	Id     string `json:"id"`
	Serial string `json:"serial"`
	Assign bool   `json:"assign"`
	Type   string `json:"type"`
}

type ContainerControlParamsJson struct {
	Uuids []string `json:"uuids"`
	Key   string   `json:"key,omitempty"`
}

type ContainerCommandJson struct {
	Uuids []string `json:"uuids"`
	Cmd   string   `json:"cmd"`
}

type ContainerChangePinJson struct {
	Uuids  []string `json:"uuids"`
	OldKey string   `json:"oldKey"`
	NewKey string   `json:"newKey"`
}

type ResponseJson struct {
	Code          string `json:"code"`
	Response      string `json:"response"`
	ExecEndReason string `json:"execEndReason"`
}

type RunResponseJson struct {
	Code          string   `json:"code"`
	Data          [][]byte `json:"data"`
	ExecEndReason string   `json:"execEndReason"`
}

func serve(c *config) {

	gin.SetMode(c.mode)

	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST"},
		ExposeHeaders: []string{"Content-Length"},
		AllowHeaders:  []string{"Content-Type"},
	}))

	router.GET("/list", getContainerStatus)
	router.GET("/list_guestos", listGuestosConfigs)
	router.GET("/reload", reloadContainers)
	router.POST("config", getContainerConfig)
	router.POST("/create", createContainer)
	router.POST("/remove", removeContainer)
	router.POST("/start", startContainer)
	router.POST("/stop", stopContainer)
	router.POST("/changepin", changePin)
	router.POST("/run", runCommand)

	log.Infof("Listening and serving on %v", c.addr)

	router.Run(c.addr)
}

func getContainerStatus(c *gin.Context) {
	msg := &ControllerToDaemon{
		Command: ControllerToDaemon_GET_CONTAINER_STATUS.Enum(),
	}

	response, err := control(msg)
	if err != nil {
		sendError(c, http.StatusInternalServerError,
			fmt.Sprintf("Internal error: cml-control failed: %v", err))
	}

	c.String(http.StatusOK, string(response))
}

func listGuestosConfigs(c *gin.Context) {
	msg := &ControllerToDaemon{
		Command: ControllerToDaemon_LIST_GUESTOS_CONFIGS.Enum(),
	}

	response, err := control(msg)
	if err != nil {
		sendError(c, http.StatusInternalServerError,
			fmt.Sprintf("Internal error: cml-control failed: %v", err))
	}

	c.String(http.StatusOK, string(response))
}

func reloadContainers(c *gin.Context) {
	msg := &ControllerToDaemon{
		Command: ControllerToDaemon_RELOAD_CONTAINERS.Enum(),
	}

	response, err := control(msg)
	if err != nil {
		sendError(c, http.StatusInternalServerError,
			fmt.Sprintf("Internal error: cml-control failed: %v", err))
	}

	log.Tracef("Received from cmld: %v", string(response))

	c.String(http.StatusOK, string(response))
}

func getContainerConfig(c *gin.Context) {

	var containerControlParams ContainerControlParamsJson

	if err := c.BindJSON(&containerControlParams); err != nil {
		sendError(c, http.StatusBadRequest, fmt.Sprintf("failed to bind json: %v", err))
		return
	}

	msg := &ControllerToDaemon{
		Command:        ControllerToDaemon_GET_CONTAINER_CONFIG.Enum(),
		ContainerUuids: containerControlParams.Uuids,
	}

	response, err := control(msg)
	if err != nil {
		sendError(c, http.StatusInternalServerError,
			fmt.Sprintf("Internal error: cml-control failed: %v", err))
	}

	c.String(http.StatusOK, string(response))
}

func createContainer(c *gin.Context) {

	var conf ContainerConfigJson

	if err := c.BindJSON(&conf); err != nil {
		sendError(c, http.StatusBadRequest, fmt.Sprintf("failed to bind json: %v", err))
		return
	}

	tmp, err := json.MarshalIndent(&conf, "", "    ")
	if err != nil {
		log.Warnf("failed to marshal: %v", err)
	}
	log.Tracef("Received container config: %v", string(tmp))

	log.Tracef("Creating container %v", conf.Name)

	confFile := fmt.Sprintf(`
name: "%v"
guest_os: "%v"
guestos_version: %v
color: %v
type: %v
token_type: %v
usb_configs {
	id: "045e:0800"
	serial: "0"
	assign: true
	type: PIN_ENTRY
}`,
		conf.Name,
		conf.GuestOs,
		conf.GuestOsVersion,
		conf.Color,
		conf.Type,
		conf.TokenType,
	)

	msg := &ControllerToDaemon{
		Command:             ControllerToDaemon_CREATE_CONTAINER.Enum(),
		ContainerConfigFile: []byte(confFile),
	}

	response, err := control(msg)
	if err != nil {
		sendError(c, http.StatusInternalServerError,
			fmt.Sprintf("Internal error: cml-control failed: %v", err))
	}

	c.String(http.StatusOK, string(response))
}

func removeContainer(c *gin.Context) {

	var containerControlParams ContainerControlParamsJson

	if err := c.BindJSON(&containerControlParams); err != nil {
		sendError(c, http.StatusBadRequest, fmt.Sprintf("failed to bind json: %v", err))
		return
	}

	msg := &ControllerToDaemon{
		Command:        ControllerToDaemon_REMOVE_CONTAINER.Enum(),
		ContainerUuids: containerControlParams.Uuids,
	}

	response, err := control(msg)
	if err != nil {
		sendError(c, http.StatusInternalServerError,
			fmt.Sprintf("Internal error: cml-control failed: %v", err))
		return
	}

	log.Tracef("Received from cmld: %v", string(response))

	c.String(http.StatusOK, string(response))
}

func startContainer(c *gin.Context) {

	var containerControlParams ContainerControlParamsJson

	if err := c.BindJSON(&containerControlParams); err != nil {
		sendError(c, http.StatusBadRequest, fmt.Sprintf("failed to bind json: %v", err))
		return
	}

	// TODO provide possibility to pass setup via frontend?
	hasSetup := false

	msg := &ControllerToDaemon{
		Command: ControllerToDaemon_CONTAINER_START.Enum(),
		ContainerStartParams: &ContainerStartParams{
			Key:      &containerControlParams.Key,
			Setup:    &hasSetup,
			NoSwitch: nil,
		},
		ContainerUuids: containerControlParams.Uuids,
	}

	response, err := control(msg)
	if err != nil {
		sendError(c, http.StatusInternalServerError,
			fmt.Sprintf("Internal error: cml-control failed: %v", err))
		return
	}

	log.Tracef("Received from cmld: %v", string(response))

	c.String(http.StatusOK, string(response))
}

func stopContainer(c *gin.Context) {

	var containerControlParams ContainerControlParamsJson

	if err := c.BindJSON(&containerControlParams); err != nil {
		sendError(c, http.StatusBadRequest, fmt.Sprintf("failed to bind json: %v", err))
		return
	}

	msg := &ControllerToDaemon{
		Command: ControllerToDaemon_CONTAINER_STOP.Enum(),
		ContainerStartParams: &ContainerStartParams{
			Key:      &containerControlParams.Key,
			Setup:    nil,
			NoSwitch: nil,
		},
		ContainerUuids: containerControlParams.Uuids,
	}

	response, err := control(msg)
	if err != nil {
		sendError(c, http.StatusInternalServerError,
			fmt.Sprintf("Internal error: cml-control failed: %v", err))
		return
	}

	log.Tracef("Received from cmld: %v", string(response))

	c.String(http.StatusOK, string(response))
}

func changePin(c *gin.Context) {
	var containerChangePin ContainerChangePinJson

	if err := c.BindJSON(&containerChangePin); err != nil {
		sendError(c, http.StatusBadRequest, fmt.Sprintf("failed to bind json: %v", err))
		return
	}

	msg := &ControllerToDaemon{
		Command:        ControllerToDaemon_CONTAINER_CHANGE_TOKEN_PIN.Enum(),
		DevicePin:      &containerChangePin.OldKey,
		DeviceNewpin:   &containerChangePin.NewKey,
		ContainerUuids: containerChangePin.Uuids,
	}

	response, err := control(msg)
	if err != nil {
		sendError(c, http.StatusInternalServerError,
			fmt.Sprintf("Internal error: cml-control failed: %v", err))
		return
	}

	log.Tracef("Received from cmld: %v", string(response))

	c.String(http.StatusOK, string(response))
}

func runCommand(c *gin.Context) {

	var containerCmd ContainerCommandJson

	if err := c.BindJSON(&containerCmd); err != nil {
		sendError(c, http.StatusBadRequest, fmt.Sprintf("failed to bind json: %v", err))
		return
	}

	log.Tracef("Received command: %v", containerCmd.Cmd)

	cmds := strings.Split(containerCmd.Cmd, " ")
	index := 0
	execPty := true
	var execCmd string
	execArgs := make([]string, 0)

	if len(cmds) > 1 && strings.EqualFold("nopty", cmds[0]) {
		execPty = false
		index++
	}

	if index < len(cmds) {
		execCmd = cmds[index]
	}

	for i := index; i < len(cmds); i++ {
		execArgs = append(execArgs, cmds[i])
	}

	msg := &ControllerToDaemon{
		Command:        ControllerToDaemon_CONTAINER_EXEC_CMD.Enum(),
		ExecPty:        &execPty,
		ExecCommand:    &execCmd,
		ExecArgs:       execArgs,
		ContainerUuids: containerCmd.Uuids,
	}

	log.Tracef("Running command '%v' with args '%v'", *msg.ExecCommand,
		strings.Join(msg.ExecArgs, ", "))

	response, err := controlRun(msg)
	if err != nil {
		sendError(c, http.StatusInternalServerError,
			fmt.Sprintf("Internal error: cml-control failed: %v", err))
		return
	}

	// TODO delete
	for _, r := range response.Data {
		log.Tracef("Received from cmld: %v", string(r))
	}

	c.IndentedJSON(http.StatusOK, response)
}

func sendError(c *gin.Context, httpStatus int, msg string) {
	log.Warn(msg)
	resp := &ResponseJson{
		Code:          "RESPONSE",
		Response:      "CMD_FAILED",
		ExecEndReason: msg,
	}
	c.IndentedJSON(httpStatus, resp)
}
