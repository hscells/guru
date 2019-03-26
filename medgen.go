package guru

/*
THIS FILE WAS AUTOMATICALLY GENERATED USING THE FOLLOWING COMMAND

	~/go/bin/chidley -G esummary.fcgi.xml > medgen.go

source: https://github.com/gnewton/chidley/
 */

import "encoding/xml"

type CAssociatedGenes struct {
	XMLName xml.Name `xml:"AssociatedGenes,omitempty" json:"AssociatedGenes,omitempty"`
}

type CCategory struct {
	XMLName          xml.Name         `xml:"Category,omitempty" json:"Category,omitempty"`
	AttrCUI          string           `xml:"CUI,attr"  json:",omitempty"`
	Attrname         string           `xml:"name,attr"  json:",omitempty"`
	CClinicalFeature CClinicalFeature `xml:"ClinicalFeature,omitempty" json:"ClinicalFeature,omitempty"`
}

type CClinicalFeature struct {
	XMLName       xml.Name      `xml:"ClinicalFeature,omitempty" json:"ClinicalFeature,omitempty"`
	AttrCUI       string        `xml:"CUI,attr"  json:",omitempty"`
	AttrSDUI      string        `xml:"SDUI,attr"  json:",omitempty"`
	AttrTUI       string        `xml:"TUI,attr"  json:",omitempty"`
	Attruid       string        `xml:"uid,attr"  json:",omitempty"`
	CDefinition   CDefinition   `xml:"Definition,omitempty" json:"Definition,omitempty"`
	CName         []CName       `xml:"Name,omitempty" json:"Name,omitempty"`
	CSemanticType CSemanticType `xml:"SemanticType,omitempty" json:"SemanticType,omitempty"`
}

type CClinicalFeatures struct {
	XMLName          xml.Name         `xml:"ClinicalFeatures,omitempty" json:"ClinicalFeatures,omitempty"`
	CClinicalFeature CClinicalFeature `xml:"ClinicalFeature,omitempty" json:"ClinicalFeature,omitempty"`
}

type CConceptId struct {
	XMLName xml.Name `xml:"ConceptId,omitempty" json:"ConceptId,omitempty"`
	Value   string   `xml:",chardata" json:",omitempty"`
}

type CConceptMeta struct {
	XMLName                  xml.Name                 `xml:"ConceptMeta,omitempty" json:"ConceptMeta,omitempty"`
	CAssociatedGenes         CAssociatedGenes         `xml:"AssociatedGenes,omitempty" json:"AssociatedGenes,omitempty"`
	CClinicalFeatures        CClinicalFeatures        `xml:"ClinicalFeatures,omitempty" json:"ClinicalFeatures,omitempty"`
	CDefinitions             CDefinitions             `xml:"Definitions,omitempty" json:"Definitions,omitempty"`
	CModesOfInheritance      CModesOfInheritance      `xml:"ModesOfInheritance,omitempty" json:"ModesOfInheritance,omitempty"`
	CNames                   CNames                   `xml:"Names,omitempty" json:"Names,omitempty"`
	COMIM                    COMIM                    `xml:"OMIM,omitempty" json:"OMIM,omitempty"`
	CORDO                    CORDO                    `xml:"ORDO,omitempty" json:"ORDO,omitempty"`
	CPharmacologicResponse   []CPharmacologicResponse `xml:"PharmacologicResponse,omitempty" json:"PharmacologicResponse,omitempty"`
	CPhenotypicAbnormalities CPhenotypicAbnormalities `xml:"PhenotypicAbnormalities,omitempty" json:"PhenotypicAbnormalities,omitempty"`
	CRelatedDisorders        CRelatedDisorders        `xml:"RelatedDisorders,omitempty" json:"RelatedDisorders,omitempty"`
	CSNOMEDCT                CSNOMEDCT                `xml:"SNOMEDCT,omitempty" json:"SNOMEDCT,omitempty"`
	CSemanticTypes           CSemanticTypes           `xml:"SemanticTypes,omitempty" json:"SemanticTypes,omitempty"`
}

type CDbBuild struct {
	XMLName xml.Name `xml:"DbBuild,omitempty" json:"DbBuild,omitempty"`
	Value   string   `xml:",chardata" json:",omitempty"`
}

type CDefinition struct {
	XMLName    xml.Name `xml:"Definition,omitempty" json:"Definition,omitempty"`
	Attrsource string   `xml:"source,attr"  json:",omitempty"`
	Value      string   `xml:",chardata" json:",omitempty"`
}

type CDefinitions struct {
	XMLName     xml.Name    `xml:"Definitions,omitempty" json:"Definitions,omitempty"`
	CDefinition CDefinition `xml:"Definition,omitempty" json:"Definition,omitempty"`
}

type CDocumentSummary struct {
	XMLName           xml.Name          `xml:"DocumentSummary,omitempty" json:"DocumentSummary,omitempty"`
	Attruid           string            `xml:"uid,attr"  json:",omitempty"`
	CConceptId        CConceptId        `xml:"ConceptId,omitempty" json:"ConceptId,omitempty"`
	CConceptMeta      CConceptMeta      `xml:"ConceptMeta,omitempty" json:"ConceptMeta,omitempty"`
	CDefinition       CDefinition       `xml:"Definition,omitempty" json:"Definition,omitempty"`
	CMerged           CMerged           `xml:"Merged,omitempty" json:"Merged,omitempty"`
	CModificationDate CModificationDate `xml:"ModificationDate,omitempty" json:"ModificationDate,omitempty"`
	CSemanticId       CSemanticId       `xml:"SemanticId,omitempty" json:"SemanticId,omitempty"`
	CSemanticType     CSemanticType     `xml:"SemanticType,omitempty" json:"SemanticType,omitempty"`
	CSuppressed       CSuppressed       `xml:"Suppressed,omitempty" json:"Suppressed,omitempty"`
	CTitle            CTitle            `xml:"Title,omitempty" json:"Title,omitempty"`
}

type CDocumentSummarySet struct {
	XMLName          xml.Name           `xml:"DocumentSummarySet,omitempty" json:"DocumentSummarySet,omitempty"`
	Attrstatus       string             `xml:"status,attr"  json:",omitempty"`
	CDbBuild         CDbBuild           `xml:"DbBuild,omitempty" json:"DbBuild,omitempty"`
	CDocumentSummary []CDocumentSummary `xml:"DocumentSummary,omitempty" json:"DocumentSummary,omitempty"`
}

type CHierarchyUrl struct {
	XMLName xml.Name `xml:"HierarchyUrl,omitempty" json:"HierarchyUrl,omitempty"`
	Value   string   `xml:",chardata" json:",omitempty"`
}

type CMerged struct {
	XMLName xml.Name `xml:"Merged,omitempty" json:"Merged,omitempty"`
}

type CModeOfInheritance struct {
	XMLName       xml.Name      `xml:"ModeOfInheritance,omitempty" json:"ModeOfInheritance,omitempty"`
	AttrCUI       string        `xml:"CUI,attr"  json:",omitempty"`
	AttrTUI       string        `xml:"TUI,attr"  json:",omitempty"`
	Attruid       string        `xml:"uid,attr"  json:",omitempty"`
	CDefinition   CDefinition   `xml:"Definition,omitempty" json:"Definition,omitempty"`
	CName         []CName       `xml:"Name,omitempty" json:"Name,omitempty"`
	CSAB          []CSAB        `xml:"SAB,omitempty" json:"SAB,omitempty"`
	CSemanticType CSemanticType `xml:"SemanticType,omitempty" json:"SemanticType,omitempty"`
}

type CModesOfInheritance struct {
	XMLName            xml.Name             `xml:"ModesOfInheritance,omitempty" json:"ModesOfInheritance,omitempty"`
	CModeOfInheritance []CModeOfInheritance `xml:"ModeOfInheritance,omitempty" json:"ModeOfInheritance,omitempty"`
}

type CModificationDate struct {
	XMLName xml.Name `xml:"ModificationDate,omitempty" json:"ModificationDate,omitempty"`
}

type CName struct {
	XMLName  xml.Name `xml:"Name,omitempty" json:"Name,omitempty"`
	AttrCODE string   `xml:"CODE,attr"  json:",omitempty"`
	AttrSAB  string   `xml:"SAB,attr"  json:",omitempty"`
	AttrSAUI string   `xml:"SAUI,attr"  json:",omitempty"`
	AttrSCUI string   `xml:"SCUI,attr"  json:",omitempty"`
	AttrSDUI string   `xml:"SDUI,attr"  json:",omitempty"`
	AttrTTY  string   `xml:"TTY,attr"  json:",omitempty"`
	Attrtype string   `xml:"type,attr"  json:",omitempty"`
	Value    string   `xml:",chardata" json:",omitempty"`
}

type CNames struct {
	XMLName xml.Name `xml:"Names,omitempty" json:"Names,omitempty"`
	CName   []CName  `xml:"Name,omitempty" json:"Name,omitempty"`
}

type COMIM struct {
	XMLName xml.Name `xml:"OMIM,omitempty" json:"OMIM,omitempty"`
}

type CORDO struct {
	XMLName       xml.Name      `xml:"ORDO,omitempty" json:"ORDO,omitempty"`
	CHierarchyUrl CHierarchyUrl `xml:"HierarchyUrl,omitempty" json:"HierarchyUrl,omitempty"`
}

type CPharmacologicResponse struct {
	XMLName xml.Name `xml:"PharmacologicResponse,omitempty" json:"PharmacologicResponse,omitempty"`
}

type CPhenotypicAbnormalities struct {
	XMLName   xml.Name  `xml:"PhenotypicAbnormalities,omitempty" json:"PhenotypicAbnormalities,omitempty"`
	CCategory CCategory `xml:"Category,omitempty" json:"Category,omitempty"`
}

type CRelatedDisorders struct {
	XMLName xml.Name `xml:"RelatedDisorders,omitempty" json:"RelatedDisorders,omitempty"`
}

type CSAB struct {
	XMLName xml.Name `xml:"SAB,omitempty" json:"SAB,omitempty"`
	Value   string   `xml:",chardata" json:",omitempty"`
}

type CSNOMEDCT struct {
	XMLName xml.Name `xml:"SNOMEDCT,omitempty" json:"SNOMEDCT,omitempty"`
	CName   []CName  `xml:"Name,omitempty" json:"Name,omitempty"`
}

type CSemanticId struct {
	XMLName xml.Name `xml:"SemanticId,omitempty" json:"SemanticId,omitempty"`
	Value   string   `xml:",chardata" json:",omitempty"`
}

type CSemanticType struct {
	XMLName xml.Name `xml:"SemanticType,omitempty" json:"SemanticType,omitempty"`
	AttrTUI string   `xml:"TUI,attr"  json:",omitempty"`
	Value   string   `xml:",chardata" json:",omitempty"`
}

type CSemanticTypes struct {
	XMLName       xml.Name      `xml:"SemanticTypes,omitempty" json:"SemanticTypes,omitempty"`
	CSemanticType CSemanticType `xml:"SemanticType,omitempty" json:"SemanticType,omitempty"`
}

type CSuppressed struct {
	XMLName xml.Name `xml:"Suppressed,omitempty" json:"Suppressed,omitempty"`
}

type CTitle struct {
	XMLName xml.Name `xml:"Title,omitempty" json:"Title,omitempty"`
	Value   string   `xml:",chardata" json:",omitempty"`
}

type CeSummaryResult struct {
	XMLName             xml.Name            `xml:"eSummaryResult,omitempty" json:"eSummaryResult,omitempty"`
	CDocumentSummarySet CDocumentSummarySet `xml:"DocumentSummarySet,omitempty" json:"DocumentSummarySet,omitempty"`
}
