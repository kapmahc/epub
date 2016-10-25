package epub

//Ncx OPS/toc.ncx
type Ncx struct {
	Points []navpoint `xml:"navMap>navPoint" json:"points"`
}
type navpoint struct {
	Text    string     `xml:"navLabel>text" json:"text"`
	Content content    `xml:"content" json:"content"`
	Points  []navpoint `xml:"navPoint" json:"points"`
}
type content struct {
	Src string `xml:"src,attr" json:"src"`
}
