package ctx

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime/debug"
	"sort"
	"strings"
	"sync"
	"time"
)

const (
	backupTimeFormat = "2006-01-02T15-04-05.000"
	defaultMaxSize   = 100
)

type FileLogger struct {
	Filename   string `json:"filename"`
	MaxSize    int    `json:"maxsize"`    // 单个日志最大的容量
	MaxAge     int    `json:"maxage"`     //
	MaxBackups int    `json:"maxbackups"` // 日志的有效时间，如30天有效，超过30天的日志删除
	LocalTime  bool   `json:"localtime" `

	size int64
	file *os.File
	mu   sync.Mutex
}

var (
	_           io.WriteCloser = (*FileLogger)(nil)
	Level                      = 1
	currentTime                = time.Now
	os_Stat                    = os.Stat
	megabyte                   = 1024 * 1024
)

func (l *FileLogger) Write(p []byte) (n int, err error) {

	l.mu.Lock()
	defer l.mu.Unlock()

	now := []byte(time.Now().Format("2006-01-02 15:04:05 - "))
	writeLen := int64(len(p)) + int64(len(now))

	if writeLen > l.max() {
		return 0, fmt.Errorf(
			"write length %d exceeds maximum file size %d", writeLen, l.max(),
		)
	}

	if l.file == nil {
		if err = l.openExistingOrNew(len(p) + len(now)); err != nil {
			return 0, err
		}
	}

	if l.size+writeLen > l.max() {
		if err := l.rotate(); err != nil {
			return 0, err
		}
	}

	n1, err := l.file.Write(now)
	n2, err := l.file.Write(p)
	l.size += int64(n1 + n2)
	fmt.Print(time.Now().Format("01-02 15:04:05"))
	fmt.Print(" - ")
	fmt.Print(string(p)) //Just print-back to stdin agan.

	return n, err

}

func (l *FileLogger) Close() error {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.close()
}

func (l *FileLogger) close() error {
	if l.file == nil {
		return nil

	}
	err := l.file.Close()
	l.file = nil
	return err

}

// 关闭现有日志文件，打开一个新的
func (l *FileLogger) Rotate() error {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.rotate()

}

// 关闭现有日志文件，打开一个新的
func (l *FileLogger) rotate() error {
	if err := l.close(); err != nil {
		return err
	}
	if err := l.openNew(); err != nil {
		return err
	}
	return l.cleanup()
}

// 打开一个新的日志文件
func (l *FileLogger) openNew() error {

	err := os.MkdirAll(l.dir(), 0744)

	if err != nil {
		return fmt.Errorf("can't make directories for new logfile: %s", err)

	}

	name := l.filename()
	mode := os.FileMode(0644)
	info, err := os_Stat(name)

	if err == nil {
		mode = info.Mode()
		newname := backupName(name, l.LocalTime)
		if err := os.Rename(name, newname); err != nil {
			return fmt.Errorf("can't rename log file: %s", err)
		}
	}

	f, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, mode)
	if err != nil {
		return fmt.Errorf("can't open new logfile: %s", err)
	}

	l.file = f
	l.size = 0
	return nil

}

// 将当前的日志文件修改为离职日志文件名称
func backupName(name string, local bool) string {
	dir := filepath.Dir(name)
	filename := filepath.Base(name)
	ext := filepath.Ext(filename)
	prefix := filename[:len(filename)-len(ext)]
	t := currentTime()
	if !local {
		t = t.UTC()
	}
	timestamp := t.Format(backupTimeFormat)
	return filepath.Join(dir, fmt.Sprintf("%s-%s%s", prefix, timestamp, ext))
}

func (l *FileLogger) openExistingOrNew(writeLen int) error {
	filename := l.filename()
	info, err := os_Stat(filename)

	if os.IsNotExist(err) {
		return l.openNew()
	}

	if err != nil {
		return fmt.Errorf("error getting log file info: %s", err)
	}

	if info.Size()+int64(writeLen) >= l.max() {
		return l.rotate()
	}

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return l.openNew()
	}

	l.file = file
	l.size = info.Size()

	return nil
}

func (l *FileLogger) filename() string {
	if l.Filename != "" {
		return l.Filename
	}

	name := filepath.Base(os.Args[0]) + ".log"
	return filepath.Join(os.TempDir(), name)
}

func (l *FileLogger) cleanup() error {
	if l.MaxBackups == 0 && l.MaxAge == 0 {
		return nil
	}
	files, err := l.oldLogFiles()
	if err != nil {
		return err
	}

	var deletes []logInfo
	if l.MaxBackups > 0 && l.MaxBackups < len(files) {
		deletes = files[l.MaxBackups:]
		files = files[:l.MaxBackups]
	}

	if l.MaxAge > 0 {
		diff := time.Duration(int64(24*time.Hour) * int64(l.MaxAge))
		cutoff := currentTime().Add(-1 * diff)
		for _, f := range files {
			if f.timestamp.Before(cutoff) {
				deletes = append(deletes, f)
			}
		}
	}

	if len(deletes) == 0 {
		return nil
	}

	go deleteAll(l.dir(), deletes)
	return nil

}

// 删除过期的日志文件
func deleteAll(dir string, files []logInfo) {
	for _, f := range files {
		_ = os.Remove(filepath.Join(dir, f.Name()))
	}
}

func (l *FileLogger) oldLogFiles() ([]logInfo, error) {
	files, err := ioutil.ReadDir(l.dir())
	if err != nil {
		return nil, fmt.Errorf("can't read log file directory: %s", err)
	}
	logFiles := []logInfo{}
	prefix, ext := l.prefixAndExt()
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		name := l.timeFromName(f.Name(), prefix, ext)
		if name == "" {
			continue
		}
		t, err := time.Parse(backupTimeFormat, name)
		if err == nil {
			logFiles = append(logFiles, logInfo{t, f})
		}
	}
	sort.Sort(byFormatTime(logFiles))
	return logFiles, nil
}

func (l *FileLogger) timeFromName(filename, prefix, ext string) string {

	if !strings.HasPrefix(filename, prefix) {
		return ""

	}

	// <压缩日志>
	//if strings.HasSuffix(filename, ext+".gz") {
	//return filename[len(prefix) : len(filename)-len(ext+".gz")]
	//}

	// <常规日志>
	if strings.HasSuffix(filename, ext) {
		return filename[len(prefix) : len(filename)-len(ext)]
	}

	return ""

}

func (l *FileLogger) max() int64 {
	if l.MaxSize == 0 {
		return int64(defaultMaxSize * megabyte)

	}
	return int64(l.MaxSize) * int64(megabyte)

}

func (l *FileLogger) dir() string {
	return filepath.Dir(l.filename())

}

func (l *FileLogger) prefixAndExt() (prefix, ext string) {
	filename := filepath.Base(l.filename())
	ext = filepath.Ext(filename)
	prefix = filename[:len(filename)-len(ext)] + "-"
	return prefix, ext

}

type logInfo struct {
	timestamp time.Time
	os.FileInfo
}

type byFormatTime []logInfo

func (b byFormatTime) Less(i, j int) bool {
	return b[i].timestamp.After(b[j].timestamp)

}

func (b byFormatTime) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]

}

func (b byFormatTime) Len() int {
	return len(b)

}

func _log(level string, format string, args ...interface{}) {
	msg := fmt.Sprintf(format+"\r\n", args...)
	log.Printf(fmt.Sprintf("[%s] - %s", level, msg))
}

// Info [INFO] 级日志输出，不带错误栈
func LogInfo(format string, args ...interface{}) {
	if Level > 1 {
		_log("INFO", format, args...)
	}
}

// Error [ERROR] 级日志输出，附带错误栈
func LogError(format string, args ...interface{}) {
	_log("ERROR", format, args...)
	log.Println(string(debug.Stack()))
}

var logBean *FileLogger

func SetLogConfig(ci IContext) {
	conf := ci.GetContext()
	Level = conf.Log.Level
	logBean = &FileLogger{
		Filename:   conf.Log.Filename,
		MaxSize:    conf.Log.MaxSize,
		MaxBackups: conf.Log.MaxBackups,
		MaxAge:     conf.Log.MaxAge,
	}
	log.SetOutput(logBean)
}

func LogBean() *FileLogger {
	return logBean
}
