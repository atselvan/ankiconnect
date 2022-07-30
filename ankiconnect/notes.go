package ankiconnect

type Fields struct {
	Front string `json:"Front"`
	Back string `json:"Back"`
}
type DuplicateScopeOptions struct {
	DeckName string `json:"deckName"`
	CheckChildren bool `json:"checkChildren"`
	CheckAllModels bool `json:"checkAllModels"`
}
type Options struct {
	AllowDuplicate bool `json:"allowDuplicate"`
	DuplicateScope string `json:"duplicateScope"`
	DuplicateScopeOptions DuplicateScopeOptions `json:"duplicateScopeOptions"`
}
type Audio struct {
	URL string `json:"url"`
	Filename string `json:"filename"`
	SkipHash string `json:"skipHash"`
	Fields []string `json:"fields"`
}
type Video struct {
	URL string `json:"url"`
	Filename string `json:"filename"`
	SkipHash string `json:"skipHash"`
	Fields []string `json:"fields"`
}
type Picture struct {
	URL string `json:"url"`
	Filename string `json:"filename"`
	SkipHash string `json:"skipHash"`
	Fields []string `json:"fields"`
}
type Note struct {
	DeckName string `json:"deckName"`
	ModelName string `json:"modelName"`
	Fields Fields `json:"fields"`
	Options Options `json:"options"`
	Tags []string `json:"tags"`
	Audio []Audio `json:"audio"`
	Video []Video `json:"video"`
	Picture []Picture `json:"picture"`
}
