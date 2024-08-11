package bbConvert

const (
	MDSurroundRegEx       = "(**|*|__|_|`)(.*?)(\\1)"
	MDLargeCodeblockRegEx = "```([\\s\\S])```"
	MDImgAndLinkRegEx     = `[!]?\[(.*?)\]\((.*?)\)`
)
