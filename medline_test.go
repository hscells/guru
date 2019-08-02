package guru_test

import (
	"bytes"
	"fmt"
	"github.com/hscells/guru"
	"testing"
)

func TestMedlineParse(t *testing.T) {
	doc := new(bytes.Buffer)
	doc.WriteString(`
PMID- 24399179
OWN - NLM
STAT- MEDLINE
DCOM- 20151117
LR  - 20181113
IS  - 1863-2661 (Electronic)
IS  - 1863-2653 (Linking)
VI  - 220
IP  - 2
DP  - 2015 Mar
TI  - Definition and characterization of an extended social-affective default network.
PG  - 1031-49
LID - 10.1007/s00429-013-0698-0 [doi]
AB  - Recent evidence suggests considerable overlap between the default mode network
      (DMN) and regions involved in social, affective and introspective processes. We
      considered these overlapping regions as the social-affective part of the DMN. In 
      this study, we established a robust mapping of the underlying brain network
      formed by these regions and those strongly connected to them (the extended
      social-affective default network). We first seeded meta-analytic connectivity
      modeling and resting-state analyses in the meta-analytically defined DMN regions 
      that showed statistical overlap with regions associated with social and affective
      processing. Consensus connectivity of each seed was subsequently delineated by a 
      conjunction across both connectivity analyses. We then functionally characterized
      the ensuing regions and performed several cluster analyses. Among the identified 
      regions, the amygdala/hippocampus formed a cluster associated with emotional
      processes and memory functions. The ventral striatum, anterior cingulum,
      subgenual cingulum and ventromedial prefrontal cortex formed a heterogeneous
      subgroup associated with motivation, reward and cognitive modulation of affect.
      Posterior cingulum/precuneus and dorsomedial prefrontal cortex were associated
      with mentalizing, self-reference and autobiographic information. The cluster
      formed by the temporo-parietal junction and anterior middle temporal sulcus/gyrus
      was associated with language and social cognition. Taken together, the current
      work highlights a robustly interconnected network that may be central to
      introspective, socio-affective, that is, self- and other-related mental
      processes.
FAU - Amft, Maren
AU  - Amft M
AD  - Institute of Clinical Neuroscience and Medical Psychology, HHU Dusseldorf,
      Dusseldorf, Germany.
FAU - Bzdok, Danilo
AU  - Bzdok D
FAU - Laird, Angela R
AU  - Laird AR
FAU - Fox, Peter T
AU  - Fox PT
FAU - Schilbach, Leonhard
AU  - Schilbach L
FAU - Eickhoff, Simon B
AU  - Eickhoff SB
LA  - eng
GR  - R01 MH074457/MH/NIMH NIH HHS/United States
GR  - R01-MH074457-01A1/MH/NIMH NIH HHS/United States
PT  - Journal Article
PT  - Meta-Analysis
PT  - Research Support, N.I.H., Extramural
PT  - Research Support, Non-U.S. Gov't
DEP - 20140108
PL  - Germany
TA  - Brain Struct Funct
JT  - Brain structure & function
JID - 101282001
SB  - IM
MH  - Adult
MH  - *Affect
MH  - Brain/cytology/*physiology
MH  - Brain Mapping/methods
MH  - Cluster Analysis
MH  - Cognition
MH  - Databases, Factual
MH  - Female
MH  - Humans
MH  - Magnetic Resonance Imaging
MH  - Male
MH  - Memory
MH  - Middle Aged
MH  - Models, Neurological
MH  - Motivation
MH  - Nerve Net/cytology/*physiology
MH  - Retrospective Studies
MH  - Self Concept
MH  - *Social Behavior
MH  - Theory of Mind
MH  - Young Adult
PMC - PMC4087104
MID - NIHMS554054
EDAT- 2014/01/09 06:00
MHDA- 2015/11/18 06:00
CRDT- 2014/01/09 06:00
PHST- 2013/08/26 00:00 [received]
PHST- 2013/12/26 00:00 [accepted]
PHST- 2014/01/09 06:00 [entrez]
PHST- 2014/01/09 06:00 [pubmed]
PHST- 2015/11/18 06:00 [medline]
AID - 10.1007/s00429-013-0698-0 [doi]
PST - ppublish
SO  - Brain Struct Funct. 2015 Mar;220(2):1031-49. doi: 10.1007/s00429-013-0698-0. Epub
      2014 Jan 8.
`)

	md := guru.UnmarshalMedline(doc)

	for _, d := range md {
		fmt.Println(d.String())
	}
}
