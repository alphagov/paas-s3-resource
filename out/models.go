package out

import (
	"github.com/alphagov/paas-s3-resource"
)

type Request struct {
	Source s3resource.Source `json:"source"`
	Params Params            `json:"params"`
}

type Params struct {
	From        string `json:"from"`
	File        string `json:"file"`
	To          string `json:"to"`
	Acl         string `json:"acl"`
	ContentType string `json:"content_type"`
}

type Response struct {
	Version  s3resource.Version        `json:"version"`
	Metadata []s3resource.MetadataPair `json:"metadata"`
}
