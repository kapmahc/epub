package epub

import "archive/zip"

//Open open a epub file
func Open(fn string) (*Book, error) {
	fd, err := zip.OpenReader(fn)
	if err != nil {
		return nil, err
	}

	bk := Book{fd: fd}
	mt, err := bk.readBytes("mimetype")
	if err == nil {
		bk.Mimetype = string(mt)
		err = bk.readXML("META-INF/container.xml", &bk.Container)
	}
	if err == nil {
		err = bk.readXML(bk.Container.Rootfile.Path, &bk.Opf)
	}

	for _, mf := range bk.Opf.Manifest {
		if mf.ID == bk.Opf.Spine.Toc {
			err = bk.readXML(bk.filename(mf.Href), &bk.Ncx)
			break
		}
	}

	if err != nil {
		fd.Close()
		return nil, err
	}

	return &bk, nil
}
