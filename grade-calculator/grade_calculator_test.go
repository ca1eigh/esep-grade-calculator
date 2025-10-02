package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}


func TestGetGradeC(t *testing.T) {
    gradeCalculator := NewGradeCalculator()
    gradeCalculator.AddGrade("assignment", 70, Assignment)
    gradeCalculator.AddGrade("exam", 70, Exam)
    gradeCalculator.AddGrade("essay", 70, Essay)

    if gradeCalculator.GetFinalGrade() != "C" {
        t.Errorf("Expected 'C', got '%s'", gradeCalculator.GetFinalGrade())
    }
}

func TestGetGradeD(t *testing.T) {
    gradeCalculator := NewGradeCalculator()
    gradeCalculator.AddGrade("assignment", 60, Assignment)
    gradeCalculator.AddGrade("exam", 60, Exam)
    gradeCalculator.AddGrade("essay", 60, Essay)

    if gradeCalculator.GetFinalGrade() != "D" {
        t.Errorf("Expected 'D', got '%s'", gradeCalculator.GetFinalGrade())
    }
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 60, Assignment)
	gradeCalculator.AddGrade("exam 1", 51, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 45, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

// Test adding grades correctly
func TestAddGradeAssignment(t *testing.T) {
    gradeCalculator := NewGradeCalculator()
    gradeCalculator.AddGrade("assignment", 100, Assignment)

    if len(gradeCalculator.assignments) != 1 {
        t.Errorf("Expected 1 assignment, got %d", len(gradeCalculator.assignments))
    }
}

// Test Grade Type String function
func TestGradeTypeString(t *testing.T) {
    tests := map[GradeType]string{
        Assignment: "assignment",
        Exam:       "exam",
        Essay:      "essay",
    }

    for gt, expected := range tests {
        if gt.String() != expected {
            t.Errorf("Expected GradeType %d to return '%s'; got '%s'", gt, expected, gt.String())
        }
    }
}

// Test compute average return 0 line
func TestComputeAverageEmptySlice(t *testing.T) {
    var emptyGrades []Grade

    avg := computeAverage(emptyGrades)

    if avg != 0 {
        t.Errorf("Expected average of empty slice to be 0, got %d", avg)
    }
}
