package esepunittests

import "math"

type GradeCalculator struct {
    grades []Grade
}

type GradeType int

const (
    Assignment GradeType = iota
    Exam
    Essay
)

var gradeTypeName = map[GradeType]string{
    Assignment: "assignment",
    Exam:       "exam",
    Essay:      "essay",
}

func (gt GradeType) String() string {
    return gradeTypeName[gt]
}

type Grade struct {
    Name  string
    Grade int
    Type  GradeType
}

func NewGradeCalculator() *GradeCalculator {
    return &GradeCalculator{
        grades: make([]Grade, 0),
    }
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
    gc.grades = append(gc.grades, Grade{
        Name:  name,
        Grade: grade,
        Type:  gradeType,
    })
}

func (gc *GradeCalculator) GetFinalGrade() string {
    numericalGrade := gc.calculateNumericalGrade()

    if numericalGrade >= 90 {
        return "A"
    } else if numericalGrade >= 80 {
        return "B"
    } else if numericalGrade >= 70 {
        return "C"
    } else if numericalGrade >= 60 {
        return "D"
    }
    return "F"
}

// New method for Pass/Fail support
func (gc *GradeCalculator) IsPassFail() string {
    if gc.calculateNumericalGrade() >= 70 {
        return "Pass"
    }
    return "Fail"
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
    var assignments, exams, essays []Grade

    for _, g := range gc.grades {
        switch g.Type {
        case Assignment:
            assignments = append(assignments, g)
        case Exam:
            exams = append(exams, g)
        case Essay:
            essays = append(essays, g)
        }
    }

    assignmentAvg := computeAverage(assignments)
    examAvg := computeAverage(exams)
    essayAvg := computeAverage(essays)

    weighted := float64(assignmentAvg)*0.5 + float64(examAvg)*0.35 + float64(essayAvg)*0.15
    return int(math.Round(weighted))
}

func computeAverage(grades []Grade) int {
    if len(grades) == 0 {
        return 0
    }

    sum := 0
    for _, g := range grades {
        sum += g.Grade
    }
    return sum / len(grades)
}
