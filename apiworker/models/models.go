package models

type CertificateTemplate struct {
	TemplateKey    string `json:"template_key,omitempty"`
	TemplateSource string `json:"template_source,omitempty"`
	SourceType     string `json:"source_type,omitempty"`
	Version        string `json:"version,omitempty"`
	IssuerID       string `json:"issuer_id,omitempty"`
	IssuerName     string `json:"issuer_name,omitempty"`
}

type CertificateDetails struct {
	CertificateID        string      `json:"certificate_id,omitempty"`
	CertificateHolder    string      `json:"certificate_holder,omitempty"`
	TemplateRef          string      `json:"template_ref,omitempty"`
	CourseName           string      `json:"course_name,omitempty"`
	ModuleName           string      `json:"module_name,omitempty"`
	Email                string      `json:"email,omitempty"`
	IssuedAt             string      `json:"issued_at,omitempty"`
	IssuerName           string      `json:"issuer_name,omitempty"`
	CertificateSignature string      `json:"certificate_signature,omitempty"`
	Extras               interface{} `json:"extras,omitempty"`
}

type CertificateBatchParam struct {
	OrganizationName string                `json:"organization_name,omitempty"`
	CertTemplate     CertificateTemplate   `json:"template,omitempty"`
	CertDetails      []*CertificateDetails `json:"certificate_details,omitempty"`
	CallbackUrl      string                `json:"callback_url,omitempty"`
}
