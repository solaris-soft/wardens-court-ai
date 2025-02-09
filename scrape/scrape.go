package scrape

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

// Decision holds the decision details as key/value pairs.
type Decision struct {
	Details map[string]string
}

func scrapeDecisionPanel() ([]Decision, error) {
	fmt.Println("Scraping decision panel...")

	// The actual endpoint that handles the form submission
	targetURL := "https://emits.dmp.wa.gov.au/emits/advert/wardenCourt/wardenCourtDecisionPanel.xhtml"

	// Create form data
	formData := url.Values{}
	formData.Set("javax.faces.partial.ajax", "true")
	formData.Set("javax.faces.source", "decisionForm:j_idt42")
	formData.Set("javax.faces.partial.execute", "@all")
	formData.Set("javax.faces.partial.render", "decisionForm:decisionTable")
	formData.Set("decisionForm:j_idt42", "decisionForm:j_idt42")
	formData.Set("decisionForm", "decisionForm")

	// Create a client with custom headers
	client := &http.Client{}
	req, err := http.NewRequest("POST", targetURL, strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Set necessary headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
	req.Header.Set("Accept", "application/xml, text/xml, */*; q=0.01")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Faces-Request", "partial/ajax")

	// Make the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	content := string(body)
	fmt.Printf("Response status: %d\n", resp.StatusCode)
	fmt.Printf("Response content length: %d\n", len(content))

	// Parse the response and extract decisions
	decisions := parseDecisions(content)
	fmt.Printf("Found %v decisions\n", decisions)

	return decisions, nil
}

func parseDecisions(content string) []Decision {
	var decisions []Decision

	// Look for decision entries in the table
	decisionBlocks := strings.Split(content, `<tr class="rowEven"`)

	for _, block := range decisionBlocks[1:] { // Skip first split as it's before the first decision
		decision := Decision{
			Details: make(map[string]string),
		}

		// Extract each field using the actual HTML structure
		fields := map[string]string{
			"Decision Number": `<span class="label">Decision Number</span></td><td class="columnRightLocal">`,
			"Date Delivered":  `<span class="label">Date Delivered</span></td><td class="columnRightLocal">`,
			"Court":           `<span class="label">Court</span></td><td class="columnRightLocal">`,
			"Warden":          `<span class="label">Warden</span></td><td class="columnRightLocal">`,
			"Tenement(s)":     `<span class="label">Tenement(s)</span></td><td class="columnRightLocal">`,
			"Section Reg No":  `<span class="label">Section Reg No</span></td><td class="columnRightLocal">`,
			"Parties":         `<span class="label">Parties</span></td><td class="columnRightLocal">`,
			"Summary":         `<span class="label">Summary</span></td><td class="columnRightLocal">`,
		}

		for key, marker := range fields {
			if value := extractAfterMarker(block, marker); value != "" {
				decision.Details[key] = cleanText(value)
			}
		}

		// Only add decisions that have at least some details
		if len(decision.Details) > 0 {
			decisions = append(decisions, decision)
		}
	}

	return decisions
}

func extractAfterMarker(content, marker string) string {
	parts := strings.Split(content, marker)
	if len(parts) < 2 {
		return ""
	}

	// Find the end of the value (closing td tag)
	endIdx := strings.Index(parts[1], "</td>")
	if endIdx == -1 {
		return ""
	}

	return parts[1][:endIdx]
}

func cleanText(text string) string {
	// Remove HTML tags
	text = strings.ReplaceAll(text, "<br>", " ")
	text = strings.ReplaceAll(text, "<br/>", " ")
	text = strings.ReplaceAll(text, "<p>", " ")
	text = strings.ReplaceAll(text, "</p>", " ")

	// Remove any remaining HTML tags
	tagRegex := regexp.MustCompile(`<[^>]*>`)
	text = tagRegex.ReplaceAllString(text, "")

	// Clean up whitespace
	text = strings.TrimSpace(text)
	text = strings.ReplaceAll(text, "\n", " ")
	text = strings.ReplaceAll(text, "\t", " ")
	text = strings.ReplaceAll(text, "&nbsp;", " ")

	// Replace multiple spaces with single space
	for strings.Contains(text, "  ") {
		text = strings.ReplaceAll(text, "  ", " ")
	}

	return text
}

// ScrapeDecisions fetches and returns decisions from the panel
func ScrapeDecisions() ([]Decision, error) {
	decisions, err := scrapeDecisionPanel()
	if err != nil {
		return nil, fmt.Errorf("error scraping decisions: %v", err)
	}

	return decisions, nil
}
