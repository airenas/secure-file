package util

// Params is application parameters
type Params struct {
	Secret   string
	File     string
	FileList string
}

// SecureFileExt is additional extension for secured file
const SecureFileExt = ".aes"
