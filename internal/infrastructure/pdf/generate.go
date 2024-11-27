package pdf

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"os/exec"
)

func generatePDF(htmlContent string, fileName string) (string, error) {
	outputPath := fmt.Sprintf("./generated-pdf/%s.pdf", fileName)

	tempHTMLFile, err := os.CreateTemp("", "*.html")
	if err != nil {
		log.Printf("Failed to create temporary HTML file: %v", err)
		return "", fmt.Errorf("failed to create temporary HTML file: %v", err)
	}
	defer os.Remove(tempHTMLFile.Name())

	_, err = tempHTMLFile.WriteString(htmlContent)
	if err != nil {
		log.Printf("Failed to write to temporary HTML file: %v", err)
		return "", fmt.Errorf("failed to write to temporary HTML file: %v", err)
	}

	err = tempHTMLFile.Close()
	if err != nil {
		log.Printf("Failed to close temporary HTML file: %v", err)
		return "", fmt.Errorf("failed to close temporary HTML file: %v", err)
	}

	cmd := exec.Command("wkhtmltopdf", tempHTMLFile.Name(), outputPath)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err = cmd.Run()
	if err != nil {
		log.Printf("Failed to generate PDF: %v\nOutput: %s", err, out.String())
		return "", fmt.Errorf("failed to generate PDF: %v\nOutput: %s", err, out.String())
	}

	return outputPath, nil
}

func GenerateAgreementLetter(attributes AgreementLetterAttributes) (string, error) {
	tmpl, err := template.ParseFiles("./internal/infrastructure/pdf/templates/agreement_letter_template.html")
	if err != nil {
		log.Println("Failed to parse template:", err)
		return "", err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, attributes)
	if err != nil {
		log.Println("Failed to execute template:", err)
		return "", err
	}

	generatedHTML := buf.String()
	fileName := fmt.Sprintf("%s_agreement_letter_%s_%s", attributes.AgreementDate, attributes.BorrowerName, attributes.InvestorName)
	filePath, err := generatePDF(generatedHTML, fileName)
	if err != nil {
		log.Println("Failed to generate PDF:", err)
		return "", err
	}

	return filePath, nil
}
