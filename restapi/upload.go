package restapi

import (
	"errors"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/zfile"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zstring"
	"github.com/sohaha/zlsgo/ztype"
)

type UploadOptions struct {
	MaxSize       uint64
	MaxFiles      int
	AllowTypes    []string
	AllowExts     []string
	FormName      string
	Dir           string
	SafeDir       bool
	ConveFilePath func(file, name, dir string) string
}

// Upload Upload file
func Upload(c *znet.Context, o ...func(options *UploadOptions)) (files []string, err error) {
	opt := UploadOptions{
		MaxSize:    1024 * 1024 * 5,
		MaxFiles:   1,
		AllowTypes: []string{"image/jpeg", "image/png", "image/gif"},
		AllowExts:  []string{"jpg", "jpeg", "png", "gif"},
		FormName:   "file",
		SafeDir:    true,
		ConveFilePath: func(file, name, dir string) string {
			n, err := zstring.Md5File(file)
			if err != nil {
				return file
			}

			ext := filepath.Ext(name)
			if len(ext) > 0 {
				n += ext
			}
			return dir + n
		},
	}
	for _, f := range o {
		f(&opt)
	}

	filesHeader, err := c.FormFiles(opt.FormName)
	if err != nil {
		return nil, err
	}

	if len(filesHeader) > opt.MaxFiles {
		return nil, errors.New("more than the maximum number of files: " + ztype.ToString(opt.MaxFiles))
	}

	files = make([]string, 0, len(filesHeader))

	defer func() {
		if err != nil {
			for _, file := range files {
				_ = zfile.Remove(file)
			}
		}
	}()

	programPath, uploadDir := zfile.ProgramPath(), zfile.RealPathMkdir(opt.Dir, true)

	if opt.SafeDir && !strings.Contains(uploadDir, programPath) {
		return nil, errors.New("upload directory is not safe")
	}

	for _, fileHeader := range filesHeader {
		path := uploadDir + fileHeader.Filename

		err = fileAllow(fileHeader, opt.MaxSize, opt.AllowTypes, opt.AllowExts)
		if err != nil {
			return
		}

		err = c.SaveUploadedFile(fileHeader, path)
		if err != nil {
			return
		}

		if opt.ConveFilePath != nil {
			path = opt.ConveFilePath(path, fileHeader.Filename, uploadDir)
		}
		files = append(files, path)
	}

	if opt.SafeDir {
		for i := range files {
			files[i] = zfile.SafePath(files[i], programPath)
		}
	}

	return files, nil
}

func fileAllow(fileHeader *multipart.FileHeader, maxSize uint64, allowTypes, allowExts []string) (err error) {
	if uint64(fileHeader.Size) > maxSize {
		return errors.New("file exceeds maximum limit: " + zfile.SizeFormat(maxSize))
	}

	ext := filepath.Ext(fileHeader.Filename)
	if len(ext) > 1 && len(allowExts) > 0 && !zarray.Contains(allowExts, ext[1:]) {
		return errors.New("file types only suppor: " + strings.Join(allowExts, ","))
	}

	if len(allowTypes) > 0 {
		buf := make([]byte, 512)
		var f multipart.File
		f, err = fileHeader.Open()
		if err != nil {
			return err
		}
		_, _ = f.Read(buf)

		allow, ctype := false, zfile.GetMimeType(fileHeader.Filename, buf)
		for i := range allowTypes {
			if strings.Contains(ctype, allowTypes[i]) {
				allow = true
				break
			}
		}
		if !allow {
			return errors.New("file types only suppor: " + strings.Join(allowTypes, ","))
		}
	}

	return
}
