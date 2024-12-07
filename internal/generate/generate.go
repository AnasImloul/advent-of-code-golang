package generate

import (
	"fmt"
	"github.com/AnasImloul/advent-of-code-golang/internal/input"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func GenerateFiles(year int, day int) error {
	// Format year and day
	yearStr := strconv.Itoa(year)
	dayStr := fmt.Sprintf("%02d", day)

	// Define paths
	basePath := filepath.Join(".", "internal", yearStr, fmt.Sprintf("day_%s", dayStr))
	templatePath := filepath.Join(".", "internal", "template")
	yearFilePath := filepath.Join(".", "internal", yearStr, "year.go")

	// Create the directory structure (idempotent)
	if err := os.MkdirAll(basePath, os.ModePerm); err != nil {
		return fmt.Errorf("error creating directory: %w", err)
	}

	// List of template files and corresponding output files
	files := []struct {
		TemplateFile string
		OutputFile   string
	}{
		{"base.txt", "base.go"},
		{"first.txt", "first.go"},
		{"second.txt", "second.go"},
	}

	// Process each template
	for _, file := range files {
		outputFilePath := filepath.Join(basePath, file.OutputFile)

		// Skip file creation if it already exists
		if _, err := os.Stat(outputFilePath); err == nil {
			fmt.Printf("File %s already exists, skipping...\n", outputFilePath)
			continue
		}

		templateFilePath := filepath.Join(templatePath, file.TemplateFile)
		templateContent, err := os.ReadFile(templateFilePath)
		if err != nil {
			return fmt.Errorf("error reading template file %s: %w", templateFilePath, err)
		}

		// Replace placeholders in the template
		customizedContent := strings.ReplaceAll(string(templateContent), "YEAR", yearStr)
		customizedContent = strings.ReplaceAll(customizedContent, "day_XX", fmt.Sprintf("day_%s", dayStr))
		customizedContent = strings.ReplaceAll(customizedContent, "Day:  XX", fmt.Sprintf("Day:  %d", day))

		// Write the customized content to the output file
		if err := os.WriteFile(outputFilePath, []byte(customizedContent), 0644); err != nil {
			return fmt.Errorf("error creating file %s: %w", outputFilePath, err)
		}
	}

	// Update the year.go file
	if err := updateYearFile(yearFilePath, year, day); err != nil {
		return fmt.Errorf("error updating year file: %w", err)
	}

	fmt.Printf("Folder and files created successfully at %s\n", basePath)
	return nil
}

func updateYearFile(yearFilePath string, year, day int) error {
	dayStr := fmt.Sprintf("%02d", day)

	// Read the existing year.go file
	content, err := os.ReadFile(yearFilePath)
	if err != nil {
		return fmt.Errorf("error reading year file: %w", err)
	}

	contentStr := string(content)

	// Prepare the import and case statements
	importStatement := fmt.Sprintf("\"github.com/AnasImloul/advent-of-code-golang/internal/%d/day_%s\"", year, dayStr)
	caseStatement := fmt.Sprintf("case %d:\n\t\treturn day_%s.Solver.Solve(part)", day, dayStr)

	// Check if the import statement exists
	importPattern := regexp.MustCompile(fmt.Sprintf(`\Q%s\E`, importStatement))
	if !importPattern.MatchString(contentStr) {
		// Add the import statement
		importStart := strings.Index(contentStr, "import (")
		if importStart == -1 {
			return fmt.Errorf("could not find import block in year.go")
		}
		importEnd := strings.Index(contentStr[importStart:], ")")
		if importEnd == -1 {
			return fmt.Errorf("could not find end of import block in year.go")
		}
		importEnd += importStart
		contentStr = contentStr[:importEnd] + "\t" + importStatement + "\n" + contentStr[importEnd:]
		fmt.Printf("Added import for day_%s\n", dayStr)
	} else {
		fmt.Printf("Import for day_%s already exists, skipping...\n", dayStr)
	}

	// Check if the case statement exists
	casePattern := regexp.MustCompile(fmt.Sprintf(`\Q%s\E`, caseStatement))
	if !casePattern.MatchString(contentStr) {
		// Add the case statement
		switchStart := strings.Index(contentStr, "switch day {")
		if switchStart == -1 {
			return fmt.Errorf("could not find switch block in year.go")
		}
		switchEnd := strings.Index(contentStr[switchStart:], "default:")
		if switchEnd == -1 {
			return fmt.Errorf("could not find default case in year.go")
		}
		switchEnd += switchStart
		contentStr = contentStr[:switchEnd] + caseStatement + "\n\t" + contentStr[switchEnd:]
		fmt.Printf("Added case for day %d\n", day)
	} else {
		fmt.Printf("Case for day %d already exists, skipping...\n", day)
	}

	// Write the updated content back to the file
	if err := os.WriteFile(yearFilePath, []byte(contentStr), 0644); err != nil {
		return fmt.Errorf("error writing year file: %w", err)
	}

	// Get input file
	input.Read(year, day)

	fmt.Printf("Updated year.go file at %s\n", yearFilePath)
	return nil
}
