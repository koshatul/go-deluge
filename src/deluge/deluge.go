// Copyright 2013-2018 Bruno Albuquerque (bga@bug-br.org.br).
// Copyright 2013-2018 Kosh (koshatul@gmail.com).
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not
// use this file except in compliance with the License. You may obtain a copy of
// the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations under
// the License.

// Package deluge implements a Go wrapper around the Deluge Remote JSON API
// (http://deluge.readthedocs.io/en/develop/core/rpc.html). This allows
// programmers to control Deluge (http://deluge-torrent.org) programatically
// from inside Go programs. Note this is a work in progress and not everything
// is implemented but adding extra RPC calls is trivial.
package deluge

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/mitchellh/mapstructure"
)

// Deluge represents an endpoint for Deluge RPC requests.
type Deluge struct {
	url      string
	password string

	client  *http.Client
	cookies []*http.Cookie

	id uint64
}

// New instantiates a new Deluge instance and authenticates with the
// server.
func New(url, password string, timeout time.Duration) (*Deluge, error) {
	client := &http.Client{
		Timeout: timeout,
	}
	d := &Deluge{
		url,
		password,
		client,
		nil,
		0,
	}

	err := d.authLogin()
	if err != nil {
		return nil, err
	}

	return d, err
}

// CoreAddTorrentFile wraps the core.add_torrent_file RPC call. fileName is the
// name of the original torrent file. fileDump is the base64 encoded contents of
// the file and options is a map with options to be set (consult the Deluge
// Torrent documentation for a list of valid options).
func (d *Deluge) CoreAddTorrentFile(fileName, fileDump string, options map[string]interface{}) (string, error) {
	response, err := d.sendJSONRequest("core.add_torrent_file",
		[]interface{}{fileName, fileDump, options})
	if err != nil {
		return "", err
	}

	return response["result"].(string), nil
}

// CoreAddTorrentMagnet wraps the core.add_torrent_magnet RPC call. magnetURL is
// the Magnet URL for the torrent and options is a map with options to be set
// (consult the Deluge Torrent documentation for a list of valid options).
func (d *Deluge) CoreAddTorrentMagnet(magnetURL string, options map[string]interface{}) (string, error) {
	response, err := d.sendJSONRequest("core.add_torrent_magnet",
		[]interface{}{magnetURL, options})
	if err != nil {
		return "", err
	}

	return response["result"].(string), nil
}

// CoreAddTorrentURL wraps the core.add_torrent_url RPC call. torrentURL is
// the URL for the torrent and options is a map with options to be set
// (consult de Deluge Torrent documentation for a list of valid options).
func (d *Deluge) CoreAddTorrentURL(torrentURL string, options map[string]interface{}) (string, error) {
	response, err := d.sendJSONRequest("core.add_torrent_url",
		[]interface{}{torrentURL, options})
	if err != nil {
		return "", err
	}

	return response["result"].(string), nil
}

// CoreGetTorrentStatus needs a comment
func (d *Deluge) CoreGetTorrentStatus(hash string) (*Torrent, error) {
	response, err := d.sendJSONRequest("core.get_torrent_status", []interface{}{
		hash,
		[]string{},
	})
	if err != nil {
		return nil, err
	}

	result := &Torrent{}
	err = mapstructure.Decode(response["result"], result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// CoreGetTorrentsStatus needs a comment
func (d *Deluge) CoreGetTorrentsStatus(filterDict map[string]string) (map[string]Torrent, error) {
	response, err := d.sendJSONRequest("core.get_torrents_status", []interface{}{
		filterDict,
		[]string{},
	})
	if err != nil {
		return nil, err
	}

	result := make(map[string]Torrent)
	err = mapstructure.Decode(response["result"], &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// CoreRemoveTorrent needs a comment
func (d *Deluge) CoreRemoveTorrent(hash string, removeData bool) (bool, error) {
	response, err := d.sendJSONRequest("core.remove_torrent",
		[]interface{}{hash, removeData})
	if err != nil {
		return false, err
	}

	return response["result"].(bool), nil
}

func (d *Deluge) authLogin() error {
	response, err := d.sendJSONRequest("auth.login",
		[]interface{}{d.password})
	if err != nil {
		return err
	}

	if response["result"] != true {
		return fmt.Errorf("authetication failed")
	}

	return nil
}

func (d *Deluge) sendJSONRequest(method string, params []interface{}) (map[string]interface{}, error) {
	atomic.AddUint64(&(d.id), 1)
	data, err := json.Marshal(map[string]interface{}{
		"method": method,
		"id":     d.id,
		"params": params,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", d.url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	if d.cookies != nil {
		for _, cookie := range d.cookies {
			req.AddCookie(cookie)
		}
	}

	resp, err := d.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf(
			"received non-ok status to http request : %d",
			resp.StatusCode)
	}

	d.cookies = resp.Cookies()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := make(map[string]interface{})
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	if result["error"] != nil {
		return nil, fmt.Errorf("json error : %v", result["error"])
	}

	return result, nil
}
