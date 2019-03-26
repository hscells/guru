package guru

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"testing"
)

var data = `<eSummaryResult>
    <DocumentSummarySet status="OK">
        <DbBuild>Build190315-0747.1</DbBuild>
        <DocumentSummary uid="428">
            <ConceptId>C0004031</ConceptId>
            <Title>Allergic bronchopulmonary aspergillosis</Title>
            <Definition></Definition>
            <SemanticId>T047</SemanticId>
            <SemanticType>Disease or Syndrome</SemanticType>
            <Suppressed/>
            <ConceptMeta>
                <Names>
                    <Name SDUI="D001229" SCUI="M0001851" CODE="D001229" SAB="MSH" TTY="PM" type="syn">Allergic Aspergilloses, Bronchopulmonary</Name>
                    <Name SDUI="D001229" SCUI="M0001851" CODE="D001229" SAB="MSH" TTY="PM" type="syn">Allergic Aspergillosis, Bronchopulmonary</Name>
                    <Name SDUI="D001229" SCUI="M0001851" CODE="D001229" SAB="MSH" TTY="PM" type="syn">Aspergilloses, Bronchopulmonary Allergic</Name>
                    <Name SDUI="D001229" SCUI="M0001851" CODE="D001229" SAB="MSH" TTY="MH" type="syn">Aspergillosis, Allergic Bronchopulmonary</Name>
                    <Name SDUI="D001229" SCUI="M0001851" CODE="D001229" SAB="MSH" TTY="PM" type="syn">Bronchopulmonary Allergic Aspergilloses</Name>
                    <Name SDUI="D001229" SCUI="M0001851" CODE="D001229" SAB="MSH" TTY="PM" type="syn">Bronchopulmonary Allergic Aspergillosis</Name>
                    <Name SCUI="C84547" CODE="C84547" SAB="NCI" TTY="PT" type="syn">Allergic Bronchopulmonary Aspergillosis</Name>
                    <Name SDUI="D001229" SCUI="M0001851" CODE="D001229" SAB="MSH" TTY="ET" type="syn">Bronchopulmonary Aspergillosis, Allergic</Name>
                    <Name SDUI="D001229" SCUI="M0001851" CODE="D001229" SAB="MSH" TTY="ET" type="syn">Allergic Bronchopulmonary Aspergillosis</Name>
                    <Name SDUI="D001229" SCUI="M0001851" CODE="D001229" SAB="MSH" TTY="ET" type="syn">Allergic Bronchopulmonary Aspergilloses</Name>
                    <Name SDUI="D001229" SCUI="M0001851" CODE="D001229" SAB="MSH" TTY="ET" type="syn">Aspergillosis, Bronchopulmonary Allergic</Name>
                    <Name SDUI="D001229" SCUI="M0001851" CODE="D001229" SAB="MSH" TTY="ET" type="syn">Bronchopulmonary Aspergilloses, Allergic</Name>
                    <Name SDUI="D001229" SCUI="M0001851" CODE="D001229" SAB="MSH" TTY="ET" type="syn">Aspergilloses, Allergic Bronchopulmonary</Name>
                    <Name SCUI="37981002" CODE="37981002" SAB="SNOMEDCT_US" TTY="PT" type="preferred">Allergic bronchopulmonary aspergillosis</Name>
                    <Name SCUI="37981002" CODE="37981002" SAB="SNOMEDCT_US" TTY="SY" type="syn">ABPA - Allergic bronchopulmonary aspergillosis</Name>
                    <Name SDUI="Orphanet_1164" CODE="AN0440356" SAB="ORDO" TTY="PT" type="preferred">Allergic bronchopulmonary aspergillosis</Name>
                    <Name SDUI="Orphanet_1164" CODE="AN0450565" SAB="ORDO" TTY="SYN" type="syn">ABPA</Name>
                    <Name SDUI="Orphanet_1164" CODE="AN0450566" SAB="ORDO" TTY="SYN" type="syn">Allergic aspergillosis</Name>
                    <Name SDUI="Orphanet_1164" CODE="AN0450567" SAB="ORDO" TTY="SYN" type="syn">Hinson-Pepys disease</Name>
                </Names>
                <Definitions>
                    <Definition source="NCI">An immune reaction to aspergillus in patients usually suffering from asthma or cystic fibrosis.</Definition>
                </Definitions>
                <ModesOfInheritance>
                    <ModeOfInheritance uid="141047" CUI="C0443147" TUI="T170">
                        <Name>Autosomal dominant inheritance</Name>
                        <SemanticType>Intellectual Product</SemanticType>
                        <Definition>A mode of inheritance that is observed for traits related to a gene encoded on one of the autosomes (i.e., the human chromosomes 1-22) in which a trait manifests in heterozygotes. In the context of medical genetics, an autosomal dominant disorder is caused when a single copy of the mutant allele is present. Males and females are affected equally, and can both transmit the disorder with a risk of 50% for each child of inheriting the mutant allele.</Definition>
                        <SAB>HPO</SAB>
                        <SAB>OMIM</SAB>
                        <SAB>ORDO</SAB>
                    </ModeOfInheritance>
                    <ModeOfInheritance uid="832438" CUI="CN227390" TUI="T170">
                        <Name>not inherited</Name>
                        <SemanticType>Intellectual Product</SemanticType>
                        <Definition>Describes a disorder that is not inherited.</Definition>
                        <SAB>ORDO</SAB>
                    </ModeOfInheritance>
                </ModesOfInheritance>
                <PharmacologicResponse/>
                <PharmacologicResponse/>
                <OMIM/>
                <ClinicalFeatures>
                    <ClinicalFeature uid="867388" CUI="C4021753" TUI="T046" SDUI="HP:0002715">
                        <Name>Abnormality of the immune system</Name>
                        <SemanticType>Pathologic Function</SemanticType>
                        <Definition>An abnormality of the immune system.</Definition>
                    </ClinicalFeature>
                </ClinicalFeatures>
                <PhenotypicAbnormalities>
                    <Category CUI="C4021753" name="Abnormality of the immune system">
                        <ClinicalFeature uid="867388" CUI="C4021753" TUI="T046">
                            <SemanticType>Pathologic Function</SemanticType>
                            <Name>Abnormality of the immune system</Name>
                            <Definition>An abnormality of the immune system.</Definition>
                        </ClinicalFeature>
                    </Category>
                </PhenotypicAbnormalities>
                <RelatedDisorders/>
                <SNOMEDCT>
                    <Name SAUI="1229371018" SCUI="37981002" SAB="SNOMEDCT_US" TTY="SY">ABPA - Allergic bronchopulmonary aspergillosis</Name>
                    <Name SAUI="63349014" SCUI="37981002" SAB="SNOMEDCT_US" TTY="PT">Allergic bronchopulmonary aspergillosis</Name>
                </SNOMEDCT>
                <AssociatedGenes/>
                <ORDO>
                </ORDO>
                <SemanticTypes>
                    <SemanticType TUI="T047">Disease or Syndrome</SemanticType>
                </SemanticTypes>
            </ConceptMeta>
            <ModificationDate/>
            <Merged/>
        </DocumentSummary>
    </DocumentSummarySet>
</eSummaryResult>`

func TestName(t *testing.T) {
	var summary CeSummaryResult
	err := xml.NewDecoder(bytes.NewBufferString(data)).Decode(&summary)
	if err != nil {
		panic(err)
	}
	fmt.Println(summary)
}
