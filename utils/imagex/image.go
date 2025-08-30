package imagex

import "strings"

func SwitchContentType(suffix string) string {
	var contentType string
	switch {
	case strings.Contains(suffix, "png"):
		contentType = "image/png"
	case strings.Contains(suffix, "jpeg"), strings.Contains(suffix, "jpg"):
		contentType = "image/jpeg"
	case strings.Contains(suffix, "gif"):
		contentType = "image/gif"
	case strings.Contains(suffix, "webp"):
		contentType = "image/webp"
	case strings.Contains(suffix, "mp4"):
		contentType = "video/mp4"
	case strings.Contains(suffix, "webm"):
		contentType = "video/webm"
	case strings.Contains(suffix, "ogg"):
		contentType = "video/ogg"
	case strings.Contains(suffix, "mpeg"):
		contentType = "video/mpeg"
	case strings.Contains(suffix, "mp3"):
		contentType = "audio/mp3"
	case strings.Contains(suffix, "wav"):
		contentType = "audio/wav"
	case strings.Contains(suffix, "ogg"):
		contentType = "audio/ogg"
	case strings.Contains(suffix, "pdf"):
		contentType = "application/pdf"
	case strings.Contains(suffix, "doc"):
		contentType = "application/msword"
	case strings.Contains(suffix, "docx"):
		contentType = "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	case strings.Contains(suffix, "xls"):
		contentType = "application/vnd.ms-excel"
	case strings.Contains(suffix, "xlsx"):
		contentType = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	case strings.Contains(suffix, "ppt"):
		contentType = "application/vnd.ms-powerpoint"
	case strings.Contains(suffix, "pptx"):
		contentType = "application/vnd.openxmlformats-officedocument.presentationml.presentation"
	case strings.Contains(suffix, "zip"):
		contentType = "application/zip"
	case strings.Contains(suffix, "rar"):
		contentType = "application/x-rar-compressed"
	case strings.Contains(suffix, "7z"):
		contentType = "application/x-7z-compressed"
	case strings.Contains(suffix, "txt"):
		contentType = "text/plain"
	case strings.Contains(suffix, "md"):
		contentType = "text/markdown"
	case strings.Contains(suffix, "xml"):
		contentType = "text/xml"
	case strings.Contains(suffix, "html"):
		contentType = "text/html"
	case strings.Contains(suffix, "css"):
		contentType = "text/css"
	case strings.Contains(suffix, "js"):
		contentType = "text/javascript"
	case strings.Contains(suffix, "json"):
		contentType = "application/json"
	case strings.Contains(suffix, "ico"):
		contentType = "image/x-icon"
	case strings.Contains(suffix, "svg"):
		contentType = "image/svg+xml"
	case strings.Contains(suffix, "swf"):
		contentType = "application/x-shockwave-flash"
	case strings.Contains(suffix, "flv"):
		contentType = "video/x-flv"
	case strings.Contains(suffix, "wmv"):
		contentType = "video/x-ms-wmv"
	case strings.Contains(suffix, "avi"):
		contentType = "video/x-msvideo"
	case strings.Contains(suffix, "rmvb"):
		contentType = "application/vnd.rn-realmedia-vbr"
	case strings.Contains(suffix, "rm"):
		contentType = "application/vnd.rn-realmedia"
	default:
		contentType = "application/octet-stream"
	}

	return contentType
}
