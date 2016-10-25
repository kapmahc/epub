package epub

//Opf content.opf
type Opf struct {
	Metadata meta       `xml:"metadata" json:"metadata"`
	Manifest []manifest `xml:"manifest>item" json:"manifest"`
	Spine    spine      `xml:"spine" json:"spine"`
}

type meta struct {
	Title       []string     `xml:"title" json:"title"`
	Language    []string     `xml:"language" json:"language"`
	Identifier  []identifier `xml:"identifier" json:"identifier"`
	Creator     []author     `xml:"creator" json:"creator"`
	Subject     []string     `xml:"subject" json:"subject"`
	Description []string     `xml:"description" json:"description"`
	Publisher   []string     `xml:"publisher" json:"publisher"`
	Contributor []author     `xml:"contributor" json:"contributor"`
	Date        []date       `xml:"date" json:"date"`
	Type        []string     `xml:"type" json:"type"`
	Format      []string     `xml:"format" json:"format"`
	Source      []string     `xml:"source" json:"source"`
	Relation    []string     `xml:"relation" json:"relation"`
	Coverage    []string     `xml:"coverage" json:"coverage"`
	Rights      []string     `xml:"rights" json:"rights"`
	Meta        []metafield  `xml:"meta" json:"meta"`
}
type identifier struct {
	Data   string `xml:",chardata" json:"data"`
	ID     string `xml:"id,attr" json:"id"`
	Scheme string `xml:"scheme,attr" json:"scheme"`
}
type author struct {
	Data   string `xml:",chardata" json:"author"`
	FileAs string `xml:"file-as,attr" json:"file_as"`
	Role   string `xml:"role,attr" json:"role"`
}
type date struct {
	Data  string `xml:",chardata" json:"data"`
	Event string `xml:"event,attr" json:"event"`
}
type metafield struct {
	Name    string `xml:"name,attr" json:"name"`
	Content string `xml:"content,attr" json:"content"`
}
type manifest struct {
	ID           string `xml:"id,attr" json:"id"`
	Href         string `xml:"href,attr" json:"href"`
	MediaType    string `xml:"media-type,attr" json:"type"`
	Fallback     string `xml:"media-fallback,attr" json:"fallback"`
	Properties   string `xml:"properties,attr" json:"properties"`
	MediaOverlay string `xml:"media-overlay,attr" json:"overlay"`
}
type spine struct {
	ID              string      `xml:"id,attr" json:"id"`
	Toc             string      `xml:"toc,attr" json:"toc"`
	PageProgression string      `xml:"page-progression-direction,attr" json:"progression"`
	Items           []spineItem `xml:"itemref" json:"items"`
}
type spineItem struct {
	IDref      string `xml:"idref,attr" json:"id_ref"`
	Linear     string `xml:"linear,attr" json:"linear"`
	ID         string `xml:"id,attr" json:"id"`
	Properties string `xml:"properties,attr" json:"properties"`
}
