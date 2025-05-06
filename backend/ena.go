package main

type ENAresponse struct {
    HostBodySite string `json:"host_body_site"`
    HostTaxID string `json:"host_tax_id"`
    Country string `"country"`
    SubmittedFTP string `json:"submitted_ftp"`
}

/*
baseURL := "https://www.ebi.ac.uk/ena/portal/api/results?dataPortal=ena"
    queryParams := url.Values{}
    queryParams.Set("result", "read_run")
    queryParams.Set("query", `country="United Kingdom" AND host_tax_id=9913 AND host_body_site="rumen"`)
    queryParams.Set("fields", "submitted_ftp,host_body_site,host_tax_id,country")
    queryParams.Set("format", "json")

    fullURL := fmt.Sprintf("%s?%s", baseURL, queryParams.Encode())
    fmt.Println(fullURL)

    resp, err := http.Get(fullURL)
    if err != nil {
        fmt.Println("Error making request:", err)
        return
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response body:", err)
        return
    }

    var results []ENAresponse
    if err := json.Unmarshal(body, &results); err != nil {
        fmt.Println("Error parsing JSON:", err)
        return
    }

    fmt.Println("Raw Read Data Sets:")
    for _, record := range results {
        fmt.Printf("Country: %s, Host Tax ID: %s, Body Site: %s, FTP Link: %s\n", record.Country, record.HostTaxID, record.HostBodySite, record.SubmittedFTP)
    }
		*/