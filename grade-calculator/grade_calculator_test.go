package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
    gc := NewGradeCalculator()
    gc.AddGrade("assignment", 100, Assignment)
    gc.AddGrade("exam", 100, Exam)
    gc.AddGrade("essay", 100, Essay)

    if gc.GetFinalGrade() != "A" {
        t.Errorf("Expected 'A', got '%s'", gc.GetFinalGrade())
    }
}

func TestGetGradeB(t *testing.T) {
    gc := NewGradeCalculator()
    gc.AddGrade("assignment", 80, Assignment)
    gc.AddGrade("exam", 81, Exam)
    gc.AddGrade("essay", 85, Essay)

    if gc.GetFinalGrade() != "B" {
        t.Errorf("Expected 'B', got '%s'", gc.GetFinalGrade())
    }
}

func TestGetGradeC(t *testing.T) {
    gc := NewGradeCalculator()
    gc.AddGrade("assignment", 70, Assignment)
    gc.AddGrade("exam", 70, Exam)
    gc.AddGrade("essay", 70, Essay)

    if gc.GetFinalGrade() != "C" {
        t.Errorf("Expected 'C', got '%s'", gc.GetFinalGrade())
    }
}

func TestGetGradeD(t *testing.T) {
    gc := NewGradeCalculator()
    gc.AddGrade("assignment", 60, Assignment)
    gc.AddGrade("exam", 60, Exam)
    gc.AddGrade("essay", 60, Essay)

    if gc.GetFinalGrade() != "D" {
        t.Errorf("Expected 'D', got '%s'", gc.GetFinalGrade())
    }
}

func TestGetGradeF(t *testing.T) {
    gc := NewGradeCalculator()
    gc.AddGrade("assignment", 60, Assignment)
    gc.AddGrade("exam", 51, Exam)
    gc.AddGrade("essay", 45, Essay)

    if gc.GetFinalGrade() != "F" {
        t.Errorf("Expected 'F', got '%s'", gc.GetFinalGrade())
    }
}

func TestAddGrade(t *testing.T) {
    gc := NewGradeCalculator()
    gc.AddGrade("assignment", 100, Assignment)
    gc.AddGrade("exam", 90, Exam)
    gc.AddGrade("essay", 80, Essay)

    if len(gc.grades) != 3 {
        t.Errorf("Expected 3 grades, got %d", len(gc.grades))
    }
}

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

func TestComputeAverageEmptySlice(t *testing.T) {
    var emptyGrades []Grade
    avg := computeAverage(emptyGrades)

    if avg != 0 {
        t.Errorf("Expected average of empty slice to be 0, got %d", avg)
    }
}
