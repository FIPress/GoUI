package filepicker

/*

 */
import "C"
import goui "github.com/fipress/GoUI"

type Settings struct {
	IsSave               bool   `json:"isSave"` //save or open, default to open
	Title                string `json:"title"`
	Message              string `json:"message"`
	Multiple             bool   `json:"multiple"`
	FileOnly             bool   `json:"fileOnly"`
	DirOnly              bool   `json:"dirOnly"`
	Accept               string `json:"accept"`
	AllowsOtherFileTypes bool   `json:"allowsOtherFileTypes"`
	StartLocation        string `json:"startLocation"`

	SuggestedFileName string `json:"suggestedFileName"`
	CanCreateDir      bool   `json:"canCreateDir"`
	ShowsHiddenFiles  bool   `json:"showsHiddenFiles"`

	//ViewMode
	//FutureAccessList
	//MostRecentlyUsedList
}

func BoolToCInt(b bool) (i C.int) {
	if b {
		i = 1
	}
	return
}

/*
func convertSettings(settings *Settings) C.Settings {
	return C.Settings{C.CString(settings.Message),
		C.CString(settings.AllowedFileTypes),
		C.CString(settings.StartLocation),
		C.CString(settings.SuggestedFileName),
		BoolToCInt(settings.AllowsMultiple),
		BoolToCInt(settings.AllowsFile),
		BoolToCInt(settings.AllowsDir),
		BoolToCInt(settings.AllowsOtherFileTypes),
		BoolToCInt(settings.CanCreateDir),
		BoolToCInt(settings.ShowsHiddenFiles),
	}
}*/

type FilePicker struct {
}

func (f *FilePicker) Register() {
	goui.Service("filepicker", func(context *goui.Context) {
		s := new(Settings)
		err := context.GetEntity(&s)
		if err != nil {
			goui.Logf("Get filepicker settings failed:", err.Error())
			context.Success("")
		} else {
			//action := context.GetParam("action")
			filename := openFilePicker(s)
			context.Success(filename)
		}
	})
}
