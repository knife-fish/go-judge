package main

import "github.com/criyle/go-judge/pkg/envexec"

type cmdFile struct {
	Src     *string `json:"src"`
	Content *string `json:"content"`
	FileID  *string `json:"fileId"`
	Name    *string `json:"name"`
	Max     *int64  `json:"max"`
}

type cmd struct {
	Args  []string   `json:"args"`
	Env   []string   `json:"env,omitempty"`
	Files []*cmdFile `json:"files,omitempty"`

	CPULimit     uint64 `json:"cpuLimit"`
	RealCPULimit uint64 `json:"realCpuLimit"`
	MemoryLimit  uint64 `json:"memoryLimit"`
	ProcLimit    uint64 `json:"procLimit"`

	CopyIn map[string]cmdFile `json:"copyIn"`

	CopyOut       []string `json:"copyOut"`
	CopyOutCached []string `json:"copyOutCached"`
	CopyOutDir    string   `json:"copyOutDir"`
}

type pipeIndex struct {
	Index int `json:"index"`
	Fd    int `json:"fd"`
}

type pipeMap struct {
	In  pipeIndex `json:"in"`
	Out pipeIndex `json:"out"`
}

type request struct {
	RequestID   string    `json:"requestId"`
	Cmd         []cmd     `json:"cmd"`
	PipeMapping []pipeMap `json:"pipeMapping"`
}

type status envexec.Status

func (s status) MarshalJSON() ([]byte, error) {
	return []byte("\"" + (envexec.Status)(s).String() + "\""), nil
}

type response struct {
	Status     status            `json:"status"`
	ExitStatus int               `json:"exitStatus"`
	Error      string            `json:"error,omitempty"`
	Time       uint64            `json:"time"`
	Memory     uint64            `json:"memory"`
	Files      map[string]string `json:"files,omitempty"`
	FileIDs    map[string]string `json:"fileIds,omitempty"`
}

type result struct {
	RequestID string     `json:"requestId"`
	Response  []response `json:"results"`
	Error     error      `json:"-"`
	ErrorMsg  string     `json:"error,omitempty"`
}