package ghost

import "fmt"

func (g *Ghost) GetPages() (Pages, error) {
	const ghostPostsURLSuffix = "%s/ghost/api/v2/content/pages/?key=%s&limit=all"
	var pages Pages
	var url = fmt.Sprintf(ghostPostsURLSuffix, g.url, g.contentAPIToken)

	if err := g.getJson(url, &pages); err != nil {
		return pages, err
	}

	return pages, nil
}
