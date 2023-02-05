package auth

import "context"

const (
	subjectKey = "subject"
)

func NewContextWithSubject(parent context.Context, subject string) context.Context {
	return context.WithValue(parent, subjectKey, subject)
}

func SubjectFromContext(ctx context.Context) (string, bool) {
	subject, ok := ctx.Value(subjectKey).(string)

	return subject, ok
}
