package models

import (
	"regexp"
	"strings"
)

// Document nó com o tipo de documento e número do documento
type Document struct {
	Type   string
	Number string
}

// IsCpf diz se o DocumentType é um CPF
func (d Document) IsCpf() bool {
	return strings.ToUpper(d.Type) == "CPF"
}

// IsCnpj diz se o DocumentType é um CNPJ
func (d Document) IsCnpj() bool {
	return strings.ToUpper(d.Type) == "CNPJ"
}

// ValidateCPF verifica se é um CPF válido
func (d *Document) ValidateCPF() error {
	re := regexp.MustCompile("(\\D+)")
	cpf := re.ReplaceAllString(string(d.Number), "")
	if len(cpf) == 11 {
		d.Number = cpf
		return nil
	}
	return NewErrorResponse("MPDocumentNumber", "CPF inválido")
}

// ValidateCNPJ verifica se é um CNPJ válido
func (d *Document) ValidateCNPJ() error {
	re := regexp.MustCompile("(\\D+)")
	cnpj := re.ReplaceAllString(string(d.Number), "")
	if len(cnpj) == 14 {
		d.Number = cnpj
		return nil
	}
	return NewErrorResponse("MPDocumentNumber", "CNPJ inválido")
}
