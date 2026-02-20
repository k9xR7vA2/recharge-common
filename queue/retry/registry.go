package retry

import "sync"

type PolicyRegistry struct {
	mu       sync.RWMutex
	policies map[string]Strategy
}

func NewPolicyRegistry() *PolicyRegistry {
	return &PolicyRegistry{policies: make(map[string]Strategy)}
}

func (r *PolicyRegistry) Register(taskType string, s Strategy) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.policies[taskType] = s
}

func (r *PolicyRegistry) Get(taskType string) (Strategy, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	s, ok := r.policies[taskType]
	return s, ok
}
