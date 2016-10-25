package epub

//Container META-INF/container.xml file
type Container struct {
	Rootfile rootfile `xml:"rootfiles>rootfile" json:"rootfile"`
}

type rootfile struct {
	Path string `xml:"full-path,attr" json:"path"`
	Type string `xml:"media-type,attr" json:"type"`
}
