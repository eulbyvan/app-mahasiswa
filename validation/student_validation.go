package validation

import (
	"fmt"
	"os"
	"strconv"

	"github.com/eulbyvan/app-mahasiswa/entity"
)

func ValidateStudent(student entity.Student) error {
	minAge, err := strconv.Atoi(os.Getenv("MIN_STUDENT_AGE"))

	if err != nil {
		return fmt.Errorf("failed to read MIN_STUDENT AGE env variable: %v", err)
	}

	if student.Age < minAge {
		return fmt.Errorf("age must be at least %d", minAge)
	}

	return nil
}