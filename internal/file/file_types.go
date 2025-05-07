package file

var FileExtensionsByCategory = map[string][]string{
	"documents": {
		".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx",
		".pdf", ".txt", ".rtf", ".md", ".odt", ".ods", ".odp",
	},
	"images": {
		".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp",
		".svg", ".tif", ".tiff", ".raw", ".cr2", ".nef", ".arw",
	},
	"videos": {
		".mp4", ".webm", ".avi", ".mov", ".mkv", ".flv", ".wmv", ".3gp",
	},
	"audio": {
		".mp3", ".wav", ".aac", ".ogg", ".flac", ".m4a",
	},
	"archives": {
		".zip", ".rar", ".tar", ".gz", ".7z",
	},
	"code": {
		".go", ".js", ".mjs", ".ts", ".py", ".java",
		".c", ".cpp", ".h", ".hpp", ".html", ".htm", ".css",
		".sh", ".bash", ".php", ".json", ".yaml", ".yml",
	},
}
