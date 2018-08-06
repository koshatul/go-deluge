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

package deluge

// TorrentFiles contains the file information for a file downloaded with a torrent
type TorrentFiles struct {
	Index  int    `json:"index" mapstructure:"index"`
	Path   string `json:"path" mapstructure:"path"`
	Offset int    `json:"offset" mapstructure:"offset"`
	Size   int64  `json:"size" mapstructure:"size"`
}

// TorrentError is the error returned by a torrent
type TorrentError struct {
	Category string `json:"category" mapstructure:"category"`
	Value    int    `json:"value" mapstructure:"value"`
}

// TorrentTracker contains the tracker information for a torrent
type TorrentTracker struct {
	SendStats        bool         `json:"send_stats" mapstructure:"send_stats"`
	Fails            int          `json:"fails" mapstructure:"fails"`
	Verified         bool         `json:"verified" mapstructure:"verified"`
	ScrapeIncomplete int          `json:"scrape_incomplete" mapstructure:"scrape_incomplete"`
	MinAnnounce      interface{}  `json:"min_announce" mapstructure:"min_announce"`
	Source           int          `json:"source" mapstructure:"source"`
	URL              string       `json:"url" mapstructure:"url"`
	LastError        TorrentError `json:"last_error" mapstructure:"last_error"`
	FailLimit        int          `json:"fail_limit" mapstructure:"fail_limit"`
	NextAnnounce     interface{}  `json:"next_announce" mapstructure:"next_announce"`
	CompleteSent     bool         `json:"complete_sent" mapstructure:"complete_sent"`
	ScrapeDownloaded int          `json:"scrape_downloaded" mapstructure:"scrape_downloaded"`
	Trackerid        string       `json:"trackerid" mapstructure:"trackerid"`
	StartSent        bool         `json:"start_sent" mapstructure:"start_sent"`
	Tier             int          `json:"tier" mapstructure:"tier"`
	ScrapeComplete   int          `json:"scrape_complete" mapstructure:"scrape_complete"`
	Message          string       `json:"message" mapstructure:"message"`
	Updating         bool         `json:"updating" mapstructure:"updating"`
}

// Torrent is the struct returned by the JSON API
type Torrent struct {
	Comment              string           `json:"comment" mapstructure:"comment"`
	ActiveTime           int              `json:"active_time" mapstructure:"active_time"`
	IsSeed               bool             `json:"is_seed" mapstructure:"is_seed"`
	Hash                 string           `json:"hash" mapstructure:"hash"`
	UploadPayloadRate    int              `json:"upload_payload_rate" mapstructure:"upload_payload_rate"`
	MoveCompletedPath    string           `json:"move_completed_path" mapstructure:"move_completed_path"`
	Private              bool             `json:"private" mapstructure:"private"`
	TotalPayloadUpload   int              `json:"total_payload_upload" mapstructure:"total_payload_upload"`
	Paused               bool             `json:"paused" mapstructure:"paused"`
	SeedRank             int              `json:"seed_rank" mapstructure:"seed_rank"`
	SeedingTime          int              `json:"seeding_time" mapstructure:"seeding_time"`
	MaxUploadSlots       int              `json:"max_upload_slots" mapstructure:"max_upload_slots"`
	PrioritizeFirstLast  bool             `json:"prioritize_first_last" mapstructure:"prioritize_first_last"`
	DistributedCopies    float64          `json:"distributed_copies" mapstructure:"distributed_copies"`
	DownloadPayloadRate  int              `json:"download_payload_rate" mapstructure:"download_payload_rate"`
	Message              string           `json:"message" mapstructure:"message"`
	NumPeers             int              `json:"num_peers" mapstructure:"num_peers"`
	MaxDownloadSpeed     int              `json:"max_download_speed" mapstructure:"max_download_speed"`
	MaxConnections       int              `json:"max_connections" mapstructure:"max_connections"`
	Compact              bool             `json:"compact" mapstructure:"compact"`
	Ratio                float64          `json:"ratio" mapstructure:"ratio"`
	TotalPeers           int              `json:"total_peers" mapstructure:"total_peers"`
	TotalSize            int64            `json:"total_size" mapstructure:"total_size"`
	TotalWanted          int64            `json:"total_wanted" mapstructure:"total_wanted"`
	State                string           `json:"state" mapstructure:"state"`
	FilePriorities       []int            `json:"file_priorities" mapstructure:"file_priorities"`
	Label                string           `json:"label" mapstructure:"label"`
	MaxUploadSpeed       int              `json:"max_upload_speed" mapstructure:"max_upload_speed"`
	RemoveAtRatio        bool             `json:"remove_at_ratio" mapstructure:"remove_at_ratio"`
	Tracker              string           `json:"tracker" mapstructure:"tracker"`
	SavePath             string           `json:"save_path" mapstructure:"save_path"`
	Progress             float64          `json:"progress" mapstructure:"progress"`
	TimeAdded            float64          `json:"time_added" mapstructure:"time_added"`
	TrackerHost          string           `json:"tracker_host" mapstructure:"tracker_host"`
	TotalUploaded        int              `json:"total_uploaded" mapstructure:"total_uploaded"`
	Files                []TorrentFiles   `json:"files" mapstructure:"files"`
	TotalDone            int64            `json:"total_done" mapstructure:"total_done"`
	NumPieces            int              `json:"num_pieces" mapstructure:"num_pieces"`
	TrackerStatus        string           `json:"tracker_status" mapstructure:"tracker_status"`
	TotalSeeds           int              `json:"total_seeds" mapstructure:"total_seeds"`
	MoveOnCompleted      bool             `json:"move_on_completed" mapstructure:"move_on_completed"`
	NextAnnounce         int              `json:"next_announce" mapstructure:"next_announce"`
	StopAtRatio          bool             `json:"stop_at_ratio" mapstructure:"stop_at_ratio"`
	FileProgress         []float64        `json:"file_progress" mapstructure:"file_progress"`
	MoveCompleted        bool             `json:"move_completed" mapstructure:"move_completed"`
	PieceLength          int              `json:"piece_length" mapstructure:"piece_length"`
	AllTimeDownload      int64            `json:"all_time_download" mapstructure:"all_time_download"`
	MoveOnCompletedPath  string           `json:"move_on_completed_path" mapstructure:"move_on_completed_path"`
	NumSeeds             int              `json:"num_seeds" mapstructure:"num_seeds"`
	Peers                []interface{}    `json:"peers" mapstructure:"peers"`
	Name                 string           `json:"name" mapstructure:"name"`
	Trackers             []TorrentTracker `json:"trackers" mapstructure:"trackers"`
	TotalPayloadDownload int              `json:"total_payload_download" mapstructure:"total_payload_download"`
	IsAutoManaged        bool             `json:"is_auto_managed" mapstructure:"is_auto_managed"`
	SeedsPeersRatio      float64          `json:"seeds_peers_ratio" mapstructure:"seeds_peers_ratio"`
	Queue                int              `json:"queue" mapstructure:"queue"`
	NumFiles             int              `json:"num_files" mapstructure:"num_files"`
	Eta                  int              `json:"eta" mapstructure:"eta"`
	StopRatio            float64          `json:"stop_ratio" mapstructure:"stop_ratio"`
	IsFinished           bool             `json:"is_finished" mapstructure:"is_finished"`
}
