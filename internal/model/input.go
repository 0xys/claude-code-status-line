package model

type Input struct {
	HookEventName  string        `json:"hook_event_name,omitempty"`
	SessionID      string        `json:"session_id,omitempty"`
	TranscriptPath string        `json:"transcript_path,omitempty"`
	Cwd            string        `json:"cwd,omitempty"`
	Model          Model         `json:"model,omitempty"`
	Workspace      Workspace     `json:"workspace,omitempty"`
	Version        string        `json:"version,omitempty"`
	OutputStyle    OutputStyle   `json:"output_style,omitempty"`
	Cost           Cost          `json:"cost,omitempty"`
	ContextWindow  ContextWindow `json:"context_window,omitempty"`
}
type Model struct {
	ID          string `json:"id,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
}
type Workspace struct {
	CurrentDir string `json:"current_dir,omitempty"`
	ProjectDir string `json:"project_dir,omitempty"`
}
type OutputStyle struct {
	Name string `json:"name,omitempty"`
}
type Cost struct {
	TotalCostUsd       float64 `json:"total_cost_usd,omitempty"`
	TotalDurationMs    int     `json:"total_duration_ms,omitempty"`
	TotalAPIDurationMs int     `json:"total_api_duration_ms,omitempty"`
	TotalLinesAdded    int     `json:"total_lines_added,omitempty"`
	TotalLinesRemoved  int     `json:"total_lines_removed,omitempty"`
}
type CurrentUsage struct {
	InputTokens              int `json:"input_tokens,omitempty"`
	OutputTokens             int `json:"output_tokens,omitempty"`
	CacheCreationInputTokens int `json:"cache_creation_input_tokens,omitempty"`
	CacheReadInputTokens     int `json:"cache_read_input_tokens,omitempty"`
}
type ContextWindow struct {
	TotalInputTokens    int          `json:"total_input_tokens,omitempty"`
	TotalOutputTokens   int          `json:"total_output_tokens,omitempty"`
	ContextWindowSize   int          `json:"context_window_size,omitempty"`
	UsedPercentage      float64      `json:"used_percentage,omitempty"`
	RemainingPercentage float64      `json:"remaining_percentage,omitempty"`
	CurrentUsage        CurrentUsage `json:"current_usage,omitempty"`
}
