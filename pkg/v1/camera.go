package v1

import (
	"encoding/json"
	"fmt"
)

type cameraHandler struct {
	printer *Printer
}

// Returns the list of active cameras.
func (c cameraHandler) List() (*[]Camera, error) {
	return parseAsJSON[[]Camera](c.printer.get("/api/v1/cameras"))
}

// Puts the cameras in the given order.
func (c cameraHandler) Order(intendedOrder []string) error {
	data, err := json.Marshal(intendedOrder)
	if err != nil {
		return err
	}
	_, err = c.printer.put("/api/v1/cameras", data, nil)
	return err
}

// Get current settings and properties of a specific camera.
func (c cameraHandler) GetInfo(id string) (*CameraConfig, error) {
	return parseAsJSON[CameraConfig](c.printer.get(fmt.Sprintf("/api/v1/cameras/%s", id)))
}

// Setup a new camera or fix a broken one.
func (c cameraHandler) Setup(id string, config CameraConfigSet) error {
	data, err := json.Marshal(config)
	if err != nil {
		return err
	}
	_, err = c.printer.post(fmt.Sprintf("/api/v1/cameras/%s", id), data)
	return err
}

// Delete a camera.
func (c cameraHandler) Delete(id string) error {
	_, err := c.printer.delete(fmt.Sprintf("/api/v1/cameras/%s", id), nil)
	return err
}

// Return a captured image from the camera with the given ID.
// If the ID is empty it will return a snapshot from the default camera.
// TODO: Verify uf MakeSnapshot should be called first.
// TODO: Figure out the data encoding of the response.
func (c cameraHandler) GetSnapshot(id string) (string, error) {
	url := "/api/v1/cameras/snap"
	if id != "" {
		url = fmt.Sprintf("/api/v1/cameras/%s/snap", id)
	}
	resp, err := c.printer.get(url)
	if err != nil {
		return "", err
	}
	return string(resp), nil
}

// Make a snapshot of the camera with the given ID.
func (c cameraHandler) MakeSnapshot(id string) error {
	_, err := c.printer.post(fmt.Sprintf("/api/v1/cameras/%s/snap", id), nil)
	return err
}

// Update the config of the camera with the given ID.
func (c cameraHandler) UpdateConfig(id string, config CameraConfigSet) error {
	data, err := json.Marshal(config)
	if err != nil {
		return err
	}
	_, err = c.printer.patch(fmt.Sprintf("/api/v1/cameras/%s/config", id), data)
	return err
}

// Reset the config for the camera with the given ID.
func (c cameraHandler) ResetConfig(id string) error {
	_, err := c.printer.delete(fmt.Sprintf("/api/v1/cameras/%s/config", id), nil)
	return err
}

// Register the camera with the given ID to connect.
func (c cameraHandler) RegisterToConnect(id string) error {
	_, err := c.printer.post(fmt.Sprintf("/api/v1/cameras/%s/connection", id), nil)
	return err
}

// Unregister the camera with the given ID from connect.
func (c cameraHandler) UnregisterCameraToConnect(id string) error {
	_, err := c.printer.delete(fmt.Sprintf("/api/v1/cameras/%s/connection", id), nil)
	return err
}
