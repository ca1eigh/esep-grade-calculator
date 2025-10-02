package esepunittests

import "testing"

func TestIsPassFail_Pass(t *testing.T) {
    gc := NewGradeCalculator()
    gc.AddGrade("assignment", 90, Assignment)
    gc.AddGrade("exam", 90, Exam)
    gc.AddGrade("essay", 90, Essay)

    if gc.IsPassFail() != "Pass" {
        t.Errorf("Expected 'Pass', got '%s'", gc.IsPassFail())
    }
}

func TestIsPassFail_Fail(t *testing.T) {
    gc := NewGradeCalculator()
    gc.AddGrade("assignment", 50, Assignment)
    gc.AddGrade("exam", 50, Exam)
    gc.AddGrade("essay", 50, Essay)

    if gc.IsPassFail() != "Fail" {
        t.Errorf("Expected 'Fail', got '%s'", gc.IsPassFail())
    }
}
