package engine

// DefaultCourseID scopes a learner who has not chosen a level: every Primary level at once.
// The corpus is 1250 concepts and about 130 of them are Primary, so an unscoped learner was
// offered calculus and topology from a product sold as Primary maths. The other concepts stay
// in the corpus; they are only out of a student's reach until a course that contains them is
// chosen.
const DefaultCourseID = "sg-p6"

// LevelScope is the set of concepts the scheduler may propose to a student: their chosen
// course's targets and every prerequisite of them. It reads the course from the student row on
// each call, so a level survives a restart and a change takes effect on the next question.
//
// A course the planner does not know falls back to DefaultCourseID, and an engine with no
// planner or no default course is unscoped, which is what tests built on a small DAG expect.
// Explicit requests for one concept do not pass through here, so a learner can still open a
// topic above their level when they ask for it.
func (e *Engine) LevelScope(studentID string) map[string]bool {
	if e.planner == nil {
		return nil
	}
	courseID := ""
	if studentID != "" {
		if st, err := e.repo.GetStudent(studentID); err == nil && st != nil {
			courseID = st.CourseID
		}
	}
	if scope := e.courseScope(courseID); scope != nil {
		return scope
	}
	return e.courseScope(DefaultCourseID)
}

func (e *Engine) courseScope(courseID string) map[string]bool {
	if courseID == "" {
		return nil
	}
	if v, ok := e.courseScopes.Load(courseID); ok {
		return v.(map[string]bool)
	}
	path, err := e.planner.PathForCourse(courseID)
	if err != nil || len(path.Concepts) == 0 {
		return nil
	}
	set := make(map[string]bool, len(path.Concepts))
	for _, c := range path.Concepts {
		set[c.ID] = true
	}
	e.courseScopes.Store(courseID, set)
	return set
}
