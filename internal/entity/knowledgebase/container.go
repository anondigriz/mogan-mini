package knowledgebase

type Container struct {
	KnowledgeBase KnowledgeBase
	Groups        map[string]Group
	Parameters    map[string]Parameter
	Patterns      map[string]Pattern
	Rules         map[string]Rule
}
