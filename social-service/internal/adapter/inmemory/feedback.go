package inmemory

import (
	"sync"

	"github.com/sherinur/doit-platform/social-service/internal/model"
)

type Feedback struct {
	feedbacks map[string]model.Feedback
	m         sync.RWMutex
}

func NewFeedback() *Feedback {
	return &Feedback{
		feedbacks: make(map[string]model.Feedback),
		m:         sync.RWMutex{},
	}
}

func (f *Feedback) Set(feedback model.Feedback) {
	f.m.Lock()
	defer f.m.Unlock()

	f.feedbacks[feedback.ID] = feedback
}

func (f *Feedback) SetMany(feedbacks []model.Feedback) {
	f.m.Lock()
	defer f.m.Unlock()

	for _, feedback := range feedbacks {
		f.feedbacks[feedback.ID] = feedback
	}
}

func (f *Feedback) Get(feedbackID string) (model.Feedback, bool) {
	f.m.RLock()
	defer f.m.RUnlock()

	feedback, ok := f.feedbacks[feedbackID]

	return feedback, ok
}

func (f *Feedback) Delete(feedbackID string) {
	f.m.Lock()
	defer f.m.Unlock()

	delete(f.feedbacks, feedbackID)
}
