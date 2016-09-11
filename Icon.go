package goalfred

// IconType describes the two possible values for the Type of an icon
type IconType string

// From https://www.alfredapp.com/help/workflows/inputs/script-filter/json/
// By omitting the "type", Alfred will load the file path itself, for example a png.
// By using "type": "fileicon", Alfred will get the icon for the specified path.
// Finally, by using "type": "filetype", you can get the icon of a specific file, for example "path": "public.png"

// NoIconType makes Alfred load the file path itself
var NoIconType IconType

// FileIconType makes Alfred get the icon for the specified path
var FileIconType IconType = "fileicon"

// FileTypeIconType makes Alfed get the icon of a specific file
var FileTypeIconType IconType = "filetype"

// Icon holds all information about an item's icon
type Icon struct {
	Type IconType `json:"type,omitempty"`
	Path string   `json:"path,omitempty"`
}
